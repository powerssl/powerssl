// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/v1/certificate_authority.proto

package api // import "powerssl.io/pkg/api/v1"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/googleapis/google/api"
import types "github.com/gogo/protobuf/types"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type CertificateAuthority struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CertificateAuthority) Reset()         { *m = CertificateAuthority{} }
func (m *CertificateAuthority) String() string { return proto.CompactTextString(m) }
func (*CertificateAuthority) ProtoMessage()    {}
func (*CertificateAuthority) Descriptor() ([]byte, []int) {
	return fileDescriptor_certificate_authority_d6b013fd7ab068a4, []int{0}
}
func (m *CertificateAuthority) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CertificateAuthority.Unmarshal(m, b)
}
func (m *CertificateAuthority) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CertificateAuthority.Marshal(b, m, deterministic)
}
func (dst *CertificateAuthority) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateAuthority.Merge(dst, src)
}
func (m *CertificateAuthority) XXX_Size() int {
	return xxx_messageInfo_CertificateAuthority.Size(m)
}
func (m *CertificateAuthority) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateAuthority.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateAuthority proto.InternalMessageInfo

type CreateCertificateAuthorityRequest struct {
	CertificateAuthority *CertificateAuthority `protobuf:"bytes,1,opt,name=certificate_authority,json=certificateAuthority" json:"certificate_authority,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CreateCertificateAuthorityRequest) Reset()         { *m = CreateCertificateAuthorityRequest{} }
func (m *CreateCertificateAuthorityRequest) String() string { return proto.CompactTextString(m) }
func (*CreateCertificateAuthorityRequest) ProtoMessage()    {}
func (*CreateCertificateAuthorityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_certificate_authority_d6b013fd7ab068a4, []int{1}
}
func (m *CreateCertificateAuthorityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCertificateAuthorityRequest.Unmarshal(m, b)
}
func (m *CreateCertificateAuthorityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCertificateAuthorityRequest.Marshal(b, m, deterministic)
}
func (dst *CreateCertificateAuthorityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCertificateAuthorityRequest.Merge(dst, src)
}
func (m *CreateCertificateAuthorityRequest) XXX_Size() int {
	return xxx_messageInfo_CreateCertificateAuthorityRequest.Size(m)
}
func (m *CreateCertificateAuthorityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCertificateAuthorityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCertificateAuthorityRequest proto.InternalMessageInfo

func (m *CreateCertificateAuthorityRequest) GetCertificateAuthority() *CertificateAuthority {
	if m != nil {
		return m.CertificateAuthority
	}
	return nil
}

type DeleteCertificateAuthorityRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteCertificateAuthorityRequest) Reset()         { *m = DeleteCertificateAuthorityRequest{} }
func (m *DeleteCertificateAuthorityRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteCertificateAuthorityRequest) ProtoMessage()    {}
func (*DeleteCertificateAuthorityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_certificate_authority_d6b013fd7ab068a4, []int{2}
}
func (m *DeleteCertificateAuthorityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCertificateAuthorityRequest.Unmarshal(m, b)
}
func (m *DeleteCertificateAuthorityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCertificateAuthorityRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteCertificateAuthorityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCertificateAuthorityRequest.Merge(dst, src)
}
func (m *DeleteCertificateAuthorityRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteCertificateAuthorityRequest.Size(m)
}
func (m *DeleteCertificateAuthorityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteCertificateAuthorityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteCertificateAuthorityRequest proto.InternalMessageInfo

func (m *DeleteCertificateAuthorityRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetCertificateAuthorityRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCertificateAuthorityRequest) Reset()         { *m = GetCertificateAuthorityRequest{} }
func (m *GetCertificateAuthorityRequest) String() string { return proto.CompactTextString(m) }
func (*GetCertificateAuthorityRequest) ProtoMessage()    {}
func (*GetCertificateAuthorityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_certificate_authority_d6b013fd7ab068a4, []int{3}
}
func (m *GetCertificateAuthorityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCertificateAuthorityRequest.Unmarshal(m, b)
}
func (m *GetCertificateAuthorityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCertificateAuthorityRequest.Marshal(b, m, deterministic)
}
func (dst *GetCertificateAuthorityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCertificateAuthorityRequest.Merge(dst, src)
}
func (m *GetCertificateAuthorityRequest) XXX_Size() int {
	return xxx_messageInfo_GetCertificateAuthorityRequest.Size(m)
}
func (m *GetCertificateAuthorityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCertificateAuthorityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCertificateAuthorityRequest proto.InternalMessageInfo

func (m *GetCertificateAuthorityRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ListCertificateAuthoritiesRequest struct {
	PageSize             int32    `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken            string   `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListCertificateAuthoritiesRequest) Reset()         { *m = ListCertificateAuthoritiesRequest{} }
func (m *ListCertificateAuthoritiesRequest) String() string { return proto.CompactTextString(m) }
func (*ListCertificateAuthoritiesRequest) ProtoMessage()    {}
func (*ListCertificateAuthoritiesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_certificate_authority_d6b013fd7ab068a4, []int{4}
}
func (m *ListCertificateAuthoritiesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCertificateAuthoritiesRequest.Unmarshal(m, b)
}
func (m *ListCertificateAuthoritiesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCertificateAuthoritiesRequest.Marshal(b, m, deterministic)
}
func (dst *ListCertificateAuthoritiesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCertificateAuthoritiesRequest.Merge(dst, src)
}
func (m *ListCertificateAuthoritiesRequest) XXX_Size() int {
	return xxx_messageInfo_ListCertificateAuthoritiesRequest.Size(m)
}
func (m *ListCertificateAuthoritiesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCertificateAuthoritiesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListCertificateAuthoritiesRequest proto.InternalMessageInfo

func (m *ListCertificateAuthoritiesRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListCertificateAuthoritiesRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

type ListCertificateAuthoritiesResponse struct {
	CertificateAuthorities []*CertificateAuthority `protobuf:"bytes,1,rep,name=certificate_authorities,json=certificateAuthorities" json:"certificate_authorities,omitempty"`
	NextPageToken          string                  `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                `json:"-"`
	XXX_unrecognized       []byte                  `json:"-"`
	XXX_sizecache          int32                   `json:"-"`
}

