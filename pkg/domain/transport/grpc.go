package domaintransport

import (
	"context"

	// "google.golang.org/grpc"

	oldcontext "golang.org/x/net/context"

	// "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"

	types "github.com/gogo/protobuf/types"

	pb "powerssl.io/api/v1"
	domainendpoint "powerssl.io/pkg/domain/endpoint"
	domainmodel "powerssl.io/pkg/domain/model"
	// domainservice "powerssl.io/pkg/domain/service"
)

type grpcServer struct {
	create grpctransport.Handler
	delete grpctransport.Handler
	get    grpctransport.Handler
	list   grpctransport.Handler
	update grpctransport.Handler
}

func NewGRPCServer(endpoints domainendpoint.Set, logger log.Logger) pb.DomainServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}

	return &grpcServer{
		create: grpctransport.NewServer(
			endpoints.CreateEndpoint,
			decodeGRPCCreateRequest,
			encodeGRPCCreateResponse,
			options...,
		),
		delete: grpctransport.NewServer(
			endpoints.DeleteEndpoint,
			decodeGRPCDeleteRequest,
			encodeGRPCDeleteResponse,
			options...,
		),
		get: grpctransport.NewServer(
			endpoints.GetEndpoint,
			decodeGRPCGetRequest,
			encodeGRPCGetResponse,
			options...,
		),
		list: grpctransport.NewServer(
			endpoints.ListEndpoint,
			decodeGRPCListRequest,
			encodeGRPCListResponse,
			options...,
		),
		update: grpctransport.NewServer(
			endpoints.UpdateEndpoint,
			decodeGRPCUpdateRequest,
			encodeGRPCUpdateResponse,
			options...,
		),
	}
}

func (s *grpcServer) CreateDomain(ctx oldcontext.Context, req *pb.CreateDomainRequest) (*pb.Domain, error) {
	_, rep, err := s.create.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.Domain), nil
}

func (s *grpcServer) DeleteDomain(ctx oldcontext.Context, req *pb.DeleteDomainRequest) (*types.Empty, error) {
	_, rep, err := s.delete.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*types.Empty), nil
}

func (s *grpcServer) GetDomain(ctx oldcontext.Context, req *pb.GetDomainRequest) (*pb.Domain, error) {
	_, rep, err := s.get.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.Domain), nil
}

func (s *grpcServer) ListDomains(ctx oldcontext.Context, req *pb.ListDomainsRequest) (*pb.ListDomainsResponse, error) {
	_, rep, err := s.list.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ListDomainsResponse), nil
}

func (s *grpcServer) UpdateDomain(ctx oldcontext.Context, req *pb.UpdateDomainRequest) (*pb.Domain, error) {
	_, rep, err := s.update.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.Domain), nil
}

/*
func NewGRPCClient(conn *grpc.ClientConn) service.Service {
	options := []grpctransport.ClientOption{}

	var getEndpoint endpoint.Endpoint
	{
		getEndpoint = grpctransport.NewClient(
			conn,
			"powerssl.v1.DomainService",
			"GetDomain",
			encodeGRPCGetRequest,
			decodeGRPCGetResponse,
			pb.Domain{},
			options...,
		).Endpoint()
	}

	var listEndpoint endpoint.Endpoint
	{
		listEndpoint = grpctransport.NewClient(
			conn,
			"powerssl.v1.DomainService",
			"ListDomains",
			encodeGRPCListRequest,
			decodeGRPCListResponse,
			pb.ListDomainsResponse{},
			options...,
		).Endpoint()
	}

	return domainendpoint.Set{
		GetEndpoint:  getEndpoint,
		ListEndpoint: listEndpoint,
	}
}
*/

func decodeGRPCCreateRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.CreateDomainRequest)
	domain := domainmodel.Domain{
		DNSName: req.Domain.DnsName,
	}
	return domainendpoint.CreateRequest{Domain: domain}, nil
}

func decodeGRPCDeleteRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.DeleteDomainRequest)
	return domainendpoint.DeleteRequest{Name: req.Name}, nil
}

func decodeGRPCGetRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.GetDomainRequest)
	return domainendpoint.GetRequest{Name: req.Name}, nil
}

func decodeGRPCListRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.ListDomainsRequest)
	return domainendpoint.ListRequest{
		PageSize:  int(req.PageSize),
		PageToken: req.PageToken,
	}, nil
}

func decodeGRPCUpdateRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.UpdateDomainRequest)
	domain := domainmodel.Domain{
		DNSName: req.Domain.DnsName,
	}
	return domainendpoint.UpdateRequest{
		Domain:     domain,
		UpdateMask: "",
		// UpdateMask: string(req.UpdateMask), // TODO: ...
	}, nil
}

// func decodeGRPCGetResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
// 	reply := grpcReply.(*pb.Domain)
// 	return domainendpoint.GetResponse{Domain: reply.Domain, Err: str2err(reply.Err)}, nil
// }

// func decodeGRPCListResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
// 	reply := grpcReply.(*pb.ListDomainsResponse)
// 	return domainendpoint.ListResponse{Domains: reply.Domains, Err: str2err(reply.Err)}, nil
// }

func encodeGRPCCreateResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(domainendpoint.CreateResponse)
	return &pb.Domain{
		Name:    resp.Domain.ExternalName(),
		DnsName: resp.Domain.DNSName,
	}, nil
}

func encodeGRPCDeleteResponse(_ context.Context, response interface{}) (interface{}, error) {
	// resp := response.(domainendpoint.DeleteResponse)
	return nil, nil // TODO: empty
}

func encodeGRPCGetResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(domainendpoint.GetResponse)
	// return &pb.Domain{Domain: resp.Domain, Err: err2str(resp.Err)}, nil
	return &pb.Domain{
		Name:    resp.Domain.ExternalName(),
		DnsName: resp.Domain.DNSName,
	}, nil
}

func encodeGRPCListResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(domainendpoint.ListResponse)
	// return &pb.ListDomainsResponse{Domains: resp.Domains, Err: err2str(resp.Err)}, nil
	var domains []*pb.Domain
	for _, domain := range resp.Domains {
		domains = append(domains, &pb.Domain{Name: domain.ExternalName(), DnsName: domain.DNSName})
	}
	return &pb.ListDomainsResponse{Domains: domains}, nil
}

func encodeGRPCUpdateResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(domainendpoint.UpdateResponse)
	// return &pb.Domain{Domain: resp.Domain, Err: err2str(resp.Err)}, nil
	return &pb.Domain{
		Name:    resp.Domain.ExternalName(),
		DnsName: resp.Domain.DNSName,
	}, nil
}

// func encodeGRPCGetRequest(_ context.Context, request interface{}) (interface{}, error) {
// 	req := request.(domainendpoint.GetRequest)
// 	return &pb.GetDomainRequest{Name: req.Name}, nil
// }

// func encodeGRPCListRequest(_ context.Context, request interface{}) (interface{}, error) {
// 	req := request.(domainendpoint.ListRequest)
// 	return &pb.ListDomainsRequest{}, nil
// }

// func str2err(s string) error {
// 	if s == "" {
// 		return nil
// 	}
// 	return errors.New(s)
// }

// func err2str(err error) string {
// 	if err == nil {
// 		return ""
// 	}
// 	return err.Error()
// }
