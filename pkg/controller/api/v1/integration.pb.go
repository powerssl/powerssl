// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: controller/api/v1/integration.proto

package api // import "powerssl.io/pkg/controller/api/v1"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

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

type IntegrationKind int32

const (
	IntegrationKind_INTEGRATION_KIND_UNSPECIFIED IntegrationKind = 0
	IntegrationKind_CA                           IntegrationKind = 1
	IntegrationKind_DNS                          IntegrationKind = 2
)

var IntegrationKind_name = map[int32]string{
	0: "INTEGRATION_KIND_UNSPECIFIED",
	1: "CA",
	2: "DNS",
}
var IntegrationKind_value = map[string]int32{
	"INTEGRATION_KIND_UNSPECIFIED": 0,
	"CA":                           1,
	"DNS":                          2,
}

func (x IntegrationKind) String() string {
	return proto.EnumName(IntegrationKind_name, int32(x))
}
func (IntegrationKind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_integration_a11df427e3d526ef, []int{0}
}

type RegisterIntegrationRequest struct {
	Kind                 IntegrationKind `protobuf:"varint,1,opt,name=kind,proto3,enum=powerssl.controller.v1.IntegrationKind" json:"kind,omitempty"`
	Name                 string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *RegisterIntegrationRequest) Reset()         { *m = RegisterIntegrationRequest{} }
