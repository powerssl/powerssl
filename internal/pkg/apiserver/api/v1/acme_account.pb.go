// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: powerssl/apiserver/v1/acme_account.proto

package api

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/googleapis/google/api"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	grpc "google.golang.org/grpc"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ACMEAccount struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CreateTime           *types.Timestamp  `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime           *types.Timestamp  `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	DisplayName          string            `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Title                string            `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	Description          string            `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	Labels               map[string]string `protobuf:"bytes,7,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	TermsOfServiceAgreed bool              `protobuf:"varint,8,opt,name=terms_of_service_agreed,json=termsOfServiceAgreed,proto3" json:"terms_of_service_agreed,omitempty"`
	Contacts             []string          `protobuf:"bytes,9,rep,name=contacts,proto3" json:"contacts,omitempty"`
	AccountUrl           string            `protobuf:"bytes,10,opt,name=account_url,json=accountUrl,proto3" json:"account_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ACMEAccount) Reset()         { *m = ACMEAccount{} }
func (m *ACMEAccount) String() string { return proto.CompactTextString(m) }
func (*ACMEAccount) ProtoMessage()    {}
func (*ACMEAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ba36e91bea5548e, []int{0}
}
func (m *ACMEAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ACMEAccount.Unmarshal(m, b)
}
func (m *ACMEAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ACMEAccount.Marshal(b, m, deterministic)
}
func (m *ACMEAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ACMEAccount.Merge(m, src)
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
	AcmeAccount          *ACMEAccount `protobuf:"bytes,2,opt,name=acme_account,json=acmeAccount,proto3" json:"acme_account,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateACMEAccountRequest) Reset()         { *m = CreateACMEAccountRequest{} }
