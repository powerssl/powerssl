// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/v1/certificate.proto

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

type DigestAlgorithm int32

const (
	DigestAlgorithm_DIGEST_ALGORITHM_UNSPECIFIED DigestAlgorithm = 0
	DigestAlgorithm_SHA1                         DigestAlgorithm = 1
	DigestAlgorithm_SHA256                       DigestAlgorithm = 2
)

var DigestAlgorithm_name = map[int32]string{
	0: "DIGEST_ALGORITHM_UNSPECIFIED",
	1: "SHA1",
	2: "SHA256",
}
var DigestAlgorithm_value = map[string]int32{
	"DIGEST_ALGORITHM_UNSPECIFIED": 0,
	"SHA1":                         1,
	"SHA256":                       2,
}

func (x DigestAlgorithm) String() string {
	return proto.EnumName(DigestAlgorithm_name, int32(x))
}
func (DigestAlgorithm) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_certificate_53162111b312b807, []int{0}
}

type KeyAlgorithm int32

const (
	KeyAlgorithm_KEY_ALGORITHM_UNSPECIFIED KeyAlgorithm = 0
	KeyAlgorithm_RSA                       KeyAlgorithm = 1
)

var KeyAlgorithm_name = map[int32]string{
	0: "KEY_ALGORITHM_UNSPECIFIED",
	1: "RSA",
}
var KeyAlgorithm_value = map[string]int32{
	"KEY_ALGORITHM_UNSPECIFIED": 0,
	"RSA":                       1,
}

func (x KeyAlgorithm) String() string {
	return proto.EnumName(KeyAlgorithm_name, int32(x))
}
func (KeyAlgorithm) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_certificate_53162111b312b807, []int{1}
}

type Certificate struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CreateTime           *types.Timestamp  `protobuf:"bytes,2,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
	UpdateTime           *types.Timestamp  `protobuf:"bytes,3,opt,name=update_time,json=updateTime" json:"update_time,omitempty"`
	DisplayName          string            `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Title                string            `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	Description          string            `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	Labels               map[string]string `protobuf:"bytes,7,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Dnsnames             []string          `protobuf:"bytes,8,rep,name=dnsnames" json:"dnsnames,omitempty"`
	KeyAlgorithm         KeyAlgorithm      `protobuf:"varint,9,opt,name=key_algorithm,json=keyAlgorithm,proto3,enum=powerssl.api.v1.KeyAlgorithm" json:"key_algorithm,omitempty"`
	KeySize              int32             `protobuf:"varint,10,opt,name=key_size,json=keySize,proto3" json:"key_size,omitempty"`
	DigestAlgorithm      DigestAlgorithm   `protobuf:"varint,11,opt,name=digest_algorithm,json=digestAlgorithm,proto3,enum=powerssl.api.v1.DigestAlgorithm" json:"digest_algorithm,omitempty"`
	AutoRenew            bool              `protobuf:"varint,12,opt,name=auto_renew,json=autoRenew,proto3" json:"auto_renew,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Certificate) Reset()         { *m = Certificate{} }
func (m *Certificate) String() string { return proto.CompactTextString(m) }
func (*Certificate) ProtoMessage()    {}
func (*Certificate) Descriptor() ([]byte, []int) {
	return fileDescriptor_certificate_53162111b312b807, []int{0}
}
func (m *Certificate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Certificate.Unmarshal(m, b)
}
func (m *Certificate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Certificate.Marshal(b, m, deterministic)
}
func (dst *Certificate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Certificate.Merge(dst, src)
}
func (m *Certificate) XXX_Size() int {
	return xxx_messageInfo_Certificate.Size(m)
}
func (m *Certificate) XXX_DiscardUnknown() {
	xxx_messageInfo_Certificate.DiscardUnknown(m)
}

var xxx_messageInfo_Certificate proto.InternalMessageInfo

func (m *Certificate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Certificate) GetCreateTime() *types.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Certificate) GetUpdateTime() *types.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

func (m *Certificate) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Certificate) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Certificate) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Certificate) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Certificate) GetDnsnames() []string {
	if m != nil {
		return m.Dnsnames
	}
	return nil
}

func (m *Certificate) GetKeyAlgorithm() KeyAlgorithm {
	if m != nil {
		return m.KeyAlgorithm
	}
	return KeyAlgorithm_KEY_ALGORITHM_UNSPECIFIED
}

func (m *Certificate) GetKeySize() int32 {
	if m != nil {
		return m.KeySize
	}
	return 0
}

