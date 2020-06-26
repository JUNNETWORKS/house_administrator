# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: room.proto

from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='room.proto',
  package='room',
  syntax='proto3',
  serialized_options=b'Z\004.;pb',
  serialized_pb=b'\n\nroom.proto\x12\x04room*\x83\x01\n\x0bRoomNumbers\x12\x0c\n\x08JUN_ROOM\x10\x00\x12\x0f\n\x0bSUBARU_ROOM\x10\x01\x12\x0f\n\x0bNORIKO_ROOM\x10\x02\x12\x0f\n\x0bLIVING_ROOM\x10\x03\x12\x0c\n\x08\x45NTRANCE\x10\x04\x12\n\n\x06STAIRS\x10\x05\x12\r\n\tBATH_ROOM\x10\x06\x12\n\n\x06TOILET\x10\x07\x42\x06Z\x04.;pbb\x06proto3'
)

_ROOMNUMBERS = _descriptor.EnumDescriptor(
  name='RoomNumbers',
  full_name='room.RoomNumbers',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='JUN_ROOM', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBARU_ROOM', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='NORIKO_ROOM', index=2, number=2,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='LIVING_ROOM', index=3, number=3,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='ENTRANCE', index=4, number=4,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='STAIRS', index=5, number=5,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='BATH_ROOM', index=6, number=6,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='TOILET', index=7, number=7,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=21,
  serialized_end=152,
)
_sym_db.RegisterEnumDescriptor(_ROOMNUMBERS)

RoomNumbers = enum_type_wrapper.EnumTypeWrapper(_ROOMNUMBERS)
JUN_ROOM = 0
SUBARU_ROOM = 1
NORIKO_ROOM = 2
LIVING_ROOM = 3
ENTRANCE = 4
STAIRS = 5
BATH_ROOM = 6
TOILET = 7


DESCRIPTOR.enum_types_by_name['RoomNumbers'] = _ROOMNUMBERS
_sym_db.RegisterFileDescriptor(DESCRIPTOR)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)