# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: style.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='style.proto',
  package='services',
  syntax='proto3',
  serialized_options=b'Z\n.;services',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x0bstyle.proto\x12\x08services\"\'\n\x08StyleReq\x12\r\n\x05\x66unds\x18\x01 \x03(\t\x12\x0c\n\x04many\x18\x02 \x01(\x08\"w\n\tStyleResp\x12+\n\x04\x64\x61ta\x18\x01 \x03(\x0b\x32\x1d.services.StyleResp.DataEntry\x1a=\n\tDataEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\x1f\n\x05value\x18\x02 \x01(\x0b\x32\x10.services.Styles:\x02\x38\x01\"h\n\x06Styles\x12&\n\x06styles\x18\x01 \x03(\x0b\x32\x16.services.Styles.Style\x1a\x36\n\x05Style\x12\x10\n\x08secucode\x18\x01 \x01(\t\x12\x0c\n\x04\x64\x61te\x18\x02 \x01(\t\x12\r\n\x05style\x18\x03 \x01(\tB\x0cZ\n.;servicesb\x06proto3'
)




_STYLEREQ = _descriptor.Descriptor(
  name='StyleReq',
  full_name='services.StyleReq',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='funds', full_name='services.StyleReq.funds', index=0,
      number=1, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='many', full_name='services.StyleReq.many', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=25,
  serialized_end=64,
)


_STYLERESP_DATAENTRY = _descriptor.Descriptor(
  name='DataEntry',
  full_name='services.StyleResp.DataEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='services.StyleResp.DataEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='value', full_name='services.StyleResp.DataEntry.value', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=b'8\001',
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=124,
  serialized_end=185,
)

_STYLERESP = _descriptor.Descriptor(
  name='StyleResp',
  full_name='services.StyleResp',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='data', full_name='services.StyleResp.data', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[_STYLERESP_DATAENTRY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=66,
  serialized_end=185,
)


_STYLES_STYLE = _descriptor.Descriptor(
  name='Style',
  full_name='services.Styles.Style',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='secucode', full_name='services.Styles.Style.secucode', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='date', full_name='services.Styles.Style.date', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='style', full_name='services.Styles.Style.style', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=237,
  serialized_end=291,
)

_STYLES = _descriptor.Descriptor(
  name='Styles',
  full_name='services.Styles',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='styles', full_name='services.Styles.styles', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[_STYLES_STYLE, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=187,
  serialized_end=291,
)

_STYLERESP_DATAENTRY.fields_by_name['value'].message_type = _STYLES
_STYLERESP_DATAENTRY.containing_type = _STYLERESP
_STYLERESP.fields_by_name['data'].message_type = _STYLERESP_DATAENTRY
_STYLES_STYLE.containing_type = _STYLES
_STYLES.fields_by_name['styles'].message_type = _STYLES_STYLE
DESCRIPTOR.message_types_by_name['StyleReq'] = _STYLEREQ
DESCRIPTOR.message_types_by_name['StyleResp'] = _STYLERESP
DESCRIPTOR.message_types_by_name['Styles'] = _STYLES
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

StyleReq = _reflection.GeneratedProtocolMessageType('StyleReq', (_message.Message,), {
  'DESCRIPTOR' : _STYLEREQ,
  '__module__' : 'style_pb2'
  # @@protoc_insertion_point(class_scope:services.StyleReq)
  })
_sym_db.RegisterMessage(StyleReq)

StyleResp = _reflection.GeneratedProtocolMessageType('StyleResp', (_message.Message,), {

  'DataEntry' : _reflection.GeneratedProtocolMessageType('DataEntry', (_message.Message,), {
    'DESCRIPTOR' : _STYLERESP_DATAENTRY,
    '__module__' : 'style_pb2'
    # @@protoc_insertion_point(class_scope:services.StyleResp.DataEntry)
    })
  ,
  'DESCRIPTOR' : _STYLERESP,
  '__module__' : 'style_pb2'
  # @@protoc_insertion_point(class_scope:services.StyleResp)
  })
_sym_db.RegisterMessage(StyleResp)
_sym_db.RegisterMessage(StyleResp.DataEntry)

Styles = _reflection.GeneratedProtocolMessageType('Styles', (_message.Message,), {

  'Style' : _reflection.GeneratedProtocolMessageType('Style', (_message.Message,), {
    'DESCRIPTOR' : _STYLES_STYLE,
    '__module__' : 'style_pb2'
    # @@protoc_insertion_point(class_scope:services.Styles.Style)
    })
  ,
  'DESCRIPTOR' : _STYLES,
  '__module__' : 'style_pb2'
  # @@protoc_insertion_point(class_scope:services.Styles)
  })
_sym_db.RegisterMessage(Styles)
_sym_db.RegisterMessage(Styles.Style)


DESCRIPTOR._options = None
_STYLERESP_DATAENTRY._options = None
# @@protoc_insertion_point(module_scope)
