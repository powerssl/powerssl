// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: powerssl/apiserver/v1/certificate_issue.proto

package api

import (
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// A CertificateIssue ...
type CertificateIssue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the certificate issue.
	// CertificateIssue names have the form
	// `certificates/{certificate_id}/issues/{issue_id}`. The name is ignored when
	// creating a certificate issue.
	Name            string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CreateTime      *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime      *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	DisplayName     string                 `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Title           string                 `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	Description     string                 `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	Labels          map[string]string      `protobuf:"bytes,7,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Dnsnames        []string               `protobuf:"bytes,8,rep,name=dnsnames,proto3" json:"dnsnames,omitempty"`
	KeyAlgorithm    string                 `protobuf:"bytes,9,opt,name=key_algorithm,json=keyAlgorithm,proto3" json:"key_algorithm,omitempty"`
	KeySize         int32                  `protobuf:"varint,10,opt,name=key_size,json=keySize,proto3" json:"key_size,omitempty"`
	DigestAlgorithm string                 `protobuf:"bytes,11,opt,name=digest_algorithm,json=digestAlgorithm,proto3" json:"digest_algorithm,omitempty"`
	AutoRenew       bool                   `protobuf:"varint,12,opt,name=auto_renew,json=autoRenew,proto3" json:"auto_renew,omitempty"`
}

func (x *CertificateIssue) Reset() {
	*x = CertificateIssue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_powerssl_apiserver_v1_certificate_issue_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CertificateIssue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertificateIssue) ProtoMessage() {}

func (x *CertificateIssue) ProtoReflect() protoreflect.Message {
	mi := &file_powerssl_apiserver_v1_certificate_issue_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertificateIssue.ProtoReflect.Descriptor instead.
func (*CertificateIssue) Descriptor() ([]byte, []int) {
	return file_powerssl_apiserver_v1_certificate_issue_proto_rawDescGZIP(), []int{0}
}

func (x *CertificateIssue) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CertificateIssue) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *CertificateIssue) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *CertificateIssue) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *CertificateIssue) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CertificateIssue) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CertificateIssue) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *CertificateIssue) GetDnsnames() []string {
	if x != nil {
		return x.Dnsnames
	}
	return nil
}

func (x *CertificateIssue) GetKeyAlgorithm() string {
	if x != nil {
		return x.KeyAlgorithm
	}
	return ""
}

func (x *CertificateIssue) GetKeySize() int32 {
	if x != nil {
		return x.KeySize
	}
	return 0
}

func (x *CertificateIssue) GetDigestAlgorithm() string {
	if x != nil {
		return x.DigestAlgorithm
	}
	return ""
}

func (x *CertificateIssue) GetAutoRenew() bool {
	if x != nil {
		return x.AutoRenew
	}
	return false
}

type CreateCertificateIssueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parent           string            `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	CertificateIssue *CertificateIssue `protobuf:"bytes,2,opt,name=certificate_issue,json=certificateIssue,proto3" json:"certificate_issue,omitempty"`
}

func (x *CreateCertificateIssueRequest) Reset() {
	*x = CreateCertificateIssueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_powerssl_apiserver_v1_certificate_issue_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCertificateIssueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCertificateIssueRequest) ProtoMessage() {}

func (x *CreateCertificateIssueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_powerssl_apiserver_v1_certificate_issue_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCertificateIssueRequest.ProtoReflect.Descriptor instead.
func (*CreateCertificateIssueRequest) Descriptor() ([]byte, []int) {
	return file_powerssl_apiserver_v1_certificate_issue_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCertificateIssueRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *CreateCertificateIssueRequest) GetCertificateIssue() *CertificateIssue {
	if x != nil {
		return x.CertificateIssue
	}
	return nil
}

type DeleteCertificateIssueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteCertificateIssueRequest) Reset() {
	*x = DeleteCertificateIssueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_powerssl_apiserver_v1_certificate_issue_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCertificateIssueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCertificateIssueRequest) ProtoMessage() {}

func (x *DeleteCertificateIssueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_powerssl_apiserver_v1_certificate_issue_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCertificateIssueRequest.ProtoReflect.Descriptor instead.
func (*DeleteCertificateIssueRequest) Descriptor() ([]byte, []int) {
	return file_powerssl_apiserver_v1_certificate_issue_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteCertificateIssueRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetCertificateIssueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetCertificateIssueRequest) Reset() {
	*x = GetCertificateIssueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_powerssl_apiserver_v1_certificate_issue_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCertificateIssueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCertificateIssueRequest) ProtoMessage() {}

func (x *GetCertificateIssueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_powerssl_apiserver_v1_certificate_issue_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCertificateIssueRequest.ProtoReflect.Descriptor instead.
func (*GetCertificateIssueRequest) Descriptor() ([]byte, []int) {
	return file_powerssl_apiserver_v1_certificate_issue_proto_rawDescGZIP(), []int{3}
}

func (x *GetCertificateIssueRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListCertificateIssuesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parent    string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	PageSize  int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListCertificateIssuesRequest) Reset() {
	*x = ListCertificateIssuesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_powerssl_apiserver_v1_certificate_issue_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCertificateIssuesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCertificateIssuesRequest) ProtoMessage() {}

func (x *ListCertificateIssuesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_powerssl_apiserver_v1_certificate_issue_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCertificateIssuesRequest.ProtoReflect.Descriptor instead.
func (*ListCertificateIssuesRequest) Descriptor() ([]byte, []int) {
	return file_powerssl_apiserver_v1_certificate_issue_proto_rawDescGZIP(), []int{4}
}

func (x *ListCertificateIssuesRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListCertificateIssuesRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListCertificateIssuesRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListCertificateIssuesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CertificateIssues []*CertificateIssue `protobuf:"bytes,1,rep,name=certificate_issues,json=certificateIssues,proto3" json:"certificate_issues,omitempty"`
	NextPageToken     string              `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListCertificateIssuesResponse) Reset() {
	*x = ListCertificateIssuesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_powerssl_apiserver_v1_certificate_issue_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCertificateIssuesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCertificateIssuesResponse) ProtoMessage() {}