func (m *RegisterIntegrationRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterIntegrationRequest) ProtoMessage()    {}
func (*RegisterIntegrationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_integration_a11df427e3d526ef, []int{0}
}
func (m *RegisterIntegrationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterIntegrationRequest.Unmarshal(m, b)
}
func (m *RegisterIntegrationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterIntegrationRequest.Marshal(b, m, deterministic)
}
func (dst *RegisterIntegrationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterIntegrationRequest.Merge(dst, src)
}
func (m *RegisterIntegrationRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterIntegrationRequest.Size(m)
}
func (m *RegisterIntegrationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterIntegrationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterIntegrationRequest proto.InternalMessageInfo

func (m *RegisterIntegrationRequest) GetKind() IntegrationKind {
	if m != nil {
		return m.Kind
	}
	return IntegrationKind_INTEGRATION_KIND_UNSPECIFIED
}

func (m *RegisterIntegrationRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*RegisterIntegrationRequest)(nil), "powerssl.controller.v1.RegisterIntegrationRequest")
	proto.RegisterEnum("powerssl.controller.v1.IntegrationKind", IntegrationKind_name, IntegrationKind_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for IntegrationService service

type IntegrationServiceClient interface {
	Register(ctx context.Context, in *RegisterIntegrationRequest, opts ...grpc.CallOption) (IntegrationService_RegisterClient, error)
}

type integrationServiceClient struct {
	cc *grpc.ClientConn
}

func NewIntegrationServiceClient(cc *grpc.ClientConn) IntegrationServiceClient {
	return &integrationServiceClient{cc}
}

func (c *integrationServiceClient) Register(ctx context.Context, in *RegisterIntegrationRequest, opts ...grpc.CallOption) (IntegrationService_RegisterClient, error) {
	stream, err := c.cc.NewStream(ctx, &_IntegrationService_serviceDesc.Streams[0], "/powerssl.controller.v1.IntegrationService/Register", opts...)
	if err != nil {
		return nil, err
	}
	x := &integrationServiceRegisterClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type IntegrationService_RegisterClient interface {
	Recv() (*Activity, error)
	grpc.ClientStream
}

type integrationServiceRegisterClient struct {
	grpc.ClientStream
}

func (x *integrationServiceRegisterClient) Recv() (*Activity, error) {
	m := new(Activity)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for IntegrationService service

type IntegrationServiceServer interface {
	Register(*RegisterIntegrationRequest, IntegrationService_RegisterServer) error
}

func RegisterIntegrationServiceServer(s *grpc.Server, srv IntegrationServiceServer) {
	s.RegisterService(&_IntegrationService_serviceDesc, srv)
}

func _IntegrationService_Register_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RegisterIntegrationRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(IntegrationServiceServer).Register(m, &integrationServiceRegisterServer{stream})
}

type IntegrationService_RegisterServer interface {
	Send(*Activity) error
	grpc.ServerStream
}

type integrationServiceRegisterServer struct {
	grpc.ServerStream
}

func (x *integrationServiceRegisterServer) Send(m *Activity) error {
	return x.ServerStream.SendMsg(m)
}

var _IntegrationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "powerssl.controller.v1.IntegrationService",
	HandlerType: (*IntegrationServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Register",
			Handler:       _IntegrationService_Register_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "controller/api/v1/integration.proto",
}

func init() {
	proto.RegisterFile("controller/api/v1/integration.proto", fileDescriptor_integration_a11df427e3d526ef)
}

var fileDescriptor_integration_a11df427e3d526ef = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xb1, 0x4f, 0x83, 0x40,
	0x14, 0xc6, 0x0b, 0x36, 0x55, 0x6f, 0x50, 0x72, 0x83, 0x69, 0x48, 0x07, 0x52, 0x63, 0xda, 0x38,
	0x80, 0xc5, 0xb1, 0x13, 0x16, 0x34, 0x97, 0x26, 0x68, 0xa0, 0x2e, 0x2e, 0x0d, 0xc2, 0x0b, 0x79,
	0x29, 0xdc, 0xe1, 0x71, 0x62, 0xf4, 0xaf, 0x37, 0xc1, 0x68, 0x1b, 0x5b, 0xb6, 0x1b, 0x7e, 0xbf,
	0xef, 0xde, 0xf7, 0x91, 0xcb, 0x54, 0x70, 0x25, 0x45, 0x51, 0x80, 0x74, 0x92, 0x0a, 0x9d, 0x66,
	0xe6, 0x20, 0x57, 0x90, 0xcb, 0x44, 0xa1, 0xe0, 0x76, 0x25, 0x85, 0x12, 0xf4, 0xa2, 0x12, 0x1f,
	0x20, 0xeb, 0xba, 0xb0, 0xb7, 0xb4, 0xdd, 0xcc, 0xcc, 0xd1, 0xbe, 0x5c, 0x82, 0x4a, 0x7e, 0xac,
	0x71, 0x49, 0xcc, 0x08, 0x72, 0xac, 0x15, 0x48, 0xb6, 0x8d, 0x8c, 0xe0, 0xed, 0x1d, 0x6a, 0x45,
	0xe7, 0xa4, 0xbf, 0x41, 0x9e, 0x0d, 0x35, 0x4b, 0x9b, 0x9e, 0xb9, 0x13, 0xfb, 0xf0, 0x17, 0xf6,
	0x8e, 0xb9, 0x44, 0x9e, 0x45, 0xad, 0x44, 0x29, 0xe9, 0xf3, 0xa4, 0x84, 0xa1, 0x6e, 0x69, 0xd3,
	0xd3, 0xa8, 0x7d, 0x5f, 0xfb, 0xe4, 0xfc, 0x1f, 0x4c, 0x2d, 0x32, 0x62, 0xe1, 0x2a, 0x78, 0x88,
	0xbc, 0x15, 0x7b, 0x0c, 0xd7, 0x4b, 0x16, 0xfa, 0xeb, 0xe7, 0x30, 0x7e, 0x0a, 0x16, 0xec, 0x9e,
	0x05, 0xbe, 0xd1, 0xa3, 0x03, 0xa2, 0x2f, 0x3c, 0x43, 0xa3, 0xc7, 0xe4, 0xc8, 0x0f, 0x63, 0x43,
	0x77, 0xbf, 0x08, 0xdd, 0x49, 0x89, 0x41, 0x36, 0x98, 0x02, 0xcd, 0xc8, 0xc9, 0x6f, 0x15, 0xea,
	0x76, 0x9d, 0xda, 0x5d, 0xd6, 0xb4, 0xba, 0x1c, 0x2f, 0x55, 0xd8, 0xa0, 0xfa, 0x1c, 0xf7, 0x6e,
	0xb4, 0xbb, 0xc9, 0xcb, 0xd5, 0x1f, 0x86, 0xc2, 0xa9, 0x36, 0xb9, 0xb3, 0x37, 0xf0, 0x3c, 0xa9,
	0xf0, 0x75, 0xd0, 0x0e, 0x7c, 0xfb, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x5e, 0xa3, 0x80, 0xbd,
	0x01, 0x00, 0x00,
}
