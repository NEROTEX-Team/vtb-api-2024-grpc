// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.19.6
// source: api/v1/user/user.proto

package user

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Firstname string `protobuf:"bytes,1,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname  string `protobuf:"bytes,2,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Email     string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	mi := &file_api_v1_user_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_user_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_api_v1_user_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserInfo) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *UserInfo) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *UserInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Info      *UserInfo              `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_api_v1_user_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_user_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_api_v1_user_user_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetId() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *User) GetInfo() *UserInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *User) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *User) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type FetchUserByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FetchUserByIdRequest) Reset() {
	*x = FetchUserByIdRequest{}
	mi := &file_api_v1_user_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FetchUserByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchUserByIdRequest) ProtoMessage() {}

func (x *FetchUserByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_user_user_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchUserByIdRequest.ProtoReflect.Descriptor instead.
func (*FetchUserByIdRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_user_user_proto_rawDescGZIP(), []int{2}
}

func (x *FetchUserByIdRequest) GetId() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type FetchUserByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *FetchUserByIdResponse) Reset() {
	*x = FetchUserByIdResponse{}
	mi := &file_api_v1_user_user_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FetchUserByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchUserByIdResponse) ProtoMessage() {}

func (x *FetchUserByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_user_user_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchUserByIdResponse.ProtoReflect.Descriptor instead.
func (*FetchUserByIdResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_user_user_proto_rawDescGZIP(), []int{3}
}

func (x *FetchUserByIdResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *UserInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	mi := &file_api_v1_user_user_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_user_user_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_user_user_proto_rawDescGZIP(), []int{4}
}

func (x *CreateUserRequest) GetInfo() *UserInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type CreateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *CreateUserResponse) Reset() {
	*x = CreateUserResponse{}
	mi := &file_api_v1_user_user_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResponse) ProtoMessage() {}

func (x *CreateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_user_user_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResponse.ProtoReflect.Descriptor instead.
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_user_user_proto_rawDescGZIP(), []int{5}
}

func (x *CreateUserResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type FetchAllUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *FetchAllUsersRequest) Reset() {
	*x = FetchAllUsersRequest{}
	mi := &file_api_v1_user_user_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FetchAllUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchAllUsersRequest) ProtoMessage() {}

