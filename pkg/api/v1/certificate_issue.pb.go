// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/v1/certificate_issue.proto

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

type CreateCertificateIssueRequest struct {
	CertificateIssue     *CertificateIssue `protobuf:"bytes,1,opt,name=certificate_issue,json=certificateIssue" json:"certificate_issue,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CreateCertificateIssueRequest) Reset()         { *m = CreateCertificateIssueRequest{} }
func (m *CreateCertificateIssueRequest) String() string { return proto.CompactTextString(m) }
func (*CreateCertificateIssueRequest) ProtoMessage()    {}
func (*CreateCertificateIssueRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_certificate_issue_df883b84c2080a5b, []int{0}
}
func (m *CreateCertificateIssueRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCertificateIssueRequest.Unmarshal(m, b)
}
func (m *CreateCertificateIssueRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCertificateIssueRequest.Marshal(b, m, deterministic)
}
func (dst *CreateCertificateIssueRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCertificateIssueRequest.Merge(dst, src)
}
func (m *CreateCertificateIssueRequest) XXX_Size() int {
	return xxx_messageInfo_CreateCertificateIssueRequest.Size(m)
}
func (m *CreateCertificateIssueRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCertificateIssueRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCertificateIssueRequest proto.InternalMessageInfo

func (m *CreateCertificateIssueRequest) GetCertificateIssue() *CertificateIssue {
	if m != nil {
		return m.CertificateIssue
	}
	return nil
}

type DeleteCertificateIssueRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteCertificateIssueRequest) Reset()         { *m = DeleteCertificateIssueRequest{} }
func (m *DeleteCertificateIssueRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteCertificateIssueRequest) ProtoMessage()    {}
func (*DeleteCertificateIssueRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_certificate_issue_df883b84c2080a5b, []int{1}
}
func (m *DeleteCertificateIssueRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCertificateIssueRequest.Unmarshal(m, b)
}
func (m *DeleteCertificateIssueRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCertificateIssueRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteCertificateIssueRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCertificateIssueRequest.Merge(dst, src)
}
func (m *DeleteCertificateIssueRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteCertificateIssueRequest.Size(m)
}
func (m *DeleteCertificateIssueRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteCertificateIssueRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteCertificateIssueRequest proto.InternalMessageInfo

func (m *DeleteCertificateIssueRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetCertificateIssueRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCertificateIssueRequest) Reset()         { *m = GetCertificateIssueRequest{} }
func (m *GetCertificateIssueRequest) String() string { return proto.CompactTextString(m) }
func (*GetCertificateIssueRequest) ProtoMessage()    {}
func (*GetCertificateIssueRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_certificate_issue_df883b84c2080a5b, []int{2}
}
func (m *GetCertificateIssueRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCertificateIssueRequest.Unmarshal(m, b)
}
func (m *GetCertificateIssueRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCertificateIssueRequest.Marshal(b, m, deterministic)
}
func (dst *GetCertificateIssueRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCertificateIssueRequest.Merge(dst, src)
}
func (m *GetCertificateIssueRequest) XXX_Size() int {
	return xxx_messageInfo_GetCertificateIssueRequest.Size(m)
}
func (m *GetCertificateIssueRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCertificateIssueRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCertificateIssueRequest proto.InternalMessageInfo

func (m *GetCertificateIssueRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ListCertificateIssuesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListCertificateIssuesRequest) Reset()         { *m = ListCertificateIssuesRequest{} }
func (m *ListCertificateIssuesRequest) String() string { return proto.CompactTextString(m) }
func (*ListCertificateIssuesRequest) ProtoMessage()    {}
func (*ListCertificateIssuesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_certificate_issue_df883b84c2080a5b, []int{3}
}
func (m *ListCertificateIssuesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCertificateIssuesRequest.Unmarshal(m, b)
}
func (m *ListCertificateIssuesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCertificateIssuesRequest.Marshal(b, m, deterministic)
}
func (dst *ListCertificateIssuesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCertificateIssuesRequest.Merge(dst, src)
}
func (m *ListCertificateIssuesRequest) XXX_Size() int {
	return xxx_messageInfo_ListCertificateIssuesRequest.Size(m)
}
func (m *ListCertificateIssuesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCertificateIssuesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListCertificateIssuesRequest proto.InternalMessageInfo

type ListCertificateIssuesResponse struct {
	TypeMeta             *TypeMeta           `protobuf:"bytes,1,opt,name=type_meta,json=typeMeta" json:"type_meta,omitempty"`
	ListMeta             *ListMeta           `protobuf:"bytes,2,opt,name=list_meta,json=listMeta" json:"list_meta,omitempty"`
	Items                []*CertificateIssue `protobuf:"bytes,3,rep,name=items" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ListCertificateIssuesResponse) Reset()         { *m = ListCertificateIssuesResponse{} }