func (m *Certificate) GetDigestAlgorithm() DigestAlgorithm {
	if m != nil {
		return m.DigestAlgorithm
	}
	return DigestAlgorithm_DIGEST_ALGORITHM_UNSPECIFIED
}

func (m *Certificate) GetAutoRenew() bool {
	if m != nil {
		return m.AutoRenew
	}
	return false
}

type CreateCertificateRequest struct {
	Certificate          *Certificate `protobuf:"bytes,1,opt,name=certificate" json:"certificate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateCertificateRequest) Reset()         { *m = CreateCertificateRequest{} }
func (m *CreateCertificateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateCertificateRequest) ProtoMessage()    {}
func (*CreateCertificateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_certificate_53162111b312b807, []int{1}
}
func (m *CreateCertificateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCertificateRequest.Unmarshal(m, b)
}
func (m *CreateCertificateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCertificateRequest.Marshal(b, m, deterministic)
}
func (dst *CreateCertificateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCertificateRequest.Merge(dst, src)
}
func (m *CreateCertificateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateCertificateRequest.Size(m)
}
func (m *CreateCertificateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCertificateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCertificateRequest proto.InternalMessageInfo

func (m *CreateCertificateRequest) GetCertificate() *Certificate {
	if m != nil {
		return m.Certificate
	}
	return nil
}

type DeleteCertificateRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteCertificateRequest) Reset()         { *m = DeleteCertificateRequest{} }
func (m *DeleteCertificateRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteCertificateRequest) ProtoMessage()    {}
func (*DeleteCertificateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_certificate_53162111b312b807, []int{2}
}
func (m *DeleteCertificateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCertificateRequest.Unmarshal(m, b)
}
func (m *DeleteCertificateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCertificateRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteCertificateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCertificateRequest.Merge(dst, src)
}
func (m *DeleteCertificateRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteCertificateRequest.Size(m)
}
func (m *DeleteCertificateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteCertificateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteCertificateRequest proto.InternalMessageInfo

func (m *DeleteCertificateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetCertificateRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCertificateRequest) Reset()         { *m = GetCertificateRequest{} }
func (m *GetCertificateRequest) String() string { return proto.CompactTextString(m) }
func (*GetCertificateRequest) ProtoMessage()    {}
func (*GetCertificateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_certificate_53162111b312b807, []int{3}
}
func (m *GetCertificateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCertificateRequest.Unmarshal(m, b)
}
func (m *GetCertificateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCertificateRequest.Marshal(b, m, deterministic)
}
func (dst *GetCertificateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCertificateRequest.Merge(dst, src)
}
func (m *GetCertificateRequest) XXX_Size() int {
	return xxx_messageInfo_GetCertificateRequest.Size(m)
}
func (m *GetCertificateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCertificateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCertificateRequest proto.InternalMessageInfo

func (m *GetCertificateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ListCertificatesRequest struct {
	PageSize             int32    `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken            string   `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListCertificatesRequest) Reset()         { *m = ListCertificatesRequest{} }
func (m *ListCertificatesRequest) String() string { return proto.CompactTextString(m) }
func (*ListCertificatesRequest) ProtoMessage()    {}
func (*ListCertificatesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_certificate_53162111b312b807, []int{4}
}
func (m *ListCertificatesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCertificatesRequest.Unmarshal(m, b)
}
func (m *ListCertificatesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCertificatesRequest.Marshal(b, m, deterministic)
}
func (dst *ListCertificatesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCertificatesRequest.Merge(dst, src)
}
func (m *ListCertificatesRequest) XXX_Size() int {
	return xxx_messageInfo_ListCertificatesRequest.Size(m)
}
func (m *ListCertificatesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCertificatesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListCertificatesRequest proto.InternalMessageInfo

func (m *ListCertificatesRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListCertificatesRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

type ListCertificatesResponse struct {
	Certificates         []*Certificate `protobuf:"bytes,1,rep,name=certificates" json:"certificates,omitempty"`
	NextPageToken        string         `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ListCertificatesResponse) Reset()         { *m = ListCertificatesResponse{} }
