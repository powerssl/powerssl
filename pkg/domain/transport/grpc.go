package transport

import (
	"context"

	// "google.golang.org/grpc"

	oldcontext "golang.org/x/net/context"

	// kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"

	pbempty "github.com/golang/protobuf/ptypes/empty"

	pb "powerssl.io/api/v1"
	domainmodel "powerssl.io/pkg/domain"
	"powerssl.io/pkg/domain/endpoint"
	// "powerssl.io/pkg/domain/service"
)

type grpcServer struct {
	create grpctransport.Handler
	delete grpctransport.Handler
	get    grpctransport.Handler
	list   grpctransport.Handler
	update grpctransport.Handler
}

func NewGRPCServer(endpoints endpoint.Set, logger log.Logger) pb.DomainServiceServer {
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

func (s *grpcServer) DeleteDomain(ctx oldcontext.Context, req *pb.DeleteDomainRequest) (*pbempty.Empty, error) {
	_, rep, err := s.delete.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pbempty.Empty), nil
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

	var getEndpoint kitendpoint.Endpoint
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

	var listEndpoint kitendpoint.Endpoint
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

	return endpoint.Set{
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
	return endpoint.CreateRequest{Domain: domain}, nil
}

func decodeGRPCDeleteRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.DeleteDomainRequest)
	return endpoint.DeleteRequest{Name: req.Name}, nil
}

func decodeGRPCGetRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.GetDomainRequest)
	return endpoint.GetRequest{Name: req.Name}, nil
}

func decodeGRPCListRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.ListDomainsRequest)
	return endpoint.ListRequest{
		PageSize:  int(req.PageSize),
		PageToken: req.PageToken,
	}, nil
}

func decodeGRPCUpdateRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.UpdateDomainRequest)
	domain := domainmodel.Domain{
		DNSName: req.Domain.DnsName,
	}
	return endpoint.UpdateRequest{
		Domain:     domain,
		UpdateMask: "",
		// UpdateMask: string(req.UpdateMask), // TODO: ...
	}, nil
}

// func decodeGRPCGetResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
// 	reply := grpcReply.(*pb.Domain)
// 	return endpoint.GetResponse{Domain: reply.Domain, Err: str2err(reply.Err)}, nil
// }

// func decodeGRPCListResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
// 	reply := grpcReply.(*pb.ListDomainsResponse)
// 	return endpoint.ListResponse{Domains: reply.Domains, Err: str2err(reply.Err)}, nil
// }

func encodeGRPCCreateResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.CreateResponse)
	return &pb.Domain{
		Name:    resp.Domain.ExternalName(),
		DnsName: resp.Domain.DNSName,
	}, nil
}

func encodeGRPCDeleteResponse(_ context.Context, response interface{}) (interface{}, error) {
	// resp := response.(endpoint.DeleteResponse)
	return nil, nil // TODO: empty
}

func encodeGRPCGetResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.GetResponse)
	// return &pb.Domain{Domain: resp.Domain, Err: err2str(resp.Err)}, nil
	return &pb.Domain{
		Name:    resp.Domain.ExternalName(),
		DnsName: resp.Domain.DNSName,
	}, nil
}

func encodeGRPCListResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.ListResponse)
	// return &pb.ListDomainsResponse{Domains: resp.Domains, Err: err2str(resp.Err)}, nil
	var domains []*pb.Domain
	for _, domain := range resp.Domains {
		domains = append(domains, &pb.Domain{Name: domain.ExternalName(), DnsName: domain.DNSName})
	}
	return &pb.ListDomainsResponse{Domains: domains}, nil
}

func encodeGRPCUpdateResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.UpdateResponse)
	// return &pb.Domain{Domain: resp.Domain, Err: err2str(resp.Err)}, nil
	return &pb.Domain{
		Name:    resp.Domain.ExternalName(),
		DnsName: resp.Domain.DNSName,
	}, nil
}

// func encodeGRPCGetRequest(_ context.Context, request interface{}) (interface{}, error) {
// 	req := request.(endpoint.GetRequest)
// 	return &pb.GetDomainRequest{Name: req.Name}, nil
// }

// func encodeGRPCListRequest(_ context.Context, request interface{}) (interface{}, error) {
// 	req := request.(endpoint.ListRequest)
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