func (m *ListCertificateIssuesResponse) String() string { return proto.CompactTextString(m) }
func (*ListCertificateIssuesResponse) ProtoMessage()    {}
func (*ListCertificateIssuesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_certificate_issue_df883b84c2080a5b, []int{4}
}
func (m *ListCertificateIssuesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCertificateIssuesResponse.Unmarshal(m, b)
}
func (m *ListCertificateIssuesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCertificateIssuesResponse.Marshal(b, m, deterministic)
}
func (dst *ListCertificateIssuesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCertificateIssuesResponse.Merge(dst, src)
}
func (m *ListCertificateIssuesResponse) XXX_Size() int {
	return xxx_messageInfo_ListCertificateIssuesResponse.Size(m)
}
func (m *ListCertificateIssuesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCertificateIssuesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListCertificateIssuesResponse proto.InternalMessageInfo

func (m *ListCertificateIssuesResponse) GetTypeMeta() *TypeMeta {
	if m != nil {
		return m.TypeMeta
	}
	return nil
}

func (m *ListCertificateIssuesResponse) GetListMeta() *ListMeta {
	if m != nil {
		return m.ListMeta
	}
	return nil
}

func (m *ListCertificateIssuesResponse) GetItems() []*CertificateIssue {
	if m != nil {
		return m.Items
	}
	return nil
}

type UpdateCertificateIssueRequest struct {
	CertificateIssue     *CertificateIssue `protobuf:"bytes,1,opt,name=certificate_issue,json=certificateIssue" json:"certificate_issue,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *UpdateCertificateIssueRequest) Reset()         { *m = UpdateCertificateIssueRequest{} }
func (m *UpdateCertificateIssueRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateCertificateIssueRequest) ProtoMessage()    {}
func (*UpdateCertificateIssueRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_certificate_issue_df883b84c2080a5b, []int{5}
}
func (m *UpdateCertificateIssueRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateCertificateIssueRequest.Unmarshal(m, b)
}
func (m *UpdateCertificateIssueRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateCertificateIssueRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateCertificateIssueRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCertificateIssueRequest.Merge(dst, src)
}
func (m *UpdateCertificateIssueRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateCertificateIssueRequest.Size(m)
}
func (m *UpdateCertificateIssueRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCertificateIssueRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCertificateIssueRequest proto.InternalMessageInfo

func (m *UpdateCertificateIssueRequest) GetCertificateIssue() *CertificateIssue {
	if m != nil {
		return m.CertificateIssue
	}
	return nil
}

type CertificateIssue struct {
	TypeMeta             *TypeMeta               `protobuf:"bytes,1,opt,name=type_meta,json=typeMeta" json:"type_meta,omitempty"`
	ObjectMeta           *ObjectMeta             `protobuf:"bytes,2,opt,name=object_meta,json=objectMeta" json:"object_meta,omitempty"`
	Spec                 *CertificateIssueSpec   `protobuf:"bytes,3,opt,name=spec" json:"spec,omitempty"`
	Status               *CertificateIssueStatus `protobuf:"bytes,4,opt,name=status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *CertificateIssue) Reset()         { *m = CertificateIssue{} }
func (m *CertificateIssue) String() string { return proto.CompactTextString(m) }
func (*CertificateIssue) ProtoMessage()    {}
func (*CertificateIssue) Descriptor() ([]byte, []int) {
	return fileDescriptor_certificate_issue_df883b84c2080a5b, []int{6}
}
func (m *CertificateIssue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CertificateIssue.Unmarshal(m, b)
}
func (m *CertificateIssue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CertificateIssue.Marshal(b, m, deterministic)
}
func (dst *CertificateIssue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateIssue.Merge(dst, src)
}
func (m *CertificateIssue) XXX_Size() int {
	return xxx_messageInfo_CertificateIssue.Size(m)
}
func (m *CertificateIssue) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateIssue.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateIssue proto.InternalMessageInfo

func (m *CertificateIssue) GetTypeMeta() *TypeMeta {
	if m != nil {
		return m.TypeMeta
	}
	return nil
}

func (m *CertificateIssue) GetObjectMeta() *ObjectMeta {
	if m != nil {
		return m.ObjectMeta
	}
	return nil
}

func (m *CertificateIssue) GetSpec() *CertificateIssueSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *CertificateIssue) GetStatus() *CertificateIssueStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type CertificateIssueSpec struct {
	CommonName           string   `protobuf:"bytes,1,opt,name=common_name,json=commonName,proto3" json:"common_name,omitempty"`
	EncryptionAlgorithm  string   `protobuf:"bytes,2,opt,name=encryption_algorithm,json=encryptionAlgorithm,proto3" json:"encryption_algorithm,omitempty"`
	KeySize              int32    `protobuf:"varint,3,opt,name=key_size,json=keySize,proto3" json:"key_size,omitempty"`
	SignatureAlgorithm   string   `protobuf:"bytes,4,opt,name=signature_algorithm,json=signatureAlgorithm,proto3" json:"signature_algorithm,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CertificateIssueSpec) Reset()         { *m = CertificateIssueSpec{} }
func (m *CertificateIssueSpec) String() string { return proto.CompactTextString(m) }
func (*CertificateIssueSpec) ProtoMessage()    {}
func (*CertificateIssueSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_certificate_issue_df883b84c2080a5b, []int{7}
}
func (m *CertificateIssueSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CertificateIssueSpec.Unmarshal(m, b)
}
func (m *CertificateIssueSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CertificateIssueSpec.Marshal(b, m, deterministic)
}
func (dst *CertificateIssueSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateIssueSpec.Merge(dst, src)
}
func (m *CertificateIssueSpec) XXX_Size() int {
	return xxx_messageInfo_CertificateIssueSpec.Size(m)
}
func (m *CertificateIssueSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateIssueSpec.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateIssueSpec proto.InternalMessageInfo

func (m *CertificateIssueSpec) GetCommonName() string {
	if m != nil {
		return m.CommonName
	}
	return ""
}

func (m *CertificateIssueSpec) GetEncryptionAlgorithm() string {
	if m != nil {
		return m.EncryptionAlgorithm
	}
	return ""
}

func (m *CertificateIssueSpec) GetKeySize() int32 {
	if m != nil {
		return m.KeySize
	}
	return 0
}

func (m *CertificateIssueSpec) GetSignatureAlgorithm() string {
	if m != nil {
		return m.SignatureAlgorithm
	}
	return ""
}

type CertificateIssueStatus struct {
	Phase                string   `protobuf:"bytes,1,opt,name=phase,proto3" json:"phase,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CertificateIssueStatus) Reset()         { *m = CertificateIssueStatus{} }
func (m *CertificateIssueStatus) String() string { return proto.CompactTextString(m) }
func (*CertificateIssueStatus) ProtoMessage()    {}
func (*CertificateIssueStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_certificate_issue_df883b84c2080a5b, []int{8}
}
func (m *CertificateIssueStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CertificateIssueStatus.Unmarshal(m, b)
}
func (m *CertificateIssueStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CertificateIssueStatus.Marshal(b, m, deterministic)
}
func (dst *CertificateIssueStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateIssueStatus.Merge(dst, src)
}
func (m *CertificateIssueStatus) XXX_Size() int {
	return xxx_messageInfo_CertificateIssueStatus.Size(m)
}
func (m *CertificateIssueStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateIssueStatus.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateIssueStatus proto.InternalMessageInfo

func (m *CertificateIssueStatus) GetPhase() string {
	if m != nil {
		return m.Phase
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateCertificateIssueRequest)(nil), "powerssl.api.v1.CreateCertificateIssueRequest")
	proto.RegisterType((*DeleteCertificateIssueRequest)(nil), "powerssl.api.v1.DeleteCertificateIssueRequest")
	proto.RegisterType((*GetCertificateIssueRequest)(nil), "powerssl.api.v1.GetCertificateIssueRequest")
	proto.RegisterType((*ListCertificateIssuesRequest)(nil), "powerssl.api.v1.ListCertificateIssuesRequest")
	proto.RegisterType((*ListCertificateIssuesResponse)(nil), "powerssl.api.v1.ListCertificateIssuesResponse")
	proto.RegisterType((*UpdateCertificateIssueRequest)(nil), "powerssl.api.v1.UpdateCertificateIssueRequest")
	proto.RegisterType((*CertificateIssue)(nil), "powerssl.api.v1.CertificateIssue")
	proto.RegisterType((*CertificateIssueSpec)(nil), "powerssl.api.v1.CertificateIssueSpec")
	proto.RegisterType((*CertificateIssueStatus)(nil), "powerssl.api.v1.CertificateIssueStatus")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CertificateIssueService service

type CertificateIssueServiceClient interface {
	Create(ctx context.Context, in *CreateCertificateIssueRequest, opts ...grpc.CallOption) (*CertificateIssue, error)
	Delete(ctx context.Context, in *DeleteCertificateIssueRequest, opts ...grpc.CallOption) (*types.Empty, error)
	Get(ctx context.Context, in *GetCertificateIssueRequest, opts ...grpc.CallOption) (*CertificateIssue, error)
	List(ctx context.Context, in *ListCertificateIssuesRequest, opts ...grpc.CallOption) (*ListCertificateIssuesResponse, error)
	Update(ctx context.Context, in *UpdateCertificateIssueRequest, opts ...grpc.CallOption) (*CertificateIssue, error)
}

type certificateIssueServiceClient struct {
	cc *grpc.ClientConn
}

func NewCertificateIssueServiceClient(cc *grpc.ClientConn) CertificateIssueServiceClient {
	return &certificateIssueServiceClient{cc}
}

func (c *certificateIssueServiceClient) Create(ctx context.Context, in *CreateCertificateIssueRequest, opts ...grpc.CallOption) (*CertificateIssue, error) {
	out := new(CertificateIssue)
	err := c.cc.Invoke(ctx, "/powerssl.api.v1.CertificateIssueService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateIssueServiceClient) Delete(ctx context.Context, in *DeleteCertificateIssueRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/powerssl.api.v1.CertificateIssueService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateIssueServiceClient) Get(ctx context.Context, in *GetCertificateIssueRequest, opts ...grpc.CallOption) (*CertificateIssue, error) {
	out := new(CertificateIssue)
	err := c.cc.Invoke(ctx, "/powerssl.api.v1.CertificateIssueService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateIssueServiceClient) List(ctx context.Context, in *ListCertificateIssuesRequest, opts ...grpc.CallOption) (*ListCertificateIssuesResponse, error) {
	out := new(ListCertificateIssuesResponse)
	err := c.cc.Invoke(ctx, "/powerssl.api.v1.CertificateIssueService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateIssueServiceClient) Update(ctx context.Context, in *UpdateCertificateIssueRequest, opts ...grpc.CallOption) (*CertificateIssue, error) {
	out := new(CertificateIssue)
	err := c.cc.Invoke(ctx, "/powerssl.api.v1.CertificateIssueService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CertificateIssueService service

type CertificateIssueServiceServer interface {
	Create(context.Context, *CreateCertificateIssueRequest) (*CertificateIssue, error)
	Delete(context.Context, *DeleteCertificateIssueRequest) (*types.Empty, error)
	Get(context.Context, *GetCertificateIssueRequest) (*CertificateIssue, error)
	List(context.Context, *ListCertificateIssuesRequest) (*ListCertificateIssuesResponse, error)
	Update(context.Context, *UpdateCertificateIssueRequest) (*CertificateIssue, error)
}

func RegisterCertificateIssueServiceServer(s *grpc.Server, srv CertificateIssueServiceServer) {
	s.RegisterService(&_CertificateIssueService_serviceDesc, srv)
}

func _CertificateIssueService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCertificateIssueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateIssueServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/powerssl.api.v1.CertificateIssueService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateIssueServiceServer).Create(ctx, req.(*CreateCertificateIssueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateIssueService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCertificateIssueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateIssueServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/powerssl.api.v1.CertificateIssueService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateIssueServiceServer).Delete(ctx, req.(*DeleteCertificateIssueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateIssueService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCertificateIssueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateIssueServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/powerssl.api.v1.CertificateIssueService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateIssueServiceServer).Get(ctx, req.(*GetCertificateIssueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateIssueService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCertificateIssuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateIssueServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/powerssl.api.v1.CertificateIssueService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateIssueServiceServer).List(ctx, req.(*ListCertificateIssuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateIssueService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCertificateIssueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateIssueServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/powerssl.api.v1.CertificateIssueService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateIssueServiceServer).Update(ctx, req.(*UpdateCertificateIssueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CertificateIssueService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "powerssl.api.v1.CertificateIssueService",
	HandlerType: (*CertificateIssueServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _CertificateIssueService_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CertificateIssueService_Delete_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _CertificateIssueService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _CertificateIssueService_List_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CertificateIssueService_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/certificate_issue.proto",
}

func init() {
	proto.RegisterFile("api/v1/certificate_issue.proto", fileDescriptor_certificate_issue_df883b84c2080a5b)
}

var fileDescriptor_certificate_issue_df883b84c2080a5b = []byte{
	// 581 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0x6d, 0x9a, 0x9f, 0xaf, 0x99, 0x5c, 0x7c, 0xed, 0x36, 0x2a, 0xa9, 0xdb, 0x94, 0x62, 0x09,
	0x81, 0x84, 0xb0, 0x49, 0x2b, 0x81, 0x10, 0x48, 0x08, 0x0a, 0xaa, 0x90, 0xa0, 0x48, 0x2e, 0x08,
	0x09, 0x09, 0x45, 0x1b, 0x33, 0x75, 0x97, 0xd8, 0xde, 0xc5, 0xbb, 0x09, 0x72, 0xaf, 0x79, 0x26,
	0x9e, 0x81, 0x97, 0xe1, 0x1d, 0x90, 0x77, 0x9d, 0xb4, 0xc4, 0xf9, 0x83, 0x0b, 0xee, 0x26, 0x3b,
	0xe7, 0x9c, 0xf1, 0x1c, 0x7b, 0x4f, 0x60, 0x8f, 0x0a, 0xe6, 0x0e, 0x3b, 0xae, 0x8f, 0x89, 0x62,
	0x67, 0xcc, 0xa7, 0x0a, 0xbb, 0x4c, 0xca, 0x01, 0x3a, 0x22, 0xe1, 0x8a, 0x93, 0xff, 0x05, 0xff,
	0x8a, 0x89, 0x94, 0xa1, 0x43, 0x05, 0x73, 0x86, 0x1d, 0x6b, 0x27, 0xe0, 0x3c, 0x08, 0xd1, 0xd5,
	0xed, 0xde, 0xe0, 0xcc, 0xc5, 0x48, 0xa8, 0xd4, 0xa0, 0xad, 0x8d, 0x5c, 0x2d, 0x42, 0x45, 0xcd,
	0x91, 0xcd, 0xa1, 0x7d, 0x94, 0x20, 0x55, 0x78, 0x74, 0x39, 0xe1, 0x65, 0x36, 0xc0, 0xc3, 0x2f,
	0x03, 0x94, 0x8a, 0x9c, 0xc0, 0x46, 0x61, 0x78, 0xab, 0xb4, 0x5f, 0xba, 0xdd, 0x38, 0xb8, 0xe1,
	0x4c, 0x4c, 0x77, 0x0a, 0x22, 0xeb, 0xfe, 0xc4, 0x89, 0x7d, 0x08, 0xed, 0xe7, 0x18, 0xe2, 0xec,
	0x81, 0x04, 0x2a, 0x31, 0x8d, 0xcc, 0x8c, 0xba, 0xa7, 0x6b, 0xfb, 0x1e, 0x58, 0xc7, 0xa8, 0xfe,
	0x84, 0xb1, 0x07, 0xbb, 0xaf, 0x98, 0x2c, 0x50, 0x64, 0xce, 0xb1, 0x7f, 0x94, 0xa0, 0x3d, 0x03,
	0x20, 0x05, 0x8f, 0x25, 0x92, 0xfb, 0x50, 0x57, 0xa9, 0xc0, 0x6e, 0x66, 0x56, 0xbe, 0xf0, 0x76,
	0x61, 0xe1, 0xb7, 0xa9, 0xc0, 0xd7, 0xa8, 0xa8, 0xb7, 0xa6, 0xf2, 0x2a, 0xe3, 0x85, 0x4c, 0x2a,
	0xc3, 0x5b, 0x9d, 0xc1, 0xcb, 0x46, 0x1b, 0x5e, 0x98, 0x57, 0xe4, 0x01, 0x54, 0x99, 0xc2, 0x48,
	0xb6, 0xca, 0xfb, 0xe5, 0xe5, 0xcc, 0x35, 0xf8, 0xec, 0x15, 0xbe, 0x13, 0x9f, 0xfe, 0xe1, 0x2b,
	0xfc, 0xb6, 0x0a, 0xeb, 0x93, 0xb0, 0xbf, 0xb6, 0xeb, 0x31, 0x34, 0x78, 0xef, 0x33, 0xfa, 0xbf,
	0x19, 0xb6, 0x53, 0x60, 0xbe, 0xd1, 0x18, 0xcd, 0x05, 0x3e, 0xae, 0xc9, 0x43, 0xa8, 0x48, 0x81,
	0x7e, 0xab, 0xac, 0x69, 0x37, 0x17, 0x6e, 0x73, 0x2a, 0xd0, 0xf7, 0x34, 0x85, 0x3c, 0x81, 0x9a,
	0x54, 0x54, 0x0d, 0x64, 0xab, 0xa2, 0xc9, 0xb7, 0x16, 0x93, 0x35, 0xdc, 0xcb, 0x69, 0xf6, 0xf7,
	0x12, 0x34, 0xa7, 0xe9, 0x93, 0xeb, 0xd0, 0xf0, 0x79, 0x14, 0xf1, 0xb8, 0x7b, 0xe5, 0xb3, 0x04,
	0x73, 0x74, 0x42, 0x23, 0x24, 0x1d, 0x68, 0x62, 0xec, 0x27, 0xa9, 0x50, 0x8c, 0xc7, 0x5d, 0x1a,
	0x06, 0x3c, 0x61, 0xea, 0x3c, 0xd2, 0xcb, 0xd7, 0xbd, 0xcd, 0xcb, 0xde, 0xd3, 0x51, 0x8b, 0x6c,
	0xc3, 0x5a, 0x1f, 0xd3, 0xae, 0x64, 0x17, 0xa8, 0x97, 0xad, 0x7a, 0xff, 0xf5, 0x31, 0x3d, 0x65,
	0x17, 0x48, 0x5c, 0xd8, 0x94, 0x2c, 0x88, 0xa9, 0x1a, 0x24, 0x78, 0x45, 0xac, 0xa2, 0xc5, 0xc8,
	0xb8, 0x35, 0xd6, 0xb2, 0x1d, 0xd8, 0x9a, 0xbe, 0x1a, 0x69, 0x42, 0x55, 0x9c, 0x53, 0x39, 0x7a,
	0x66, 0xf3, 0xe3, 0xe0, 0x67, 0x19, 0xae, 0x15, 0x08, 0x98, 0x0c, 0x99, 0x8f, 0xe4, 0x23, 0xd4,
	0x4c, 0x7e, 0x10, 0xa7, 0xe8, 0xdf, 0xbc, 0x60, 0xb1, 0x16, 0x7f, 0x7a, 0xf6, 0x0a, 0xf1, 0xa0,
	0x66, 0xd2, 0x62, 0x8a, 0xfc, 0xdc, 0x18, 0xb1, 0xb6, 0x1c, 0x93, 0x84, 0xce, 0x28, 0x09, 0x9d,
	0x17, 0x59, 0x12, 0xda, 0x2b, 0xe4, 0x3d, 0x94, 0x8f, 0x51, 0x91, 0x3b, 0x05, 0xc1, 0xd9, 0x11,
	0xb3, 0xdc, 0xc3, 0x06, 0x50, 0xc9, 0xee, 0x35, 0xb9, 0x3b, 0xf5, 0xba, 0xcf, 0x8a, 0x22, 0xcb,
	0x59, 0x16, 0x6e, 0x82, 0xc9, 0x5e, 0xc9, 0x4c, 0x37, 0x37, 0x7e, 0x8a, 0x2b, 0x73, 0xa3, 0x60,
	0xa9, 0x3d, 0x9e, 0xed, 0x7e, 0xb0, 0xc6, 0x28, 0xc6, 0x5d, 0xd1, 0x0f, 0x5c, 0xf3, 0xc7, 0xf1,
	0x88, 0x0a, 0xd6, 0xab, 0x69, 0x43, 0x0f, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x63, 0xd2, 0xb4,
	0x54, 0x9b, 0x06, 0x00, 0x00,
}
