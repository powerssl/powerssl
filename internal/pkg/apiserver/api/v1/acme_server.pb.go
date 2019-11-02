// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: powerssl/apiserver/v1/acme_server.proto

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

type ACMEServer struct {
	Name                 string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CreateTime           *types.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime           *types.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	DisplayName          string           `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	DirectoryUrl         string           `protobuf:"bytes,8,opt,name=directory_url,json=directoryUrl,proto3" json:"directory_url,omitempty"`
	IntegrationName      string           `protobuf:"bytes,9,opt,name=integration_name,json=integrationName,proto3" json:"integration_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ACMEServer) Reset()         { *m = ACMEServer{} }
func (m *ACMEServer) String() string { return proto.CompactTextString(m) }
func (*ACMEServer) ProtoMessage()    {}
func (*ACMEServer) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fb4bed17a7f2272, []int{0}
}
func (m *ACMEServer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ACMEServer.Unmarshal(m, b)
}
func (m *ACMEServer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ACMEServer.Marshal(b, m, deterministic)
}
func (m *ACMEServer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ACMEServer.Merge(m, src)
}
func (m *ACMEServer) XXX_Size() int {
	return xxx_messageInfo_ACMEServer.Size(m)
}
func (m *ACMEServer) XXX_DiscardUnknown() {
	xxx_messageInfo_ACMEServer.DiscardUnknown(m)
}

var xxx_messageInfo_ACMEServer proto.InternalMessageInfo

func (m *ACMEServer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ACMEServer) GetCreateTime() *types.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *ACMEServer) GetUpdateTime() *types.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

func (m *ACMEServer) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *ACMEServer) GetDirectoryUrl() string {
	if m != nil {
		return m.DirectoryUrl
	}
	return ""
}

func (m *ACMEServer) GetIntegrationName() string {
	if m != nil {
		return m.IntegrationName
	}
	return ""
}

