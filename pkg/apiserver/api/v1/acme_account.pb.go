// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: apiserver/api/v1/acme_account.proto

package api // import "powerssl.io/pkg/apiserver/api/v1"

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

type ACMEAccount struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CreateTime           *types.Timestamp  `protobuf:"bytes,2,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
	UpdateTime           *types.Timestamp  `protobuf:"bytes,3,opt,name=update_time,json=updateTime" json:"update_time,omitempty"`
	DisplayName          string            `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Title                string            `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	Description          string            `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	Labels               map[string]string `protobuf:"bytes,7,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	TermsOfServiceAgreed bool              `protobuf:"varint,8,opt,name=terms_of_service_agreed,json=termsOfServiceAgreed,proto3" json:"terms_of_service_agreed,omitempty"`
	Contacts             []string          `protobuf:"bytes,9,rep,name=contacts" json:"contacts,omitempty"`
	AccountUrl           string            `protobuf:"bytes,10,opt,name=account_url,json=accountUrl,proto3" json:"account_url,omitempty"`
	DirectoryUrl         string            `protobuf:"bytes,11,opt,name=directory_url,json=directoryUrl,proto3" json:"directory_url,omitempty"`
	IntegrationName      string            `protobuf:"bytes,12,opt,name=integration_name,json=integrationName,proto3" json:"integration_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ACMEAccount) Reset()         { *m = ACMEAccount{} }
func (m *ACMEAccount) String() string { return proto.CompactTextString(m) }
func (*ACMEAccount) ProtoMessage()    {}
func (*ACMEAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_acme_account_9e6597a900df71f2, []int{0}
}
func (m *ACMEAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ACMEAccount.Unmarshal(m, b)
}
func (m *ACMEAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ACMEAccount.Marshal(b, m, deterministic)
}
func (dst *ACMEAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ACMEAccount.Merge(dst, src)
}
func (m *ACMEAccount) XXX_Size() int {
	return xxx_messageInfo_ACMEAccount.Size(m)
}
func (m *ACMEAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_ACMEAccount.DiscardUnknown(m)
}

var xxx_messageInfo_ACMEAccount proto.InternalMessageInfo

func (m *ACMEAccount) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ACMEAccount) GetCreateTime() *types.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *ACMEAccount) GetUpdateTime() *types.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

func (m *ACMEAccount) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *ACMEAccount) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ACMEAccount) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ACMEAccount) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *ACMEAccount) GetTermsOfServiceAgreed() bool {
	if m != nil {
		return m.TermsOfServiceAgreed
	}
	return false
}

func (m *ACMEAccount) GetContacts() []string {
	if m != nil {
		return m.Contacts
	}
	return nil
}

func (m *ACMEAccount) GetAccountUrl() string {
	if m != nil {
		return m.AccountUrl
	}
	return ""
}

func (m *ACMEAccount) GetDirectoryUrl() string {
	if m != nil {
		return m.DirectoryUrl
	}
	return ""
}

func (m *ACMEAccount) GetIntegrationName() string {
	if m != nil {
		return m.IntegrationName
	}
	return ""
}

