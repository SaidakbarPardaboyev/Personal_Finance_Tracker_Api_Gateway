// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: users_service.proto

package users

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int32  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit    int64  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	FullName string `protobuf:"bytes,3,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Email    string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	UserRole string `protobuf:"bytes,5,opt,name=user_role,json=userRole,proto3" json:"user_role,omitempty"`
}

func (x *GetListRequest) Reset() {
	*x = GetListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListRequest) ProtoMessage() {}

func (x *GetListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_users_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListRequest.ProtoReflect.Descriptor instead.
func (*GetListRequest) Descriptor() ([]byte, []int) {
	return file_users_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetListRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetListRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetListRequest) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *GetListRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetListRequest) GetUserRole() string {
	if x != nil {
		return x.UserRole
	}
	return ""
}

type Users struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	Page  int32   `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Limit int64   `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Count int32   `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *Users) Reset() {
	*x = Users{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Users) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Users) ProtoMessage() {}

func (x *Users) ProtoReflect() protoreflect.Message {
	mi := &file_users_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Users.ProtoReflect.Descriptor instead.
func (*Users) Descriptor() ([]byte, []int) {
	return file_users_service_proto_rawDescGZIP(), []int{1}
}

func (x *Users) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *Users) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *Users) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *Users) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type UpdateUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email        string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	FullName     string `protobuf:"bytes,3,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	PasswordHash string `protobuf:"bytes,4,opt,name=password_hash,json=passwordHash,proto3" json:"password_hash,omitempty"`
}

func (x *UpdateUser) Reset() {
	*x = UpdateUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUser) ProtoMessage() {}

func (x *UpdateUser) ProtoReflect() protoreflect.Message {
	mi := &file_users_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUser.ProtoReflect.Descriptor instead.
func (*UpdateUser) Descriptor() ([]byte, []int) {
	return file_users_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateUser) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateUser) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdateUser) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *UpdateUser) GetPasswordHash() string {
	if x != nil {
		return x.PasswordHash
	}
	return ""
}

type UpdatedUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FullName  string `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Email     string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password  string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	UserRole  string `protobuf:"bytes,5,opt,name=user_role,json=userRole,proto3" json:"user_role,omitempty"`
	UpdatedAt string `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *UpdatedUser) Reset() {
	*x = UpdatedUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatedUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatedUser) ProtoMessage() {}

func (x *UpdatedUser) ProtoReflect() protoreflect.Message {
	mi := &file_users_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatedUser.ProtoReflect.Descriptor instead.
func (*UpdatedUser) Descriptor() ([]byte, []int) {
	return file_users_service_proto_rawDescGZIP(), []int{3}
}

func (x *UpdatedUser) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdatedUser) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *UpdatedUser) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdatedUser) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UpdatedUser) GetUserRole() string {
	if x != nil {
		return x.UserRole
	}
	return ""
}

func (x *UpdatedUser) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type ChangePassword struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrentPassword string `protobuf:"bytes,1,opt,name=current_password,json=currentPassword,proto3" json:"current_password,omitempty"`
	NewPassword     string `protobuf:"bytes,2,opt,name=new_password,json=newPassword,proto3" json:"new_password,omitempty"`
	UserId          string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *ChangePassword) Reset() {
	*x = ChangePassword{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangePassword) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangePassword) ProtoMessage() {}

func (x *ChangePassword) ProtoReflect() protoreflect.Message {
	mi := &file_users_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangePassword.ProtoReflect.Descriptor instead.
func (*ChangePassword) Descriptor() ([]byte, []int) {
	return file_users_service_proto_rawDescGZIP(), []int{4}
}

func (x *ChangePassword) GetCurrentPassword() string {
	if x != nil {
		return x.CurrentPassword
	}
	return ""
}

func (x *ChangePassword) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

func (x *ChangePassword) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

var File_users_service_proto protoreflect.FileDescriptor

var file_users_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x1a, 0x0b, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x01, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x22, 0x6a, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12,
	0x21, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x74, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f,
	0x68, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x48, 0x61, 0x73, 0x68, 0x22, 0xa8, 0x01, 0x0a, 0x0b, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c,
	0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x72, 0x6f, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x52, 0x6f, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0x77, 0x0a, 0x0e, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x32, 0xab, 0x02, 0x0a,
	0x0c, 0x55, 0x73, 0x65, 0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x29, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x2e, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x0b, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x12, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x2f, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x1a, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x56, 0x6f,
	0x69, 0x64, 0x12, 0x34, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x1a, 0x0b, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x11, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x0b, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x42, 0x10, 0x5a, 0x0e, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_users_service_proto_rawDescOnce sync.Once
	file_users_service_proto_rawDescData = file_users_service_proto_rawDesc
)

func file_users_service_proto_rawDescGZIP() []byte {
	file_users_service_proto_rawDescOnce.Do(func() {
		file_users_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_users_service_proto_rawDescData)
	})
	return file_users_service_proto_rawDescData
}

var file_users_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_users_service_proto_goTypes = []interface{}{
	(*GetListRequest)(nil), // 0: users.GetListRequest
	(*Users)(nil),          // 1: users.users
	(*UpdateUser)(nil),     // 2: users.updateUser
	(*UpdatedUser)(nil),    // 3: users.UpdatedUser
	(*ChangePassword)(nil), // 4: users.changePassword
	(*User)(nil),           // 5: users.user
	(*PrimaryKey)(nil),     // 6: users.PrimaryKey
	(*Void)(nil),           // 7: users.Void
}
var file_users_service_proto_depIdxs = []int32{
	5, // 0: users.users.users:type_name -> users.user
	6, // 1: users.UsersService.GetById:input_type -> users.PrimaryKey
	0, // 2: users.UsersService.GetAll:input_type -> users.GetListRequest
	2, // 3: users.UsersService.Update:input_type -> users.updateUser
	6, // 4: users.UsersService.Delete:input_type -> users.PrimaryKey
	4, // 5: users.UsersService.ChangePassword:input_type -> users.changePassword
	6, // 6: users.UsersService.ChangeUserRole:input_type -> users.PrimaryKey
	5, // 7: users.UsersService.GetById:output_type -> users.user
	1, // 8: users.UsersService.GetAll:output_type -> users.users
	3, // 9: users.UsersService.Update:output_type -> users.UpdatedUser
	7, // 10: users.UsersService.Delete:output_type -> users.Void
	7, // 11: users.UsersService.ChangePassword:output_type -> users.Void
	7, // 12: users.UsersService.ChangeUserRole:output_type -> users.Void
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_users_service_proto_init() }
func file_users_service_proto_init() {
	if File_users_service_proto != nil {
		return
	}
	file_users_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_users_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_users_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Users); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_users_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUser); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_users_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatedUser); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_users_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangePassword); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_users_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_users_service_proto_goTypes,
		DependencyIndexes: file_users_service_proto_depIdxs,
		MessageInfos:      file_users_service_proto_msgTypes,
	}.Build()
	File_users_service_proto = out.File
	file_users_service_proto_rawDesc = nil
	file_users_service_proto_goTypes = nil
	file_users_service_proto_depIdxs = nil
}