func (x *ListCertificateIssuesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_powerssl_apiserver_v1_certificate_issue_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCertificateIssuesResponse.ProtoReflect.Descriptor instead.
func (*ListCertificateIssuesResponse) Descriptor() ([]byte, []int) {
	return file_powerssl_apiserver_v1_certificate_issue_proto_rawDescGZIP(), []int{5}
}

func (x *ListCertificateIssuesResponse) GetCertificateIssues() []*CertificateIssue {
	if x != nil {
		return x.CertificateIssues
	}
	return nil
}

func (x *ListCertificateIssuesResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type UpdateCertificateIssueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name             string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	UpdateMask       *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	CertificateIssue *CertificateIssue      `protobuf:"bytes,3,opt,name=certificate_issue,json=certificateIssue,proto3" json:"certificate_issue,omitempty"`
}

func (x *UpdateCertificateIssueRequest) Reset() {
	*x = UpdateCertificateIssueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_powerssl_apiserver_v1_certificate_issue_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCertificateIssueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCertificateIssueRequest) ProtoMessage() {}

func (x *UpdateCertificateIssueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_powerssl_apiserver_v1_certificate_issue_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCertificateIssueRequest.ProtoReflect.Descriptor instead.
func (*UpdateCertificateIssueRequest) Descriptor() ([]byte, []int) {
	return file_powerssl_apiserver_v1_certificate_issue_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateCertificateIssueRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateCertificateIssueRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

func (x *UpdateCertificateIssueRequest) GetCertificateIssue() *CertificateIssue {
	if x != nil {
		return x.CertificateIssue
	}
	return nil
}

var File_powerssl_apiserver_v1_certificate_issue_proto protoreflect.FileDescriptor

var file_powerssl_apiserver_v1_certificate_issue_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x73, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x5f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x15, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x73, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa9, 0x04, 0x0a, 0x10, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x49, 0x73, 0x73, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a,
	0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x4b, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x33, 0x2e, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x73, 0x6c, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x49, 0x73, 0x73, 0x75, 0x65, 0x2e, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x6e, 0x73, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x08, 0x64, 0x6e, 0x73, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x6b,
	0x65, 0x79, 0x5f, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x6b, 0x65, 0x79, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d,
	0x12, 0x19, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x64,
	0x69, 0x67, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x41, 0x6c, 0x67,
	0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x72,
	0x65, 0x6e, 0x65, 0x77, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x61, 0x75, 0x74, 0x6f,
	0x52, 0x65, 0x6e, 0x65, 0x77, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x8d, 0x01, 0x0a, 0x1d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x54, 0x0a, 0x11, 0x63, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x73, 0x6c,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x10,
	0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x49, 0x73, 0x73, 0x75, 0x65,
	0x22, 0x33, 0x0a, 0x1d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x30, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x43, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x72, 0x0a, 0x1c, 0x4c, 0x69, 0x73, 0x74, 0x43,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x49, 0x73, 0x73, 0x75, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x9f, 0x01, 0x0a, 0x1d,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x49,
	0x73, 0x73, 0x75, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a,
	0x12, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x73, 0x73,
	0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x70, 0x6f, 0x77, 0x65,
	0x72, 0x73, 0x73, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x49, 0x73, 0x73,
	0x75, 0x65, 0x52, 0x11, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x49,
	0x73, 0x73, 0x75, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xc6, 0x01,
	0x0a, 0x1d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61,
	0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b,
	0x12, 0x54, 0x0a, 0x11, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f,
	0x69, 0x73, 0x73, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x70, 0x6f,
	0x77, 0x65, 0x72, 0x73, 0x73, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x49,
	0x73, 0x73, 0x75, 0x65, 0x52, 0x10, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x49, 0x73, 0x73, 0x75, 0x65, 0x32, 0x80, 0x06, 0x0a, 0x17, 0x43, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x49, 0x73, 0x73, 0x75, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x96, 0x01, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x34, 0x2e,
	0x70, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x73, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x73, 0x6c, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x49, 0x73, 0x73, 0x75, 0x65, 0x22, 0x2d, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x27, 0x22, 0x22, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x3d, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x2a,
	0x7d, 0x2f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x82, 0x01, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x34, 0x2e, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x73,
	0x6c, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x2a, 0x22, 0x2f, 0x76,
	0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x73, 0x2f, 0x2a, 0x2f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x73, 0x2f, 0x2a, 0x7d,
	0x12, 0x8d, 0x01, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x31, 0x2e, 0x70, 0x6f, 0x77, 0x65, 0x72,
	0x73, 0x73, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x49,
	0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x70, 0x6f,
	0x77, 0x65, 0x72, 0x73, 0x73, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x49,
	0x73, 0x73, 0x75, 0x65, 0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x12, 0x22, 0x2f, 0x76,
	0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x73, 0x2f, 0x2a, 0x2f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x73, 0x2f, 0x2a, 0x7d,
	0x12, 0x9d, 0x01, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x33, 0x2e, 0x70, 0x6f, 0x77, 0x65,
	0x72, 0x73, 0x73, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x49, 0x73, 0x73, 0x75, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34,
	0x2e, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x73, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x49, 0x73, 0x73, 0x75, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x12, 0x22, 0x2f, 0x76,
	0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x73,
	0x12, 0x96, 0x01, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x34, 0x2e, 0x70, 0x6f,
	0x77, 0x65, 0x72, 0x73, 0x73, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x27, 0x2e, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x73, 0x6c, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x49, 0x73, 0x73, 0x75, 0x65, 0x22, 0x2d, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x27, 0x1a, 0x22, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x63, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x2a, 0x2f, 0x69, 0x73, 0x73,
	0x75, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x3a, 0x01, 0x2a, 0x42, 0x23, 0x5a, 0x21, 0x70, 0x6f, 0x77,
	0x65, 0x72, 0x73, 0x73, 0x6c, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70,
	0x69, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_powerssl_apiserver_v1_certificate_issue_proto_rawDescOnce sync.Once
	file_powerssl_apiserver_v1_certificate_issue_proto_rawDescData = file_powerssl_apiserver_v1_certificate_issue_proto_rawDesc
)

func file_powerssl_apiserver_v1_certificate_issue_proto_rawDescGZIP() []byte {
	file_powerssl_apiserver_v1_certificate_issue_proto_rawDescOnce.Do(func() {
		file_powerssl_apiserver_v1_certificate_issue_proto_rawDescData = protoimpl.X.CompressGZIP(file_powerssl_apiserver_v1_certificate_issue_proto_rawDescData)
	})
	return file_powerssl_apiserver_v1_certificate_issue_proto_rawDescData
}

var file_powerssl_apiserver_v1_certificate_issue_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_powerssl_apiserver_v1_certificate_issue_proto_goTypes = []interface{}{
	(*CertificateIssue)(nil),              // 0: powerssl.apiserver.v1.CertificateIssue
	(*CreateCertificateIssueRequest)(nil), // 1: powerssl.apiserver.v1.CreateCertificateIssueRequest
	(*DeleteCertificateIssueRequest)(nil), // 2: powerssl.apiserver.v1.DeleteCertificateIssueRequest
	(*GetCertificateIssueRequest)(nil),    // 3: powerssl.apiserver.v1.GetCertificateIssueRequest
	(*ListCertificateIssuesRequest)(nil),  // 4: powerssl.apiserver.v1.ListCertificateIssuesRequest
	(*ListCertificateIssuesResponse)(nil), // 5: powerssl.apiserver.v1.ListCertificateIssuesResponse
	(*UpdateCertificateIssueRequest)(nil), // 6: powerssl.apiserver.v1.UpdateCertificateIssueRequest
	nil,                                   // 7: powerssl.apiserver.v1.CertificateIssue.LabelsEntry
	(*timestamppb.Timestamp)(nil),         // 8: google.protobuf.Timestamp
	(*fieldmaskpb.FieldMask)(nil),         // 9: google.protobuf.FieldMask
	(*emptypb.Empty)(nil),                 // 10: google.protobuf.Empty
}
var file_powerssl_apiserver_v1_certificate_issue_proto_depIdxs = []int32{
	8,  // 0: powerssl.apiserver.v1.CertificateIssue.create_time:type_name -> google.protobuf.Timestamp
	8,  // 1: powerssl.apiserver.v1.CertificateIssue.update_time:type_name -> google.protobuf.Timestamp
	7,  // 2: powerssl.apiserver.v1.CertificateIssue.labels:type_name -> powerssl.apiserver.v1.CertificateIssue.LabelsEntry
	0,  // 3: powerssl.apiserver.v1.CreateCertificateIssueRequest.certificate_issue:type_name -> powerssl.apiserver.v1.CertificateIssue
	0,  // 4: powerssl.apiserver.v1.ListCertificateIssuesResponse.certificate_issues:type_name -> powerssl.apiserver.v1.CertificateIssue
	9,  // 5: powerssl.apiserver.v1.UpdateCertificateIssueRequest.update_mask:type_name -> google.protobuf.FieldMask
	0,  // 6: powerssl.apiserver.v1.UpdateCertificateIssueRequest.certificate_issue:type_name -> powerssl.apiserver.v1.CertificateIssue
	1,  // 7: powerssl.apiserver.v1.CertificateIssueService.Create:input_type -> powerssl.apiserver.v1.CreateCertificateIssueRequest
	2,  // 8: powerssl.apiserver.v1.CertificateIssueService.Delete:input_type -> powerssl.apiserver.v1.DeleteCertificateIssueRequest
	3,  // 9: powerssl.apiserver.v1.CertificateIssueService.Get:input_type -> powerssl.apiserver.v1.GetCertificateIssueRequest
	4,  // 10: powerssl.apiserver.v1.CertificateIssueService.List:input_type -> powerssl.apiserver.v1.ListCertificateIssuesRequest
	6,  // 11: powerssl.apiserver.v1.CertificateIssueService.Update:input_type -> powerssl.apiserver.v1.UpdateCertificateIssueRequest
	0,  // 12: powerssl.apiserver.v1.CertificateIssueService.Create:output_type -> powerssl.apiserver.v1.CertificateIssue
	10, // 13: powerssl.apiserver.v1.CertificateIssueService.Delete:output_type -> google.protobuf.Empty
	0,  // 14: powerssl.apiserver.v1.CertificateIssueService.Get:output_type -> powerssl.apiserver.v1.CertificateIssue
	5,  // 15: powerssl.apiserver.v1.CertificateIssueService.List:output_type -> powerssl.apiserver.v1.ListCertificateIssuesResponse
	0,  // 16: powerssl.apiserver.v1.CertificateIssueService.Update:output_type -> powerssl.apiserver.v1.CertificateIssue
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_powerssl_apiserver_v1_certificate_issue_proto_init() }
func file_powerssl_apiserver_v1_certificate_issue_proto_init() {
	if File_powerssl_apiserver_v1_certificate_issue_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_powerssl_apiserver_v1_certificate_issue_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CertificateIssue); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_powerssl_apiserver_v1_certificate_issue_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCertificateIssueRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_powerssl_apiserver_v1_certificate_issue_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCertificateIssueRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_powerssl_apiserver_v1_certificate_issue_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCertificateIssueRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_powerssl_apiserver_v1_certificate_issue_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCertificateIssuesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_powerssl_apiserver_v1_certificate_issue_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCertificateIssuesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_powerssl_apiserver_v1_certificate_issue_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCertificateIssueRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_powerssl_apiserver_v1_certificate_issue_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_powerssl_apiserver_v1_certificate_issue_proto_goTypes,
		DependencyIndexes: file_powerssl_apiserver_v1_certificate_issue_proto_depIdxs,
		MessageInfos:      file_powerssl_apiserver_v1_certificate_issue_proto_msgTypes,
	}.Build()
	File_powerssl_apiserver_v1_certificate_issue_proto = out.File
	file_powerssl_apiserver_v1_certificate_issue_proto_rawDesc = nil
	file_powerssl_apiserver_v1_certificate_issue_proto_goTypes = nil
	file_powerssl_apiserver_v1_certificate_issue_proto_depIdxs = nil
}