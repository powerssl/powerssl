// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: powerssl/controller/v1/workflow.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type WorkflowKind int32

const (
	WorkflowKind_WORKFLOW_KIND_UNSPECIFIED WorkflowKind = 0
	WorkflowKind_CREATE_ACME_ACCOUNT       WorkflowKind = 1
	WorkflowKind_REQUEST_ACME_CERTIFICATE  WorkflowKind = 2
)

var WorkflowKind_name = map[int32]string{
	0: "WORKFLOW_KIND_UNSPECIFIED",
	1: "CREATE_ACME_ACCOUNT",
	2: "REQUEST_ACME_CERTIFICATE",
}

var WorkflowKind_value = map[string]int32{
	"WORKFLOW_KIND_UNSPECIFIED": 0,
	"CREATE_ACME_ACCOUNT":       1,
	"REQUEST_ACME_CERTIFICATE":  2,
}

func (x WorkflowKind) String() string {
	return proto.EnumName(WorkflowKind_name, int32(x))
}

func (WorkflowKind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cd8a016b5c91effc, []int{0}
}

type CreateWorkflowRequest struct {
	Workflow             *Workflow `protobuf:"bytes,1,opt,name=workflow,proto3" json:"workflow,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CreateWorkflowRequest) Reset()         { *m = CreateWorkflowRequest{} }
func (m *CreateWorkflowRequest) String() string { return proto.CompactTextString(m) }
func (*CreateWorkflowRequest) ProtoMessage()    {}
func (*CreateWorkflowRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd8a016b5c91effc, []int{0}
}
func (m *CreateWorkflowRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateWorkflowRequest.Unmarshal(m, b)
}
func (m *CreateWorkflowRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateWorkflowRequest.Marshal(b, m, deterministic)
}
func (m *CreateWorkflowRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateWorkflowRequest.Merge(m, src)
}
func (m *CreateWorkflowRequest) XXX_Size() int {
	return xxx_messageInfo_CreateWorkflowRequest.Size(m)
}
func (m *CreateWorkflowRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateWorkflowRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateWorkflowRequest proto.InternalMessageInfo

func (m *CreateWorkflowRequest) GetWorkflow() *Workflow {
	if m != nil {
		return m.Workflow
	}
	return nil
}

type Workflow struct {
	Name               string                        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Kind               WorkflowKind                  `protobuf:"varint,2,opt,name=kind,proto3,enum=powerssl.controller.v1.WorkflowKind" json:"kind,omitempty"`
	IntegrationFilters []*Workflow_IntegrationFilter `protobuf:"bytes,3,rep,name=integration_filters,json=integrationFilters,proto3" json:"integration_filters,omitempty"`
	// Types that are valid to be assigned to Input:
	//	*Workflow_CreateAcmeAccountInput
	//	*Workflow_RequestAcmeCertificateInput
	Input                isWorkflow_Input `protobuf_oneof:"input"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Workflow) Reset()         { *m = Workflow{} }
func (m *Workflow) String() string { return proto.CompactTextString(m) }
func (*Workflow) ProtoMessage()    {}
func (*Workflow) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd8a016b5c91effc, []int{1}
}
func (m *Workflow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Workflow.Unmarshal(m, b)
}
func (m *Workflow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Workflow.Marshal(b, m, deterministic)
}
func (m *Workflow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Workflow.Merge(m, src)
}
func (m *Workflow) XXX_Size() int {
	return xxx_messageInfo_Workflow.Size(m)
}
func (m *Workflow) XXX_DiscardUnknown() {
	xxx_messageInfo_Workflow.DiscardUnknown(m)
}

var xxx_messageInfo_Workflow proto.InternalMessageInfo

type isWorkflow_Input interface {
	isWorkflow_Input()
}

type Workflow_CreateAcmeAccountInput struct {
	CreateAcmeAccountInput *CreateACMEAccountInput `protobuf:"bytes,11,opt,name=create_acme_account_input,json=createAcmeAccountInput,proto3,oneof"`
}
type Workflow_RequestAcmeCertificateInput struct {
	RequestAcmeCertificateInput *RequestACMECertificateInput `protobuf:"bytes,12,opt,name=request_acme_certificate_input,json=requestAcmeCertificateInput,proto3,oneof"`
}