type CreateACMEAccountRequest struct {
	AcmeAccount          *ACMEAccount `protobuf:"bytes,1,opt,name=acme_account,json=acmeAccount" json:"acme_account,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateACMEAccountRequest) Reset()         { *m = CreateACMEAccountRequest{} }
func (m *CreateACMEAccountRequest) String() string { return proto.CompactTextString(m) }
func (*CreateACMEAccountRequest) ProtoMessage()    {}
func (*CreateACMEAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_acme_account_9e6597a900df71f2, []int{1}
}
func (m *CreateACMEAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateACMEAccountRequest.Unmarshal(m, b)
}
func (m *CreateACMEAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateACMEAccountRequest.Marshal(b, m, deterministic)
}
func (dst *CreateACMEAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateACMEAccountRequest.Merge(dst, src)
}
func (m *CreateACMEAccountRequest) XXX_Size() int {
	return xxx_messageInfo_CreateACMEAccountRequest.Size(m)
}
func (m *CreateACMEAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateACMEAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateACMEAccountRequest proto.InternalMessageInfo

func (m *CreateACMEAccountRequest) GetAcmeAccount() *ACMEAccount {
	if m != nil {
		return m.AcmeAccount
	}
	return nil
}

type DeleteACMEAccountRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteACMEAccountRequest) Reset()         { *m = DeleteACMEAccountRequest{} }
func (m *DeleteACMEAccountRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteACMEAccountRequest) ProtoMessage()    {}
func (*DeleteACMEAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_acme_account_9e6597a900df71f2, []int{2}
}
func (m *DeleteACMEAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteACMEAccountRequest.Unmarshal(m, b)
}
func (m *DeleteACMEAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteACMEAccountRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteACMEAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteACMEAccountRequest.Merge(dst, src)
}
func (m *DeleteACMEAccountRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteACMEAccountRequest.Size(m)
}
func (m *DeleteACMEAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteACMEAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteACMEAccountRequest proto.InternalMessageInfo

func (m *DeleteACMEAccountRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetACMEAccountRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetACMEAccountRequest) Reset()         { *m = GetACMEAccountRequest{} }
func (m *GetACMEAccountRequest) String() string { return proto.CompactTextString(m) }
func (*GetACMEAccountRequest) ProtoMessage()    {}
func (*GetACMEAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_acme_account_9e6597a900df71f2, []int{3}
}
func (m *GetACMEAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetACMEAccountRequest.Unmarshal(m, b)
}
func (m *GetACMEAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetACMEAccountRequest.Marshal(b, m, deterministic)
}
func (dst *GetACMEAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetACMEAccountRequest.Merge(dst, src)
}
func (m *GetACMEAccountRequest) XXX_Size() int {
	return xxx_messageInfo_GetACMEAccountRequest.Size(m)
}
func (m *GetACMEAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetACMEAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetACMEAccountRequest proto.InternalMessageInfo

func (m *GetACMEAccountRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ListACMEAccountsRequest struct {
	PageSize             int32    `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken            string   `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListACMEAccountsRequest) Reset()         { *m = ListACMEAccountsRequest{} }
func (m *ListACMEAccountsRequest) String() string { return proto.CompactTextString(m) }
func (*ListACMEAccountsRequest) ProtoMessage()    {}
func (*ListACMEAccountsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_acme_account_9e6597a900df71f2, []int{4}
}
func (m *ListACMEAccountsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListACMEAccountsRequest.Unmarshal(m, b)
}
func (m *ListACMEAccountsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListACMEAccountsRequest.Marshal(b, m, deterministic)
}
func (dst *ListACMEAccountsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListACMEAccountsRequest.Merge(dst, src)
}
func (m *ListACMEAccountsRequest) XXX_Size() int {
	return xxx_messageInfo_ListACMEAccountsRequest.Size(m)
}
func (m *ListACMEAccountsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListACMEAccountsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListACMEAccountsRequest proto.InternalMessageInfo

func (m *ListACMEAccountsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListACMEAccountsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

type ListACMEAccountsResponse struct {
	AcmeAccounts         []*ACMEAccount `protobuf:"bytes,1,rep,name=acme_accounts,json=acmeAccounts" json:"acme_accounts,omitempty"`
	NextPageToken        string         `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ListACMEAccountsResponse) Reset()         { *m = ListACMEAccountsResponse{} }
