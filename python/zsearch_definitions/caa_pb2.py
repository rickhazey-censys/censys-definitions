# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: caa.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='caa.proto',
  package='zsearch',
  syntax='proto3',
  serialized_pb=_b('\n\tcaa.proto\x12\x07zsearch\"D\n\x0b\x43\x41\x41TagValue\x12\x0c\n\x04\x66lag\x18\x01 \x01(\r\x12\x0b\n\x03tag\x18\x02 \x01(\t\x12\r\n\x05value\x18\x03 \x01(\t\x12\x0b\n\x03ttl\x18\x04 \x01(\r\"k\n\tCAARecord\x12\x0e\n\x06\x64omain\x18\x01 \x01(\t\x12(\n\x06result\x18\x02 \x01(\x0e\x32\x18.zsearch.CAADomainStatus\x12$\n\x06values\x18\x03 \x03(\x0b\x32\x14.zsearch.CAATagValue\"g\n\tCAALookup\x12\x11\n\ttimestamp\x18\x01 \x01(\x03\x12#\n\x07records\x18\x02 \x03(\x0b\x32\x12.zsearch.CAARecord\x12\"\n\x06result\x18\x03 \x01(\x0e\x32\x12.zsearch.CAAResult*\xb2\x01\n\tCAAResult\x12\x17\n\x13\x43\x41\x41_RESULT_RESERVED\x10\x00\x12!\n\x1d\x43\x41\x41_RESULT_VALIDATION_SUCCESS\x10\x01\x12\x1e\n\x1a\x43\x41\x41_RESULT_VALIDATION_FAIL\x10\x02\x12!\n\x1d\x43\x41\x41_RESULT_VALIDATION_SKIPPED\x10\x03\x12&\n\"CAA_RESULT_VALIDATION_NOT_REQUIRED\x10\x04*\xa9\x05\n\x0f\x43\x41\x41\x44omainStatus\x12\x1e\n\x1a\x43\x41\x41_DOMAIN_STATUS_RESERVED\x10\x00\x12(\n$CAA_DOMAIN_STATUS_VALIDATION_SUCCESS\x10\x01\x12%\n!CAA_DOMAIN_STATUS_VALIDATION_FAIL\x10\x02\x12(\n$CAA_DOMAIN_STATUS_VALIDATION_SKIPPED\x10\x03\x12\x1f\n\x1b\x43\x41\x41_DOMAIN_STATUS_DNS_ERROR\x10\x05\x12(\n$CAA_DOMAIN_STATUS_DNS_ERROR_SERVFAIL\x10\x06\x12(\n$CAA_DOMAIN_STATUS_DNS_ERROR_AUTHFAIL\x10\x07\x12)\n%CAA_DOMAIN_STATUS_DNS_ERROR_NO_RECORD\x10\x08\x12)\n%CAA_DOMAIN_STATUS_DNS_ERROR_BLACKLIST\x10\t\x12)\n%CAA_DOMAIN_STATUS_DNS_ERROR_NO_OUTPUT\x10\n\x12)\n%CAA_DOMAIN_STATUS_DNS_ERROR_NO_ANSWER\x10\x0b\x12-\n)CAA_DOMAIN_STATUS_DNS_ERROR_ILLEGAL_INPUT\x10\x0c\x12\'\n#CAA_DOMAIN_STATUS_DNS_ERROR_TIMEOUT\x10\r\x12,\n(CAA_DOMAIN_STATUS_DNS_ERROR_ITER_TIMEOUT\x10\x0e\x12)\n%CAA_DOMAIN_STATUS_DNS_ERROR_TEMPORARY\x10\x0f\x12)\n%CAA_DOMAIN_STATUS_DNS_ERROR_TRUNCATED\x10\x10\x62\x06proto3')
)
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

_CAARESULT = _descriptor.EnumDescriptor(
  name='CAAResult',
  full_name='zsearch.CAAResult',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='CAA_RESULT_RESERVED', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CAA_RESULT_VALIDATION_SUCCESS', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CAA_RESULT_VALIDATION_FAIL', index=2, number=2,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CAA_RESULT_VALIDATION_SKIPPED', index=3, number=3,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CAA_RESULT_VALIDATION_NOT_REQUIRED', index=4, number=4,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=307,
  serialized_end=485,
)
_sym_db.RegisterEnumDescriptor(_CAARESULT)

CAAResult = enum_type_wrapper.EnumTypeWrapper(_CAARESULT)
_CAADOMAINSTATUS = _descriptor.EnumDescriptor(
  name='CAADomainStatus',
  full_name='zsearch.CAADomainStatus',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='CAA_DOMAIN_STATUS_RESERVED', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CAA_DOMAIN_STATUS_VALIDATION_SUCCESS', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CAA_DOMAIN_STATUS_VALIDATION_FAIL', index=2, number=2,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CAA_DOMAIN_STATUS_VALIDATION_SKIPPED', index=3, number=3,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CAA_DOMAIN_STATUS_DNS_ERROR', index=4, number=5,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CAA_DOMAIN_STATUS_DNS_ERROR_SERVFAIL', index=5, number=6,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CAA_DOMAIN_STATUS_DNS_ERROR_AUTHFAIL', index=6, number=7,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CAA_DOMAIN_STATUS_DNS_ERROR_NO_RECORD', index=7, number=8,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CAA_DOMAIN_STATUS_DNS_ERROR_BLACKLIST', index=8, number=9,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CAA_DOMAIN_STATUS_DNS_ERROR_NO_OUTPUT', index=9, number=10,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CAA_DOMAIN_STATUS_DNS_ERROR_NO_ANSWER', index=10, number=11,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CAA_DOMAIN_STATUS_DNS_ERROR_ILLEGAL_INPUT', index=11, number=12,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CAA_DOMAIN_STATUS_DNS_ERROR_TIMEOUT', index=12, number=13,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CAA_DOMAIN_STATUS_DNS_ERROR_ITER_TIMEOUT', index=13, number=14,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CAA_DOMAIN_STATUS_DNS_ERROR_TEMPORARY', index=14, number=15,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CAA_DOMAIN_STATUS_DNS_ERROR_TRUNCATED', index=15, number=16,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=488,
  serialized_end=1169,
)
_sym_db.RegisterEnumDescriptor(_CAADOMAINSTATUS)