func (m *ListCertificatesResponse) String() string { return proto.CompactTextString(m) }
func (*ListCertificatesResponse) ProtoMessage()    {}
func (*ListCertificatesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_certificate_53162111b312b807, []int{5}
}
func (m *ListCertificatesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCertificatesResponse.Unmarshal(m, b)
}
func (m *ListCertificatesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCertificatesResponse.Marshal(b, m, deterministic)
}
func (dst *ListCertificatesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCertificatesResponse.Merge(dst, src)
}
func (m *ListCertificatesResponse) XXX_Size() int {
	return xxx_messageInfo_ListCertificatesResponse.Size(m)
}
func (m *ListCertificatesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCertificatesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListCertificatesResponse proto.InternalMessageInfo

func (m *ListCertificatesResponse) GetCertificates() []*Certificate {
	if m != nil {
		return m.Certificates
	}
	return nil
}

func (m *ListCertificatesResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

type UpdateCertificateRequest struct {
	Certificate          *Certificate `protobuf:"bytes,1,opt,name=certificate" json:"certificate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UpdateCertificateRequest) Reset()         { *m = UpdateCertificateRequest{} }
func (m *UpdateCertificateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateCertificateRequest) ProtoMessage()    {}
func (*UpdateCertificateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_certificate_53162111b312b807, []int{6}
}
func (m *UpdateCertificateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateCertificateRequest.Unmarshal(m, b)
}
func (m *UpdateCertificateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateCertificateRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateCertificateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCertificateRequest.Merge(dst, src)
}
func (m *UpdateCertificateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateCertificateRequest.Size(m)
}
func (m *UpdateCertificateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCertificateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCertificateRequest proto.InternalMessageInfo

func (m *UpdateCertificateRequest) GetCertificate() *Certificate {
	if m != nil {
		return m.Certificate
	}
	return nil
}

func init() {
	proto.RegisterType((*Certificate)(nil), "powerssl.api.v1.Certificate")
	proto.RegisterMapType((map[string]string)(nil), "powerssl.api.v1.Certificate.LabelsEntry")
	proto.RegisterType((*CreateCertificateRequest)(nil), "powerssl.api.v1.CreateCertificateRequest")
	proto.RegisterType((*DeleteCertificateRequest)(nil), "powerssl.api.v1.DeleteCertificateRequest")
	proto.RegisterType((*GetCertificateRequest)(nil), "powerssl.api.v1.GetCertificateRequest")
	proto.RegisterType((*ListCertificatesRequest)(nil), "powerssl.api.v1.ListCertificatesRequest")
	proto.RegisterType((*ListCertificatesResponse)(nil), "powerssl.api.v1.ListCertificatesResponse")
	proto.RegisterType((*UpdateCertificateRequest)(nil), "powerssl.api.v1.UpdateCertificateRequest")
	proto.RegisterEnum("powerssl.api.v1.DigestAlgorithm", DigestAlgorithm_name, DigestAlgorithm_value)
	proto.RegisterEnum("powerssl.api.v1.KeyAlgorithm", KeyAlgorithm_name, KeyAlgorithm_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CertificateService service

type CertificateServiceClient interface {
	Create(ctx context.Context, in *CreateCertificateRequest, opts ...grpc.CallOption) (*Certificate, error)
	Delete(ctx context.Context, in *DeleteCertificateRequest, opts ...grpc.CallOption) (*types.Empty, error)
	Get(ctx context.Context, in *GetCertificateRequest, opts ...grpc.CallOption) (*Certificate, error)
	List(ctx context.Context, in *ListCertificatesRequest, opts ...grpc.CallOption) (*ListCertificatesResponse, error)
	Update(ctx context.Context, in *UpdateCertificateRequest, opts ...grpc.CallOption) (*Certificate, error)
}

type certificateServiceClient struct {
	cc *grpc.ClientConn
}

func NewCertificateServiceClient(cc *grpc.ClientConn) CertificateServiceClient {
	return &certificateServiceClient{cc}
}

func (c *certificateServiceClient) Create(ctx context.Context, in *CreateCertificateRequest, opts ...grpc.CallOption) (*Certificate, error) {
	out := new(Certificate)
	err := c.cc.Invoke(ctx, "/powerssl.api.v1.CertificateService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateServiceClient) Delete(ctx context.Context, in *DeleteCertificateRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/powerssl.api.v1.CertificateService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateServiceClient) Get(ctx context.Context, in *GetCertificateRequest, opts ...grpc.CallOption) (*Certificate, error) {
	out := new(Certificate)
	err := c.cc.Invoke(ctx, "/powerssl.api.v1.CertificateService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateServiceClient) List(ctx context.Context, in *ListCertificatesRequest, opts ...grpc.CallOption) (*ListCertificatesResponse, error) {
	out := new(ListCertificatesResponse)
	err := c.cc.Invoke(ctx, "/powerssl.api.v1.CertificateService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateServiceClient) Update(ctx context.Context, in *UpdateCertificateRequest, opts ...grpc.CallOption) (*Certificate, error) {
	out := new(Certificate)
	err := c.cc.Invoke(ctx, "/powerssl.api.v1.CertificateService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CertificateService service

type CertificateServiceServer interface {
	Create(context.Context, *CreateCertificateRequest) (*Certificate, error)
	Delete(context.Context, *DeleteCertificateRequest) (*types.Empty, error)
	Get(context.Context, *GetCertificateRequest) (*Certificate, error)
	List(context.Context, *ListCertificatesRequest) (*ListCertificatesResponse, error)
	Update(context.Context, *UpdateCertificateRequest) (*Certificate, error)
}

func RegisterCertificateServiceServer(s *grpc.Server, srv CertificateServiceServer) {
	s.RegisterService(&_CertificateService_serviceDesc, srv)
}

func _CertificateService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/powerssl.api.v1.CertificateService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateServiceServer).Create(ctx, req.(*CreateCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/powerssl.api.v1.CertificateService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateServiceServer).Delete(ctx, req.(*DeleteCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/powerssl.api.v1.CertificateService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateServiceServer).Get(ctx, req.(*GetCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCertificatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/powerssl.api.v1.CertificateService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateServiceServer).List(ctx, req.(*ListCertificatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/powerssl.api.v1.CertificateService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateServiceServer).Update(ctx, req.(*UpdateCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CertificateService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "powerssl.api.v1.CertificateService",
	HandlerType: (*CertificateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _CertificateService_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CertificateService_Delete_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _CertificateService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _CertificateService_List_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CertificateService_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/certificate.proto",
}

func init() {
	proto.RegisterFile("api/v1/certificate.proto", fileDescriptor_certificate_53162111b312b807)
}

var fileDescriptor_certificate_53162111b312b807 = []byte{
	// 827 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x2e, 0x2d, 0x59, 0x96, 0x86, 0x4a, 0x2d, 0x2c, 0xfa, 0xb3, 0xa6, 0x6d, 0x84, 0x21, 0xd0,
	0x80, 0x71, 0x01, 0x0a, 0x56, 0xd1, 0xa0, 0x4d, 0xd0, 0x22, 0x8a, 0xad, 0xda, 0x82, 0xdd, 0x34,
	0xa0, 0xec, 0x43, 0x7d, 0x11, 0x68, 0x69, 0xa2, 0x2e, 0x44, 0x91, 0x2c, 0x77, 0xa5, 0x44, 0x29,
	0x7a, 0x68, 0xd1, 0x37, 0xe8, 0xa3, 0xf5, 0xd6, 0x73, 0x2f, 0x7d, 0x8b, 0x62, 0x97, 0x94, 0x4c,
	0x8b, 0x92, 0xec, 0x43, 0x6f, 0xdc, 0xd9, 0x6f, 0xe6, 0x9b, 0x9f, 0x6f, 0x96, 0x40, 0xbd, 0x88,
	0xd5, 0x27, 0x87, 0xf5, 0x1e, 0xc6, 0x82, 0xbd, 0x61, 0x3d, 0x4f, 0xa0, 0x13, 0xc5, 0xa1, 0x08,
	0xc9, 0x76, 0x14, 0xbe, 0xc5, 0x98, 0x73, 0xdf, 0xf1, 0x22, 0xe6, 0x4c, 0x0e, 0x8d, 0xbd, 0x41,
	0x18, 0x0e, 0x7c, 0xac, 0x4b, 0x0f, 0x2f, 0x08, 0x42, 0xe1, 0x09, 0x16, 0x06, 0x3c, 0x81, 0x1b,
	0xbb, 0xe9, 0xad, 0x3a, 0x5d, 0x8f, 0xdf, 0xd4, 0x71, 0x14, 0x89, 0x69, 0x7a, 0xf9, 0x70, 0xf1,
	0x52, 0xb0, 0x11, 0x72, 0xe1, 0x8d, 0xa2, 0x04, 0x60, 0xfd, 0x5d, 0x04, 0xfd, 0xe8, 0x26, 0x05,
	0x42, 0xa0, 0x18, 0x78, 0x23, 0xa4, 0x9a, 0xa9, 0xd9, 0x15, 0x57, 0x7d, 0x93, 0xe7, 0xa0, 0xf7,
	0x62, 0xf4, 0x04, 0x76, 0xa5, 0x37, 0xdd, 0x30, 0x35, 0x5b, 0x6f, 0x18, 0x4e, 0x12, 0xda, 0x99,
	0x85, 0x76, 0x2e, 0x66, 0xa1, 0x5d, 0x48, 0xe0, 0xd2, 0x20, 0x9d, 0xc7, 0x51, 0x7f, 0xee, 0x5c,
	0xb8, 0xdb, 0x39, 0x81, 0x2b, 0xe7, 0x47, 0x50, 0xed, 0x33, 0x1e, 0xf9, 0xde, 0xb4, 0xab, 0xb2,
	0x2a, 0xaa, 0xac, 0xf4, 0xd4, 0xf6, 0x4a, 0x26, 0xf7, 0x11, 0x6c, 0x0a, 0x26, 0x7c, 0xa4, 0x9b,
	0xea, 0x2e, 0x39, 0x10, 0x13, 0xf4, 0x3e, 0xf2, 0x5e, 0xcc, 0x22, 0xd9, 0x2a, 0x5a, 0x4a, 0xfd,
	0x6e, 0x4c, 0xe4, 0x05, 0x94, 0x7c, 0xef, 0x1a, 0x7d, 0x4e, 0xb7, 0xcc, 0x82, 0xad, 0x37, 0x6c,
	0x67, 0xa1, 0xed, 0x4e, 0xa6, 0x2d, 0xce, 0xb9, 0x82, 0xb6, 0x02, 0x11, 0x4f, 0xdd, 0xd4, 0x8f,
	0x18, 0x50, 0xee, 0x07, 0x5c, 0xe6, 0xc5, 0x69, 0xd9, 0x2c, 0xd8, 0x15, 0x77, 0x7e, 0x26, 0x2f,
	0xe1, 0xc1, 0x10, 0xa7, 0x5d, 0xcf, 0x1f, 0x84, 0x31, 0x13, 0x3f, 0x8d, 0x68, 0xc5, 0xd4, 0xec,
	0x0f, 0x1b, 0xfb, 0x39, 0x92, 0x33, 0x9c, 0x36, 0x67, 0x20, 0xb7, 0x3a, 0xcc, 0x9c, 0xc8, 0x0e,
	0x94, 0x65, 0x0c, 0xce, 0xde, 0x23, 0x05, 0x53, 0xb3, 0x37, 0xdd, 0xad, 0x21, 0x4e, 0x3b, 0xec,
	0x3d, 0x92, 0x33, 0xa8, 0xf5, 0xd9, 0x00, 0xb9, 0xc8, 0x30, 0xe8, 0x8a, 0xc1, 0xcc, 0x31, 0x1c,
	0x2b, 0xe0, 0x0d, 0xc9, 0x76, 0xff, 0xb6, 0x81, 0xec, 0x03, 0x78, 0x63, 0x11, 0x76, 0x63, 0x0c,
	0xf0, 0x2d, 0xad, 0x9a, 0x9a, 0x5d, 0x76, 0x2b, 0xd2, 0xe2, 0x4a, 0x83, 0xf1, 0x35, 0xe8, 0x99,
	0xea, 0x49, 0x0d, 0x0a, 0x43, 0x9c, 0xa6, 0xfa, 0x90, 0x9f, 0x72, 0x02, 0x13, 0xcf, 0x1f, 0x27,
	0xc2, 0xa8, 0xb8, 0xc9, 0xe1, 0xd9, 0xc6, 0x57, 0x9a, 0x75, 0x05, 0xf4, 0x48, 0x29, 0x21, 0xd3,
	0x4a, 0x17, 0x7f, 0x1e, 0x23, 0x17, 0xe4, 0x5b, 0xd0, 0x33, 0xd2, 0x57, 0xf1, 0xf4, 0xc6, 0xde,
	0xba, 0x21, 0xb8, 0x59, 0x07, 0xcb, 0x01, 0x7a, 0x8c, 0x3e, 0x2e, 0x8d, 0xbd, 0x44, 0xc4, 0xd6,
	0xe7, 0xf0, 0xf1, 0x09, 0x8a, 0x7b, 0x82, 0x2f, 0xe1, 0xd3, 0x73, 0xc6, 0xb3, 0x68, 0x3e, 0x83,
	0xef, 0x42, 0x25, 0xf2, 0x06, 0x98, 0x8c, 0x45, 0x53, 0x63, 0x29, 0x4b, 0x83, 0x9a, 0xcb, 0x3e,
	0x80, 0xba, 0x14, 0xe1, 0x10, 0x83, 0xb4, 0x1f, 0x0a, 0x7e, 0x21, 0x0d, 0xd6, 0x1f, 0x1a, 0xd0,
	0x7c, 0x5c, 0x1e, 0x85, 0x01, 0x47, 0xf2, 0x02, 0xaa, 0x99, 0xfa, 0x38, 0xd5, 0x94, 0x2c, 0xd7,
	0x77, 0xe4, 0x96, 0x07, 0x79, 0x0c, 0xdb, 0x01, 0xbe, 0x13, 0xdd, 0x5c, 0x0a, 0x0f, 0xa4, 0xf9,
	0xf5, 0x3c, 0x8d, 0x2b, 0xa0, 0x97, 0x6a, 0xc7, 0xfe, 0xff, 0xb1, 0x1c, 0xb4, 0x61, 0x7b, 0x41,
	0x70, 0xc4, 0x84, 0xbd, 0xe3, 0xf6, 0x49, 0xab, 0x73, 0xd1, 0x6d, 0x9e, 0x9f, 0xfc, 0xe0, 0xb6,
	0x2f, 0x4e, 0xbf, 0xef, 0x5e, 0xbe, 0xea, 0xbc, 0x6e, 0x1d, 0xb5, 0xbf, 0x6b, 0xb7, 0x8e, 0x6b,
	0x1f, 0x90, 0x32, 0x14, 0x3b, 0xa7, 0xcd, 0xc3, 0x9a, 0x46, 0x00, 0x4a, 0x9d, 0xd3, 0x66, 0xe3,
	0xcb, 0xa7, 0xb5, 0x8d, 0x83, 0xa7, 0x50, 0xcd, 0x6e, 0x07, 0xd9, 0x87, 0x9d, 0xb3, 0xd6, 0x8f,
	0x2b, 0x83, 0x6c, 0x41, 0xc1, 0xed, 0x34, 0x6b, 0x5a, 0xe3, 0xdf, 0x22, 0x90, 0x4c, 0x7e, 0x1d,
	0x8c, 0x27, 0xac, 0x87, 0xe4, 0x1d, 0x94, 0x12, 0x31, 0x92, 0x27, 0xf9, 0x72, 0x56, 0xa8, 0xd4,
	0x58, 0x5b, 0xb9, 0xf5, 0xd9, 0xef, 0x7f, 0xfd, 0xf3, 0xe7, 0xc6, 0x43, 0xab, 0xb6, 0xf0, 0x90,
	0xf3, 0x67, 0xd9, 0x9e, 0x90, 0x00, 0x4a, 0x89, 0x54, 0x97, 0x30, 0xaf, 0xd2, 0xb0, 0xf1, 0x49,
	0xee, 0x89, 0x6c, 0xc9, 0x77, 0xdd, 0x7a, 0xa4, 0x38, 0x77, 0x0f, 0x76, 0x24, 0xe7, 0x2f, 0x52,
	0xad, 0xdf, 0x64, 0x99, 0xeb, 0x07, 0xbf, 0x92, 0x00, 0x0a, 0x27, 0x28, 0xc8, 0xe3, 0x1c, 0xd9,
	0xd2, 0x05, 0xb8, 0xa3, 0xc6, 0x94, 0x8f, 0xac, 0xe1, 0x1b, 0x43, 0x51, 0xaa, 0x9a, 0xe4, 0x9f,
	0xd0, 0x15, 0x4b, 0x64, 0x3c, 0xb9, 0x07, 0x32, 0x59, 0x0b, 0x8b, 0x2a, 0x7e, 0x42, 0x72, 0x3d,
	0x26, 0xbf, 0x69, 0x50, 0x4a, 0x74, 0xbc, 0xa4, 0xaf, 0xab, 0x04, 0x7e, 0x47, 0xb5, 0x8e, 0x62,
	0xb3, 0x8d, 0xd5, 0xd5, 0xde, 0x1a, 0xed, 0xcb, 0xbd, 0x2b, 0x63, 0x1e, 0x8e, 0x85, 0xf5, 0x68,
	0x38, 0xa8, 0x27, 0xff, 0xf5, 0xe7, 0x5e, 0xc4, 0xae, 0x4b, 0x6a, 0x76, 0x5f, 0xfc, 0x17, 0x00,
	0x00, 0xff, 0xff, 0x87, 0x26, 0x11, 0xa7, 0xec, 0x07, 0x00, 0x00,
}
