// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: controller/api/v1/meta.proto

package api // import "powerssl.io/pkg/controller/api/v1"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Activity_Name int32

const (
	Activity_NAME_UNSPECIFIED       Activity_Name = 0
	Activity_CA_AUTHORIZE_DOMAIN    Activity_Name = 101
	Activity_CA_REQUEST_CERTIFICATE Activity_Name = 102
	Activity_CA_REVOKE_CERTIFICATE  Activity_Name = 103
	Activity_CA_VERIFY_DOMAIN       Activity_Name = 104
	Activity_DNS_CREATE_RECORD      Activity_Name = 201
	Activity_DNS_DELETE_RECORD      Activity_Name = 202
	Activity_DNS_VERIFY_DOMAIN      Activity_Name = 203
)

var Activity_Name_name = map[int32]string{
	0:   "NAME_UNSPECIFIED",
	101: "CA_AUTHORIZE_DOMAIN",
	102: "CA_REQUEST_CERTIFICATE",
	103: "CA_REVOKE_CERTIFICATE",
	104: "CA_VERIFY_DOMAIN",
	201: "DNS_CREATE_RECORD",
	202: "DNS_DELETE_RECORD",
	203: "DNS_VERIFY_DOMAIN",
}
var Activity_Name_value = map[string]int32{
	"NAME_UNSPECIFIED":       0,
	"CA_AUTHORIZE_DOMAIN":    101,
	"CA_REQUEST_CERTIFICATE": 102,
	"CA_REVOKE_CERTIFICATE":  103,
	"CA_VERIFY_DOMAIN":       104,
	"DNS_CREATE_RECORD":      201,
	"DNS_DELETE_RECORD":      202,
	"DNS_VERIFY_DOMAIN":      203,
}

func (x Activity_Name) String() string {
	return proto.EnumName(Activity_Name_name, int32(x))
}
func (Activity_Name) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_meta_b03dac0b310a1143, []int{0, 0}
}

type Activity struct {
	Token                string             `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Name                 Activity_Name      `protobuf:"varint,2,opt,name=name,proto3,enum=powerssl.controller.v1.Activity_Name" json:"name,omitempty"`
	Workflow             *Activity_Workflow `protobuf:"bytes,3,opt,name=workflow" json:"workflow,omitempty"`
	Signature            string             `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Activity) Reset()         { *m = Activity{} }
func (m *Activity) String() string { return proto.CompactTextString(m) }
func (*Activity) ProtoMessage()    {}
func (*Activity) Descriptor() ([]byte, []int) {
	return fileDescriptor_meta_b03dac0b310a1143, []int{0}
}
func (m *Activity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Activity.Unmarshal(m, b)
}
func (m *Activity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Activity.Marshal(b, m, deterministic)
}
func (dst *Activity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Activity.Merge(dst, src)
}
func (m *Activity) XXX_Size() int {
	return xxx_messageInfo_Activity.Size(m)
}
func (m *Activity) XXX_DiscardUnknown() {
	xxx_messageInfo_Activity.DiscardUnknown(m)
}

var xxx_messageInfo_Activity proto.InternalMessageInfo

func (m *Activity) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Activity) GetName() Activity_Name {
	if m != nil {
		return m.Name
	}
	return Activity_NAME_UNSPECIFIED
}

func (m *Activity) GetWorkflow() *Activity_Workflow {
	if m != nil {
		return m.Workflow
	}
	return nil
}

func (m *Activity) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

