// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: powerssl/apiserver/v1/certificate_issue.proto

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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// A CertificateIssue ...
type CertificateIssue struct {
	// The resource name of the certificate issue.
	// CertificateIssue names have the form
	// `certificates/{certificate_id}/issues/{issue_id}`. The name is ignored when
	// creating a certificate issue.
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CreateTime           *types.Timestamp  `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime           *types.Timestamp  `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	DisplayName          string            `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Title                string            `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	Description          string            `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	Labels               map[string]string `protobuf:"bytes,7,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Dnsnames             []string          `protobuf:"bytes,8,rep,name=dnsnames,proto3" json:"dnsnames,omitempty"`
	KeyAlgorithm         string            `protobuf:"bytes,9,opt,name=key_algorithm,json=keyAlgorithm,proto3" json:"key_algorithm,omitempty"`
	KeySize              int32             `protobuf:"varint,10,opt,name=key_size,json=keySize,proto3" json:"key_size,omitempty"`
	DigestAlgorithm      string            `protobuf:"bytes,11,opt,name=digest_algorithm,json=digestAlgorithm,proto3" json:"digest_algorithm,omitempty"`
	AutoRenew            bool              `protobuf:"varint,12,opt,name=auto_renew,json=autoRenew,proto3" json:"auto_renew,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CertificateIssue) Reset()         { *m = CertificateIssue{} }
func (m *CertificateIssue) String() string { return proto.CompactTextString(m) }
func (*CertificateIssue) ProtoMessage()    {}
func (*CertificateIssue) Descriptor() ([]byte, []int) {
	return fileDescriptor_58b50effdd0def26, []int{0}
}
func (m *CertificateIssue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CertificateIssue.Unmarshal(m, b)
}
func (m *CertificateIssue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CertificateIssue.Marshal(b, m, deterministic)
}
func (m *CertificateIssue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateIssue.Merge(m, src)
}
func (m *CertificateIssue) XXX_Size() int {
	return xxx_messageInfo_CertificateIssue.Size(m)
}
func (m *CertificateIssue) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateIssue.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateIssue proto.InternalMessageInfo

func (m *CertificateIssue) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CertificateIssue) GetCreateTime() *types.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *CertificateIssue) GetUpdateTime() *types.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

func (m *CertificateIssue) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *CertificateIssue) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *CertificateIssue) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CertificateIssue) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *CertificateIssue) GetDnsnames() []string {
	if m != nil {
		return m.Dnsnames
	}
	return nil
}

func (m *CertificateIssue) GetKeyAlgorithm() string {
	if m != nil {
		return m.KeyAlgorithm
	}
	return ""
}

func (m *CertificateIssue) GetKeySize() int32 {
	if m != nil {
		return m.KeySize
	}
	return 0
}

func (m *CertificateIssue) GetDigestAlgorithm() string {
	if m != nil {
		return m.DigestAlgorithm
	}
	return ""
}

func (m *CertificateIssue) GetAutoRenew() bool {
	if m != nil {
		return m.AutoRenew
	}
	return false
}

type CreateCertificateIssueRequest struct {
	Parent               string            `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	CertificateIssue     *CertificateIssue `protobuf:"bytes,2,opt,name=certificate_issue,json=certificateIssue,proto3" json:"certificate_issue,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CreateCertificateIssueRequest) Reset()         { *m = CreateCertificateIssueRequest{} }
func (m *CreateCertificateIssueRequest) String() string { return proto.CompactTextString(m) }
func (*CreateCertificateIssueRequest) ProtoMessage()    {}
func (*CreateCertificateIssueRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_58b50effdd0def26, []int{1}
}
func (m *CreateCertificateIssueRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCertificateIssueRequest.Unmarshal(m, b)
}
func (m *CreateCertificateIssueRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCertificateIssueRequest.Marshal(b, m, deterministic)
}
func (m *CreateCertificateIssueRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCertificateIssueRequest.Merge(m, src)
}
func (m *CreateCertificateIssueRequest) XXX_Size() int {
	return xxx_messageInfo_CreateCertificateIssueRequest.Size(m)
}
func (m *CreateCertificateIssueRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCertificateIssueRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCertificateIssueRequest proto.InternalMessageInfo

func (m *CreateCertificateIssueRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

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
	return fileDescriptor_58b50effdd0def26, []int{2}
}
func (m *DeleteCertificateIssueRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCertificateIssueRequest.Unmarshal(m, b)
}
func (m *DeleteCertificateIssueRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCertificateIssueRequest.Marshal(b, m, deterministic)
}
func (m *DeleteCertificateIssueRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCertificateIssueRequest.Merge(m, src)
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
	return fileDescriptor_58b50effdd0def26, []int{3}
}
func (m *GetCertificateIssueRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCertificateIssueRequest.Unmarshal(m, b)
}
func (m *GetCertificateIssueRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCertificateIssueRequest.Marshal(b, m, deterministic)
}
func (m *GetCertificateIssueRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCertificateIssueRequest.Merge(m, src)
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
	Parent               string   `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	PageSize             int32    `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken            string   `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListCertificateIssuesRequest) Reset()         { *m = ListCertificateIssuesRequest{} }