type CreateACMEServerRequest struct {
	AcmeServer           *ACMEServer `protobuf:"bytes,1,opt,name=acme_server,json=acmeServer,proto3" json:"acme_server,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CreateACMEServerRequest) Reset()         { *m = CreateACMEServerRequest{} }
func (m *CreateACMEServerRequest) String() string { return proto.CompactTextString(m) }
func (*CreateACMEServerRequest) ProtoMessage()    {}
func (*CreateACMEServerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fb4bed17a7f2272, []int{1}
}
func (m *CreateACMEServerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateACMEServerRequest.Unmarshal(m, b)
}
func (m *CreateACMEServerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateACMEServerRequest.Marshal(b, m, deterministic)
}
func (m *CreateACMEServerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateACMEServerRequest.Merge(m, src)
}
func (m *CreateACMEServerRequest) XXX_Size() int {
	return xxx_messageInfo_CreateACMEServerRequest.Size(m)
}
func (m *CreateACMEServerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateACMEServerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateACMEServerRequest proto.InternalMessageInfo

func (m *CreateACMEServerRequest) GetAcmeServer() *ACMEServer {
	if m != nil {
		return m.AcmeServer
	}
	return nil
}

type DeleteACMEServerRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteACMEServerRequest) Reset()         { *m = DeleteACMEServerRequest{} }
func (m *DeleteACMEServerRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteACMEServerRequest) ProtoMessage()    {}
func (*DeleteACMEServerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fb4bed17a7f2272, []int{2}
}
func (m *DeleteACMEServerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteACMEServerRequest.Unmarshal(m, b)
}
func (m *DeleteACMEServerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteACMEServerRequest.Marshal(b, m, deterministic)
}
func (m *DeleteACMEServerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteACMEServerRequest.Merge(m, src)
}
func (m *DeleteACMEServerRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteACMEServerRequest.Size(m)
}
func (m *DeleteACMEServerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteACMEServerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteACMEServerRequest proto.InternalMessageInfo

func (m *DeleteACMEServerRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetACMEServerRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetACMEServerRequest) Reset()         { *m = GetACMEServerRequest{} }
func (m *GetACMEServerRequest) String() string { return proto.CompactTextString(m) }
func (*GetACMEServerRequest) ProtoMessage()    {}
func (*GetACMEServerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fb4bed17a7f2272, []int{3}
}
func (m *GetACMEServerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetACMEServerRequest.Unmarshal(m, b)
}
func (m *GetACMEServerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetACMEServerRequest.Marshal(b, m, deterministic)
}
func (m *GetACMEServerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetACMEServerRequest.Merge(m, src)
}
func (m *GetACMEServerRequest) XXX_Size() int {
	return xxx_messageInfo_GetACMEServerRequest.Size(m)
}
func (m *GetACMEServerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetACMEServerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetACMEServerRequest proto.InternalMessageInfo

func (m *GetACMEServerRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ListACMEServersRequest struct {
	PageSize             int32    `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken            string   `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListACMEServersRequest) Reset()         { *m = ListACMEServersRequest{} }
func (m *ListACMEServersRequest) String() string { return proto.CompactTextString(m) }
func (*ListACMEServersRequest) ProtoMessage()    {}
func (*ListACMEServersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fb4bed17a7f2272, []int{4}
}
func (m *ListACMEServersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListACMEServersRequest.Unmarshal(m, b)
}
func (m *ListACMEServersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListACMEServersRequest.Marshal(b, m, deterministic)
}
func (m *ListACMEServersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListACMEServersRequest.Merge(m, src)
}
func (m *ListACMEServersRequest) XXX_Size() int {
	return xxx_messageInfo_ListACMEServersRequest.Size(m)
}
func (m *ListACMEServersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListACMEServersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListACMEServersRequest proto.InternalMessageInfo

func (m *ListACMEServersRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListACMEServersRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

type ListACMEServersResponse struct {
	AcmeServers          []*ACMEServer `protobuf:"bytes,1,rep,name=acme_servers,json=acmeServers,proto3" json:"acme_servers,omitempty"`
	NextPageToken        string        `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ListACMEServersResponse) Reset()         { *m = ListACMEServersResponse{} }
