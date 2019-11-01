// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: powerssl/apiserver/v1/user.proto

package api

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/googleapis/google/api"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// A User ...
type User struct {
	// The resource name of the user.
	// User names have the form `users/{user_id}`.
	// The name is ignored when creating a user.
	Name                 string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CreateTime           *types.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime           *types.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	DisplayName          string           `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	UserName             string           `protobuf:"bytes,5,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_869f7cd01c33efef, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetCreateTime() *types.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *User) GetUpdateTime() *types.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

func (m *User) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *User) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

// Request message for UserService.Create.
type CreateUserRequest struct {
	// The user to create.
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_869f7cd01c33efef, []int{1}
}
func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

// Request message for UserService.Delete.
type DeleteUserRequest struct {
	// The name of the user to delete.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUserRequest) Reset()         { *m = DeleteUserRequest{} }
func (m *DeleteUserRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteUserRequest) ProtoMessage()    {}
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_869f7cd01c33efef, []int{2}
}
func (m *DeleteUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUserRequest.Unmarshal(m, b)
}
func (m *DeleteUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUserRequest.Marshal(b, m, deterministic)
}
func (m *DeleteUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUserRequest.Merge(m, src)
}
func (m *DeleteUserRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteUserRequest.Size(m)
}
func (m *DeleteUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUserRequest proto.InternalMessageInfo

func (m *DeleteUserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Request message for UserService.Get.
type GetUserRequest struct {
	// The name of the user to retrieve.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_869f7cd01c33efef, []int{3}
}
func (m *GetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRequest.Unmarshal(m, b)
}
func (m *GetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRequest.Marshal(b, m, deterministic)
}
func (m *GetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRequest.Merge(m, src)
}
func (m *GetUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRequest.Size(m)
}
func (m *GetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRequest proto.InternalMessageInfo

func (m *GetUserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Request message for UserService.List.
type ListUsersRequest struct {
	// Requested page size. Server may return fewer users than requested.
	// If unspecified, server will pick an appropriate default.
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// A token identifying a page of results the server should return.
	// Typically, this is the value of
	// [ListUsersResponse.next_page_token][powerssl.apiserver.v1.ListUsersResponse.next_page_token].
	// returned from the previous call to `List` method.
	PageToken            string   `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUsersRequest) Reset()         { *m = ListUsersRequest{} }
func (m *ListUsersRequest) String() string { return proto.CompactTextString(m) }
func (*ListUsersRequest) ProtoMessage()    {}
func (*ListUsersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_869f7cd01c33efef, []int{4}
}
func (m *ListUsersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUsersRequest.Unmarshal(m, b)
}
func (m *ListUsersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUsersRequest.Marshal(b, m, deterministic)
}
func (m *ListUsersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUsersRequest.Merge(m, src)
}
func (m *ListUsersRequest) XXX_Size() int {
	return xxx_messageInfo_ListUsersRequest.Size(m)
}
func (m *ListUsersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUsersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListUsersRequest proto.InternalMessageInfo

func (m *ListUsersRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListUsersRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// Response message for UserService.List.
type ListUsersResponse struct {
	// The list of users.
	Users []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	// A token to retrieve next page of results.
	// Pass this value in the
	// [ListUsersRequest.page_token][powerssl.apiserver.v1.ListUsersRequest.page_token]
	// field in the subsequent call to `List` method to retrieve the next
	// page of results.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUsersResponse) Reset()         { *m = ListUsersResponse{} }
func (m *ListUsersResponse) String() string { return proto.CompactTextString(m) }
func (*ListUsersResponse) ProtoMessage()    {}
func (*ListUsersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_869f7cd01c33efef, []int{5}
}
func (m *ListUsersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUsersResponse.Unmarshal(m, b)
}
func (m *ListUsersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUsersResponse.Marshal(b, m, deterministic)
}
func (m *ListUsersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUsersResponse.Merge(m, src)
}
func (m *ListUsersResponse) XXX_Size() int {
	return xxx_messageInfo_ListUsersResponse.Size(m)
}
func (m *ListUsersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUsersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListUsersResponse proto.InternalMessageInfo

func (m *ListUsersResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *ListUsersResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// Request message for UserService.Update.
type UpdateUserRequest struct {
	// The name of the user to update.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The field mask with paths to update.
	UpdateMask *types.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The user to update with. The name must match or be empty.
	User                 *User    `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserRequest) Reset()         { *m = UpdateUserRequest{} }
func (m *UpdateUserRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateUserRequest) ProtoMessage()    {}
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_869f7cd01c33efef, []int{6}
}
func (m *UpdateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserRequest.Unmarshal(m, b)
}
func (m *UpdateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserRequest.Marshal(b, m, deterministic)
}
func (m *UpdateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserRequest.Merge(m, src)
}
func (m *UpdateUserRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateUserRequest.Size(m)
}
func (m *UpdateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserRequest proto.InternalMessageInfo

func (m *UpdateUserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateUserRequest) GetUpdateMask() *types.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

func (m *UpdateUserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "powerssl.apiserver.v1.User")
	proto.RegisterType((*CreateUserRequest)(nil), "powerssl.apiserver.v1.CreateUserRequest")
	proto.RegisterType((*DeleteUserRequest)(nil), "powerssl.apiserver.v1.DeleteUserRequest")
	proto.RegisterType((*GetUserRequest)(nil), "powerssl.apiserver.v1.GetUserRequest")
	proto.RegisterType((*ListUsersRequest)(nil), "powerssl.apiserver.v1.ListUsersRequest")
	proto.RegisterType((*ListUsersResponse)(nil), "powerssl.apiserver.v1.ListUsersResponse")
	proto.RegisterType((*UpdateUserRequest)(nil), "powerssl.apiserver.v1.UpdateUserRequest")
}

func init() { proto.RegisterFile("powerssl/apiserver/v1/user.proto", fileDescriptor_869f7cd01c33efef) }

var fileDescriptor_869f7cd01c33efef = []byte{
	// 600 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x96, 0x1b, 0x37, 0x6a, 0xc6, 0x50, 0xc8, 0xaa, 0x40, 0xe4, 0x50, 0x11, 0x2c, 0xa0, 0x51,
	0x0e, 0xb6, 0x12, 0x0e, 0x88, 0x56, 0x5c, 0xa0, 0xd0, 0x0b, 0x54, 0x28, 0x6d, 0x2f, 0x5c, 0xa2,
	0x6d, 0x33, 0x89, 0xac, 0x38, 0xb6, 0xf1, 0x6e, 0x0c, 0x2d, 0xe2, 0xc2, 0x2b, 0x20, 0x1e, 0x80,
	0x67, 0x82, 0x47, 0xe0, 0x41, 0xd0, 0xec, 0xc6, 0x26, 0x4d, 0xe2, 0x84, 0x9b, 0x3d, 0xf3, 0x7d,
	0xf3, 0xcd, 0xcf, 0x67, 0x43, 0x23, 0x8e, 0x3e, 0x61, 0x22, 0x44, 0xe0, 0xf1, 0xd8, 0x17, 0x98,
	0xa4, 0x98, 0x78, 0x69, 0xdb, 0x9b, 0x08, 0x4c, 0xdc, 0x38, 0x89, 0x64, 0xc4, 0xee, 0x64, 0x08,
	0x37, 0x47, 0xb8, 0x69, 0xdb, 0xbe, 0x3f, 0x8c, 0xa2, 0x61, 0x80, 0x44, 0xf3, 0x78, 0x18, 0x46,
	0x92, 0x4b, 0x3f, 0x0a, 0x85, 0x26, 0xd9, 0xf5, 0x69, 0x56, 0xbd, 0x9d, 0x4f, 0x06, 0x1e, 0x8e,
	0x63, 0x79, 0x39, 0x4d, 0x36, 0xe6, 0x93, 0x03, 0x1f, 0x83, 0x7e, 0x6f, 0xcc, 0xc5, 0x68, 0x8a,
	0x78, 0x30, 0x8f, 0x90, 0xfe, 0x18, 0x85, 0xe4, 0xe3, 0x58, 0x03, 0x9c, 0xdf, 0x06, 0x98, 0x67,
	0x02, 0x13, 0xc6, 0xc0, 0x0c, 0xf9, 0x18, 0x6b, 0x46, 0xc3, 0x68, 0x56, 0xba, 0xea, 0x99, 0x1d,
	0x80, 0x75, 0x91, 0x20, 0x97, 0xd8, 0x23, 0x5a, 0x6d, 0xa3, 0x61, 0x34, 0xad, 0x8e, 0xed, 0xea,
	0x9a, 0x6e, 0x56, 0xd3, 0x3d, 0xcd, 0x6a, 0x76, 0x41, 0xc3, 0x29, 0x40, 0xe4, 0x49, 0xdc, 0xcf,
	0xc9, 0xa5, 0xf5, 0x64, 0x0d, 0x57, 0xe4, 0x87, 0x70, 0xa3, 0xef, 0x8b, 0x38, 0xe0, 0x97, 0x3d,
	0xd5, 0x95, 0xa9, 0xba, 0xb2, 0xa6, 0xb1, 0x63, 0x6a, 0xae, 0x0e, 0x15, 0x5a, 0xae, 0xce, 0x6f,
	0xaa, 0xfc, 0x16, 0x05, 0x28, 0xe9, 0x1c, 0x42, 0xf5, 0x95, 0x6a, 0x85, 0x66, 0xeb, 0xe2, 0xc7,
	0x09, 0x0a, 0xc9, 0x3c, 0x30, 0x09, 0xa0, 0x46, 0xb4, 0x3a, 0x75, 0x77, 0xe9, 0x3d, 0x5c, 0xc5,
	0x50, 0x40, 0x67, 0x0f, 0xaa, 0x87, 0x18, 0xe0, 0xf5, 0x2a, 0x4b, 0x16, 0xe5, 0x3c, 0x82, 0xed,
	0x23, 0x94, 0xeb, 0x50, 0xc7, 0x70, 0xfb, 0xad, 0x2f, 0x14, 0x4c, 0x64, 0xb8, 0x3a, 0x54, 0x62,
	0x3e, 0xc4, 0x9e, 0xf0, 0xaf, 0x34, 0x78, 0xb3, 0xbb, 0x45, 0x81, 0x13, 0xff, 0x0a, 0xd9, 0x2e,
	0x80, 0x4a, 0xca, 0x68, 0x84, 0xa1, 0x5a, 0x7f, 0xa5, 0xab, 0xe0, 0xa7, 0x14, 0x70, 0x42, 0xa8,
	0xce, 0xd4, 0x13, 0x71, 0x14, 0x0a, 0x64, 0x6d, 0xd8, 0xa4, 0xde, 0x45, 0xcd, 0x68, 0x94, 0xd6,
	0x4d, 0xa9, 0x91, 0xec, 0x09, 0xdc, 0x0a, 0xf1, 0xb3, 0xec, 0x2d, 0x68, 0xdd, 0xa4, 0xf0, 0xfb,
	0x5c, 0xef, 0x87, 0x01, 0xd5, 0x33, 0x75, 0xa3, 0x35, 0x93, 0xce, 0xdc, 0x9e, 0xbc, 0x58, 0x68,
	0x9c, 0x37, 0x64, 0xd7, 0x77, 0x5c, 0x8c, 0xb2, 0xdb, 0xd3, 0x73, 0x7e, 0xa6, 0xd2, 0x7f, 0x9e,
	0xa9, 0xf3, 0xd3, 0x04, 0x8b, 0x5e, 0x4f, 0x30, 0x49, 0xfd, 0x0b, 0x64, 0x08, 0x65, 0x7d, 0x7c,
	0xd6, 0x2c, 0x20, 0x2f, 0x78, 0xc3, 0x5e, 0x25, 0xe3, 0xec, 0x7c, 0xfb, 0xf5, 0xe7, 0xfb, 0xc6,
	0xb6, 0x53, 0xc9, 0xbe, 0x68, 0xb1, 0x6f, 0xb4, 0xd8, 0x00, 0xca, 0xda, 0x1d, 0x85, 0x32, 0x0b,
	0xe6, 0xb1, 0xef, 0x2e, 0xec, 0xe0, 0x35, 0x7d, 0xcf, 0x8e, 0xad, 0x14, 0x76, 0x5a, 0x8c, 0x14,
	0xbe, 0xd0, 0x0e, 0x5f, 0x28, 0x1d, 0xaf, 0xf5, 0x95, 0x21, 0x94, 0x8e, 0x50, 0xb2, 0xc7, 0x05,
	0x22, 0xd7, 0x8d, 0xb7, 0x7a, 0x90, 0xa9, 0x0c, 0x5b, 0x26, 0x13, 0x80, 0x49, 0x6e, 0x62, 0x7b,
	0x05, 0x05, 0xe6, 0xad, 0x6b, 0x37, 0xd7, 0x03, 0xb5, 0x27, 0x9d, 0xaa, 0x92, 0xb5, 0xd8, 0xbf,
	0xfd, 0xb1, 0x10, 0xca, 0xda, 0x4a, 0x85, 0xcb, 0x5b, 0x70, 0xda, 0xea, 0xd1, 0x76, 0x95, 0xc6,
	0x3d, 0x7b, 0xc9, 0x68, 0xfb, 0x46, 0xeb, 0xe5, 0xf3, 0x0f, 0xcf, 0x72, 0x72, 0x1f, 0x53, 0x2f,
	0xff, 0x5b, 0xfb, 0xa1, 0xc4, 0x24, 0xe4, 0x81, 0x17, 0x8f, 0x86, 0x33, 0xbf, 0x6e, 0xfa, 0x1b,
	0xa7, 0xed, 0x03, 0x1e, 0xfb, 0xe7, 0x65, 0x75, 0xab, 0xa7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x29, 0x6b, 0xeb, 0xeb, 0xe2, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	// Creates a user, and returns the new User.
	Create(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error)
	// Deletes a user. Returns NOT_FOUND if the user does not exist.
	Delete(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*types.Empty, error)
	// Gets a user. Returns NOT_FOUND if the user does not exist.
	Get(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error)
	// Lists users. The order is unspecified but deterministic. Newly created
	// users will not necessarily be added to the end of this list.
	List(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error)
	// Updates a user. Returns INVALID_ARGUMENT if the name of the user
	// is non-empty and does equal the previous name.
	Update(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Create(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/powerssl.apiserver.v1.UserService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Delete(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/powerssl.apiserver.v1.UserService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Get(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/powerssl.apiserver.v1.UserService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) List(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error) {
	out := new(ListUsersResponse)
	err := c.cc.Invoke(ctx, "/powerssl.apiserver.v1.UserService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Update(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/powerssl.apiserver.v1.UserService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	// Creates a user, and returns the new User.
	Create(context.Context, *CreateUserRequest) (*User, error)
	// Deletes a user. Returns NOT_FOUND if the user does not exist.
	Delete(context.Context, *DeleteUserRequest) (*types.Empty, error)
	// Gets a user. Returns NOT_FOUND if the user does not exist.
	Get(context.Context, *GetUserRequest) (*User, error)
	// Lists users. The order is unspecified but deterministic. Newly created
	// users will not necessarily be added to the end of this list.
	List(context.Context, *ListUsersRequest) (*ListUsersResponse, error)
	// Updates a user. Returns INVALID_ARGUMENT if the name of the user
	// is non-empty and does equal the previous name.
	Update(context.Context, *UpdateUserRequest) (*User, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) Create(ctx context.Context, req *CreateUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedUserServiceServer) Delete(ctx context.Context, req *DeleteUserRequest) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedUserServiceServer) Get(ctx context.Context, req *GetUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedUserServiceServer) List(ctx context.Context, req *ListUsersRequest) (*ListUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedUserServiceServer) Update(ctx context.Context, req *UpdateUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/powerssl.apiserver.v1.UserService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Create(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/powerssl.apiserver.v1.UserService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Delete(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/powerssl.apiserver.v1.UserService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Get(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/powerssl.apiserver.v1.UserService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).List(ctx, req.(*ListUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/powerssl.apiserver.v1.UserService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Update(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "powerssl.apiserver.v1.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _UserService_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _UserService_Delete_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _UserService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _UserService_List_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _UserService_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "powerssl/apiserver/v1/user.proto",
}