type Activity_Workflow struct {
	Activities           []string `protobuf:"bytes,1,rep,name=activities" json:"activities,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Activity_Workflow) Reset()         { *m = Activity_Workflow{} }
func (m *Activity_Workflow) String() string { return proto.CompactTextString(m) }
func (*Activity_Workflow) ProtoMessage()    {}
func (*Activity_Workflow) Descriptor() ([]byte, []int) {
	return fileDescriptor_meta_b03dac0b310a1143, []int{0, 0}
}
func (m *Activity_Workflow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Activity_Workflow.Unmarshal(m, b)
}
func (m *Activity_Workflow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Activity_Workflow.Marshal(b, m, deterministic)
}
func (dst *Activity_Workflow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Activity_Workflow.Merge(dst, src)
}
func (m *Activity_Workflow) XXX_Size() int {
	return xxx_messageInfo_Activity_Workflow.Size(m)
}
func (m *Activity_Workflow) XXX_DiscardUnknown() {
	xxx_messageInfo_Activity_Workflow.DiscardUnknown(m)
}

var xxx_messageInfo_Activity_Workflow proto.InternalMessageInfo

func (m *Activity_Workflow) GetActivities() []string {
	if m != nil {
		return m.Activities
	}
	return nil
}

type Error struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_meta_b03dac0b310a1143, []int{1}
}
func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (dst *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(dst, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Activity)(nil), "powerssl.controller.v1.Activity")
	proto.RegisterType((*Activity_Workflow)(nil), "powerssl.controller.v1.Activity.Workflow")
	proto.RegisterType((*Error)(nil), "powerssl.controller.v1.Error")
	proto.RegisterEnum("powerssl.controller.v1.Activity_Name", Activity_Name_name, Activity_Name_value)
}

func init() { proto.RegisterFile("controller/api/v1/meta.proto", fileDescriptor_meta_b03dac0b310a1143) }

var fileDescriptor_meta_b03dac0b310a1143 = []byte{
	// 388 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x51, 0x6f, 0xd3, 0x30,
	0x10, 0xc7, 0x31, 0xed, 0xa0, 0x3d, 0x24, 0x14, 0xcc, 0x28, 0x61, 0x9a, 0x50, 0xa9, 0x34, 0x51,
	0x78, 0x48, 0xb4, 0xf1, 0x84, 0x78, 0x32, 0xce, 0x55, 0x44, 0xb0, 0x14, 0xdc, 0x74, 0x88, 0xbd,
	0x58, 0x66, 0xf2, 0x42, 0xd4, 0x24, 0x8e, 0x9c, 0xd0, 0x8a, 0x8f, 0xc6, 0x47, 0x60, 0x7c, 0x29,
	0xb4, 0x6c, 0x69, 0xa9, 0x40, 0xda, 0xe3, 0xfd, 0xcf, 0xbf, 0x9f, 0x4f, 0xf6, 0xc1, 0xfe, 0x99,
	0x29, 0x6a, 0x6b, 0xb2, 0x4c, 0x5b, 0x5f, 0x95, 0xa9, 0xbf, 0x3c, 0xf4, 0x73, 0x5d, 0x2b, 0xaf,
	0xb4, 0xa6, 0x36, 0x74, 0x50, 0x9a, 0x95, 0xb6, 0x55, 0x95, 0x79, 0x9b, 0x63, 0xde, 0xf2, 0x70,
	0xf4, 0xb3, 0x03, 0x3d, 0x76, 0x56, 0xa7, 0xcb, 0xb4, 0xfe, 0x41, 0x77, 0x61, 0xa7, 0x36, 0x0b,
	0x5d, 0xb8, 0x64, 0x48, 0xc6, 0x7d, 0x71, 0x55, 0xd0, 0xd7, 0xd0, 0x2d, 0x54, 0xae, 0xdd, 0xdb,
	0x43, 0x32, 0xbe, 0x7f, 0x74, 0xe0, 0xfd, 0xdf, 0xe4, 0xb5, 0x16, 0x2f, 0x52, 0xb9, 0x16, 0x0d,
	0x42, 0x11, 0x7a, 0x2b, 0x63, 0x17, 0xe7, 0x99, 0x59, 0xb9, 0x9d, 0x21, 0x19, 0xdf, 0x3b, 0x7a,
	0x71, 0x23, 0xfe, 0xf9, 0x1a, 0x10, 0x6b, 0x94, 0xee, 0x43, 0xbf, 0x4a, 0x93, 0x42, 0xd5, 0xdf,
	0xad, 0x76, 0xbb, 0xcd, 0x6c, 0x9b, 0x60, 0xef, 0x25, 0xf4, 0x5a, 0x86, 0x3e, 0x05, 0x50, 0x57,
	0xa2, 0x54, 0x57, 0x2e, 0x19, 0x76, 0xc6, 0x7d, 0xf1, 0x57, 0x32, 0xba, 0x20, 0xd0, 0xbd, 0x9c,
	0x8f, 0xee, 0x82, 0x13, 0xb1, 0x63, 0x94, 0xf3, 0x68, 0xf6, 0x11, 0x79, 0x38, 0x09, 0x31, 0x70,
	0x6e, 0xd1, 0xc7, 0xf0, 0x90, 0x33, 0xc9, 0xe6, 0xf1, 0xbb, 0xa9, 0x08, 0x4f, 0x51, 0x06, 0xd3,
	0x63, 0x16, 0x46, 0x8e, 0xa6, 0x7b, 0x30, 0xe0, 0x4c, 0x0a, 0xfc, 0x34, 0xc7, 0x59, 0x2c, 0x39,
	0x8a, 0x38, 0x9c, 0x84, 0x9c, 0xc5, 0xe8, 0x9c, 0xd3, 0x27, 0xf0, 0xa8, 0xe9, 0x9d, 0x4c, 0xdf,
	0xe3, 0x56, 0x2b, 0xb9, 0xbc, 0x85, 0x33, 0x79, 0x82, 0x22, 0x9c, 0x7c, 0x69, 0x65, 0xdf, 0xe8,
	0x00, 0x1e, 0x04, 0xd1, 0x4c, 0x72, 0x81, 0x2c, 0x46, 0x29, 0x90, 0x4f, 0x45, 0xe0, 0xfc, 0x22,
	0x6d, 0x1e, 0xe0, 0x07, 0xdc, 0xe4, 0x17, 0xeb, 0x7c, 0x5b, 0xf3, 0x9b, 0x8c, 0x9e, 0xc1, 0x0e,
	0x5a, 0x6b, 0x2c, 0x75, 0xe1, 0x6e, 0xae, 0xab, 0x4a, 0x25, 0xfa, 0xfa, 0xe7, 0xda, 0xf2, 0xed,
	0xf3, 0xd3, 0x83, 0xf5, 0x7b, 0xa7, 0xc6, 0x2f, 0x17, 0x89, 0xff, 0xcf, 0x9a, 0xbc, 0x51, 0x65,
	0xfa, 0xf5, 0x4e, 0xb3, 0x26, 0xaf, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0xcc, 0x99, 0x1f, 0x7a,
	0x46, 0x02, 0x00, 0x00,
}