// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/v1/certificate_authority.proto

package api // import "powerssl.io/pkg/api/v1"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
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

type CertificateAuthorityRequest struct {
	CertificateAuthority *CertificateAuthority `protobuf:"bytes,1,opt,name=certificate_authority,json=certificateAuthority" json:"certificate_authority,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CertificateAuthorityRequest) Reset()         { *m = CertificateAuthorityRequest{} }
func (m *CertificateAuthorityRequest) String() string { return proto.CompactTextString(m) }
func (*CertificateAuthorityRequest) ProtoMessage()    {}
func (*CertificateAuthorityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_certificate_authority_4d0ee494c146a995, []int{0}
}
func (m *CertificateAuthorityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CertificateAuthorityRequest.Unmarshal(m, b)
}
func (m *CertificateAuthorityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CertificateAuthorityRequest.Marshal(b, m, deterministic)
}
func (dst *CertificateAuthorityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateAuthorityRequest.Merge(dst, src)
}
func (m *CertificateAuthorityRequest) XXX_Size() int {
	return xxx_messageInfo_CertificateAuthorityRequest.Size(m)
}
func (m *CertificateAuthorityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateAuthorityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateAuthorityRequest proto.InternalMessageInfo

func (m *CertificateAuthorityRequest) GetCertificateAuthority() *CertificateAuthority {
	if m != nil {
		return m.CertificateAuthority
	}
	return nil
}

type DeleteCertificatAuthorityRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteCertificatAuthorityRequest) Reset()         { *m = DeleteCertificatAuthorityRequest{} }
func (m *DeleteCertificatAuthorityRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteCertificatAuthorityRequest) ProtoMessage()    {}
func (*DeleteCertificatAuthorityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_certificate_authority_4d0ee494c146a995, []int{1}
}
func (m *DeleteCertificatAuthorityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCertificatAuthorityRequest.Unmarshal(m, b)
}
func (m *DeleteCertificatAuthorityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCertificatAuthorityRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteCertificatAuthorityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCertificatAuthorityRequest.Merge(dst, src)
}
func (m *DeleteCertificatAuthorityRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteCertificatAuthorityRequest.Size(m)
}
func (m *DeleteCertificatAuthorityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteCertificatAuthorityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteCertificatAuthorityRequest proto.InternalMessageInfo

func (m *DeleteCertificatAuthorityRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CertificateAuthority struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CertificateAuthority) Reset()         { *m = CertificateAuthority{} }
func (m *CertificateAuthority) String() string { return proto.CompactTextString(m) }
func (*CertificateAuthority) ProtoMessage()    {}
func (*CertificateAuthority) Descriptor() ([]byte, []int) {
	return fileDescriptor_certificate_authority_4d0ee494c146a995, []int{2}
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

func (m *CertificateAuthority) GetName() string {
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
	return fileDescriptor_certificate_authority_4d0ee494c146a995, []int{3}
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
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListCertificateAuthoritiesRequest) Reset()         { *m = ListCertificateAuthoritiesRequest{} }
func (m *ListCertificateAuthoritiesRequest) String() string { return proto.CompactTextString(m) }
func (*ListCertificateAuthoritiesRequest) ProtoMessage()    {}
func (*ListCertificateAuthoritiesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_certificate_authority_4d0ee494c146a995, []int{4}
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

type ListCertificateAuthoritiesResponse struct {
	CertificateAuthority []*CertificateAuthority `protobuf:"bytes,1,rep,name=certificate_authority,json=certificateAuthority" json:"certificate_authority,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ListCertificateAuthoritiesResponse) Reset()         { *m = ListCertificateAuthoritiesResponse{} }
func (m *ListCertificateAuthoritiesResponse) String() string { return proto.CompactTextString(m) }
func (*ListCertificateAuthoritiesResponse) ProtoMessage()    {}
func (*ListCertificateAuthoritiesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_certificate_authority_4d0ee494c146a995, []int{5}
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

func (m *ListCertificateAuthoritiesResponse) GetCertificateAuthority() []*CertificateAuthority {
	if m != nil {
		return m.CertificateAuthority
	}
	return nil
}

type UpdateCertificateAuthorityRequest struct {
	CertificateAuthority *CertificateAuthority `protobuf:"bytes,1,opt,name=certificate_authority,json=certificateAuthority" json:"certificate_authority,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateCertificateAuthorityRequest) Reset()         { *m = UpdateCertificateAuthorityRequest{} }
func (m *UpdateCertificateAuthorityRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateCertificateAuthorityRequest) ProtoMessage()    {}
func (*UpdateCertificateAuthorityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_certificate_authority_4d0ee494c146a995, []int{6}
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

func (m *UpdateCertificateAuthorityRequest) GetCertificateAuthority() *CertificateAuthority {
	if m != nil {
		return m.CertificateAuthority
	}
	return nil
}

func init() {
	proto.RegisterType((*CertificateAuthorityRequest)(nil), "powerssl.v1.CertificateAuthorityRequest")
	proto.RegisterType((*DeleteCertificatAuthorityRequest)(nil), "powerssl.v1.DeleteCertificatAuthorityRequest")
	proto.RegisterType((*CertificateAuthority)(nil), "powerssl.v1.CertificateAuthority")
	proto.RegisterType((*GetCertificateAuthorityRequest)(nil), "powerssl.v1.GetCertificateAuthorityRequest")
	proto.RegisterType((*ListCertificateAuthoritiesRequest)(nil), "powerssl.v1.ListCertificateAuthoritiesRequest")
	proto.RegisterType((*ListCertificateAuthoritiesResponse)(nil), "powerssl.v1.ListCertificateAuthoritiesResponse")
	proto.RegisterType((*UpdateCertificateAuthorityRequest)(nil), "powerssl.v1.UpdateCertificateAuthorityRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CertificateAuthorityService service

type CertificateAuthorityServiceClient interface {
	CreateCertificateAuthority(ctx context.Context, in *CertificateAuthorityRequest, opts ...grpc.CallOption) (*CertificateAuthority, error)
	DeleteCertificatAuthority(ctx context.Context, in *DeleteCertificatAuthorityRequest, opts ...grpc.CallOption) (*types.Empty, error)
	GetCertificateAuthority(ctx context.Context, in *GetCertificateAuthorityRequest, opts ...grpc.CallOption) (*CertificateAuthority, error)
	ListCertificateAuthorities(ctx context.Context, in *ListCertificateAuthoritiesRequest, opts ...grpc.CallOption) (*ListCertificateAuthoritiesResponse, error)
	UpdateCertificateAuthority(ctx context.Context, in *UpdateCertificateAuthorityRequest, opts ...grpc.CallOption) (*CertificateAuthority, error)
}

type certificateAuthorityServiceClient struct {
	cc *grpc.ClientConn
}

func NewCertificateAuthorityServiceClient(cc *grpc.ClientConn) CertificateAuthorityServiceClient {
	return &certificateAuthorityServiceClient{cc}
}

func (c *certificateAuthorityServiceClient) CreateCertificateAuthority(ctx context.Context, in *CertificateAuthorityRequest, opts ...grpc.CallOption) (*CertificateAuthority, error) {
	out := new(CertificateAuthority)
	err := c.cc.Invoke(ctx, "/powerssl.v1.CertificateAuthorityService/CreateCertificateAuthority", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAuthorityServiceClient) DeleteCertificatAuthority(ctx context.Context, in *DeleteCertificatAuthorityRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/powerssl.v1.CertificateAuthorityService/DeleteCertificatAuthority", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAuthorityServiceClient) GetCertificateAuthority(ctx context.Context, in *GetCertificateAuthorityRequest, opts ...grpc.CallOption) (*CertificateAuthority, error) {
	out := new(CertificateAuthority)
	err := c.cc.Invoke(ctx, "/powerssl.v1.CertificateAuthorityService/GetCertificateAuthority", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAuthorityServiceClient) ListCertificateAuthorities(ctx context.Context, in *ListCertificateAuthoritiesRequest, opts ...grpc.CallOption) (*ListCertificateAuthoritiesResponse, error) {
	out := new(ListCertificateAuthoritiesResponse)
	err := c.cc.Invoke(ctx, "/powerssl.v1.CertificateAuthorityService/ListCertificateAuthorities", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAuthorityServiceClient) UpdateCertificateAuthority(ctx context.Context, in *UpdateCertificateAuthorityRequest, opts ...grpc.CallOption) (*CertificateAuthority, error) {
	out := new(CertificateAuthority)
	err := c.cc.Invoke(ctx, "/powerssl.v1.CertificateAuthorityService/UpdateCertificateAuthority", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CertificateAuthorityService service

type CertificateAuthorityServiceServer interface {
	CreateCertificateAuthority(context.Context, *CertificateAuthorityRequest) (*CertificateAuthority, error)
	DeleteCertificatAuthority(context.Context, *DeleteCertificatAuthorityRequest) (*types.Empty, error)
	GetCertificateAuthority(context.Context, *GetCertificateAuthorityRequest) (*CertificateAuthority, error)
	ListCertificateAuthorities(context.Context, *ListCertificateAuthoritiesRequest) (*ListCertificateAuthoritiesResponse, error)
	UpdateCertificateAuthority(context.Context, *UpdateCertificateAuthorityRequest) (*CertificateAuthority, error)
}

func RegisterCertificateAuthorityServiceServer(s *grpc.Server, srv CertificateAuthorityServiceServer) {
	s.RegisterService(&_CertificateAuthorityService_serviceDesc, srv)
}

func _CertificateAuthorityService_CreateCertificateAuthority_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateAuthorityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAuthorityServiceServer).CreateCertificateAuthority(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/powerssl.v1.CertificateAuthorityService/CreateCertificateAuthority",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAuthorityServiceServer).CreateCertificateAuthority(ctx, req.(*CertificateAuthorityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateAuthorityService_DeleteCertificatAuthority_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCertificatAuthorityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAuthorityServiceServer).DeleteCertificatAuthority(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/powerssl.v1.CertificateAuthorityService/DeleteCertificatAuthority",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAuthorityServiceServer).DeleteCertificatAuthority(ctx, req.(*DeleteCertificatAuthorityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateAuthorityService_GetCertificateAuthority_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCertificateAuthorityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAuthorityServiceServer).GetCertificateAuthority(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/powerssl.v1.CertificateAuthorityService/GetCertificateAuthority",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAuthorityServiceServer).GetCertificateAuthority(ctx, req.(*GetCertificateAuthorityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateAuthorityService_ListCertificateAuthorities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCertificateAuthoritiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAuthorityServiceServer).ListCertificateAuthorities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/powerssl.v1.CertificateAuthorityService/ListCertificateAuthorities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAuthorityServiceServer).ListCertificateAuthorities(ctx, req.(*ListCertificateAuthoritiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateAuthorityService_UpdateCertificateAuthority_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCertificateAuthorityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAuthorityServiceServer).UpdateCertificateAuthority(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/powerssl.v1.CertificateAuthorityService/UpdateCertificateAuthority",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAuthorityServiceServer).UpdateCertificateAuthority(ctx, req.(*UpdateCertificateAuthorityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CertificateAuthorityService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "powerssl.v1.CertificateAuthorityService",
	HandlerType: (*CertificateAuthorityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCertificateAuthority",
			Handler:    _CertificateAuthorityService_CreateCertificateAuthority_Handler,
		},
		{
			MethodName: "DeleteCertificatAuthority",
			Handler:    _CertificateAuthorityService_DeleteCertificatAuthority_Handler,
		},
		{
			MethodName: "GetCertificateAuthority",
			Handler:    _CertificateAuthorityService_GetCertificateAuthority_Handler,
		},
		{
			MethodName: "ListCertificateAuthorities",
			Handler:    _CertificateAuthorityService_ListCertificateAuthorities_Handler,
		},
		{
			MethodName: "UpdateCertificateAuthority",
			Handler:    _CertificateAuthorityService_UpdateCertificateAuthority_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/certificate_authority.proto",
}

func init() {
	proto.RegisterFile("api/v1/certificate_authority.proto", fileDescriptor_certificate_authority_4d0ee494c146a995)
}

var fileDescriptor_certificate_authority_4d0ee494c146a995 = []byte{
	// 359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0x4f, 0x4f, 0xc2, 0x40,
	0x10, 0xc5, 0x21, 0x12, 0x13, 0x87, 0xdb, 0x06, 0xff, 0x2d, 0xc6, 0xc0, 0x7a, 0x21, 0x1a, 0x77,
	0x03, 0x1a, 0x2f, 0x9e, 0x14, 0x8d, 0x17, 0x4f, 0x18, 0x3d, 0x78, 0x31, 0xa5, 0x0c, 0xb8, 0xe1,
	0xcf, 0x2e, 0xdd, 0x2d, 0xc6, 0x68, 0xe2, 0x67, 0xf4, 0x1b, 0x19, 0x5b, 0x11, 0xd0, 0x6d, 0x4b,
	0x62, 0xe2, 0xad, 0x69, 0xdf, 0xbc, 0x37, 0x3b, 0xfb, 0x9b, 0x02, 0xf3, 0xb4, 0x14, 0x93, 0xba,
	0xf0, 0x31, 0xb0, 0xb2, 0x2b, 0x7d, 0xcf, 0xe2, 0x83, 0x17, 0xda, 0x47, 0x15, 0x48, 0xfb, 0xcc,
	0x75, 0xa0, 0xac, 0x22, 0x45, 0xad, 0x9e, 0x30, 0x30, 0x66, 0xc0, 0x27, 0x75, 0x5a, 0xee, 0x29,
	0xd5, 0x1b, 0xa0, 0x88, 0x3e, 0xb5, 0xc3, 0xae, 0xc0, 0xa1, 0x9e, 0x2a, 0x59, 0x08, 0xe5, 0xe6,
	0xcc, 0xe8, 0x6c, 0xea, 0xd3, 0xc2, 0x71, 0x88, 0xc6, 0x92, 0x3b, 0x58, 0x77, 0xe6, 0x6c, 0xe5,
	0x2b, 0xf9, 0x5a, 0xb1, 0x51, 0xe5, 0x73, 0x41, 0xdc, 0x69, 0x54, 0xf2, 0x1d, 0x6f, 0xd9, 0x09,
	0x54, 0x2e, 0x70, 0x80, 0x16, 0x67, 0x35, 0xbf, 0xb2, 0x09, 0x14, 0x46, 0xde, 0x10, 0xa3, 0xa8,
	0xb5, 0x56, 0xf4, 0xcc, 0xf6, 0xa1, 0xe4, 0x4a, 0x71, 0x6a, 0x8f, 0x61, 0xf7, 0x0a, 0x6d, 0xda,
	0xe9, 0x5c, 0x55, 0x7b, 0x50, 0xbd, 0x96, 0xc6, 0x55, 0x26, 0xd1, 0x7c, 0x15, 0xb2, 0x57, 0x60,
	0x69, 0x22, 0xa3, 0xd5, 0xc8, 0x60, 0xda, 0xf0, 0x56, 0xfe, 0x32, 0xbc, 0x17, 0xa8, 0xde, 0xea,
	0x8e, 0x37, 0x3f, 0xbc, 0x7f, 0xbb, 0xb9, 0xc6, 0x7b, 0xc1, 0x4d, 0xcc, 0x0d, 0x06, 0x13, 0xe9,
	0x23, 0xe9, 0x03, 0x6d, 0x06, 0x98, 0xd0, 0x1c, 0xa9, 0x65, 0xc7, 0xc6, 0xfd, 0xd3, 0xec, 0x06,
	0x59, 0x8e, 0x74, 0x60, 0x3b, 0x11, 0x23, 0x72, 0xb8, 0xe0, 0x90, 0x85, 0x1b, 0xdd, 0xe0, 0xf1,
	0x9e, 0xf0, 0xe9, 0x9e, 0xf0, 0xcb, 0xcf, 0x3d, 0x61, 0x39, 0xd2, 0x87, 0xcd, 0x04, 0x90, 0xc8,
	0xc1, 0x42, 0x46, 0x3a, 0x6e, 0xcb, 0x1d, 0xe9, 0x0d, 0x68, 0x32, 0x5a, 0x84, 0x2f, 0x58, 0x64,
	0x82, 0x4a, 0xc5, 0xd2, 0xfa, 0x98, 0x59, 0x96, 0x23, 0x63, 0xa0, 0xc9, 0x74, 0xfd, 0x68, 0x20,
	0x13, 0xc3, 0xa5, 0xce, 0x7c, 0xbe, 0x73, 0x4f, 0xbf, 0x55, 0x52, 0x09, 0xdd, 0xef, 0x89, 0xf8,
	0x27, 0x77, 0xea, 0x69, 0xd9, 0x5e, 0x8d, 0x2e, 0xe4, 0xe8, 0x23, 0x00, 0x00, 0xff, 0xff, 0xac,
	0xf5, 0x64, 0xe6, 0xf9, 0x04, 0x00, 0x00,
}