func (*Workflow_CreateAcmeAccountInput) isWorkflow_Input()      {}
func (*Workflow_RequestAcmeCertificateInput) isWorkflow_Input() {}

func (m *Workflow) GetInput() isWorkflow_Input {
	if m != nil {
		return m.Input
	}
	return nil
}

func (m *Workflow) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Workflow) GetKind() WorkflowKind {
	if m != nil {
		return m.Kind
	}
	return WorkflowKind_WORKFLOW_KIND_UNSPECIFIED
}

func (m *Workflow) GetIntegrationFilters() []*Workflow_IntegrationFilter {
	if m != nil {
		return m.IntegrationFilters
	}
	return nil
}

func (m *Workflow) GetCreateAcmeAccountInput() *CreateACMEAccountInput {
	if x, ok := m.GetInput().(*Workflow_CreateAcmeAccountInput); ok {
		return x.CreateAcmeAccountInput
	}
	return nil
}

func (m *Workflow) GetRequestAcmeCertificateInput() *RequestACMECertificateInput {
	if x, ok := m.GetInput().(*Workflow_RequestAcmeCertificateInput); ok {
		return x.RequestAcmeCertificateInput
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Workflow) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Workflow_OneofMarshaler, _Workflow_OneofUnmarshaler, _Workflow_OneofSizer, []interface{}{
		(*Workflow_CreateAcmeAccountInput)(nil),
		(*Workflow_RequestAcmeCertificateInput)(nil),
	}
}

func _Workflow_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Workflow)
	// input
	switch x := m.Input.(type) {
	case *Workflow_CreateAcmeAccountInput:
		_ = b.EncodeVarint(11<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CreateAcmeAccountInput); err != nil {
			return err
		}
	case *Workflow_RequestAcmeCertificateInput:
		_ = b.EncodeVarint(12<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RequestAcmeCertificateInput); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Workflow.Input has unexpected type %T", x)
	}
	return nil
}