func (x *FetchAllUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_user_user_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchAllUsersRequest.ProtoReflect.Descriptor instead.
func (*FetchAllUsersRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_user_user_proto_rawDescGZIP(), []int{6}
}

func (x *FetchAllUsersRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *FetchAllUsersRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type FetchListUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int32   `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Users []*User `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *FetchListUsersResponse) Reset() {
	*x = FetchListUsersResponse{}
	mi := &file_api_v1_user_user_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FetchListUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchListUsersResponse) ProtoMessage() {}

func (x *FetchListUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_user_user_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchListUsersResponse.ProtoReflect.Descriptor instead.
func (*FetchListUsersResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_user_user_proto_rawDescGZIP(), []int{7}
}

func (x *FetchListUsersResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *FetchListUsersResponse) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

type UpdateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UpdateUserResponse) Reset() {
	*x = UpdateUserResponse{}
	mi := &file_api_v1_user_user_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserResponse) ProtoMessage() {}

func (x *UpdateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_user_user_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserResponse.ProtoReflect.Descriptor instead.
func (*UpdateUserResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_user_user_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateUserResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type FetchUserByEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *FetchUserByEmailRequest) Reset() {
	*x = FetchUserByEmailRequest{}
	mi := &file_api_v1_user_user_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FetchUserByEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchUserByEmailRequest) ProtoMessage() {}

func (x *FetchUserByEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_user_user_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchUserByEmailRequest.ProtoReflect.Descriptor instead.
func (*FetchUserByEmailRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_user_user_proto_rawDescGZIP(), []int{9}
}

func (x *FetchUserByEmailRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type FetchUserByEmailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *FetchUserByEmailResponse) Reset() {
	*x = FetchUserByEmailResponse{}
	mi := &file_api_v1_user_user_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FetchUserByEmailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchUserByEmailResponse) ProtoMessage() {}

func (x *FetchUserByEmailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_user_user_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchUserByEmailResponse.ProtoReflect.Descriptor instead.
func (*FetchUserByEmailResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_user_user_proto_rawDescGZIP(), []int{10}
}

func (x *FetchUserByEmailResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type DeleteUserByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteUserByIdRequest) Reset() {
	*x = DeleteUserByIdRequest{}
	mi := &file_api_v1_user_user_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteUserByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserByIdRequest) ProtoMessage() {}

func (x *DeleteUserByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_user_user_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserByIdRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserByIdRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_user_user_proto_rawDescGZIP(), []int{11}
}

func (x *DeleteUserByIdRequest) GetId() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type UpdateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Info *UserInfo `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *UpdateUserRequest) Reset() {
	*x = UpdateUserRequest{}
	mi := &file_api_v1_user_user_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserRequest) ProtoMessage() {}

func (x *UpdateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_user_user_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_user_user_proto_rawDescGZIP(), []int{12}
}

func (x *UpdateUserRequest) GetId() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *UpdateUserRequest) GetInfo() *UserInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_api_v1_user_user_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_user_user_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_api_v1_user_user_proto_rawDescGZIP(), []int{13}
}

var File_api_v1_user_user_proto protoreflect.FileDescriptor

var file_api_v1_user_user_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x5a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73,
	0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73,
	0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0xb0, 0x01, 0x0a, 0x04,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x26,
	0x0a, 0x14, 0x46, 0x65, 0x74, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x37, 0x0a, 0x15, 0x46, 0x65, 0x74, 0x63, 0x68, 0x55,
	0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1e, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22,
	0x37, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x34, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e,
	0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x44,
	0x0a, 0x14, 0x46, 0x65, 0x74, 0x63, 0x68, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x22, 0x50, 0x0a, 0x16, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4c, 0x69, 0x73,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x12, 0x20, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x34, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x2f, 0x0a, 0x17,
	0x46, 0x65, 0x74, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x3a, 0x0a,
	0x18, 0x46, 0x65, 0x74, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x27, 0x0a, 0x15, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x47, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x32, 0xab, 0x03, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x56, 0x31, 0x12,
	0x48, 0x0a, 0x0d, 0x46, 0x65, 0x74, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64,
	0x12, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x55, 0x73, 0x65,
	0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x0e, 0x46, 0x65, 0x74,
	0x63, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x1a, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x46,
	0x65, 0x74, 0x63, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3f,
	0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x51, 0x0a, 0x10, 0x46, 0x65, 0x74, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x1d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x55,
	0x73, 0x65, 0x72, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x15, 0x5a, 0x13, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x3b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_v1_user_user_proto_rawDescOnce sync.Once
	file_api_v1_user_user_proto_rawDescData = file_api_v1_user_user_proto_rawDesc
)

func file_api_v1_user_user_proto_rawDescGZIP() []byte {
	file_api_v1_user_user_proto_rawDescOnce.Do(func() {
		file_api_v1_user_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_user_user_proto_rawDescData)
	})
	return file_api_v1_user_user_proto_rawDescData
}

var file_api_v1_user_user_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_api_v1_user_user_proto_goTypes = []any{
	(*UserInfo)(nil),                 // 0: user.UserInfo
	(*User)(nil),                     // 1: user.User
	(*FetchUserByIdRequest)(nil),     // 2: user.FetchUserByIdRequest
	(*FetchUserByIdResponse)(nil),    // 3: user.FetchUserByIdResponse
	(*CreateUserRequest)(nil),        // 4: user.CreateUserRequest
	(*CreateUserResponse)(nil),       // 5: user.CreateUserResponse
	(*FetchAllUsersRequest)(nil),     // 6: user.FetchAllUsersRequest
	(*FetchListUsersResponse)(nil),   // 7: user.FetchListUsersResponse
	(*UpdateUserResponse)(nil),       // 8: user.UpdateUserResponse
	(*FetchUserByEmailRequest)(nil),  // 9: user.FetchUserByEmailRequest
	(*FetchUserByEmailResponse)(nil), // 10: user.FetchUserByEmailResponse
	(*DeleteUserByIdRequest)(nil),    // 11: user.DeleteUserByIdRequest
	(*UpdateUserRequest)(nil),        // 12: user.UpdateUserRequest
	(*Empty)(nil),                    // 13: user.Empty
	(*timestamppb.Timestamp)(nil),    // 14: google.protobuf.Timestamp
}
var file_api_v1_user_user_proto_depIdxs = []int32{
	0,  // 0: user.User.info:type_name -> user.UserInfo
	14, // 1: user.User.created_at:type_name -> google.protobuf.Timestamp
	14, // 2: user.User.updated_at:type_name -> google.protobuf.Timestamp
	1,  // 3: user.FetchUserByIdResponse.user:type_name -> user.User
	0,  // 4: user.CreateUserRequest.info:type_name -> user.UserInfo
	1,  // 5: user.CreateUserResponse.user:type_name -> user.User
	1,  // 6: user.FetchListUsersResponse.users:type_name -> user.User
	1,  // 7: user.UpdateUserResponse.user:type_name -> user.User
	1,  // 8: user.FetchUserByEmailResponse.user:type_name -> user.User
	0,  // 9: user.UpdateUserRequest.info:type_name -> user.UserInfo
	2,  // 10: user.UserV1.FetchUserById:input_type -> user.FetchUserByIdRequest
	6,  // 11: user.UserV1.FetchListUsers:input_type -> user.FetchAllUsersRequest
	12, // 12: user.UserV1.UpdateUser:input_type -> user.UpdateUserRequest
	11, // 13: user.UserV1.DeleteUser:input_type -> user.DeleteUserByIdRequest
	4,  // 14: user.UserV1.CreateUser:input_type -> user.CreateUserRequest
	9,  // 15: user.UserV1.FetchUserByEmail:input_type -> user.FetchUserByEmailRequest
	3,  // 16: user.UserV1.FetchUserById:output_type -> user.FetchUserByIdResponse
	7,  // 17: user.UserV1.FetchListUsers:output_type -> user.FetchListUsersResponse
	8,  // 18: user.UserV1.UpdateUser:output_type -> user.UpdateUserResponse
	13, // 19: user.UserV1.DeleteUser:output_type -> user.Empty
	5,  // 20: user.UserV1.CreateUser:output_type -> user.CreateUserResponse
	10, // 21: user.UserV1.FetchUserByEmail:output_type -> user.FetchUserByEmailResponse
	16, // [16:22] is the sub-list for method output_type
	10, // [10:16] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_api_v1_user_user_proto_init() }
func file_api_v1_user_user_proto_init() {
	if File_api_v1_user_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_user_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_user_user_proto_goTypes,
		DependencyIndexes: file_api_v1_user_user_proto_depIdxs,
		MessageInfos:      file_api_v1_user_user_proto_msgTypes,
	}.Build()
	File_api_v1_user_user_proto = out.File
	file_api_v1_user_user_proto_rawDesc = nil
	file_api_v1_user_user_proto_goTypes = nil
	file_api_v1_user_user_proto_depIdxs = nil
}
