// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// CertificateIssueServiceClient is the client API for CertificateIssueService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CertificateIssueServiceClient interface {
	// Creates a certificate, and returns the new CertificateIssue.
	Create(ctx context.Context, in *CreateCertificateIssueRequest, opts ...grpc.CallOption) (*CertificateIssue, error)
	// Deletes a certificate. Returns NOT_FOUND if the certificate does not exist.
	Delete(ctx context.Context, in *DeleteCertificateIssueRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
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
	cc grpc.ClientConnInterface
}

func NewCertificateIssueServiceClient(cc grpc.ClientConnInterface) CertificateIssueServiceClient {
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

func (c *certificateIssueServiceClient) Delete(ctx context.Context, in *DeleteCertificateIssueRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
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
// All implementations must embed UnimplementedCertificateIssueServiceServer
// for forward compatibility
type CertificateIssueServiceServer interface {
	// Creates a certificate, and returns the new CertificateIssue.
	Create(context.Context, *CreateCertificateIssueRequest) (*CertificateIssue, error)
	// Deletes a certificate. Returns NOT_FOUND if the certificate does not exist.
	Delete(context.Context, *DeleteCertificateIssueRequest) (*emptypb.Empty, error)
	// Gets a certificate. Returns NOT_FOUND if the certificate does not exist.
	Get(context.Context, *GetCertificateIssueRequest) (*CertificateIssue, error)
	// Lists certificates. The order is unspecified but deterministic. Newly
	// created certificates will not necessarily be added to the end of this list.
	List(context.Context, *ListCertificateIssuesRequest) (*ListCertificateIssuesResponse, error)
	// Updates a certificate. Returns INVALID_ARGUMENT if the name of the
	// certificate is non-empty and does equal the previous name.
	Update(context.Context, *UpdateCertificateIssueRequest) (*CertificateIssue, error)
	mustEmbedUnimplementedCertificateIssueServiceServer()
}

// UnimplementedCertificateIssueServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCertificateIssueServiceServer struct {
}

func (UnimplementedCertificateIssueServiceServer) Create(context.Context, *CreateCertificateIssueRequest) (*CertificateIssue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCertificateIssueServiceServer) Delete(context.Context, *DeleteCertificateIssueRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCertificateIssueServiceServer) Get(context.Context, *GetCertificateIssueRequest) (*CertificateIssue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCertificateIssueServiceServer) List(context.Context, *ListCertificateIssuesRequest) (*ListCertificateIssuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedCertificateIssueServiceServer) Update(context.Context, *UpdateCertificateIssueRequest) (*CertificateIssue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedCertificateIssueServiceServer) mustEmbedUnimplementedCertificateIssueServiceServer() {
}

// UnsafeCertificateIssueServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CertificateIssueServiceServer will
// result in compilation errors.
type UnsafeCertificateIssueServiceServer interface {
	mustEmbedUnimplementedCertificateIssueServiceServer()
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