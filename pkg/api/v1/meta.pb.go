// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/v1/meta.proto

package api // import "powerssl.io/pkg/api/v1"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import types "github.com/gogo/protobuf/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ListMeta struct {
	ResourceVersion      string   `protobuf:"bytes,1,opt,name=resource_version,json=resourceVersion,proto3" json:"resource_version,omitempty"`
	Continue             string   `protobuf:"bytes,2,opt,name=continue,proto3" json:"continue,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListMeta) Reset()         { *m = ListMeta{} }
func (m *ListMeta) String() string { return proto.CompactTextString(m) }
func (*ListMeta) ProtoMessage()    {}
func (*ListMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_meta_296d89e4b963e72b, []int{0}
}
func (m *ListMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMeta.Unmarshal(m, b)
}
func (m *ListMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMeta.Marshal(b, m, deterministic)
}
func (dst *ListMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMeta.Merge(dst, src)
}
func (m *ListMeta) XXX_Size() int {
	return xxx_messageInfo_ListMeta.Size(m)
}
func (m *ListMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMeta.DiscardUnknown(m)
}

var xxx_messageInfo_ListMeta proto.InternalMessageInfo

func (m *ListMeta) GetResourceVersion() string {
	if m != nil {
		return m.ResourceVersion
	}
	return ""
}

func (m *ListMeta) GetContinue() string {
	if m != nil {
		return m.Continue
	}
	return ""
}

type ObjectMeta struct {
	ResourceVersion      string            `protobuf:"bytes,1,opt,name=resource_version,json=resourceVersion,proto3" json:"resource_version,omitempty"`
	Name                 string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Uid                  string            `protobuf:"bytes,3,opt,name=uid,proto3" json:"uid,omitempty"`
	CreationTimestamp    *types.Timestamp  `protobuf:"bytes,4,opt,name=creation_timestamp,json=creationTimestamp" json:"creation_timestamp,omitempty"`
	DeletionTimestamp    *types.Timestamp  `protobuf:"bytes,5,opt,name=deletion_timestamp,json=deletionTimestamp" json:"deletion_timestamp,omitempty"`
	Labels               map[string]string `protobuf:"bytes,6,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ObjectMeta) Reset()         { *m = ObjectMeta{} }
func (m *ObjectMeta) String() string { return proto.CompactTextString(m) }
func (*ObjectMeta) ProtoMessage()    {}
func (*ObjectMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_meta_296d89e4b963e72b, []int{1}
}
func (m *ObjectMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectMeta.Unmarshal(m, b)
}
func (m *ObjectMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectMeta.Marshal(b, m, deterministic)
}
func (dst *ObjectMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectMeta.Merge(dst, src)
}
func (m *ObjectMeta) XXX_Size() int {
	return xxx_messageInfo_ObjectMeta.Size(m)
}
func (m *ObjectMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectMeta.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectMeta proto.InternalMessageInfo

func (m *ObjectMeta) GetResourceVersion() string {
	if m != nil {
		return m.ResourceVersion
	}
	return ""
}

func (m *ObjectMeta) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ObjectMeta) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ObjectMeta) GetCreationTimestamp() *types.Timestamp {
	if m != nil {
		return m.CreationTimestamp
	}
	return nil
}

func (m *ObjectMeta) GetDeletionTimestamp() *types.Timestamp {
	if m != nil {
		return m.DeletionTimestamp
	}
	return nil
}

func (m *ObjectMeta) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type TypeMeta struct {
	ApiVersion           string   `protobuf:"bytes,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	Kind                 string   `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TypeMeta) Reset()         { *m = TypeMeta{} }
func (m *TypeMeta) String() string { return proto.CompactTextString(m) }
func (*TypeMeta) ProtoMessage()    {}
func (*TypeMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_meta_296d89e4b963e72b, []int{2}
}
func (m *TypeMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TypeMeta.Unmarshal(m, b)
}
func (m *TypeMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TypeMeta.Marshal(b, m, deterministic)
}
func (dst *TypeMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TypeMeta.Merge(dst, src)
}
func (m *TypeMeta) XXX_Size() int {
	return xxx_messageInfo_TypeMeta.Size(m)
}
func (m *TypeMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_TypeMeta.DiscardUnknown(m)
}

var xxx_messageInfo_TypeMeta proto.InternalMessageInfo

func (m *TypeMeta) GetApiVersion() string {
	if m != nil {
		return m.ApiVersion
	}
	return ""
}

func (m *TypeMeta) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func init() {
	proto.RegisterType((*ListMeta)(nil), "powerssl.api.v1.ListMeta")
	proto.RegisterType((*ObjectMeta)(nil), "powerssl.api.v1.ObjectMeta")
	proto.RegisterMapType((map[string]string)(nil), "powerssl.api.v1.ObjectMeta.LabelsEntry")
	proto.RegisterType((*TypeMeta)(nil), "powerssl.api.v1.TypeMeta")
}

