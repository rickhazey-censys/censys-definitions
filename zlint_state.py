#!/usr/bin/env python

import itertools
import json
import sys
from google.protobuf import descriptor_pb2
from google.protobuf.compiler import plugin_pb2


def traverse_messages(proto_file):
    def _traverse(items):
        for item in items:
            yield item
            if isinstance(item, descriptor_pb2.DescriptorProto):
                for enum in item.enum_type: 
                    yield enum
                for nested in item.nested_type:
                    for nested_item in _traverse(nested):
                        yield nested_item
    chain = list()
    for file_descriptor_proto in proto_file:
        chain.append(_traverse(file_descriptor_proto.enum_type))
        chain.append(_traverse(file_descriptor_proto.message_type))
    return itertools.chain(*chain)


def generate_code(request):
    existing_lints = dict()
    response = plugin_pb2.CodeGeneratorResponse()
    for item in traverse_messages(request.proto_file):
        if item.name != "Lints":
            continue
        if not isinstance(item, descriptor_pb2.DescriptorProto):
            raise Exception()
        for field in item.field:
            field_name = field.name
            if field.type != descriptor_pb2.FieldDescriptorProto.TYPE_MESSAGE:
                continue
            if field.type_name not in {"LintResult", ".zsearch.LintResult"}:
                continue
            field_id = field.number
            existing_lints[field_name] = field_id
    f = response.file.add()
    f.name = 'zlint-state.json'
    f.content = json.dumps(existing_lints) + "\n"
    return response


if __name__ == '__main__':
    data = sys.stdin.read()
    request = plugin_pb2.CodeGeneratorRequest()
    request.ParseFromString(data)
    response = generate_code(request)
    output = response.SerializeToString()
    sys.stdout.write(output)