func (m *ListACMEAccountsResponse) String() string { return proto.CompactTextString(m) }
func (*ListACMEAccountsResponse) ProtoMessage()    {}
func (*ListACMEAccountsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_acme_account_9e6597a900df71f2, []int{5}
}
func (m *ListACMEAccountsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListACMEAccountsResponse.Unmarshal(m, b)
}
func (m *ListACMEAccountsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListACMEAccountsResponse.Marshal(b, m, deterministic)
}
func (dst *ListACMEAccountsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListACMEAccountsResponse.Merge(dst, src)
}
func (m *ListACMEAccountsResponse) XXX_Size() int {
	return xxx_messageInfo_ListACMEAccountsResponse.Size(m)
}
func (m *ListACMEAccountsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListACMEAccountsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListACMEAccountsResponse proto.InternalMessageInfo

func (m *ListACMEAccountsResponse) GetAcmeAccounts() []*ACMEAccount {
	if m != nil {
		return m.AcmeAccounts
	}
	return nil
}

func (m *ListACMEAccountsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

type UpdateACMEAccountRequest struct {
	Name                 string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	AcmeAccount          *ACMEAccount `protobuf:"bytes,2,opt,name=acme_account,json=acmeAccount" json:"acme_account,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UpdateACMEAccountRequest) Reset()         { *m = UpdateACMEAccountRequest{} }
func (m *UpdateACMEAccountRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateACMEAccountRequest) ProtoMessage()    {}
func (*UpdateACMEAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_acme_account_9e6597a900df71f2, []int{6}
}
func (m *UpdateACMEAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateACMEAccountRequest.Unmarshal(m, b)
}
func (m *UpdateACMEAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateACMEAccountRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateACMEAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateACMEAccountRequest.Merge(dst, src)
}
func (m *UpdateACMEAccountRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateACMEAccountRequest.Size(m)
}
func (m *UpdateACMEAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateACMEAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateACMEAccountRequest proto.InternalMessageInfo

func (m *UpdateACMEAccountRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateACMEAccountRequest) GetAcmeAccount() *ACMEAccount {
	if m != nil {
		return m.AcmeAccount
	}
	return nil
}

func init() {
	proto.RegisterType((*ACMEAccount)(nil), "powerssl.apiserver.v1.ACMEAccount")
	proto.RegisterMapType((map[string]string)(nil), "powerssl.apiserver.v1.ACMEAccount.LabelsEntry")
	proto.RegisterType((*CreateACMEAccountRequest)(nil), "powerssl.apiserver.v1.CreateACMEAccountRequest")
	proto.RegisterType((*DeleteACMEAccountRequest)(nil), "powerssl.apiserver.v1.DeleteACMEAccountRequest")
	proto.RegisterType((*GetACMEAccountRequest)(nil), "powerssl.apiserver.v1.GetACMEAccountRequest")
	proto.RegisterType((*ListACMEAccountsRequest)(nil), "powerssl.apiserver.v1.ListACMEAccountsRequest")
	proto.RegisterType((*ListACMEAccountsResponse)(nil), "powerssl.apiserver.v1.ListACMEAccountsResponse")
	proto.RegisterType((*UpdateACMEAccountRequest)(nil), "powerssl.apiserver.v1.UpdateACMEAccountRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ACMEAccountService service

type ACMEAccountServiceClient interface {
	Create(ctx context.Context, in *CreateACMEAccountRequest, opts ...grpc.CallOption) (*ACMEAccount, error)
	Delete(ctx context.Context, in *DeleteACMEAccountRequest, opts ...grpc.CallOption) (*types.Empty, error)
	Get(ctx context.Context, in *GetACMEAccountRequest, opts ...grpc.CallOption) (*ACMEAccount, error)
	List(ctx context.Context, in *ListACMEAccountsRequest, opts ...grpc.CallOption) (*ListACMEAccountsResponse, error)
	Update(ctx context.Context, in *UpdateACMEAccountRequest, opts ...grpc.CallOption) (*ACMEAccount, error)
}

type aCMEAccountServiceClient struct {
	cc *grpc.ClientConn
}

func NewACMEAccountServiceClient(cc *grpc.ClientConn) ACMEAccountServiceClient {
	return &aCMEAccountServiceClient{cc}
}

func (c *aCMEAccountServiceClient) Create(ctx context.Context, in *CreateACMEAccountRequest, opts ...grpc.CallOption) (*ACMEAccount, error) {
	out := new(ACMEAccount)
	err := c.cc.Invoke(ctx, "/powerssl.apiserver.v1.ACMEAccountService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aCMEAccountServiceClient) Delete(ctx context.Context, in *DeleteACMEAccountRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/powerssl.apiserver.v1.ACMEAccountService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aCMEAccountServiceClient) Get(ctx context.Context, in *GetACMEAccountRequest, opts ...grpc.CallOption) (*ACMEAccount, error) {
	out := new(ACMEAccount)
	err := c.cc.Invoke(ctx, "/powerssl.apiserver.v1.ACMEAccountService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aCMEAccountServiceClient) List(ctx context.Context, in *ListACMEAccountsRequest, opts ...grpc.CallOption) (*ListACMEAccountsResponse, error) {
	out := new(ListACMEAccountsResponse)
	err := c.cc.Invoke(ctx, "/powerssl.apiserver.v1.ACMEAccountService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aCMEAccountServiceClient) Update(ctx context.Context, in *UpdateACMEAccountRequest, opts ...grpc.CallOption) (*ACMEAccount, error) {
	out := new(ACMEAccount)
	err := c.cc.Invoke(ctx, "/powerssl.apiserver.v1.ACMEAccountService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ACMEAccountService service

type ACMEAccountServiceServer interface {
	Create(context.Context, *CreateACMEAccountRequest) (*ACMEAccount, error)
	Delete(context.Context, *DeleteACMEAccountRequest) (*types.Empty, error)
	Get(context.Context, *GetACMEAccountRequest) (*ACMEAccount, error)
	List(context.Context, *ListACMEAccountsRequest) (*ListACMEAccountsResponse, error)
	Update(context.Context, *UpdateACMEAccountRequest) (*ACMEAccount, error)
}

func RegisterACMEAccountServiceServer(s *grpc.Server, srv ACMEAccountServiceServer) {
	s.RegisterService(&_ACMEAccountService_serviceDesc, srv)
}

func _ACMEAccountService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateACMEAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ACMEAccountServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/powerssl.apiserver.v1.ACMEAccountService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ACMEAccountServiceServer).Create(ctx, req.(*CreateACMEAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ACMEAccountService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteACMEAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ACMEAccountServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/powerssl.apiserver.v1.ACMEAccountService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ACMEAccountServiceServer).Delete(ctx, req.(*DeleteACMEAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ACMEAccountService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetACMEAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ACMEAccountServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/powerssl.apiserver.v1.ACMEAccountService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ACMEAccountServiceServer).Get(ctx, req.(*GetACMEAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ACMEAccountService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListACMEAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ACMEAccountServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/powerssl.apiserver.v1.ACMEAccountService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ACMEAccountServiceServer).List(ctx, req.(*ListACMEAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ACMEAccountService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateACMEAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ACMEAccountServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/powerssl.apiserver.v1.ACMEAccountService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ACMEAccountServiceServer).Update(ctx, req.(*UpdateACMEAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ACMEAccountService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "powerssl.apiserver.v1.ACMEAccountService",
	HandlerType: (*ACMEAccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ACMEAccountService_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ACMEAccountService_Delete_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ACMEAccountService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ACMEAccountService_List_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ACMEAccountService_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apiserver/api/v1/acme_account.proto",
}

func init() {
	proto.RegisterFile("apiserver/api/v1/acme_account.proto", fileDescriptor_acme_account_9e6597a900df71f2)
}

var fileDescriptor_acme_account_9e6597a900df71f2 = []byte{
	// 767 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xcd, 0x4e, 0xdb, 0x4a,
	0x14, 0x96, 0xf3, 0x77, 0x93, 0xe3, 0x20, 0xb8, 0x73, 0xe1, 0xe2, 0x6b, 0xb8, 0x22, 0x98, 0x8a,
	0xa6, 0xb4, 0xb5, 0x4b, 0xaa, 0x4a, 0x2d, 0xa8, 0x0b, 0x4a, 0x53, 0x36, 0xf4, 0x47, 0x06, 0x36,
	0xdd, 0x58, 0x83, 0x33, 0x44, 0x16, 0x8e, 0xed, 0x7a, 0xc6, 0x69, 0x03, 0xea, 0x86, 0x45, 0x37,
	0x5d, 0x74, 0xd1, 0x97, 0xe8, 0x4b, 0xf4, 0x29, 0xfa, 0x0a, 0x7d, 0x90, 0x6a, 0x66, 0x9c, 0xd4,
	0x90, 0x18, 0xa2, 0xee, 0x32, 0xe7, 0x7c, 0xdf, 0x39, 0xc7, 0xdf, 0xf9, 0x66, 0x02, 0x6b, 0x38,
	0xf2, 0x28, 0x89, 0xfb, 0x24, 0xb6, 0x70, 0xe4, 0x59, 0xfd, 0x4d, 0x0b, 0xbb, 0x3d, 0xe2, 0x60,
	0xd7, 0x0d, 0x93, 0x80, 0x99, 0x51, 0x1c, 0xb2, 0x10, 0x2d, 0x44, 0xe1, 0x7b, 0x12, 0x53, 0xea,
	0x9b, 0x23, 0xb4, 0xd9, 0xdf, 0xd4, 0x97, 0xbb, 0x61, 0xd8, 0xf5, 0x89, 0x20, 0xe2, 0x20, 0x08,
	0x19, 0x66, 0x5e, 0x18, 0x50, 0x49, 0xd2, 0x97, 0xd2, 0xac, 0x38, 0x1d, 0x27, 0x27, 0x16, 0xe9,
	0x45, 0x6c, 0x90, 0x26, 0x57, 0xae, 0x26, 0x99, 0xd7, 0x23, 0x94, 0xe1, 0x5e, 0x24, 0x01, 0xc6,
	0xf7, 0x12, 0xa8, 0x3b, 0xbb, 0x2f, 0xdb, 0x3b, 0x72, 0x10, 0x84, 0xa0, 0x14, 0xe0, 0x1e, 0xd1,
	0x94, 0x86, 0xd2, 0xac, 0xd9, 0xe2, 0x37, 0xda, 0x06, 0xd5, 0x8d, 0x09, 0x66, 0xc4, 0xe1, 0x6c,
	0xad, 0xd0, 0x50, 0x9a, 0x6a, 0x4b, 0x37, 0x65, 0x69, 0x73, 0x58, 0xda, 0x3c, 0x1c, 0x96, 0xb6,
	0x41, 0xc2, 0x79, 0x80, 0x93, 0x93, 0xa8, 0x33, 0x22, 0x17, 0x6f, 0x26, 0x4b, 0xb8, 0x20, 0xaf,
	0x42, 0xbd, 0xe3, 0xd1, 0xc8, 0xc7, 0x03, 0x47, 0x4c, 0x55, 0x12, 0x53, 0xa9, 0x69, 0xec, 0x15,
	0x1f, 0x6e, 0x1e, 0xca, 0xcc, 0x63, 0x3e, 0xd1, 0xca, 0x22, 0x27, 0x0f, 0xa8, 0x01, 0x6a, 0x87,
	0x50, 0x37, 0xf6, 0x22, 0x2e, 0x95, 0x56, 0x49, 0x79, 0xbf, 0x43, 0xe8, 0x05, 0x54, 0x7c, 0x7c,
	0x4c, 0x7c, 0xaa, 0xfd, 0xd5, 0x28, 0x36, 0xd5, 0x96, 0x69, 0x4e, 0x14, 0xdf, 0xcc, 0x88, 0x63,
	0xee, 0x0b, 0x42, 0x3b, 0x60, 0xf1, 0xc0, 0x4e, 0xd9, 0xe8, 0x11, 0x2c, 0x32, 0x12, 0xf7, 0xa8,
	0x13, 0x9e, 0x38, 0x9c, 0xe5, 0xb9, 0xc4, 0xc1, 0xdd, 0x98, 0x90, 0x8e, 0x56, 0x6d, 0x28, 0xcd,
	0xaa, 0x3d, 0x2f, 0xd2, 0xaf, 0x4f, 0x0e, 0x64, 0x72, 0x47, 0xe4, 0x90, 0x0e, 0x55, 0x37, 0x0c,
	0x18, 0x76, 0x19, 0xd5, 0x6a, 0x8d, 0x62, 0xb3, 0x66, 0x8f, 0xce, 0x68, 0x05, 0xd4, 0xd4, 0x17,
	0x4e, 0x12, 0xfb, 0x1a, 0x88, 0xe1, 0x21, 0x0d, 0x1d, 0xc5, 0x3e, 0x5a, 0x83, 0x99, 0x8e, 0x17,
	0x13, 0x97, 0x85, 0xf1, 0x40, 0x40, 0x54, 0x01, 0xa9, 0x8f, 0x82, 0x1c, 0x74, 0x07, 0xe6, 0xbc,
	0x80, 0x91, 0x6e, 0x2c, 0xdc, 0x22, 0xf5, 0xab, 0x0b, 0xdc, 0x6c, 0x26, 0xce, 0x35, 0xd4, 0x9f,
	0x80, 0x9a, 0xf9, 0x34, 0x34, 0x07, 0xc5, 0x53, 0x32, 0x48, 0x2d, 0xc0, 0x7f, 0x72, 0x91, 0xfb,
	0xd8, 0x4f, 0xe4, 0xee, 0x6b, 0xb6, 0x3c, 0x6c, 0x15, 0x1e, 0x2b, 0x06, 0x06, 0x6d, 0x57, 0x2c,
	0x3b, 0xa3, 0x93, 0x4d, 0xde, 0x25, 0x84, 0x32, 0xd4, 0x86, 0x7a, 0xd6, 0xe4, 0xa2, 0xa0, 0xda,
	0x32, 0x6e, 0x16, 0xda, 0x56, 0x39, 0x2f, 0x3d, 0x18, 0x26, 0x68, 0xcf, 0x89, 0x4f, 0x26, 0xb6,
	0x98, 0x60, 0x57, 0xe3, 0x2e, 0x2c, 0xec, 0x11, 0x36, 0x25, 0xf8, 0x08, 0x16, 0xf7, 0x3d, 0x9a,
	0x45, 0xd3, 0x21, 0x7c, 0x09, 0x6a, 0x11, 0xee, 0x12, 0x87, 0x7a, 0x67, 0x92, 0x53, 0xb6, 0xab,
	0x3c, 0x70, 0xe0, 0x9d, 0x11, 0xf4, 0x3f, 0x80, 0x48, 0xb2, 0xf0, 0x94, 0x04, 0xa9, 0x2c, 0x02,
	0x7e, 0xc8, 0x03, 0xc6, 0x67, 0x05, 0xb4, 0xf1, 0xba, 0x34, 0x0a, 0x03, 0x4a, 0xd0, 0x1e, 0xcc,
	0x64, 0x75, 0xa1, 0x9a, 0x22, 0x1c, 0x38, 0x8d, 0x30, 0xf5, 0x8c, 0x30, 0x14, 0xad, 0xc3, 0x6c,
	0x40, 0x3e, 0x30, 0x67, 0x6c, 0x92, 0x19, 0x1e, 0x7e, 0x33, 0x9a, 0x26, 0x01, 0xed, 0x48, 0x5c,
	0xaa, 0xe9, 0x44, 0x19, 0x5b, 0x5c, 0xe1, 0x8f, 0x16, 0xd7, 0xfa, 0x56, 0x06, 0x94, 0x49, 0xa6,
	0x17, 0x00, 0x7d, 0x52, 0xa0, 0x22, 0x3d, 0x83, 0xac, 0x9c, 0x92, 0x79, 0x96, 0xd2, 0xa7, 0x98,
	0xc1, 0xb8, 0x7d, 0xf1, 0xe3, 0xe7, 0xd7, 0xc2, 0xaa, 0xf1, 0xf7, 0xf0, 0x95, 0xbd, 0x3f, 0x14,
	0x7a, 0xab, 0x9e, 0x3d, 0xa2, 0x04, 0x2a, 0xd2, 0x58, 0xb9, 0x73, 0xe4, 0xf9, 0x4e, 0xff, 0x77,
	0xec, 0x01, 0x6b, 0xf3, 0x57, 0xd7, 0x30, 0x44, 0xef, 0xe5, 0x0d, 0x9d, 0xf7, 0x3e, 0xe7, 0x62,
	0x3e, 0xbd, 0x34, 0x81, 0xb5, 0xf1, 0x11, 0x9d, 0x43, 0x71, 0x8f, 0x30, 0x74, 0x2f, 0xa7, 0xe7,
	0x44, 0xef, 0x4e, 0xf5, 0xe1, 0x69, 0x73, 0x74, 0x5d, 0xf3, 0x0b, 0x05, 0x4a, 0xdc, 0x98, 0x28,
	0xef, 0xbd, 0xcb, 0xb9, 0x0d, 0xba, 0x35, 0x35, 0x5e, 0xba, 0xdc, 0xf8, 0x4f, 0x4c, 0xf3, 0x0f,
	0x1a, 0x5f, 0x03, 0xfa, 0xa2, 0x40, 0x45, 0x1a, 0x32, 0x57, 0xf9, 0x3c, 0xbf, 0x4e, 0x25, 0xc4,
	0x03, 0xd1, 0x7a, 0x43, 0xbf, 0x46, 0x88, 0xcb, 0x56, 0x78, 0xb6, 0xfe, 0xf6, 0xd6, 0xa8, 0xac,
	0x17, 0x5a, 0xd1, 0x69, 0xd7, 0xba, 0xfa, 0x87, 0xbd, 0x8d, 0x23, 0xef, 0xb8, 0x22, 0xf6, 0xfd,
	0xf0, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdb, 0x97, 0xf5, 0x2e, 0xcf, 0x07, 0x00, 0x00,
}