func init() { proto.RegisterFile("api/v1/meta.proto", fileDescriptor_meta_296d89e4b963e72b) }

var fileDescriptor_meta_296d89e4b963e72b = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4d, 0x6b, 0xf2, 0x40,
	0x14, 0x85, 0x89, 0x5f, 0xf8, 0xde, 0x2c, 0xd4, 0xe1, 0x5d, 0x84, 0x50, 0x50, 0xdc, 0xd4, 0x6e,
	0x26, 0x68, 0x37, 0xfd, 0x58, 0x08, 0x85, 0x2e, 0x0a, 0x96, 0x52, 0x91, 0x2e, 0xba, 0x91, 0x49,
	0xbc, 0x95, 0xa9, 0x49, 0x66, 0xc8, 0x4c, 0x52, 0xfc, 0xd5, 0xfd, 0x0b, 0x25, 0x13, 0x27, 0x8a,
	0x9b, 0xd2, 0xdd, 0xbd, 0x27, 0xf7, 0x1c, 0x72, 0x9e, 0x04, 0x06, 0x4c, 0xf2, 0xa0, 0x98, 0x06,
	0x09, 0x6a, 0x46, 0x65, 0x26, 0xb4, 0x20, 0x3d, 0x29, 0xbe, 0x30, 0x53, 0x2a, 0xa6, 0x4c, 0x72,
	0x5a, 0x4c, 0xfd, 0xe1, 0x56, 0x88, 0x6d, 0x8c, 0x81, 0x79, 0x1c, 0xe6, 0x1f, 0x81, 0xe6, 0x09,
	0x2a, 0xcd, 0x12, 0x59, 0x39, 0xc6, 0xaf, 0xd0, 0x5d, 0x70, 0xa5, 0x9f, 0x51, 0x33, 0x72, 0x05,
	0xfd, 0x0c, 0x95, 0xc8, 0xb3, 0x08, 0xd7, 0x05, 0x66, 0x8a, 0x8b, 0xd4, 0x73, 0x46, 0xce, 0xe4,
	0xdf, 0xb2, 0x67, 0xf5, 0xb7, 0x4a, 0x26, 0x3e, 0x74, 0x23, 0x91, 0x6a, 0x9e, 0xe6, 0xe8, 0x35,
	0xcc, 0x49, 0xbd, 0x8f, 0xbf, 0x1b, 0x00, 0x2f, 0xe1, 0x27, 0x46, 0x7f, 0x4e, 0x25, 0xd0, 0x4a,
	0x59, 0x62, 0x13, 0xcd, 0x4c, 0xfa, 0xd0, 0xcc, 0xf9, 0xc6, 0x6b, 0x1a, 0xa9, 0x1c, 0xc9, 0x13,
	0x90, 0x28, 0x43, 0xa6, 0xb9, 0x48, 0xd7, 0x75, 0x1d, 0xaf, 0x35, 0x72, 0x26, 0xee, 0xcc, 0xa7,
	0x55, 0x61, 0x6a, 0x0b, 0xd3, 0x95, 0xbd, 0x58, 0x0e, 0xac, 0xab, 0x96, 0xca, 0xa8, 0x0d, 0xc6,
	0x78, 0x16, 0xd5, 0xfe, 0x3d, 0xca, 0xba, 0x8e, 0x51, 0x73, 0xe8, 0xc4, 0x2c, 0xc4, 0x58, 0x79,
	0x9d, 0x51, 0x73, 0xe2, 0xce, 0x2e, 0xe9, 0xd9, 0xb7, 0xa0, 0x47, 0x26, 0x74, 0x61, 0x2e, 0x1f,
	0x53, 0x9d, 0xed, 0x97, 0x07, 0x9b, 0x7f, 0x0b, 0xee, 0x89, 0x5c, 0xf6, 0xde, 0xe1, 0xfe, 0x40,
	0xaa, 0x1c, 0xc9, 0x7f, 0x68, 0x17, 0x2c, 0xae, 0x81, 0x57, 0xcb, 0x5d, 0xe3, 0xc6, 0x19, 0xcf,
	0xa1, 0xbb, 0xda, 0x4b, 0x34, 0xb8, 0x87, 0xe0, 0x32, 0xc9, 0xcf, 0x48, 0x03, 0x93, 0xfc, 0x04,
	0xf2, 0x8e, 0xa7, 0x1b, 0x0b, 0xb9, 0x9c, 0x1f, 0x2e, 0xde, 0xfd, 0xfa, 0x6d, 0xb9, 0x08, 0xe4,
	0x6e, 0x1b, 0x54, 0x3f, 0xd7, 0x3d, 0x93, 0x3c, 0xec, 0x18, 0x02, 0xd7, 0x3f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x87, 0x8e, 0x93, 0xf0, 0x71, 0x02, 0x00, 0x00,
}