func _Workflow_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Workflow)
	switch tag {
	case 11: // input.create_acme_account_input
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CreateACMEAccountInput)
		err := b.DecodeMessage(msg)
		m.Input = &Workflow_CreateAcmeAccountInput{msg}
		return true, err
	case 12: // input.request_acme_certificate_input
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RequestACMECertificateInput)
		err := b.DecodeMessage(msg)
		m.Input = &Workflow_RequestAcmeCertificateInput{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Workflow_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Workflow)
	// input
	switch x := m.Input.(type) {
	case *Workflow_CreateAcmeAccountInput:
		s := proto.Size(x.CreateAcmeAccountInput)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Workflow_RequestAcmeCertificateInput:
		s := proto.Size(x.RequestAcmeCertificateInput)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Workflow_IntegrationFilter struct {
	Kind                 IntegrationKind `protobuf:"varint,1,opt,name=kind,proto3,enum=powerssl.controller.v1.IntegrationKind" json:"kind,omitempty"`
	Name                 string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Workflow_IntegrationFilter) Reset()         { *m = Workflow_IntegrationFilter{} }
func (m *Workflow_IntegrationFilter) String() string { return proto.CompactTextString(m) }
func (*Workflow_IntegrationFilter) ProtoMessage()    {}
func (*Workflow_IntegrationFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd8a016b5c91effc, []int{1, 0}
}
func (m *Workflow_IntegrationFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Workflow_IntegrationFilter.Unmarshal(m, b)
}
func (m *Workflow_IntegrationFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Workflow_IntegrationFilter.Marshal(b, m, deterministic)
}
func (m *Workflow_IntegrationFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Workflow_IntegrationFilter.Merge(m, src)
}
func (m *Workflow_IntegrationFilter) XXX_Size() int {
	return xxx_messageInfo_Workflow_IntegrationFilter.Size(m)
}
func (m *Workflow_IntegrationFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_Workflow_IntegrationFilter.DiscardUnknown(m)
}

var xxx_messageInfo_Workflow_IntegrationFilter proto.InternalMessageInfo

func (m *Workflow_IntegrationFilter) GetKind() IntegrationKind {
	if m != nil {
		return m.Kind
	}
	return IntegrationKind_INTEGRATION_KIND_UNSPECIFIED
}

func (m *Workflow_IntegrationFilter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateACMEAccountInput struct {
	Account              string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	DirectoryUrl         string   `protobuf:"bytes,2,opt,name=directory_url,json=directoryUrl,proto3" json:"directory_url,omitempty"`
	TermsOfServiceAgreed bool     `protobuf:"varint,3,opt,name=terms_of_service_agreed,json=termsOfServiceAgreed,proto3" json:"terms_of_service_agreed,omitempty"`
	Contacts             []string `protobuf:"bytes,4,rep,name=contacts,proto3" json:"contacts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateACMEAccountInput) Reset()         { *m = CreateACMEAccountInput{} }
func (m *CreateACMEAccountInput) String() string { return proto.CompactTextString(m) }
func (*CreateACMEAccountInput) ProtoMessage()    {}
func (*CreateACMEAccountInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd8a016b5c91effc, []int{2}
}
func (m *CreateACMEAccountInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateACMEAccountInput.Unmarshal(m, b)
}
func (m *CreateACMEAccountInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateACMEAccountInput.Marshal(b, m, deterministic)
}
func (m *CreateACMEAccountInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateACMEAccountInput.Merge(m, src)
}
func (m *CreateACMEAccountInput) XXX_Size() int {
	return xxx_messageInfo_CreateACMEAccountInput.Size(m)
}
func (m *CreateACMEAccountInput) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateACMEAccountInput.DiscardUnknown(m)
}

var xxx_messageInfo_CreateACMEAccountInput proto.InternalMessageInfo

func (m *CreateACMEAccountInput) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *CreateACMEAccountInput) GetDirectoryUrl() string {
	if m != nil {
		return m.DirectoryUrl
	}
	return ""
}

func (m *CreateACMEAccountInput) GetTermsOfServiceAgreed() bool {
	if m != nil {
		return m.TermsOfServiceAgreed
	}
	return false
}

func (m *CreateACMEAccountInput) GetContacts() []string {
	if m != nil {
		return m.Contacts
	}
	return nil
}

type RequestACMECertificateInput struct {
	DirectoryUrl         string   `protobuf:"bytes,1,opt,name=directory_url,json=directoryUrl,proto3" json:"directory_url,omitempty"`
	AccountUrl           string   `protobuf:"bytes,2,opt,name=account_url,json=accountUrl,proto3" json:"account_url,omitempty"`
	Dnsnames             []string `protobuf:"bytes,3,rep,name=dnsnames,proto3" json:"dnsnames,omitempty"`
	NotBefore            string   `protobuf:"bytes,4,opt,name=not_before,json=notBefore,proto3" json:"not_before,omitempty"`
	NotAfter             string   `protobuf:"bytes,5,opt,name=not_after,json=notAfter,proto3" json:"not_after,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestACMECertificateInput) Reset()         { *m = RequestACMECertificateInput{} }
func (m *RequestACMECertificateInput) String() string { return proto.CompactTextString(m) }
func (*RequestACMECertificateInput) ProtoMessage()    {}
func (*RequestACMECertificateInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd8a016b5c91effc, []int{3}
}
func (m *RequestACMECertificateInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestACMECertificateInput.Unmarshal(m, b)
}
func (m *RequestACMECertificateInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestACMECertificateInput.Marshal(b, m, deterministic)
}
func (m *RequestACMECertificateInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestACMECertificateInput.Merge(m, src)
}
func (m *RequestACMECertificateInput) XXX_Size() int {
	return xxx_messageInfo_RequestACMECertificateInput.Size(m)
}
func (m *RequestACMECertificateInput) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestACMECertificateInput.DiscardUnknown(m)
}

var xxx_messageInfo_RequestACMECertificateInput proto.InternalMessageInfo

func (m *RequestACMECertificateInput) GetDirectoryUrl() string {
	if m != nil {
		return m.DirectoryUrl
	}
	return ""
}

func (m *RequestACMECertificateInput) GetAccountUrl() string {
	if m != nil {
		return m.AccountUrl
	}
	return ""
}

func (m *RequestACMECertificateInput) GetDnsnames() []string {
	if m != nil {
		return m.Dnsnames
	}
	return nil
}

func (m *RequestACMECertificateInput) GetNotBefore() string {
	if m != nil {
		return m.NotBefore
	}
	return ""
}

func (m *RequestACMECertificateInput) GetNotAfter() string {
	if m != nil {
		return m.NotAfter
	}
	return ""
}

func init() {
	proto.RegisterEnum("powerssl.controller.v1.WorkflowKind", WorkflowKind_name, WorkflowKind_value)
	proto.RegisterType((*CreateWorkflowRequest)(nil), "powerssl.controller.v1.CreateWorkflowRequest")
	proto.RegisterType((*Workflow)(nil), "powerssl.controller.v1.Workflow")
	proto.RegisterType((*Workflow_IntegrationFilter)(nil), "powerssl.controller.v1.Workflow.IntegrationFilter")
	proto.RegisterType((*CreateACMEAccountInput)(nil), "powerssl.controller.v1.CreateACMEAccountInput")
	proto.RegisterType((*RequestACMECertificateInput)(nil), "powerssl.controller.v1.RequestACMECertificateInput")
}

func init() {
	proto.RegisterFile("powerssl/controller/v1/workflow.proto", fileDescriptor_cd8a016b5c91effc)
}

var fileDescriptor_cd8a016b5c91effc = []byte{
	// 646 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xe1, 0x6e, 0xd3, 0x3c,
	0x14, 0x5d, 0xd6, 0x6e, 0x6b, 0x6f, 0xf7, 0x7d, 0x0c, 0x0f, 0xb6, 0xac, 0x63, 0x50, 0x15, 0x10,
	0x15, 0x12, 0x29, 0xeb, 0x84, 0x84, 0x34, 0xfe, 0x64, 0x59, 0x2a, 0xa2, 0xc1, 0x0a, 0x59, 0xab,
	0x49, 0xf0, 0xc3, 0xca, 0x12, 0x67, 0xb2, 0x9a, 0xda, 0xc5, 0x71, 0x3b, 0xc1, 0x7b, 0xf0, 0x12,
	0xbc, 0x02, 0x2f, 0x87, 0xe2, 0x24, 0x6d, 0xb4, 0xad, 0xdb, 0x9f, 0xaa, 0xf7, 0x5e, 0x9f, 0x73,
	0x7c, 0xcf, 0x75, 0x2e, 0xbc, 0x1c, 0xf3, 0x2b, 0x22, 0xe2, 0x38, 0x6a, 0xfb, 0x9c, 0x49, 0xc1,
	0xa3, 0x88, 0x88, 0xf6, 0x74, 0xbf, 0x7d, 0xc5, 0xc5, 0x30, 0x8c, 0xf8, 0x95, 0x31, 0x16, 0x5c,
	0x72, 0xb4, 0x95, 0x1f, 0x33, 0xe6, 0xc7, 0x8c, 0xe9, 0x7e, 0xbd, 0xb5, 0x00, 0x4e, 0x99, 0x24,
	0x97, 0xc2, 0x93, 0x94, 0xb3, 0x94, 0xa1, 0x39, 0x80, 0xc7, 0x96, 0x20, 0x9e, 0x24, 0xe7, 0x19,
	0xb3, 0x4b, 0x7e, 0x4c, 0x48, 0x2c, 0xd1, 0x07, 0xa8, 0xe4, 0x62, 0xba, 0xd6, 0xd0, 0x5a, 0xb5,
	0x4e, 0xc3, 0xb8, 0x5d, 0xcd, 0x98, 0x41, 0x67, 0x88, 0xe6, 0xef, 0x32, 0x54, 0xf2, 0x34, 0x42,
	0x50, 0x66, 0xde, 0x88, 0x28, 0x9a, 0xaa, 0xab, 0xfe, 0xa3, 0xf7, 0x50, 0x1e, 0x52, 0x16, 0xe8,
	0xcb, 0x0d, 0xad, 0xf5, 0x7f, 0xe7, 0xc5, 0x7d, 0xd4, 0x27, 0x94, 0x05, 0xae, 0x42, 0x20, 0x1f,
	0x36, 0x0b, 0x6d, 0xe0, 0x90, 0x46, 0x92, 0x88, 0x58, 0x2f, 0x35, 0x4a, 0xad, 0x5a, 0xa7, 0x73,
	0x1f, 0x91, 0xe1, 0xcc, 0xb1, 0x5d, 0x05, 0x75, 0x11, 0xbd, 0x9e, 0x8a, 0xd1, 0x10, 0x76, 0x7c,
	0x65, 0x0b, 0xf6, 0xfc, 0x51, 0xf2, 0xe3, 0xf3, 0x09, 0x93, 0x98, 0xb2, 0xf1, 0x44, 0xea, 0x35,
	0x65, 0x87, 0xb1, 0x48, 0x2a, 0xf5, 0xd3, 0xb4, 0x3e, 0xdb, 0x66, 0x0a, 0x73, 0x12, 0xd4, 0xc7,
	0x25, 0x77, 0x2b, 0xa5, 0x34, 0xfd, 0x11, 0x29, 0x56, 0xd0, 0x2f, 0x78, 0x2a, 0x52, 0xd7, 0x53,
	0x35, 0x9f, 0x08, 0x49, 0x43, 0xea, 0x27, 0xf2, 0xa9, 0xe2, 0xba, 0x52, 0x3c, 0x58, 0xa4, 0x98,
	0xcd, 0x2c, 0x91, 0xb4, 0xe6, 0xd8, 0x5c, 0x76, 0x37, 0x23, 0x4f, 0x74, 0xaf, 0x97, 0xeb, 0x01,
	0x3c, 0xbc, 0xe1, 0x08, 0x3a, 0xcc, 0x86, 0xa3, 0xa9, 0xe1, 0xbc, 0x5a, 0x24, 0x5b, 0x00, 0x16,
	0xe6, 0x93, 0x4f, 0x7b, 0x79, 0x3e, 0xed, 0xa3, 0x35, 0x58, 0x51, 0x8d, 0x34, 0xff, 0x68, 0xb0,
	0x75, 0xbb, 0x3f, 0x48, 0x87, 0xb5, 0xcc, 0xe6, 0xec, 0xa1, 0xe4, 0x21, 0x7a, 0x0e, 0xff, 0x05,
	0x54, 0x10, 0x5f, 0x72, 0xf1, 0x13, 0x4f, 0x44, 0x94, 0x51, 0xaf, 0xcf, 0x92, 0x03, 0x11, 0xa1,
	0x77, 0xb0, 0x2d, 0x89, 0x18, 0xc5, 0x98, 0x87, 0x38, 0x26, 0x62, 0x4a, 0x7d, 0x82, 0xbd, 0x4b,
	0x41, 0x48, 0xa0, 0x97, 0x1a, 0x5a, 0xab, 0xe2, 0x3e, 0x52, 0xe5, 0x5e, 0x78, 0x96, 0x16, 0x4d,
	0x55, 0x43, 0x75, 0xa8, 0x24, 0x4d, 0x79, 0xbe, 0x8c, 0xf5, 0x72, 0xa3, 0xd4, 0xaa, 0xba, 0xb3,
	0xb8, 0xf9, 0x57, 0x83, 0xdd, 0x3b, 0xac, 0xbd, 0x79, 0x2f, 0xed, 0x96, 0x7b, 0x3d, 0x83, 0x5a,
	0xfe, 0x7a, 0xe6, 0x57, 0x87, 0x2c, 0x95, 0x1c, 0xa8, 0x43, 0x25, 0x60, 0x71, 0x62, 0x53, 0xfa,
	0x88, 0xab, 0xee, 0x2c, 0x46, 0x7b, 0x00, 0x8c, 0x4b, 0x7c, 0x41, 0x42, 0x2e, 0x88, 0x5e, 0x56,
	0xd8, 0x2a, 0xe3, 0xf2, 0x48, 0x25, 0xd0, 0x2e, 0x24, 0x01, 0xf6, 0x42, 0x49, 0x84, 0xbe, 0xa2,
	0xaa, 0x15, 0xc6, 0xa5, 0x99, 0xc4, 0xaf, 0x03, 0x58, 0x2f, 0x7e, 0x3d, 0x68, 0x0f, 0x76, 0xce,
	0x7b, 0xee, 0x49, 0xf7, 0x53, 0xef, 0x1c, 0x9f, 0x38, 0xa7, 0xc7, 0x78, 0x70, 0x7a, 0xf6, 0xc5,
	0xb6, 0x9c, 0xae, 0x63, 0x1f, 0x6f, 0x2c, 0xa1, 0x6d, 0xd8, 0xb4, 0x5c, 0xdb, 0xec, 0xdb, 0x38,
	0xe9, 0x15, 0x9b, 0x96, 0xd5, 0x1b, 0x9c, 0xf6, 0x37, 0x34, 0xf4, 0x04, 0x74, 0xd7, 0xfe, 0x3a,
	0xb0, 0xcf, 0xfa, 0x69, 0xc5, 0xb2, 0xdd, 0xbe, 0xd3, 0x75, 0x2c, 0xb3, 0x6f, 0x6f, 0x2c, 0x77,
	0x18, 0x3c, 0xc8, 0x55, 0x32, 0x63, 0xd1, 0x77, 0x58, 0x4d, 0x47, 0x8c, 0xde, 0xdc, 0xfd, 0x89,
	0x5c, 0x5b, 0x39, 0xf5, 0x7b, 0x17, 0x4c, 0x73, 0xe9, 0xe8, 0xed, 0xb7, 0xf9, 0x21, 0xca, 0xd5,
	0x42, 0x13, 0xcc, 0x8b, 0xda, 0xe3, 0xe1, 0x65, 0x71, 0xd7, 0x79, 0x63, 0xda, 0x9e, 0xee, 0x1f,
	0x7a, 0x63, 0x7a, 0xb1, 0xaa, 0x16, 0xdd, 0xc1, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe4, 0x04,
	0x26, 0xf5, 0x53, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WorkflowServiceClient is the client API for WorkflowService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WorkflowServiceClient interface {
	Create(ctx context.Context, in *CreateWorkflowRequest, opts ...grpc.CallOption) (*Workflow, error)
}

type workflowServiceClient struct {
	cc *grpc.ClientConn
}

func NewWorkflowServiceClient(cc *grpc.ClientConn) WorkflowServiceClient {
	return &workflowServiceClient{cc}
}

func (c *workflowServiceClient) Create(ctx context.Context, in *CreateWorkflowRequest, opts ...grpc.CallOption) (*Workflow, error) {
	out := new(Workflow)
	err := c.cc.Invoke(ctx, "/powerssl.controller.v1.WorkflowService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkflowServiceServer is the server API for WorkflowService service.
type WorkflowServiceServer interface {
	Create(context.Context, *CreateWorkflowRequest) (*Workflow, error)
}

func RegisterWorkflowServiceServer(s *grpc.Server, srv WorkflowServiceServer) {
	s.RegisterService(&_WorkflowService_serviceDesc, srv)
}

func _WorkflowService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/powerssl.controller.v1.WorkflowService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServiceServer).Create(ctx, req.(*CreateWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WorkflowService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "powerssl.controller.v1.WorkflowService",
	HandlerType: (*WorkflowServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _WorkflowService_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "powerssl/controller/v1/workflow.proto",
}