func (m *ListCertificateIssuesRequest) String() string { return proto.CompactTextString(m) }
func (*ListCertificateIssuesRequest) ProtoMessage()    {}
func (*ListCertificateIssuesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_58b50effdd0def26, []int{4}
}
func (m *ListCertificateIssuesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCertificateIssuesRequest.Unmarshal(m, b)
}
func (m *ListCertificateIssuesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCertificateIssuesRequest.Marshal(b, m, deterministic)
}
func (m *ListCertificateIssuesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCertificateIssuesRequest.Merge(m, src)
}
func (m *ListCertificateIssuesRequest) XXX_Size() int {
	return xxx_messageInfo_ListCertificateIssuesRequest.Size(m)
}
func (m *ListCertificateIssuesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCertificateIssuesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListCertificateIssuesRequest proto.InternalMessageInfo

func (m *ListCertificateIssuesRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ListCertificateIssuesRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListCertificateIssuesRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

type ListCertificateIssuesResponse struct {
	CertificateIssues    []*CertificateIssue `protobuf:"bytes,1,rep,name=certificate_issues,json=certificateIssues,proto3" json:"certificate_issues,omitempty"`
	NextPageToken        string              `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ListCertificateIssuesResponse) Reset()         { *m = ListCertificateIssuesResponse{} }
func (m *ListCertificateIssuesResponse) String() string { return proto.CompactTextString(m) }
func (*ListCertificateIssuesResponse) ProtoMessage()    {}
func (*ListCertificateIssuesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_58b50effdd0def26, []int{5}
}
func (m *ListCertificateIssuesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCertificateIssuesResponse.Unmarshal(m, b)
}
func (m *ListCertificateIssuesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCertificateIssuesResponse.Marshal(b, m, deterministic)
}
func (m *ListCertificateIssuesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCertificateIssuesResponse.Merge(m, src)
}
func (m *ListCertificateIssuesResponse) XXX_Size() int {
	return xxx_messageInfo_ListCertificateIssuesResponse.Size(m)
}
func (m *ListCertificateIssuesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCertificateIssuesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListCertificateIssuesResponse proto.InternalMessageInfo

func (m *ListCertificateIssuesResponse) GetCertificateIssues() []*CertificateIssue {
	if m != nil {
		return m.CertificateIssues
	}
	return nil
}

func (m *ListCertificateIssuesResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

type UpdateCertificateIssueRequest struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	UpdateMask           *types.FieldMask  `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	CertificateIssue     *CertificateIssue `protobuf:"bytes,3,opt,name=certificate_issue,json=certificateIssue,proto3" json:"certificate_issue,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *UpdateCertificateIssueRequest) Reset()         { *m = UpdateCertificateIssueRequest{} }
func (m *UpdateCertificateIssueRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateCertificateIssueRequest) ProtoMessage()    {}
func (*UpdateCertificateIssueRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_58b50effdd0def26, []int{6}
}
func (m *UpdateCertificateIssueRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateCertificateIssueRequest.Unmarshal(m, b)
}
func (m *UpdateCertificateIssueRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateCertificateIssueRequest.Marshal(b, m, deterministic)
}
func (m *UpdateCertificateIssueRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCertificateIssueRequest.Merge(m, src)
}
func (m *UpdateCertificateIssueRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateCertificateIssueRequest.Size(m)
}
func (m *UpdateCertificateIssueRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCertificateIssueRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCertificateIssueRequest proto.InternalMessageInfo

func (m *UpdateCertificateIssueRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateCertificateIssueRequest) GetUpdateMask() *types.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

func (m *UpdateCertificateIssueRequest) GetCertificateIssue() *CertificateIssue {
	if m != nil {
		return m.CertificateIssue
	}
	return nil
}

func init() {
	proto.RegisterType((*CertificateIssue)(nil), "powerssl.apiserver.v1.CertificateIssue")
	proto.RegisterMapType((map[string]string)(nil), "powerssl.apiserver.v1.CertificateIssue.LabelsEntry")
	proto.RegisterType((*CreateCertificateIssueRequest)(nil), "powerssl.apiserver.v1.CreateCertificateIssueRequest")
	proto.RegisterType((*DeleteCertificateIssueRequest)(nil), "powerssl.apiserver.v1.DeleteCertificateIssueRequest")
	proto.RegisterType((*GetCertificateIssueRequest)(nil), "powerssl.apiserver.v1.GetCertificateIssueRequest")
	proto.RegisterType((*ListCertificateIssuesRequest)(nil), "powerssl.apiserver.v1.ListCertificateIssuesRequest")
	proto.RegisterType((*ListCertificateIssuesResponse)(nil), "powerssl.apiserver.v1.ListCertificateIssuesResponse")
	proto.RegisterType((*UpdateCertificateIssueRequest)(nil), "powerssl.apiserver.v1.UpdateCertificateIssueRequest")
}

func init() {
	proto.RegisterFile("powerssl/apiserver/v1/certificate_issue.proto", fileDescriptor_58b50effdd0def26)
}

var fileDescriptor_58b50effdd0def26 = []byte{
	// 826 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0x4f, 0x6f, 0xeb, 0x44,
	0x10, 0x97, 0xf3, 0xef, 0x25, 0x93, 0x3e, 0xbd, 0xbc, 0x15, 0x3c, 0x8c, 0xdf, 0x8b, 0x30, 0x06,
	0x41, 0x08, 0xc2, 0x26, 0xed, 0x3b, 0x40, 0x2b, 0x0e, 0x50, 0x4a, 0x85, 0x28, 0x08, 0xb9, 0x85,
	0x03, 0x97, 0x68, 0x9b, 0x4c, 0xc3, 0xca, 0x8e, 0x6d, 0xbc, 0x9b, 0x94, 0x14, 0xf5, 0xd2, 0x13,
	0x97, 0x9e, 0xb8, 0x21, 0x21, 0xce, 0x7c, 0x11, 0xbe, 0x00, 0x5f, 0x81, 0x0f, 0x82, 0x76, 0xd7,
	0x49, 0x23, 0xa7, 0x4e, 0x13, 0x71, 0xf2, 0xee, 0xec, 0xcc, 0xf8, 0xe7, 0xdf, 0xfc, 0x66, 0xbc,
	0xf0, 0x41, 0x12, 0x5f, 0x62, 0xca, 0x79, 0xe8, 0xd1, 0x84, 0x71, 0x4c, 0xa7, 0x98, 0x7a, 0xd3,
	0x9e, 0x37, 0xc0, 0x54, 0xb0, 0x0b, 0x36, 0xa0, 0x02, 0xfb, 0x8c, 0xf3, 0x09, 0xba, 0x49, 0x1a,
	0x8b, 0x98, 0xbc, 0x3a, 0x77, 0x77, 0x17, 0xee, 0xee, 0xb4, 0x67, 0xbd, 0x18, 0xc5, 0xf1, 0x28,
	0x44, 0x99, 0xc3, 0xa3, 0x51, 0x14, 0x0b, 0x2a, 0x58, 0x1c, 0x71, 0x1d, 0x64, 0x3d, 0xcf, 0x4e,
	0xd5, 0xee, 0x7c, 0x72, 0xe1, 0xe1, 0x38, 0x11, 0xb3, 0xec, 0xd0, 0xce, 0x1f, 0x5e, 0x30, 0x0c,
	0x87, 0xfd, 0x31, 0xe5, 0x41, 0xe6, 0xf1, 0x46, 0xde, 0x43, 0xb0, 0x31, 0x72, 0x41, 0xc7, 0x89,
	0x76, 0x70, 0xfe, 0xaa, 0x40, 0xeb, 0xf0, 0x0e, 0xf0, 0x97, 0x12, 0x2f, 0x21, 0x50, 0x89, 0xe8,
	0x18, 0x4d, 0xc3, 0x36, 0x3a, 0x0d, 0x5f, 0xad, 0xc9, 0x01, 0x34, 0x07, 0x29, 0xca, 0x6f, 0x92,
	0x29, 0xcc, 0x92, 0x6d, 0x74, 0x9a, 0xbb, 0x96, 0xab, 0xf3, 0xbb, 0xf3, 0xfc, 0xee, 0xd9, 0x3c,
	0xbf, 0x0f, 0xda, 0x5d, 0x1a, 0x64, 0xf0, 0x24, 0x19, 0x2e, 0x82, 0xcb, 0x0f, 0x07, 0x6b, 0x77,
	0x15, 0xfc, 0x26, 0xec, 0x0c, 0x19, 0x4f, 0x42, 0x3a, 0xeb, 0x2b, 0x54, 0x15, 0x85, 0xaa, 0x99,
	0xd9, 0xbe, 0x91, 0xe0, 0x5e, 0x81, 0xaa, 0x60, 0x22, 0x44, 0xb3, 0xaa, 0xce, 0xf4, 0x86, 0xd8,
	0xd0, 0x1c, 0x22, 0x1f, 0xa4, 0x2c, 0x91, 0x8c, 0x9a, 0xb5, 0x2c, 0xee, 0xce, 0x44, 0xbe, 0x82,
	0x5a, 0x48, 0xcf, 0x31, 0xe4, 0xe6, 0x23, 0xbb, 0xdc, 0x69, 0xee, 0xee, 0xb9, 0xf7, 0xd6, 0xc8,
	0xcd, 0x33, 0xe4, 0x9e, 0xa8, 0xa8, 0xa3, 0x48, 0xa4, 0x33, 0x3f, 0x4b, 0x41, 0x2c, 0xa8, 0x0f,
	0x23, 0x2e, 0x21, 0x72, 0xb3, 0x6e, 0x97, 0x3b, 0x0d, 0x7f, 0xb1, 0x27, 0x6f, 0xc1, 0xe3, 0x00,
	0x67, 0x7d, 0x1a, 0x8e, 0xe2, 0x94, 0x89, 0x1f, 0xc7, 0x66, 0x43, 0x81, 0xd9, 0x09, 0x70, 0xf6,
	0xe9, 0xdc, 0x46, 0x5e, 0x87, 0xba, 0x74, 0xe2, 0xec, 0x0a, 0x4d, 0xb0, 0x8d, 0x4e, 0xd5, 0x7f,
	0x14, 0xe0, 0xec, 0x94, 0x5d, 0x21, 0x79, 0x0f, 0x5a, 0x43, 0x36, 0x42, 0x2e, 0x96, 0x52, 0x34,
	0x55, 0x8a, 0x27, 0xda, 0x7e, 0x97, 0xa5, 0x0d, 0x40, 0x27, 0x22, 0xee, 0xa7, 0x18, 0xe1, 0xa5,
	0xb9, 0x63, 0x1b, 0x9d, 0xba, 0xdf, 0x90, 0x16, 0x5f, 0x1a, 0xac, 0x8f, 0xa1, 0xb9, 0x04, 0x9e,
	0xb4, 0xa0, 0x1c, 0xe0, 0x2c, 0xab, 0xb4, 0x5c, 0x4a, 0x2e, 0xa7, 0x34, 0x9c, 0xe8, 0x12, 0x37,
	0x7c, 0xbd, 0xd9, 0x2f, 0x7d, 0x64, 0x38, 0xb7, 0x06, 0xb4, 0x0f, 0x55, 0x51, 0xf3, 0x7c, 0xf8,
	0xf8, 0xd3, 0x04, 0xb9, 0x20, 0xcf, 0xa0, 0x96, 0xd0, 0x14, 0x23, 0x91, 0x25, 0xcc, 0x76, 0xe4,
	0x0c, 0x9e, 0xae, 0x74, 0x45, 0x26, 0xa1, 0x77, 0x37, 0xa4, 0xdc, 0x6f, 0x0d, 0x72, 0x16, 0x67,
	0x0f, 0xda, 0x9f, 0x63, 0x88, 0xc5, 0x70, 0xee, 0xd1, 0xb1, 0xf3, 0x21, 0x58, 0xc7, 0x28, 0xb6,
	0x89, 0x48, 0xe1, 0xc5, 0x09, 0xe3, 0x2b, 0x21, 0xfc, 0xa1, 0x8f, 0x7e, 0x0e, 0x8d, 0x84, 0x8e,
	0x50, 0xd7, 0xb3, 0xa4, 0xea, 0x59, 0x97, 0x06, 0x55, 0xd0, 0x36, 0x80, 0x3a, 0x14, 0x71, 0x80,
	0x91, 0x6a, 0x88, 0x86, 0xaf, 0xdc, 0xcf, 0xa4, 0xc1, 0xf9, 0xd3, 0x80, 0x76, 0xc1, 0x4b, 0x79,
	0x12, 0x47, 0x1c, 0xc9, 0xf7, 0x40, 0x56, 0x28, 0xe5, 0xa6, 0xa1, 0x64, 0xbc, 0x31, 0xa7, 0x4f,
	0xf3, 0x9c, 0x72, 0xf2, 0x0e, 0x3c, 0x89, 0xf0, 0x67, 0xd1, 0x5f, 0x42, 0xa7, 0x85, 0xf0, 0x58,
	0x9a, 0xbf, 0x5d, 0x20, 0xfc, 0xdb, 0x80, 0xf6, 0x77, 0xaa, 0x49, 0xb7, 0xe0, 0x72, 0x69, 0x10,
	0xc8, 0x21, 0x55, 0x38, 0x45, 0xbe, 0x90, 0x73, 0xec, 0x6b, 0xca, 0x83, 0xf9, 0x20, 0x90, 0xeb,
	0xfb, 0x55, 0x54, 0xfe, 0x9f, 0x2a, 0xda, 0xfd, 0xb5, 0x06, 0xaf, 0xe5, 0xdd, 0x4e, 0x31, 0x9d,
	0xb2, 0x01, 0x92, 0xdf, 0x0d, 0xa8, 0x69, 0xc5, 0x93, 0x97, 0x45, 0x6f, 0x58, 0xd7, 0x10, 0xd6,
	0xa6, 0xb8, 0x9c, 0xde, 0xcd, 0x3f, 0xff, 0xfe, 0x56, 0x7a, 0xdf, 0x71, 0xe4, 0x1f, 0xe4, 0x17,
	0xad, 0xa0, 0x4f, 0x96, 0xb0, 0x72, 0xaf, 0x7b, 0xed, 0xe9, 0x22, 0xef, 0x57, 0xd5, 0x93, 0xdc,
	0x18, 0x50, 0xd3, 0xfa, 0x2f, 0x04, 0xb7, 0xb6, 0x3d, 0xac, 0x67, 0x2b, 0xbc, 0x1f, 0xc9, 0x9f,
	0x8b, 0xd3, 0x55, 0x58, 0xde, 0xee, 0x6a, 0x2c, 0xb2, 0x6e, 0x39, 0x24, 0x19, 0x10, 0xaf, 0x7b,
	0x4d, 0x6e, 0x0d, 0x28, 0x1f, 0xa3, 0x20, 0xbd, 0x02, 0x04, 0xc5, 0xbd, 0xb6, 0x39, 0x37, 0x19,
	0x1e, 0xb2, 0x09, 0x9e, 0x3f, 0x0c, 0xa8, 0xc8, 0xc6, 0x21, 0x45, 0xa3, 0x7c, 0x5d, 0x2b, 0x5b,
	0x2f, 0xb7, 0x0b, 0xd2, 0xad, 0x98, 0xc3, 0xb7, 0xb6, 0x76, 0x4a, 0x51, 0xba, 0x6d, 0x0a, 0x8b,
	0xb6, 0xb6, 0xab, 0xb6, 0x56, 0x94, 0xb5, 0x01, 0x6b, 0x99, 0xa2, 0x3e, 0xf3, 0x7e, 0x58, 0x5c,
	0x69, 0x5c, 0x16, 0x7b, 0x2c, 0x12, 0x98, 0x46, 0x34, 0xf4, 0x92, 0x60, 0xb4, 0x74, 0xc5, 0x91,
	0x17, 0x95, 0x69, 0xef, 0x80, 0x26, 0xec, 0xbc, 0xa6, 0x94, 0xb3, 0xf7, 0x5f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x94, 0x30, 0x15, 0xa9, 0x0a, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CertificateIssueServiceClient is the client API for CertificateIssueService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CertificateIssueServiceClient interface {
	// Creates a certificate, and returns the new CertificateIssue.
	Create(ctx context.Context, in *CreateCertificateIssueRequest, opts ...grpc.CallOption) (*CertificateIssue, error)
	// Deletes a certificate. Returns NOT_FOUND if the certificate does not exist.
	Delete(ctx context.Context, in *DeleteCertificateIssueRequest, opts ...grpc.CallOption) (*types.Empty, error)
	// Gets a certificate. Returns NOT_FOUND if the certificate does not exist.
	Get(ctx context.Context, in *GetCertificateIssueRequest, opts ...grpc.CallOption) (*CertificateIssue, error)
	// Lists certificates. The order is unspecified but deterministic. Newly
	// created certificates will not necessarily be added to the end of this list.
	List(ctx context.Context, in *ListCertificateIssuesRequest, opts ...grpc.CallOption) (*ListCertificateIssuesResponse, error)
	// Updates a certificate. Returns INVALID_ARGUMENT if the name of the
	// certificate is non-empty and does equal the previous name.
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
	err := c.cc.Invoke(ctx, "/powerssl.apiserver.v1.CertificateIssueService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateIssueServiceClient) Delete(ctx context.Context, in *DeleteCertificateIssueRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/powerssl.apiserver.v1.CertificateIssueService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateIssueServiceClient) Get(ctx context.Context, in *GetCertificateIssueRequest, opts ...grpc.CallOption) (*CertificateIssue, error) {
	out := new(CertificateIssue)
	err := c.cc.Invoke(ctx, "/powerssl.apiserver.v1.CertificateIssueService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateIssueServiceClient) List(ctx context.Context, in *ListCertificateIssuesRequest, opts ...grpc.CallOption) (*ListCertificateIssuesResponse, error) {
	out := new(ListCertificateIssuesResponse)
	err := c.cc.Invoke(ctx, "/powerssl.apiserver.v1.CertificateIssueService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateIssueServiceClient) Update(ctx context.Context, in *UpdateCertificateIssueRequest, opts ...grpc.CallOption) (*CertificateIssue, error) {
	out := new(CertificateIssue)
	err := c.cc.Invoke(ctx, "/powerssl.apiserver.v1.CertificateIssueService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CertificateIssueServiceServer is the server API for CertificateIssueService service.
type CertificateIssueServiceServer interface {
	// Creates a certificate, and returns the new CertificateIssue.
	Create(context.Context, *CreateCertificateIssueRequest) (*CertificateIssue, error)
	// Deletes a certificate. Returns NOT_FOUND if the certificate does not exist.
	Delete(context.Context, *DeleteCertificateIssueRequest) (*types.Empty, error)
	// Gets a certificate. Returns NOT_FOUND if the certificate does not exist.
	Get(context.Context, *GetCertificateIssueRequest) (*CertificateIssue, error)
	// Lists certificates. The order is unspecified but deterministic. Newly
	// created certificates will not necessarily be added to the end of this list.
	List(context.Context, *ListCertificateIssuesRequest) (*ListCertificateIssuesResponse, error)
	// Updates a certificate. Returns INVALID_ARGUMENT if the name of the
	// certificate is non-empty and does equal the previous name.
	Update(context.Context, *UpdateCertificateIssueRequest) (*CertificateIssue, error)
}

// UnimplementedCertificateIssueServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCertificateIssueServiceServer struct {
}

func (*UnimplementedCertificateIssueServiceServer) Create(ctx context.Context, req *CreateCertificateIssueRequest) (*CertificateIssue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedCertificateIssueServiceServer) Delete(ctx context.Context, req *DeleteCertificateIssueRequest) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedCertificateIssueServiceServer) Get(ctx context.Context, req *GetCertificateIssueRequest) (*CertificateIssue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedCertificateIssueServiceServer) List(ctx context.Context, req *ListCertificateIssuesRequest) (*ListCertificateIssuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedCertificateIssueServiceServer) Update(ctx context.Context, req *UpdateCertificateIssueRequest) (*CertificateIssue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
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
		FullMethod: "/powerssl.apiserver.v1.CertificateIssueService/Create",
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
		FullMethod: "/powerssl.apiserver.v1.CertificateIssueService/Delete",
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
		FullMethod: "/powerssl.apiserver.v1.CertificateIssueService/Get",
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
		FullMethod: "/powerssl.apiserver.v1.CertificateIssueService/List",
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
		FullMethod: "/powerssl.apiserver.v1.CertificateIssueService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateIssueServiceServer).Update(ctx, req.(*UpdateCertificateIssueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CertificateIssueService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "powerssl.apiserver.v1.CertificateIssueService",
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
	Metadata: "powerssl/apiserver/v1/certificate_issue.proto",
}
