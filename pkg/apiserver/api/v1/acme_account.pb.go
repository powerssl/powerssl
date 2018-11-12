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
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ACMEAccount) Reset()         { *m = ACMEAccount{} }
func (m *ACMEAccount) String() string { return proto.CompactTextString(m) }
func (*ACMEAccount) ProtoMessage()    {}
func (*ACMEAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_acme_account_527807d8f147f9d3, []int{0}
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

type CreateACMEAccountRequest struct {
	Parent               string       `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	AcmeAccount          *ACMEAccount `protobuf:"bytes,2,opt,name=acme_account,json=acmeAccount" json:"acme_account,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateACMEAccountRequest) Reset()         { *m = CreateACMEAccountRequest{} }
func (m *CreateACMEAccountRequest) String() string { return proto.CompactTextString(m) }
func (*CreateACMEAccountRequest) ProtoMessage()    {}
func (*CreateACMEAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_acme_account_527807d8f147f9d3, []int{1}
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

func (m *CreateACMEAccountRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

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
	return fileDescriptor_acme_account_527807d8f147f9d3, []int{2}
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
	return fileDescriptor_acme_account_527807d8f147f9d3, []int{3}
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
	Parent               string   `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	PageSize             int32    `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken            string   `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListACMEAccountsRequest) Reset()         { *m = ListACMEAccountsRequest{} }
func (m *ListACMEAccountsRequest) String() string { return proto.CompactTextString(m) }
func (*ListACMEAccountsRequest) ProtoMessage()    {}
func (*ListACMEAccountsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_acme_account_527807d8f147f9d3, []int{4}
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

func (m *ListACMEAccountsRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

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
	return fileDescriptor_acme_account_527807d8f147f9d3, []int{5}
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
	return fileDescriptor_acme_account_527807d8f147f9d3, []int{6}
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
	proto.RegisterFile("apiserver/api/v1/acme_account.proto", fileDescriptor_acme_account_527807d8f147f9d3)
}

var fileDescriptor_acme_account_527807d8f147f9d3 = []byte{
	// 748 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x4d, 0x4f, 0xdb, 0x48,
	0x18, 0x96, 0x09, 0x64, 0x93, 0xd7, 0x41, 0xbb, 0x1a, 0xf1, 0x61, 0x99, 0x5d, 0x91, 0xf5, 0xae,
	0x50, 0x16, 0x16, 0x5b, 0x50, 0x55, 0x2a, 0xa0, 0xaa, 0xa2, 0x34, 0xe5, 0x42, 0x3f, 0x64, 0xe0,
	0xd2, 0x8b, 0x35, 0x38, 0x2f, 0x91, 0x85, 0x63, 0xbb, 0x9e, 0x49, 0xda, 0x50, 0x71, 0xa9, 0x2a,
	0x55, 0x6a, 0x8f, 0xbd, 0xf5, 0xd0, 0x3f, 0xd5, 0xbf, 0xd0, 0x6b, 0xff, 0x43, 0x35, 0x33, 0x26,
	0x35, 0x21, 0x16, 0x46, 0xea, 0x2d, 0xef, 0xc7, 0x33, 0xf3, 0xf8, 0x79, 0x9f, 0x79, 0x03, 0xff,
	0xd0, 0x24, 0x60, 0x98, 0x0e, 0x30, 0x75, 0x68, 0x12, 0x38, 0x83, 0x0d, 0x87, 0xfa, 0x3d, 0xf4,
	0xa8, 0xef, 0xc7, 0xfd, 0x88, 0xdb, 0x49, 0x1a, 0xf3, 0x98, 0xcc, 0x27, 0xf1, 0x2b, 0x4c, 0x19,
	0x0b, 0xed, 0x51, 0xb7, 0x3d, 0xd8, 0x30, 0xff, 0xec, 0xc6, 0x71, 0x37, 0x44, 0x09, 0xa4, 0x51,
	0x14, 0x73, 0xca, 0x83, 0x38, 0x62, 0x0a, 0x64, 0x2e, 0x65, 0x55, 0x19, 0x9d, 0xf4, 0x4f, 0x1d,
	0xec, 0x25, 0x7c, 0x98, 0x15, 0x97, 0xc7, 0x8b, 0x3c, 0xe8, 0x21, 0xe3, 0xb4, 0x97, 0xa8, 0x06,
	0xeb, 0x7b, 0x05, 0xf4, 0xdd, 0xbd, 0x27, 0xed, 0x5d, 0x45, 0x84, 0x10, 0x98, 0x8e, 0x68, 0x0f,
	0x0d, 0xad, 0xa9, 0xb5, 0xea, 0xae, 0xfc, 0x4d, 0x76, 0x40, 0xf7, 0x53, 0xa4, 0x1c, 0x3d, 0x81,
	0x36, 0xa6, 0x9a, 0x5a, 0x4b, 0xdf, 0x34, 0x6d, 0x75, 0xb4, 0x7d, 0x79, 0xb4, 0x7d, 0x74, 0x79,
	0xb4, 0x0b, 0xaa, 0x5d, 0x24, 0x04, 0xb8, 0x9f, 0x74, 0x46, 0xe0, 0xca, 0xcd, 0x60, 0xd5, 0x2e,
	0xc1, 0x7f, 0x43, 0xa3, 0x13, 0xb0, 0x24, 0xa4, 0x43, 0x4f, 0xb2, 0x9a, 0x96, 0xac, 0xf4, 0x2c,
	0xf7, 0x54, 0x90, 0x9b, 0x83, 0x19, 0x1e, 0xf0, 0x10, 0x8d, 0x19, 0x59, 0x53, 0x01, 0x69, 0x82,
	0xde, 0x41, 0xe6, 0xa7, 0x41, 0x22, 0xa4, 0x32, 0xaa, 0x19, 0xee, 0x67, 0x8a, 0x3c, 0x86, 0x6a,
	0x48, 0x4f, 0x30, 0x64, 0xc6, 0x6f, 0xcd, 0x4a, 0x4b, 0xdf, 0xb4, 0xed, 0x89, 0xe2, 0xdb, 0x39,
	0x71, 0xec, 0x03, 0x09, 0x68, 0x47, 0x3c, 0x1d, 0xba, 0x19, 0x9a, 0xdc, 0x85, 0x45, 0x8e, 0x69,
	0x8f, 0x79, 0xf1, 0xa9, 0x27, 0x50, 0x81, 0x8f, 0x1e, 0xed, 0xa6, 0x88, 0x1d, 0xa3, 0xd6, 0xd4,
	0x5a, 0x35, 0x77, 0x4e, 0x96, 0x9f, 0x9d, 0x1e, 0xaa, 0xe2, 0xae, 0xac, 0x11, 0x13, 0x6a, 0x7e,
	0x1c, 0x71, 0xea, 0x73, 0x66, 0xd4, 0x9b, 0x95, 0x56, 0xdd, 0x1d, 0xc5, 0x64, 0x19, 0xf4, 0xcc,
	0x17, 0x5e, 0x3f, 0x0d, 0x0d, 0x90, 0xe4, 0x21, 0x4b, 0x1d, 0xa7, 0xa1, 0xb9, 0x05, 0x7a, 0x8e,
	0x0a, 0xf9, 0x03, 0x2a, 0x67, 0x38, 0xcc, 0x46, 0x26, 0x7e, 0x0a, 0x51, 0x06, 0x34, 0xec, 0xab,
	0x59, 0xd5, 0x5d, 0x15, 0x6c, 0x4f, 0xdd, 0xd3, 0xac, 0x21, 0x18, 0x7b, 0x72, 0x38, 0xb9, 0xef,
	0x72, 0xf1, 0x65, 0x1f, 0x19, 0x27, 0x0b, 0x50, 0x4d, 0x68, 0x8a, 0x11, 0xcf, 0x8e, 0xca, 0x22,
	0xd2, 0x86, 0x46, 0xde, 0xac, 0x99, 0x01, 0xac, 0x9b, 0x05, 0x73, 0x75, 0x81, 0xcb, 0x02, 0xcb,
	0x06, 0xe3, 0x11, 0x86, 0x38, 0xf1, 0xea, 0x09, 0xb6, 0xb3, 0xd6, 0x60, 0x7e, 0x1f, 0x79, 0xc9,
	0xe6, 0x1e, 0x2c, 0x1e, 0x04, 0x2c, 0xdf, 0xcd, 0x6e, 0xfa, 0xac, 0x25, 0xa8, 0x27, 0xb4, 0x8b,
	0x1e, 0x0b, 0xce, 0x95, 0x50, 0x33, 0x6e, 0x4d, 0x24, 0x0e, 0x83, 0x73, 0x24, 0x7f, 0x01, 0xc8,
	0x22, 0x8f, 0xcf, 0x30, 0x92, 0xae, 0xad, 0xbb, 0xb2, 0xfd, 0x48, 0x24, 0xac, 0x8f, 0x1a, 0x18,
	0xd7, 0xef, 0x63, 0x49, 0x1c, 0x31, 0x24, 0xfb, 0x30, 0x9b, 0xd7, 0x8b, 0x19, 0x9a, 0x74, 0x58,
	0x19, 0xc1, 0x1a, 0x39, 0xc1, 0x18, 0x59, 0x81, 0xdf, 0x23, 0x7c, 0xcd, 0xbd, 0x1c, 0x13, 0x35,
	0xd0, 0x59, 0x91, 0x7e, 0x3e, 0x62, 0xd3, 0x07, 0xe3, 0x58, 0x3e, 0x9a, 0x72, 0x62, 0xfd, 0xa2,
	0x81, 0x6e, 0xbe, 0xaf, 0x02, 0xc9, 0x15, 0x33, 0x83, 0x93, 0x2f, 0x1a, 0x54, 0x95, 0xc7, 0x88,
	0x53, 0x70, 0x64, 0x91, 0x05, 0xcd, 0x12, 0x1c, 0xac, 0x07, 0x6f, 0xbf, 0x7e, 0xfb, 0x34, 0xb5,
	0x65, 0xfd, 0x27, 0xb6, 0xe8, 0x1b, 0x35, 0xcc, 0xfb, 0x82, 0xd7, 0xba, 0xea, 0x67, 0xce, 0xea,
	0x85, 0xdc, 0xae, 0xeb, 0x97, 0x03, 0xd8, 0x6e, 0xe4, 0x43, 0xf2, 0x4e, 0x83, 0xaa, 0x72, 0x62,
	0x21, 0xc1, 0x22, 0xa3, 0x9a, 0x0b, 0xd7, 0x36, 0x57, 0x5b, 0xac, 0x5b, 0x6b, 0x43, 0x92, 0x5a,
	0x5b, 0x55, 0xa4, 0x84, 0xca, 0x63, 0x94, 0xae, 0x32, 0x72, 0x56, 0x2f, 0xc8, 0x07, 0x0d, 0x2a,
	0xfb, 0xc8, 0xc9, 0xff, 0x05, 0x1c, 0x26, 0x9a, 0xbf, 0x94, 0x42, 0x19, 0x19, 0x72, 0x0b, 0x32,
	0x9f, 0x35, 0x98, 0x16, 0x86, 0x26, 0x45, 0x7b, 0xb0, 0xe0, 0x75, 0x99, 0x4e, 0xe9, 0x7e, 0xf5,
	0x3a, 0xc6, 0xc8, 0x95, 0x19, 0x9f, 0x74, 0x94, 0x32, 0x78, 0xe1, 0xc0, 0x8a, 0xfc, 0x7f, 0x1b,
	0x47, 0x99, 0xe5, 0xf5, 0xba, 0xea, 0xa8, 0x87, 0x2b, 0x2f, 0xfe, 0x1d, 0xdd, 0x12, 0xc4, 0x4e,
	0x72, 0xd6, 0x75, 0xc6, 0xff, 0xef, 0x77, 0x68, 0x12, 0x9c, 0x54, 0xa5, 0x6b, 0xee, 0xfc, 0x08,
	0x00, 0x00, 0xff, 0xff, 0xaf, 0xbd, 0xb9, 0xbd, 0x0e, 0x08, 0x00, 0x00,
}
