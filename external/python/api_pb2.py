# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: api.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


import classify_pb2 as classify__pb2
import basic_pb2 as basic__pb2
import style_pb2 as style__pb2
import index_pb2 as index__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='api.proto',
  package='services',
  syntax='proto3',
  serialized_options=b'Z\n.;services',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\tapi.proto\x12\x08services\x1a\x0e\x63lassify.proto\x1a\x0b\x62\x61sic.proto\x1a\x0bstyle.proto\x1a\x0bindex.proto2\xe6\x02\n\x0fScreenRpcServer\x12=\n\x0c\x46undCategory\x12\x15.services.ClassifyReq\x1a\x16.services.ClassifyResp\x12W\n\x14\x46undBasicInfoHandler\x12\x1e.services.FundBasicInfoRequest\x1a\x1f.services.FundBasicInfoResponse\x12:\n\x0f\x46undScaleNature\x12\x12.services.StyleReq\x1a\x13.services.StyleResp\x12:\n\x0f\x46undStyleNature\x12\x12.services.StyleReq\x1a\x13.services.StyleResp\x12\x43\n\x10\x46undRelatedIndex\x12\x16.services.IndexCorrReq\x1a\x17.services.IndexCorrRespB\x0cZ\n.;servicesb\x06proto3'
  ,
  dependencies=[classify__pb2.DESCRIPTOR,basic__pb2.DESCRIPTOR,style__pb2.DESCRIPTOR,index__pb2.DESCRIPTOR,])



_sym_db.RegisterFileDescriptor(DESCRIPTOR)


DESCRIPTOR._options = None

_SCREENRPCSERVER = _descriptor.ServiceDescriptor(
  name='ScreenRpcServer',
  full_name='services.ScreenRpcServer',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=79,
  serialized_end=437,
  methods=[
  _descriptor.MethodDescriptor(
    name='FundCategory',
    full_name='services.ScreenRpcServer.FundCategory',
    index=0,
    containing_service=None,
    input_type=classify__pb2._CLASSIFYREQ,
    output_type=classify__pb2._CLASSIFYRESP,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='FundBasicInfoHandler',
    full_name='services.ScreenRpcServer.FundBasicInfoHandler',
    index=1,
    containing_service=None,
    input_type=basic__pb2._FUNDBASICINFOREQUEST,
    output_type=basic__pb2._FUNDBASICINFORESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='FundScaleNature',
    full_name='services.ScreenRpcServer.FundScaleNature',
    index=2,
    containing_service=None,
    input_type=style__pb2._STYLEREQ,
    output_type=style__pb2._STYLERESP,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='FundStyleNature',
    full_name='services.ScreenRpcServer.FundStyleNature',
    index=3,
    containing_service=None,
    input_type=style__pb2._STYLEREQ,
    output_type=style__pb2._STYLERESP,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='FundRelatedIndex',
    full_name='services.ScreenRpcServer.FundRelatedIndex',
    index=4,
    containing_service=None,
    input_type=index__pb2._INDEXCORRREQ,
    output_type=index__pb2._INDEXCORRRESP,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_SCREENRPCSERVER)

DESCRIPTOR.services_by_name['ScreenRpcServer'] = _SCREENRPCSERVER

# @@protoc_insertion_point(module_scope)