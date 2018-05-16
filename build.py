#!/usr/bin/env python
import argparse
import itertools
import json
import os
import os.path
import sys
import tempfile

import sh
from sh import cd, mkdir, glob, which, cp, mv

TARGET_GO = "go"
TARGET_PYTHON = "python"
TARGET_CPP = "cpp"

TARGETS = [
    TARGET_GO,
    TARGET_PYTHON,
    TARGET_CPP,
]
TARGET_ALL = "all"

protobuf_prefixes = [
    "anonstore",
    "certificate",
    "common",
    "ct",
    "caa",
    "hoststore",
    "protocols",
    "pubkey",
    "zlint",
]

grpc_prefixes = [
    "esloader",
    "rpc",
    "search",
]

all_prefixes = protobuf_prefixes + grpc_prefixes

# Protobuf Command Wrappers
protoc = sh.Command("protoc")

directory = os.path.dirname(os.path.realpath(__file__))
proto_directory = os.path.join(directory, 'proto')
zlint_proto_file = os.path.join(proto_directory, 'zlint.proto')

# Generate relative filepaths for the inputs
protobuf_inputs = [x + ".proto" for x in protobuf_prefixes]
grpc_inputs = [x + ".proto" for x in grpc_prefixes]
all_inputs = protobuf_inputs + grpc_inputs

# Generate absolute filepaths for the inputs
abs_protobuf_inputs = [os.path.join(directory, x) for x in protobuf_inputs]
abs_grpc_inputs = [os.path.join(directory, x) for x in grpc_inputs]
abs_all_inputs = abs_protobuf_inputs + abs_grpc_inputs


def build_go():
    protoc_gen_go_cmd = which('protoc-gen-go')
    go_directory = os.path.join(os.path.join(directory, 'go'), 'censys-definitions')
    cd(proto_directory)
    protoc("--go_out=plugins=grpc,import_path=censys_definitions:" + go_directory, glob('*.proto'))


def build_python():
    python_dir = os.path.join(os.path.join(directory, 'python'), 'zsearch_definitions')
    grpc_python_plugin_cmd = which('grpc_python_plugin')
    grpc_python_plugin = '='.join(['protoc-gen-grpc', grpc_python_plugin_cmd])
    cd(proto_directory)
    protoc("--python_out=" + python_dir, *all_inputs)
    protoc(
        "--python_out=" + python_dir,
        "--grpc_out=" + python_dir,
        "--plugin=" + grpc_python_plugin,
        *grpc_inputs
    )


def build_cpp():
    cpp_dir = os.path.join(directory, 'cpp')
    grpc_cpp_plugin_cmd = which('grpc_cpp_plugin')
    grpc_cpp_plugin = '='.join(['protoc-gen-grpc', grpc_cpp_plugin_cmd])
    cd(proto_directory)
    protoc("--cpp_out=" + cpp_dir, *all_inputs)
    protoc(
        "--grpc_out=" + cpp_dir,
        "--plugin=" + grpc_cpp_plugin,
        *grpc_inputs
    )


def generate_zlint_state():
    zlint_protoc_plugin_path = os.path.join(directory, 'zlint_state.py')
    cd(proto_directory)
    protoc(
        "--plugin=protoc-gen-custom=" + zlint_protoc_plugin_path,
        "--custom_out=" + directory,
        'zlint.proto',
    )


def generate_zlint_proto(zlint_state_path, zlint_list_path):
    sys.stderr.write("Using ZLint lint list file at {}\n".format(zlint_list_path))
    sys.stderr.write("Using ZLint state file at {}\n".format(zlint_state_path))
    fields = dict()
    fields['__start__'] = 0
    try:
        with open(zlint_state_path) as fd:
            obj = json.loads(fd.read())
            fields.update(obj)
    except IOError as e:
        if e.errno != 2:
            raise
    next_id = max(fields.itervalues()) + 1
    current_lints = set()
    with open(zlint_list_path) as fd:
        for line in fd:
            lint = json.loads(line.strip())
            lint_name = lint['name']
            field_id = fields.get(lint_name)
            if field_id is None:
                field_id = next_id
                next_id += 1
                fields[lint_name] = field_id
            current_lints.add(lint_name)
    with open(zlint_proto_file) as fd:
        with tempfile.NamedTemporaryFile(prefix='zlint.proto_', delete=False) as tmpfd:
            for line in fd:
                tmpfd.write(line)
                if line.startswith('//') and line[2:].strip() == '__begin_zlint_proto':
                    break
            tmpfd.write("\n")
            tmpfd.write("message Lints {\n")
            for field_name, field_id in sorted(fields.iteritems(), key=lambda x: x[1]):
                if field_name.startswith('__'):
                    continue
                if field_name not in current_lints:
                    continue
                field_def = "    LintResult {} = {};\n".format(field_name, field_id)
                tmpfd.write(field_def)
            id_comment = "    // Highest ID in use is currently {}\n".format(next_id - 1)
            tmpfd.write(id_comment)
            tmpfd.write("}\n")
    mv(tmpfd.name, zlint_proto_file)


def main():
    parser = argparse.ArgumentParser(description='Compile some protobufs.')
    parser.add_argument(
        'target',
        metavar='TARGET',
        choices=TARGETS + [TARGET_ALL],
        nargs='?',
        default='all',
        help='one of {go, python, cpp, all}',
    )
    parser.add_argument(
            '--generate-zlint-state',
            help='generate zlint-state.json based on zlint.proto',
            type=bool,
            nargs='?',
            default=False,
            const=True,
    )
    parser.add_argument(
            '--generate-zlint-proto',
            help='generate zlint.proto based on --zlint-lints and --zlint-state',
            type=bool,
            nargs='?',
            default=False,
            const=True,
    )
    parser.add_argument(
            '--zlint-lints',
            help='path to ZLint JSON description of lints',
            type=str,
            default='zlint-lints.json',
    )
    parser.add_argument(
            '--zlint-state',
            help='path to state file describing protobuf field numbers for lints',
            type=str,
            default='zlint-state.json',
    )

    args = parser.parse_args()
    target = args.target

    if args.generate_zlint_state:
        generate_zlint_state()
        return

    if args.generate_zlint_proto:
        generate_zlint_proto(args.zlint_state, args.zlint_lints)
        return

    build_all = False
    if target == TARGET_ALL:
        build_all = True

    if build_all or target == TARGET_GO:
        build_go()
    if build_all or target == TARGET_PYTHON:
        build_python()
    if build_all or target == TARGET_CPP:
        build_cpp()
    return


if __name__ == '__main__':
    main()