func (m *CreateACMEAccountRequest) String() string { return proto.CompactTextString(m) }
func (*CreateACMEAccountRequest) ProtoMessage()    {}
func (*CreateACMEAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ba36e91bea5548e, []int{1}
}
func (m *CreateACMEAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateACMEAccountRequest.Unmarshal(m, b)
}
func (m *CreateACMEAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateACMEAccountRequest.Marshal(b, m, deterministic)
}
func (m *CreateACMEAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateACMEAccountRequest.Merge(m, src)
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
	return fileDescriptor_7ba36e91bea5548e, []int{2}
}
func (m *DeleteACMEAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteACMEAccountRequest.Unmarshal(m, b)
}
func (m *DeleteACMEAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteACMEAccountRequest.Marshal(b, m, deterministic)
}
func (m *DeleteACMEAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteACMEAccountRequest.Merge(m, src)
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
	return fileDescriptor_7ba36e91bea5548e, []int{3}
}
func (m *GetACMEAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetACMEAccountRequest.Unmarshal(m, b)
}
func (m *GetACMEAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetACMEAccountRequest.Marshal(b, m, deterministic)
}
func (m *GetACMEAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetACMEAccountRequest.Merge(m, src)
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
	return fileDescriptor_7ba36e91bea5548e, []int{4}
}
func (m *ListACMEAccountsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListACMEAccountsRequest.Unmarshal(m, b)
}
func (m *ListACMEAccountsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListACMEAccountsRequest.Marshal(b, m, deterministic)
}
func (m *ListACMEAccountsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListACMEAccountsRequest.Merge(m, src)
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
	AcmeAccounts         []*ACMEAccount `protobuf:"bytes,1,rep,name=acme_accounts,json=acmeAccounts,proto3" json:"acme_accounts,omitempty"`
	NextPageToken        string         `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ListACMEAccountsResponse) Reset()         { *m = ListACMEAccountsResponse{} }
func (m *ListACMEAccountsResponse) String() string { return proto.CompactTextString(m) }
func (*ListACMEAccountsResponse) ProtoMessage()    {}
func (*ListACMEAccountsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ba36e91bea5548e, []int{5}
}
func (m *ListACMEAccountsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListACMEAccountsResponse.Unmarshal(m, b)
}
func (m *ListACMEAccountsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListACMEAccountsResponse.Marshal(b, m, deterministic)
}
func (m *ListACMEAccountsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListACMEAccountsResponse.Merge(m, src)
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
	Name                 string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	UpdateMask           *types.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	AcmeAccount          *ACMEAccount     `protobuf:"bytes,3,opt,name=acme_account,json=acmeAccount,proto3" json:"acme_account,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *UpdateACMEAccountRequest) Reset()         { *m = UpdateACMEAccountRequest{} }
func (m *UpdateACMEAccountRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateACMEAccountRequest) ProtoMessage()    {}
func (*UpdateACMEAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ba36e91bea5548e, []int{6}
}
func (m *UpdateACMEAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateACMEAccountRequest.Unmarshal(m, b)
}
func (m *UpdateACMEAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateACMEAccountRequest.Marshal(b, m, deterministic)
}
func (m *UpdateACMEAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateACMEAccountRequest.Merge(m, src)
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

func (m *UpdateACMEAccountRequest) GetUpdateMask() *types.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
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

func init() {
	proto.RegisterFile("powerssl/apiserver/v1/acme_account.proto", fileDescriptor_7ba36e91bea5548e)
}

var fileDescriptor_7ba36e91bea5548e = []byte{
	// 792 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xcd, 0x6e, 0xeb, 0x44,
	0x14, 0x96, 0x9b, 0x26, 0x24, 0xc7, 0xad, 0x40, 0xa3, 0xfe, 0x58, 0x2e, 0xa8, 0xc1, 0x0b, 0x08,
	0xa1, 0xb2, 0xd5, 0x20, 0x04, 0x6d, 0xc5, 0xa2, 0x94, 0xb4, 0x9b, 0x16, 0x90, 0xdb, 0x6e, 0xd8,
	0x58, 0x53, 0xe7, 0x24, 0xb2, 0xe2, 0x3f, 0x3c, 0x93, 0x40, 0x8a, 0xba, 0xe9, 0x0a, 0x89, 0x15,
	0x62, 0xc3, 0x73, 0xf0, 0x2a, 0xbc, 0x02, 0x12, 0xab, 0xfb, 0x0e, 0x57, 0x33, 0xe3, 0xe4, 0xba,
	0x6d, 0xac, 0xf8, 0xde, 0x9d, 0xe7, 0x9c, 0xf3, 0x1d, 0x7f, 0xfe, 0xce, 0x77, 0x46, 0x86, 0x4e,
	0x9a, 0xfc, 0x82, 0x19, 0x63, 0xa1, 0x43, 0xd3, 0x80, 0x61, 0x36, 0xc5, 0xcc, 0x99, 0x1e, 0x3a,
	0xd4, 0x8f, 0xd0, 0xa3, 0xbe, 0x9f, 0x4c, 0x62, 0x6e, 0xa7, 0x59, 0xc2, 0x13, 0xb2, 0x3d, 0xaf,
	0xb4, 0x17, 0x95, 0xf6, 0xf4, 0xd0, 0xfc, 0x70, 0x94, 0x24, 0xa3, 0x10, 0x05, 0xdc, 0xa1, 0x71,
	0x9c, 0x70, 0xca, 0x83, 0x24, 0x66, 0x0a, 0x64, 0xee, 0xe5, 0x59, 0x79, 0xba, 0x9b, 0x0c, 0x1d,
	0x8c, 0x52, 0x3e, 0xcb, 0x93, 0xed, 0xe7, 0xc9, 0x61, 0x80, 0xe1, 0xc0, 0x8b, 0x28, 0x1b, 0xe7,
	0x15, 0xfb, 0xcf, 0x2b, 0x78, 0x10, 0x21, 0xe3, 0x34, 0x4a, 0x55, 0x81, 0xf5, 0xaa, 0x06, 0xfa,
	0xe9, 0xd9, 0x55, 0xff, 0x54, 0x51, 0x25, 0x04, 0xd6, 0x63, 0x1a, 0xa1, 0xa1, 0xb5, 0xb5, 0x4e,
	0xcb, 0x95, 0xcf, 0xe4, 0x04, 0x74, 0x3f, 0x43, 0xca, 0xd1, 0x13, 0x68, 0x63, 0xad, 0xad, 0x75,
	0xf4, 0x9e, 0x69, 0xab, 0xd6, 0xf6, 0xbc, 0xb5, 0x7d, 0x33, 0x6f, 0xed, 0x82, 0x2a, 0x17, 0x01,
	0x01, 0x9e, 0xa4, 0x83, 0x05, 0xb8, 0xb6, 0x1a, 0xac, 0xca, 0x25, 0xf8, 0x63, 0xd8, 0x18, 0x04,
	0x2c, 0x0d, 0xe9, 0xcc, 0x93, 0xac, 0xd6, 0x25, 0x2b, 0x3d, 0x8f, 0x7d, 0x2f, 0xc8, 0x6d, 0x41,
	0x9d, 0x07, 0x3c, 0x44, 0xa3, 0x2e, 0x73, 0xea, 0x40, 0xda, 0xa0, 0x0f, 0x90, 0xf9, 0x59, 0x90,
	0x0a, 0x31, 0x8d, 0x46, 0x8e, 0x7b, 0x13, 0x22, 0xe7, 0xd0, 0x08, 0xe9, 0x1d, 0x86, 0xcc, 0x78,
	0xaf, 0x5d, 0xeb, 0xe8, 0x3d, 0xdb, 0x5e, 0x3a, 0x1e, 0xbb, 0x20, 0x8e, 0x7d, 0x29, 0x01, 0xfd,
	0x98, 0x67, 0x33, 0x37, 0x47, 0x93, 0x2f, 0x61, 0x97, 0x63, 0x16, 0x31, 0x2f, 0x19, 0x7a, 0x02,
	0x15, 0xf8, 0xe8, 0xd1, 0x51, 0x86, 0x38, 0x30, 0x9a, 0x6d, 0xad, 0xd3, 0x74, 0xb7, 0x64, 0xfa,
	0x87, 0xe1, 0xb5, 0x4a, 0x9e, 0xca, 0x1c, 0x31, 0xa1, 0xe9, 0x27, 0x31, 0xa7, 0x3e, 0x67, 0x46,
	0xab, 0x5d, 0xeb, 0xb4, 0xdc, 0xc5, 0x99, 0xec, 0x83, 0x9e, 0x3b, 0xc7, 0x9b, 0x64, 0xa1, 0x01,
	0x92, 0x3c, 0xe4, 0xa1, 0xdb, 0x2c, 0x34, 0x8f, 0x40, 0x2f, 0x50, 0x21, 0x1f, 0x40, 0x6d, 0x8c,
	0xb3, 0x7c, 0x64, 0xe2, 0x51, 0x88, 0x32, 0xa5, 0xe1, 0x44, 0xcd, 0xaa, 0xe5, 0xaa, 0xc3, 0xf1,
	0xda, 0xd7, 0x9a, 0x35, 0x03, 0xe3, 0x4c, 0x0e, 0xa7, 0xf0, 0x5d, 0x2e, 0xfe, 0x3c, 0x41, 0xc6,
	0xc9, 0x0e, 0x34, 0x52, 0x9a, 0x61, 0xcc, 0xf3, 0x56, 0xf9, 0x89, 0xf4, 0x61, 0xa3, 0x68, 0xe7,
	0xdc, 0x00, 0xd6, 0x6a, 0xc1, 0x5c, 0x5d, 0xe0, 0xf2, 0x83, 0x65, 0x83, 0xf1, 0x1d, 0x86, 0xb8,
	0xf4, 0xd5, 0x4b, 0x6c, 0x67, 0x7d, 0x0e, 0xdb, 0x17, 0xc8, 0x2b, 0x16, 0x47, 0xb0, 0x7b, 0x19,
	0xb0, 0x62, 0x35, 0x5b, 0xf5, 0x59, 0x7b, 0xd0, 0x4a, 0xe9, 0x08, 0x3d, 0x16, 0xdc, 0x2b, 0xa1,
	0xea, 0x6e, 0x53, 0x04, 0xae, 0x83, 0x7b, 0x24, 0x1f, 0x01, 0xc8, 0x24, 0x4f, 0xc6, 0x18, 0x4b,
	0xd7, 0xb6, 0x5c, 0x59, 0x7e, 0x23, 0x02, 0xd6, 0x1f, 0x1a, 0x18, 0x2f, 0xdf, 0xc7, 0xd2, 0x24,
	0x66, 0x48, 0x2e, 0x60, 0xb3, 0xa8, 0x17, 0x33, 0x34, 0xe9, 0xb0, 0x2a, 0x82, 0x6d, 0x14, 0x04,
	0x63, 0xe4, 0x13, 0x78, 0x3f, 0xc6, 0x5f, 0xb9, 0x57, 0x60, 0xa2, 0x06, 0xba, 0x29, 0xc2, 0x3f,
	0x2e, 0xd8, 0xfc, 0xa3, 0x81, 0x71, 0x2b, 0xb7, 0xa6, 0x9a, 0x5a, 0x85, 0xa5, 0x14, 0x77, 0x45,
	0xe9, 0x46, 0x9f, 0x8b, 0xeb, 0xe4, 0x8a, 0xb2, 0xf1, 0x7c, 0x29, 0xc5, 0xf3, 0x0b, 0x3b, 0xd4,
	0xde, 0xc9, 0x0e, 0xbd, 0xff, 0xeb, 0x40, 0x0a, 0xc9, 0x7c, 0x3d, 0xc8, 0x9f, 0x1a, 0x34, 0x94,
	0x43, 0x89, 0x53, 0xd2, 0xb2, 0xcc, 0xc0, 0x66, 0x05, 0x0e, 0x56, 0xef, 0xf1, 0xdf, 0xff, 0xfe,
	0x5a, 0x3b, 0xb0, 0x3e, 0x15, 0xb7, 0xf4, 0x6f, 0xca, 0x0a, 0xdf, 0x08, 0x5e, 0xd7, 0xb2, 0x9c,
	0x39, 0xdd, 0x07, 0xa7, 0x38, 0x84, 0x63, 0xad, 0x4b, 0x1e, 0x35, 0x68, 0x28, 0xeb, 0x96, 0x72,
	0x2a, 0x73, 0xb6, 0xb9, 0xf3, 0x42, 0xd5, 0xbe, 0xb8, 0xc1, 0x2d, 0x47, 0xf2, 0xf8, 0xac, 0xab,
	0x78, 0x88, 0xa9, 0x3c, 0x65, 0xf1, 0x84, 0x84, 0xd3, 0x7d, 0x20, 0xbf, 0x6b, 0x50, 0xbb, 0x40,
	0x4e, 0x0e, 0x4a, 0x18, 0x2c, 0xdd, 0x95, 0x4a, 0x92, 0xe4, 0x54, 0x48, 0x65, 0x2a, 0x7f, 0x6b,
	0xb0, 0x2e, 0xdc, 0x4f, 0xca, 0x2e, 0xcd, 0x92, 0x55, 0x34, 0x9d, 0xca, 0xf5, 0x6a, 0x95, 0x9e,
	0x51, 0x5b, 0x3d, 0x2d, 0x69, 0x1f, 0xb5, 0x0a, 0xa5, 0xa3, 0x2a, 0xdb, 0x94, 0xb7, 0xb1, 0x8f,
	0x59, 0x55, 0xab, 0x63, 0xad, 0xfb, 0xed, 0xd1, 0x4f, 0x5f, 0x2d, 0x1a, 0x0f, 0x70, 0xea, 0x2c,
	0xfe, 0x18, 0x82, 0x98, 0x63, 0x16, 0xd3, 0xd0, 0x49, 0xc7, 0xa3, 0xc2, 0xef, 0x83, 0xf8, 0x13,
	0x98, 0x1e, 0x9e, 0xd0, 0x34, 0xb8, 0x6b, 0x48, 0xd7, 0x7c, 0xf1, 0x3a, 0x00, 0x00, 0xff, 0xff,
	0xff, 0x05, 0xf5, 0x75, 0x66, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ACMEAccountServiceClient is the client API for ACMEAccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
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

// ACMEAccountServiceServer is the server API for ACMEAccountService service.
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
	Metadata: "powerssl/apiserver/v1/acme_account.proto",
}