CAADomainStatus = enum_type_wrapper.EnumTypeWrapper(_CAADOMAINSTATUS)
CAA_RESULT_RESERVED = 0
CAA_RESULT_VALIDATION_SUCCESS = 1
CAA_RESULT_VALIDATION_FAIL = 2
CAA_RESULT_VALIDATION_SKIPPED = 3
CAA_RESULT_VALIDATION_NOT_REQUIRED = 4
CAA_DOMAIN_STATUS_RESERVED = 0
CAA_DOMAIN_STATUS_VALIDATION_SUCCESS = 1
CAA_DOMAIN_STATUS_VALIDATION_FAIL = 2
CAA_DOMAIN_STATUS_VALIDATION_SKIPPED = 3
CAA_DOMAIN_STATUS_DNS_ERROR = 5
CAA_DOMAIN_STATUS_DNS_ERROR_SERVFAIL = 6
CAA_DOMAIN_STATUS_DNS_ERROR_AUTHFAIL = 7
CAA_DOMAIN_STATUS_DNS_ERROR_NO_RECORD = 8
CAA_DOMAIN_STATUS_DNS_ERROR_BLACKLIST = 9
CAA_DOMAIN_STATUS_DNS_ERROR_NO_OUTPUT = 10
CAA_DOMAIN_STATUS_DNS_ERROR_NO_ANSWER = 11
CAA_DOMAIN_STATUS_DNS_ERROR_ILLEGAL_INPUT = 12
CAA_DOMAIN_STATUS_DNS_ERROR_TIMEOUT = 13
CAA_DOMAIN_STATUS_DNS_ERROR_ITER_TIMEOUT = 14
CAA_DOMAIN_STATUS_DNS_ERROR_TEMPORARY = 15
CAA_DOMAIN_STATUS_DNS_ERROR_TRUNCATED = 16



_CAATAGVALUE = _descriptor.Descriptor(
  name='CAATagValue',
  full_name='zsearch.CAATagValue',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='flag', full_name='zsearch.CAATagValue.flag', index=0,
      number=1, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='tag', full_name='zsearch.CAATagValue.tag', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='value', full_name='zsearch.CAATagValue.value', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='ttl', full_name='zsearch.CAATagValue.ttl', index=3,
      number=4, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=22,
  serialized_end=90,
)


_CAARECORD = _descriptor.Descriptor(
  name='CAARecord',
  full_name='zsearch.CAARecord',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='domain', full_name='zsearch.CAARecord.domain', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='result', full_name='zsearch.CAARecord.result', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='values', full_name='zsearch.CAARecord.values', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=92,
  serialized_end=199,
)


_CAALOOKUP = _descriptor.Descriptor(
  name='CAALookup',
  full_name='zsearch.CAALookup',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='timestamp', full_name='zsearch.CAALookup.timestamp', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='records', full_name='zsearch.CAALookup.records', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='result', full_name='zsearch.CAALookup.result', index=2,
      number=3, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=201,
  serialized_end=304,
)

_CAARECORD.fields_by_name['result'].enum_type = _CAADOMAINSTATUS
_CAARECORD.fields_by_name['values'].message_type = _CAATAGVALUE
_CAALOOKUP.fields_by_name['records'].message_type = _CAARECORD
_CAALOOKUP.fields_by_name['result'].enum_type = _CAARESULT
DESCRIPTOR.message_types_by_name['CAATagValue'] = _CAATAGVALUE
DESCRIPTOR.message_types_by_name['CAARecord'] = _CAARECORD
DESCRIPTOR.message_types_by_name['CAALookup'] = _CAALOOKUP
DESCRIPTOR.enum_types_by_name['CAAResult'] = _CAARESULT
DESCRIPTOR.enum_types_by_name['CAADomainStatus'] = _CAADOMAINSTATUS

CAATagValue = _reflection.GeneratedProtocolMessageType('CAATagValue', (_message.Message,), dict(
  DESCRIPTOR = _CAATAGVALUE,
  __module__ = 'caa_pb2'
  # @@protoc_insertion_point(class_scope:zsearch.CAATagValue)
  ))
_sym_db.RegisterMessage(CAATagValue)

CAARecord = _reflection.GeneratedProtocolMessageType('CAARecord', (_message.Message,), dict(
  DESCRIPTOR = _CAARECORD,
  __module__ = 'caa_pb2'
  # @@protoc_insertion_point(class_scope:zsearch.CAARecord)
  ))
_sym_db.RegisterMessage(CAARecord)

CAALookup = _reflection.GeneratedProtocolMessageType('CAALookup', (_message.Message,), dict(
  DESCRIPTOR = _CAALOOKUP,
  __module__ = 'caa_pb2'
  # @@protoc_insertion_point(class_scope:zsearch.CAALookup)
  ))
_sym_db.RegisterMessage(CAALookup)


# @@protoc_insertion_point(module_scope)