func (m *ListCertificateAuthoritiesResponse) Reset()         { *m = ListCertificateAuthoritiesResponse{} }
func (m *ListCertificateAuthoritiesResponse) String() string { return proto.CompactTextString(m) }
func (*ListCertificateAuthoritiesResponse) ProtoMessage()    {}
func (*ListCertificateAuthoritiesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_certificate_authority_d6b013fd7ab068a4, []int{5}
}
func (m *ListCertificateAuthoritiesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCertificateAuthoritiesResponse.Unmarshal(m, b)
}
func (m *ListCertificateAuthoritiesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCertificateAuthoritiesResponse.Marshal(b, m, deterministic)
}
func (dst *ListCertificateAuthoritiesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCertificateAuthoritiesResponse.Merge(dst, src)
}
func (m *ListCertificateAuthoritiesResponse) XXX_Size() int {
	return xxx_messageInfo_ListCertificateAuthoritiesResponse.Size(m)
}
func (m *ListCertificateAuthoritiesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCertificateAuthoritiesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListCertificateAuthoritiesResponse proto.InternalMessageInfo

func (m *ListCertificateAuthoritiesResponse) GetCertificateAuthorities() []*CertificateAuthority {
	if m != nil {
		return m.CertificateAuthorities
	}
	return nil
}

func (m *ListCertificateAuthoritiesResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

type UpdateCertificateAuthorityRequest struct {
	Name                 string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CertificateAuthority *CertificateAuthority `protobuf:"bytes,2,opt,name=certificate_authority,json=certificateAuthority" json:"certificate_authority,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateCertificateAuthorityRequest) Reset()         { *m = UpdateCertificateAuthorityRequest{} }
func (m *UpdateCertificateAuthorityRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateCertificateAuthorityRequest) ProtoMessage()    {}
func (*UpdateCertificateAuthorityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_certificate_authority_d6b013fd7ab068a4, []int{6}
}
func (m *UpdateCertificateAuthorityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateCertificateAuthorityRequest.Unmarshal(m, b)
}
func (m *UpdateCertificateAuthorityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateCertificateAuthorityRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateCertificateAuthorityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCertificateAuthorityRequest.Merge(dst, src)
}
func (m *UpdateCertificateAuthorityRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateCertificateAuthorityRequest.Size(m)
}
func (m *UpdateCertificateAuthorityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCertificateAuthorityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCertificateAuthorityRequest proto.InternalMessageInfo

func (m *UpdateCertificateAuthorityRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateCertificateAuthorityRequest) GetCertificateAuthority() *CertificateAuthority {
	if m != nil {
		return m.CertificateAuthority
	}
	return nil
}

func init() {
	proto.RegisterType((*CertificateAuthority)(nil), "powerssl.apiserver.v1.CertificateAuthority")
	proto.RegisterType((*CreateCertificateAuthorityRequest)(nil), "powerssl.apiserver.v1.CreateCertificateAuthorityRequest")
	proto.RegisterType((*DeleteCertificateAuthorityRequest)(nil), "powerssl.apiserver.v1.DeleteCertificateAuthorityRequest")
	proto.RegisterType((*GetCertificateAuthorityRequest)(nil), "powerssl.apiserver.v1.GetCertificateAuthorityRequest")
	proto.RegisterType((*ListCertificateAuthoritiesRequest)(nil), "powerssl.apiserver.v1.ListCertificateAuthoritiesRequest")
	proto.RegisterType((*ListCertificateAuthoritiesResponse)(nil), "powerssl.apiserver.v1.ListCertificateAuthoritiesResponse")
	proto.RegisterType((*UpdateCertificateAuthorityRequest)(nil), "powerssl.apiserver.v1.UpdateCertificateAuthorityRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CertificateAuthorityService service

type CertificateAuthorityServiceClient interface {
	Create(ctx context.Context, in *CreateCertificateAuthorityRequest, opts ...grpc.CallOption) (*CertificateAuthority, error)
	Delete(ctx context.Context, in *DeleteCertificateAuthorityRequest, opts ...grpc.CallOption) (*types.Empty, error)
	Get(ctx context.Context, in *GetCertificateAuthorityRequest, opts ...grpc.CallOption) (*CertificateAuthority, error)
	List(ctx context.Context, in *ListCertificateAuthoritiesRequest, opts ...grpc.CallOption) (*ListCertificateAuthoritiesResponse, error)
	Update(ctx context.Context, in *UpdateCertificateAuthorityRequest, opts ...grpc.CallOption) (*CertificateAuthority, error)
}

type certificateAuthorityServiceClient struct {
	cc *grpc.ClientConn
}

func NewCertificateAuthorityServiceClient(cc *grpc.ClientConn) CertificateAuthorityServiceClient {
	return &certificateAuthorityServiceClient{cc}
}

func (c *certificateAuthorityServiceClient) Create(ctx context.Context, in *CreateCertificateAuthorityRequest, opts ...grpc.CallOption) (*CertificateAuthority, error) {
	out := new(CertificateAuthority)
	err := c.cc.Invoke(ctx, "/powerssl.apiserver.v1.CertificateAuthorityService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAuthorityServiceClient) Delete(ctx context.Context, in *DeleteCertificateAuthorityRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/powerssl.apiserver.v1.CertificateAuthorityService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAuthorityServiceClient) Get(ctx context.Context, in *GetCertificateAuthorityRequest, opts ...grpc.CallOption) (*CertificateAuthority, error) {
	out := new(CertificateAuthority)
	err := c.cc.Invoke(ctx, "/powerssl.apiserver.v1.CertificateAuthorityService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAuthorityServiceClient) List(ctx context.Context, in *ListCertificateAuthoritiesRequest, opts ...grpc.CallOption) (*ListCertificateAuthoritiesResponse, error) {
	out := new(ListCertificateAuthoritiesResponse)
	err := c.cc.Invoke(ctx, "/powerssl.apiserver.v1.CertificateAuthorityService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAuthorityServiceClient) Update(ctx context.Context, in *UpdateCertificateAuthorityRequest, opts ...grpc.CallOption) (*CertificateAuthority, error) {
	out := new(CertificateAuthority)
	err := c.cc.Invoke(ctx, "/powerssl.apiserver.v1.CertificateAuthorityService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CertificateAuthorityService service

type CertificateAuthorityServiceServer interface {
	Create(context.Context, *CreateCertificateAuthorityRequest) (*CertificateAuthority, error)
	Delete(context.Context, *DeleteCertificateAuthorityRequest) (*types.Empty, error)
	Get(context.Context, *GetCertificateAuthorityRequest) (*CertificateAuthority, error)
	List(context.Context, *ListCertificateAuthoritiesRequest) (*ListCertificateAuthoritiesResponse, error)
	Update(context.Context, *UpdateCertificateAuthorityRequest) (*CertificateAuthority, error)
}

func RegisterCertificateAuthorityServiceServer(s *grpc.Server, srv CertificateAuthorityServiceServer) {
	s.RegisterService(&_CertificateAuthorityService_serviceDesc, srv)
}

func _CertificateAuthorityService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCertificateAuthorityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAuthorityServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/powerssl.apiserver.v1.CertificateAuthorityService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAuthorityServiceServer).Create(ctx, req.(*CreateCertificateAuthorityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateAuthorityService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCertificateAuthorityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAuthorityServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/powerssl.apiserver.v1.CertificateAuthorityService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAuthorityServiceServer).Delete(ctx, req.(*DeleteCertificateAuthorityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateAuthorityService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCertificateAuthorityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAuthorityServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/powerssl.apiserver.v1.CertificateAuthorityService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAuthorityServiceServer).Get(ctx, req.(*GetCertificateAuthorityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateAuthorityService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCertificateAuthoritiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAuthorityServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/powerssl.apiserver.v1.CertificateAuthorityService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAuthorityServiceServer).List(ctx, req.(*ListCertificateAuthoritiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateAuthorityService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCertificateAuthorityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAuthorityServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/powerssl.apiserver.v1.CertificateAuthorityService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAuthorityServiceServer).Update(ctx, req.(*UpdateCertificateAuthorityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CertificateAuthorityService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "powerssl.apiserver.v1.CertificateAuthorityService",
	HandlerType: (*CertificateAuthorityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _CertificateAuthorityService_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CertificateAuthorityService_Delete_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _CertificateAuthorityService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _CertificateAuthorityService_List_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CertificateAuthorityService_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/certificate_authority.proto",
}

func init() {
	proto.RegisterFile("api/v1/certificate_authority.proto", fileDescriptor_certificate_authority_d6b013fd7ab068a4)
}

var fileDescriptor_certificate_authority_d6b013fd7ab068a4 = []byte{
	// 513 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0x41, 0x4f, 0x13, 0x41,
	0x1c, 0xc5, 0x33, 0x05, 0x37, 0xf2, 0x37, 0x86, 0x64, 0x02, 0xb5, 0xd9, 0x82, 0xa1, 0xa3, 0x12,
	0x43, 0x74, 0x37, 0x14, 0x12, 0x13, 0x0c, 0x07, 0xad, 0x86, 0x8b, 0x07, 0x53, 0xf4, 0xc2, 0xc1,
	0x66, 0x58, 0xff, 0xd4, 0x09, 0x65, 0x67, 0xdc, 0x99, 0x56, 0xc1, 0x18, 0x13, 0xae, 0x1e, 0xe1,
	0xee, 0xd5, 0x83, 0xdf, 0xc6, 0xaf, 0xe0, 0x07, 0x31, 0x3b, 0x43, 0xd1, 0xd0, 0x69, 0x77, 0x1b,
	0x6e, 0xed, 0xcc, 0xfc, 0xe7, 0xfd, 0xf2, 0xe6, 0xbd, 0x2c, 0x30, 0xae, 0x44, 0x3c, 0x58, 0x8f,
	0x13, 0xcc, 0x8c, 0x38, 0x10, 0x09, 0x37, 0xd8, 0xe1, 0x7d, 0xf3, 0x41, 0x66, 0xc2, 0x1c, 0x47,
	0x2a, 0x93, 0x46, 0xd2, 0x79, 0x25, 0x3f, 0x61, 0xa6, 0x75, 0x2f, 0xe2, 0x4a, 0x44, 0x83, 0xf5,
	0x70, 0xa9, 0x2b, 0x65, 0xb7, 0x87, 0x71, 0x3e, 0xcb, 0xd3, 0x54, 0x1a, 0x6e, 0x84, 0x4c, 0xb5,
	0x3b, 0x1e, 0xd6, 0x2f, 0x76, 0xed, 0xbf, 0xfd, 0xfe, 0x41, 0x8c, 0x47, 0x6a, 0x78, 0x17, 0xab,
	0xc2, 0x42, 0xeb, 0x9f, 0xd4, 0xb3, 0xa1, 0x12, 0xfb, 0x06, 0x8d, 0x56, 0x86, 0xdc, 0xa0, 0x6f,
	0xb7, 0x8d, 0x1f, 0xfb, 0xa8, 0x0d, 0xdd, 0x83, 0x45, 0x2f, 0x67, 0x8d, 0xac, 0x90, 0x87, 0xb7,
	0x9a, 0x0f, 0xa2, 0x2b, 0xa0, 0x91, 0xf7, 0xb2, 0x85, 0xc4, 0x07, 0xf0, 0x04, 0x1a, 0x2f, 0xb0,
	0x87, 0x93, 0x01, 0x28, 0xcc, 0xa6, 0xfc, 0x08, 0xad, 0xde, 0x5c, 0xdb, 0xfe, 0x66, 0x9b, 0x70,
	0x77, 0x07, 0xcd, 0xb4, 0x53, 0x1d, 0x68, 0xbc, 0x12, 0xda, 0x37, 0x26, 0x50, 0x0f, 0x07, 0xeb,
	0x30, 0xa7, 0x78, 0x17, 0x3b, 0x5a, 0x9c, 0xb8, 0xe9, 0x1b, 0xed, 0x9b, 0xf9, 0xc2, 0xae, 0x38,
	0x41, 0xba, 0x0c, 0x60, 0x37, 0x8d, 0x3c, 0xc4, 0xb4, 0x56, 0xb1, 0x77, 0xdb, 0xe3, 0x6f, 0xf2,
	0x05, 0xf6, 0x8b, 0x00, 0x9b, 0xa4, 0xa0, 0x95, 0x4c, 0x35, 0xd2, 0x77, 0x70, 0xc7, 0x67, 0xa9,
	0x40, 0x5d, 0x23, 0x2b, 0x33, 0xe5, 0x4d, 0xad, 0x26, 0x5e, 0x1d, 0xba, 0x0a, 0xf3, 0x29, 0x7e,
	0x36, 0x9d, 0x11, 0xd4, 0xdb, 0xf9, 0xf2, 0xeb, 0x4b, 0xdc, 0x33, 0x02, 0x8d, 0xb7, 0xea, 0x3d,
	0x9f, 0xda, 0xff, 0xf1, 0xa1, 0xa8, 0x5c, 0x3b, 0x14, 0xcd, 0xf3, 0x00, 0xea, 0xbe, 0xe3, 0xbb,
	0x98, 0x0d, 0x44, 0x82, 0xf4, 0x07, 0x81, 0xc0, 0xc5, 0x96, 0x36, 0x47, 0x75, 0x8a, 0xf2, 0x1c,
	0x96, 0x63, 0x63, 0x5b, 0xa7, 0xbf, 0xff, 0x9c, 0x55, 0x36, 0x59, 0xfd, 0x4a, 0x51, 0x1f, 0xff,
	0xf7, 0x5a, 0x5b, 0x8b, 0xbe, 0x8d, 0x63, 0x7a, 0x4a, 0x20, 0x70, 0xb9, 0xf6, 0x10, 0x16, 0x06,
	0x3e, 0xac, 0x46, 0xae, 0xcc, 0xd1, 0xb0, 0xcc, 0xd1, 0xcb, 0xbc, 0xcc, 0xec, 0x91, 0x45, 0x5a,
	0x5d, 0xbb, 0x9f, 0x23, 0x7d, 0xc9, 0xdf, 0x61, 0x7b, 0x0c, 0x58, 0xbc, 0xf6, 0x95, 0x7e, 0x27,
	0x30, 0xb3, 0x83, 0x86, 0xc6, 0x23, 0x04, 0x93, 0x9b, 0x53, 0xd6, 0xa0, 0x0b, 0x1a, 0x5a, 0x8e,
	0xe6, 0x9c, 0xc0, 0x6c, 0xde, 0x0c, 0x8f, 0x21, 0x85, 0x95, 0x0c, 0x37, 0xa6, 0x9a, 0x71, 0x25,
	0x63, 0xf7, 0x2c, 0xdf, 0x32, 0x9d, 0xf4, 0x80, 0xf4, 0x27, 0x81, 0xc0, 0x35, 0xc0, 0x03, 0x56,
	0x58, 0x8d, 0xb2, 0x56, 0xb5, 0x2c, 0xca, 0x76, 0x58, 0xca, 0xaa, 0x31, 0xa1, 0x7a, 0xbe, 0xb4,
	0x17, 0x5e, 0x8a, 0x09, 0x19, 0xab, 0xc3, 0x6e, 0xec, 0x3e, 0x23, 0x4f, 0xb9, 0x12, 0xfb, 0x81,
	0xcd, 0xca, 0xc6, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x54, 0x0a, 0x85, 0x13, 0x5b, 0x06, 0x00,
	0x00,
}
