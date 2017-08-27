#!/usr/bin/env python
import argparse
import itertools
import sys
import os
import os.path

import sh
from sh import cd, mkdir, glob, which, cp

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
    "hoststore",
    "protocols",
    "pubkey",
    "zlint",
]

grpc_prefixes = [
    "rpc",
    "search",
]

all_prefixes = protobuf_prefixes + grpc_prefixes

# Protobuf Command Wrappers
protoc = sh.Command("protoc")

directory = os.path.dirname(os.path.realpath(__file__))
proto_directory = os.path.join(directory, 'proto')

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


def main():
    parser = argparse.ArgumentParser(description='Compile some protobufs.')
    parser.add_argument(
        'target',
        metavar='TARGET',
        choices=TARGETS + [TARGET_ALL],
        nargs=1,
        help='one of {go, python, cpp, all}',
    )
    args = parser.parse_args()
    target = args.target[0]

    build_all = False
    if target == TARGET_ALL:
        build_all = True

    if build_all or target == TARGET_GO:
        build_go()
    elif build_all or target == TARGET_PYTHON:
        build_python()
    elif build_all or target == TARGET_CPP:
        build_cpp()
    return

if __name__ == '__main__':
    main()