func (m *ListACMEServersResponse) String() string { return proto.CompactTextString(m) }
func (*ListACMEServersResponse) ProtoMessage()    {}
func (*ListACMEServersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fb4bed17a7f2272, []int{5}
}
func (m *ListACMEServersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListACMEServersResponse.Unmarshal(m, b)
}
func (m *ListACMEServersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListACMEServersResponse.Marshal(b, m, deterministic)
}
func (m *ListACMEServersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListACMEServersResponse.Merge(m, src)
}
func (m *ListACMEServersResponse) XXX_Size() int {
	return xxx_messageInfo_ListACMEServersResponse.Size(m)
}
func (m *ListACMEServersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListACMEServersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListACMEServersResponse proto.InternalMessageInfo

func (m *ListACMEServersResponse) GetAcmeServers() []*ACMEServer {
	if m != nil {
		return m.AcmeServers
	}
	return nil
}

func (m *ListACMEServersResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

type UpdateACMEServerRequest struct {
	Name                 string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	UpdateMask           *types.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	AcmeServer           *ACMEServer      `protobuf:"bytes,3,opt,name=acme_server,json=acmeServer,proto3" json:"acme_server,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *UpdateACMEServerRequest) Reset()         { *m = UpdateACMEServerRequest{} }
func (m *UpdateACMEServerRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateACMEServerRequest) ProtoMessage()    {}
func (*UpdateACMEServerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fb4bed17a7f2272, []int{6}
}
func (m *UpdateACMEServerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateACMEServerRequest.Unmarshal(m, b)
}
func (m *UpdateACMEServerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateACMEServerRequest.Marshal(b, m, deterministic)
}
func (m *UpdateACMEServerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateACMEServerRequest.Merge(m, src)
}
func (m *UpdateACMEServerRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateACMEServerRequest.Size(m)
}
func (m *UpdateACMEServerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateACMEServerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateACMEServerRequest proto.InternalMessageInfo

func (m *UpdateACMEServerRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateACMEServerRequest) GetUpdateMask() *types.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

func (m *UpdateACMEServerRequest) GetAcmeServer() *ACMEServer {
	if m != nil {
		return m.AcmeServer
	}
	return nil
}

func init() {
	proto.RegisterType((*ACMEServer)(nil), "powerssl.apiserver.v1.ACMEServer")
	proto.RegisterType((*CreateACMEServerRequest)(nil), "powerssl.apiserver.v1.CreateACMEServerRequest")
	proto.RegisterType((*DeleteACMEServerRequest)(nil), "powerssl.apiserver.v1.DeleteACMEServerRequest")
	proto.RegisterType((*GetACMEServerRequest)(nil), "powerssl.apiserver.v1.GetACMEServerRequest")
	proto.RegisterType((*ListACMEServersRequest)(nil), "powerssl.apiserver.v1.ListACMEServersRequest")
	proto.RegisterType((*ListACMEServersResponse)(nil), "powerssl.apiserver.v1.ListACMEServersResponse")
	proto.RegisterType((*UpdateACMEServerRequest)(nil), "powerssl.apiserver.v1.UpdateACMEServerRequest")
}

func init() {
	proto.RegisterFile("powerssl/apiserver/v1/acme_server.proto", fileDescriptor_4fb4bed17a7f2272)
}

var fileDescriptor_4fb4bed17a7f2272 = []byte{
	// 656 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x4f, 0x4f, 0x13, 0x41,
	0x14, 0xcf, 0xd2, 0xda, 0xd0, 0x57, 0x08, 0x32, 0x41, 0xda, 0x2c, 0x12, 0xcb, 0x92, 0x28, 0xd6,
	0xb0, 0x1b, 0xf0, 0x60, 0x84, 0x78, 0x90, 0x3f, 0x72, 0x11, 0x63, 0x0a, 0x5c, 0x4c, 0xcc, 0x66,
	0x68, 0x1f, 0xcd, 0xa4, 0xdb, 0xdd, 0x75, 0x66, 0x5a, 0x04, 0xa3, 0x31, 0x5e, 0x3c, 0x79, 0xf2,
	0x7b, 0xf8, 0x65, 0xfc, 0x0a, 0x7e, 0x0c, 0x0f, 0x66, 0x66, 0xda, 0x65, 0xd3, 0x76, 0xc3, 0xc6,
	0xcb, 0x66, 0xf6, 0xbd, 0xdf, 0xfb, 0x33, 0xbf, 0xf7, 0x9b, 0x19, 0x78, 0x14, 0x47, 0x97, 0xc8,
	0x85, 0x08, 0x3c, 0x1a, 0x33, 0x81, 0x7c, 0x80, 0xdc, 0x1b, 0x6c, 0x79, 0xb4, 0xd5, 0x43, 0xdf,
	0xfc, 0xba, 0x31, 0x8f, 0x64, 0x44, 0xee, 0x8d, 0x80, 0x6e, 0x02, 0x74, 0x07, 0x5b, 0xf6, 0xfd,
	0x4e, 0x14, 0x75, 0x02, 0x54, 0xd1, 0x1e, 0x0d, 0xc3, 0x48, 0x52, 0xc9, 0xa2, 0x50, 0x98, 0x20,
	0x7b, 0x65, 0xe8, 0xd5, 0x7f, 0xe7, 0xfd, 0x0b, 0x0f, 0x7b, 0xb1, 0xbc, 0x1a, 0x3a, 0xeb, 0xe3,
	0xce, 0x0b, 0x86, 0x41, 0xdb, 0xef, 0x51, 0xd1, 0x1d, 0x22, 0x1e, 0x8c, 0x23, 0x24, 0xeb, 0xa1,
	0x90, 0xb4, 0x17, 0x1b, 0x80, 0xf3, 0x63, 0x06, 0xe0, 0xe5, 0xfe, 0xf1, 0xe1, 0x89, 0xee, 0x87,
	0x10, 0x28, 0x86, 0xb4, 0x87, 0x35, 0xab, 0x6e, 0x6d, 0x94, 0x9b, 0x7a, 0x4d, 0x76, 0xa1, 0xd2,
	0xe2, 0x48, 0x25, 0xfa, 0x2a, 0xb8, 0x36, 0x53, 0xb7, 0x36, 0x2a, 0xdb, 0xb6, 0x6b, 0x32, 0xbb,
	0xa3, 0xcc, 0xee, 0xe9, 0x28, 0x73, 0x13, 0x0c, 0x5c, 0x19, 0x54, 0x70, 0x3f, 0x6e, 0x27, 0xc1,
	0x85, 0xdb, 0x83, 0x0d, 0x5c, 0x07, 0xaf, 0xc1, 0x5c, 0x9b, 0x89, 0x38, 0xa0, 0x57, 0xbe, 0xee,
	0xaa, 0xa8, 0xbb, 0xaa, 0x0c, 0x6d, 0x6f, 0x54, 0x73, 0xeb, 0x30, 0xdf, 0x66, 0x1c, 0x5b, 0x32,
	0xe2, 0x57, 0x7e, 0x9f, 0x07, 0xb5, 0x59, 0x8d, 0x99, 0x4b, 0x8c, 0x67, 0x3c, 0x20, 0x8f, 0xe1,
	0x2e, 0x0b, 0x25, 0x76, 0xb8, 0xa6, 0xd6, 0xe4, 0x2a, 0x6b, 0xdc, 0x42, 0xca, 0xae, 0xf2, 0x39,
	0xef, 0xa1, 0xba, 0xaf, 0xbb, 0xbf, 0x21, 0xa5, 0x89, 0x1f, 0xfa, 0x28, 0x24, 0xd9, 0x83, 0x4a,
	0x6a, 0xa8, 0x9a, 0xa2, 0xca, 0xf6, 0x9a, 0x3b, 0x75, 0xaa, 0x6e, 0x2a, 0x1c, 0x54, 0x94, 0x59,
	0x3b, 0x9b, 0x50, 0x3d, 0xc0, 0x00, 0xa7, 0xa5, 0x9f, 0x42, 0xbd, 0xd3, 0x80, 0xa5, 0x23, 0x94,
	0xf9, 0xb0, 0xa7, 0xb0, 0xfc, 0x9a, 0x89, 0x14, 0x58, 0x8c, 0xd0, 0x2b, 0x50, 0x8e, 0x69, 0x07,
	0x7d, 0xc1, 0xae, 0x4d, 0xc8, 0x9d, 0xe6, 0xac, 0x32, 0x9c, 0xb0, 0x6b, 0x24, 0xab, 0x00, 0xda,
	0x29, 0xa3, 0x2e, 0x86, 0x7a, 0xb8, 0xe5, 0xa6, 0x86, 0x9f, 0x2a, 0x83, 0xf3, 0xdd, 0x82, 0xea,
	0x44, 0x5a, 0x11, 0x47, 0xa1, 0x40, 0x72, 0x00, 0x73, 0x29, 0x42, 0x44, 0xcd, 0xaa, 0x17, 0xf2,
	0x31, 0x52, 0xb9, 0x61, 0x44, 0x90, 0x87, 0xb0, 0x10, 0xe2, 0x47, 0xe9, 0x4f, 0x74, 0x31, 0xaf,
	0xcc, 0x6f, 0x93, 0x4e, 0x7e, 0x59, 0x50, 0x3d, 0xd3, 0xda, 0xc8, 0xc5, 0x47, 0x4a, 0x79, 0xea,
	0x3c, 0x64, 0xca, 0xf6, 0x95, 0x3a, 0x32, 0xc7, 0x54, 0x74, 0x47, 0xca, 0x53, 0xeb, 0xf1, 0x59,
	0x17, 0xfe, 0x63, 0xd6, 0xdb, 0x7f, 0x8b, 0xb0, 0x78, 0xe3, 0x52, 0x5f, 0xd6, 0x42, 0x72, 0x09,
	0x25, 0x23, 0x30, 0xe2, 0x66, 0xa4, 0xcb, 0xd0, 0x9f, 0x7d, 0x7b, 0x79, 0xc7, 0xfe, 0xf6, 0xfb,
	0xcf, 0xcf, 0x99, 0x25, 0x67, 0x61, 0x74, 0x03, 0x0d, 0x49, 0xde, 0xb1, 0x1a, 0x84, 0x43, 0xc9,
	0x48, 0x2f, 0xb3, 0x70, 0x86, 0x32, 0xed, 0xe5, 0x09, 0xd2, 0x0e, 0xd5, 0x25, 0xe4, 0xd4, 0x75,
	0x35, 0xbb, 0x51, 0x53, 0xd5, 0x3e, 0x29, 0xd2, 0x5f, 0xa4, 0x6a, 0x7a, 0x8d, 0xcf, 0xe4, 0x12,
	0x0a, 0x47, 0x28, 0xc9, 0x93, 0x8c, 0x82, 0xd3, 0xb4, 0x9d, 0x67, 0x9b, 0xc3, 0xc2, 0x24, 0xbb,
	0xf0, 0x17, 0x28, 0x2a, 0xd5, 0x92, 0xcd, 0x8c, 0x64, 0xd3, 0x4f, 0x8a, 0xed, 0xe6, 0x85, 0x9b,
	0x13, 0xe0, 0x54, 0x75, 0x23, 0x8b, 0x64, 0x9c, 0x6f, 0xf2, 0xd5, 0x82, 0x92, 0x11, 0x6b, 0x26,
	0xdb, 0x19, 0x5a, 0xce, 0xb3, 0xff, 0x75, 0x5d, 0x76, 0xd5, 0xce, 0xdc, 0xff, 0x8e, 0xd5, 0xd8,
	0x7b, 0xfe, 0xee, 0x59, 0x92, 0xa8, 0x8d, 0x03, 0x2f, 0x79, 0xa6, 0xd4, 0x95, 0xc7, 0x43, 0x1a,
	0x78, 0x71, 0xb7, 0x93, 0x7a, 0xb3, 0xd4, 0xfb, 0x33, 0xd8, 0xda, 0xa5, 0x31, 0x3b, 0x2f, 0xe9,
	0x41, 0x3f, 0xfd, 0x17, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x52, 0x7f, 0xad, 0xdb, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ACMEServerServiceClient is the client API for ACMEServerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ACMEServerServiceClient interface {
	Create(ctx context.Context, in *CreateACMEServerRequest, opts ...grpc.CallOption) (*ACMEServer, error)
	Delete(ctx context.Context, in *DeleteACMEServerRequest, opts ...grpc.CallOption) (*types.Empty, error)
	Get(ctx context.Context, in *GetACMEServerRequest, opts ...grpc.CallOption) (*ACMEServer, error)
	List(ctx context.Context, in *ListACMEServersRequest, opts ...grpc.CallOption) (*ListACMEServersResponse, error)
	Update(ctx context.Context, in *UpdateACMEServerRequest, opts ...grpc.CallOption) (*ACMEServer, error)
}

type aCMEServerServiceClient struct {
	cc *grpc.ClientConn
}

func NewACMEServerServiceClient(cc *grpc.ClientConn) ACMEServerServiceClient {
	return &aCMEServerServiceClient{cc}
}

func (c *aCMEServerServiceClient) Create(ctx context.Context, in *CreateACMEServerRequest, opts ...grpc.CallOption) (*ACMEServer, error) {
	out := new(ACMEServer)
	err := c.cc.Invoke(ctx, "/powerssl.apiserver.v1.ACMEServerService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aCMEServerServiceClient) Delete(ctx context.Context, in *DeleteACMEServerRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/powerssl.apiserver.v1.ACMEServerService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aCMEServerServiceClient) Get(ctx context.Context, in *GetACMEServerRequest, opts ...grpc.CallOption) (*ACMEServer, error) {
	out := new(ACMEServer)
	err := c.cc.Invoke(ctx, "/powerssl.apiserver.v1.ACMEServerService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aCMEServerServiceClient) List(ctx context.Context, in *ListACMEServersRequest, opts ...grpc.CallOption) (*ListACMEServersResponse, error) {
	out := new(ListACMEServersResponse)
	err := c.cc.Invoke(ctx, "/powerssl.apiserver.v1.ACMEServerService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aCMEServerServiceClient) Update(ctx context.Context, in *UpdateACMEServerRequest, opts ...grpc.CallOption) (*ACMEServer, error) {
	out := new(ACMEServer)
	err := c.cc.Invoke(ctx, "/powerssl.apiserver.v1.ACMEServerService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ACMEServerServiceServer is the server API for ACMEServerService service.
type ACMEServerServiceServer interface {
	Create(context.Context, *CreateACMEServerRequest) (*ACMEServer, error)
	Delete(context.Context, *DeleteACMEServerRequest) (*types.Empty, error)
	Get(context.Context, *GetACMEServerRequest) (*ACMEServer, error)
	List(context.Context, *ListACMEServersRequest) (*ListACMEServersResponse, error)
	Update(context.Context, *UpdateACMEServerRequest) (*ACMEServer, error)
}

// UnimplementedACMEServerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedACMEServerServiceServer struct {
}

func (*UnimplementedACMEServerServiceServer) Create(ctx context.Context, req *CreateACMEServerRequest) (*ACMEServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedACMEServerServiceServer) Delete(ctx context.Context, req *DeleteACMEServerRequest) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedACMEServerServiceServer) Get(ctx context.Context, req *GetACMEServerRequest) (*ACMEServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedACMEServerServiceServer) List(ctx context.Context, req *ListACMEServersRequest) (*ListACMEServersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedACMEServerServiceServer) Update(ctx context.Context, req *UpdateACMEServerRequest) (*ACMEServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}

func RegisterACMEServerServiceServer(s *grpc.Server, srv ACMEServerServiceServer) {
	s.RegisterService(&_ACMEServerService_serviceDesc, srv)
}

func _ACMEServerService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateACMEServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ACMEServerServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/powerssl.apiserver.v1.ACMEServerService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ACMEServerServiceServer).Create(ctx, req.(*CreateACMEServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ACMEServerService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteACMEServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ACMEServerServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/powerssl.apiserver.v1.ACMEServerService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ACMEServerServiceServer).Delete(ctx, req.(*DeleteACMEServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ACMEServerService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetACMEServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ACMEServerServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/powerssl.apiserver.v1.ACMEServerService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ACMEServerServiceServer).Get(ctx, req.(*GetACMEServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ACMEServerService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListACMEServersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ACMEServerServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/powerssl.apiserver.v1.ACMEServerService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ACMEServerServiceServer).List(ctx, req.(*ListACMEServersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ACMEServerService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateACMEServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ACMEServerServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/powerssl.apiserver.v1.ACMEServerService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ACMEServerServiceServer).Update(ctx, req.(*UpdateACMEServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ACMEServerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "powerssl.apiserver.v1.ACMEServerService",
	HandlerType: (*ACMEServerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ACMEServerService_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ACMEServerService_Delete_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ACMEServerService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ACMEServerService_List_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ACMEServerService_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "powerssl/apiserver/v1/acme_server.proto",
}
