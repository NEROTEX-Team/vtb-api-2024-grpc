# type: ignore
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: user.proto
# Protobuf Python Version: 5.28.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    28,
    1,
    '',
    'user.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\nuser.proto\x12\x04user\x1a\x1fgoogle/protobuf/timestamp.proto\"\xa8\x01\n\x04User\x12\n\n\x02id\x18\x01 \x01(\t\x12\r\n\x05\x65mail\x18\x02 \x01(\t\x12\x12\n\nfirst_name\x18\x03 \x01(\t\x12\x11\n\tlast_name\x18\x04 \x01(\t\x12.\n\ncreated_at\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12.\n\nupdated_at\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"\"\n\x14\x46\x65tchUserByIdRequest\x12\n\n\x02id\x18\x01 \x01(\t\"1\n\x15\x46\x65tchUserByIdResponse\x12\x18\n\x04user\x18\x01 \x01(\x0b\x32\n.user.User\"[\n\x11\x43reateUserRequest\x12\r\n\x05\x65mail\x18\x01 \x01(\t\x12\x10\n\x08password\x18\x02 \x01(\t\x12\x12\n\nfirst_name\x18\x03 \x01(\t\x12\x11\n\tlast_name\x18\x04 \x01(\t\".\n\x12\x43reateUserResponse\x12\x18\n\x04user\x18\x01 \x01(\x0b\x32\n.user.User\"5\n\x14\x46\x65tchUserListRequest\x12\r\n\x05limit\x18\x01 \x01(\x05\x12\x0e\n\x06offset\x18\x02 \x01(\x05\"A\n\x15\x46\x65tchUserListResponse\x12\r\n\x05total\x18\x01 \x01(\x05\x12\x19\n\x05users\x18\x02 \x03(\x0b\x32\n.user.User\".\n\x12UpdateUserResponse\x12\x18\n\x04user\x18\x01 \x01(\x0b\x32\n.user.User\"(\n\x17\x46\x65tchUserByEmailRequest\x12\r\n\x05\x65mail\x18\x01 \x01(\t\"4\n\x18\x46\x65tchUserByEmailResponse\x12\x18\n\x04user\x18\x01 \x01(\x0b\x32\n.user.User\"#\n\x15\x44\x65leteUserByIdRequest\x12\n\n\x02id\x18\x01 \x01(\t\"U\n\x11UpdateUserRequest\x12\n\n\x02id\x18\x01 \x01(\t\x12\r\n\x05\x65mail\x18\x02 \x01(\t\x12\x12\n\nfirst_name\x18\x03 \x01(\t\x12\x11\n\tlast_name\x18\x04 \x01(\t\"\x07\n\x05\x45mpty2\xa9\x03\n\x06UserV1\x12H\n\rFetchUserById\x12\x1a.user.FetchUserByIdRequest\x1a\x1b.user.FetchUserByIdResponse\x12H\n\rFetchUserList\x12\x1a.user.FetchUserListRequest\x1a\x1b.user.FetchUserListResponse\x12?\n\nUpdateUser\x12\x17.user.UpdateUserRequest\x1a\x18.user.UpdateUserResponse\x12\x36\n\nDeleteUser\x12\x1b.user.DeleteUserByIdRequest\x1a\x0b.user.Empty\x12?\n\nCreateUser\x12\x17.user.CreateUserRequest\x1a\x18.user.CreateUserResponse\x12Q\n\x10\x46\x65tchUserByEmail\x12\x1d.user.FetchUserByEmailRequest\x1a\x1e.user.FetchUserByEmailResponseB\x15Z\x13./pkg/v1/user;user;b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'user_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\023./pkg/v1/user;user;'
  _globals['_USER']._serialized_start=54
  _globals['_USER']._serialized_end=222
  _globals['_FETCHUSERBYIDREQUEST']._serialized_start=224
  _globals['_FETCHUSERBYIDREQUEST']._serialized_end=258
  _globals['_FETCHUSERBYIDRESPONSE']._serialized_start=260
  _globals['_FETCHUSERBYIDRESPONSE']._serialized_end=309
  _globals['_CREATEUSERREQUEST']._serialized_start=311
  _globals['_CREATEUSERREQUEST']._serialized_end=402
  _globals['_CREATEUSERRESPONSE']._serialized_start=404
  _globals['_CREATEUSERRESPONSE']._serialized_end=450
  _globals['_FETCHUSERLISTREQUEST']._serialized_start=452
  _globals['_FETCHUSERLISTREQUEST']._serialized_end=505
  _globals['_FETCHUSERLISTRESPONSE']._serialized_start=507
  _globals['_FETCHUSERLISTRESPONSE']._serialized_end=572
  _globals['_UPDATEUSERRESPONSE']._serialized_start=574
  _globals['_UPDATEUSERRESPONSE']._serialized_end=620
  _globals['_FETCHUSERBYEMAILREQUEST']._serialized_start=622
  _globals['_FETCHUSERBYEMAILREQUEST']._serialized_end=662
  _globals['_FETCHUSERBYEMAILRESPONSE']._serialized_start=664
  _globals['_FETCHUSERBYEMAILRESPONSE']._serialized_end=716
  _globals['_DELETEUSERBYIDREQUEST']._serialized_start=718
  _globals['_DELETEUSERBYIDREQUEST']._serialized_end=753
  _globals['_UPDATEUSERREQUEST']._serialized_start=755
  _globals['_UPDATEUSERREQUEST']._serialized_end=840
  _globals['_EMPTY']._serialized_start=842
  _globals['_EMPTY']._serialized_end=849
  _globals['_USERV1']._serialized_start=852
  _globals['_USERV1']._serialized_end=1277
# @@protoc_insertion_point(module_scope)
