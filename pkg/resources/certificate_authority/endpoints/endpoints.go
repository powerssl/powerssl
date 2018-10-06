package endpoints

import (
	"context"
	//"fmt"
	//"reflect"
	//"strings"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"

	"powerssl.io/pkg/api"
	"powerssl.io/pkg/resources/certificate_authority/service"
	"powerssl.io/pkg/resources/endpoints"
)

type Endpoints struct {
	CreateEndpoint endpoint.Endpoint `endpoint:"Create,logging,instrumenting"`
	DeleteEndpoint endpoint.Endpoint `endpoint:"Delete,logging,instrumenting"`
	GetEndpoint    endpoint.Endpoint `endpoint:"Get,logging,instrumenting"`
	ListEndpoint   endpoint.Endpoint `endpoint:"List,logging,instrumenting"`
	UpdateEndpoint endpoint.Endpoint `endpoint:"Update,logging,instrumenting"`
}

const tagName = "endpoint"

func New(svc service.Service, logger log.Logger, duration metrics.Histogram) Endpoints {
	var createEndpoint endpoint.Endpoint
	{
		createEndpoint = makeCreateEndpoint(svc)
		createEndpoint = endpoints.LoggingMiddleware(log.With(logger, "method", "Create"))(createEndpoint)
		createEndpoint = endpoints.InstrumentingMiddleware(duration.With("method", "Create"))(createEndpoint)
	}

	var deleteEndpoint endpoint.Endpoint
	{
		deleteEndpoint = makeDeleteEndpoint(svc)
		deleteEndpoint = endpoints.LoggingMiddleware(log.With(logger, "method", "Delete"))(deleteEndpoint)
		deleteEndpoint = endpoints.InstrumentingMiddleware(duration.With("method", "Delete"))(deleteEndpoint)
	}

	var getEndpoint endpoint.Endpoint
	{
		getEndpoint = makeGetEndpoint(svc)
		getEndpoint = endpoints.LoggingMiddleware(log.With(logger, "method", "Get"))(getEndpoint)
		getEndpoint = endpoints.InstrumentingMiddleware(duration.With("method", "Get"))(getEndpoint)
	}

	var listEndpoint endpoint.Endpoint
	{
		listEndpoint = makeListEndpoint(svc)
		listEndpoint = endpoints.LoggingMiddleware(log.With(logger, "method", "List"))(listEndpoint)
		listEndpoint = endpoints.InstrumentingMiddleware(duration.With("method", "List"))(listEndpoint)
	}

	var updateEndpoint endpoint.Endpoint
	{
		updateEndpoint = makeUpdateEndpoint(svc)
		updateEndpoint = endpoints.LoggingMiddleware(log.With(logger, "method", "Update"))(updateEndpoint)
		updateEndpoint = endpoints.InstrumentingMiddleware(duration.With("method", "Update"))(updateEndpoint)
	}

	endpointss := Endpoints{
		CreateEndpoint: createEndpoint,
		DeleteEndpoint: deleteEndpoint,
		GetEndpoint:    getEndpoint,
		ListEndpoint:   listEndpoint,
		UpdateEndpoint: updateEndpoint,
	}

	//v := reflect.ValueOf(endpointss)
	//for i := 0; i < v.NumField(); i++ {
	//	tag := v.Type().Field(i).Tag.Get(tagName)
	//	value := v.Field(i).Interface().(endpoint.Endpoint)

	//	if tag == "" || tag == "-" {
	//		continue
	//	}

	//	args := strings.Split(tag, ",")
	//	name := args[0]
	//	fmt.Println(name)
	//	if len(args) > 1 {
	//		for ii := 1; ii < len(args); ii++ {
	//			switch args[ii] {
	//			case "logging":
	//				//value = endpoints.LoggingMiddleware(log.With(logger, "method", name))(value)
	//				fmt.Println("LOGGING")
	//			case "instrumenting":
	//				//value = endpoints.InstrumentingMiddleware(duration.With("method", name))(value)
	//				fmt.Println("INST")
	//			}
	//		}
	//	}
	//}

	return endpointss
}

type CreateRequest struct {
	CertificateAuthority *api.CertificateAuthority
}

type CreateResponse struct {
	CertificateAuthority *api.CertificateAuthority
}

func makeCreateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		ca, err := s.Create(ctx, req.CertificateAuthority)
		if err != nil {
			return nil, err
		}
		return CreateResponse{
			CertificateAuthority: ca,
		}, nil
	}
}

type DeleteRequest struct {
	Name string
}

type DeleteResponse struct {
}

func makeDeleteEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRequest)
		if err := s.Delete(ctx, req.Name); err != nil {
			return nil, err
		}
		return DeleteResponse{}, nil
	}
}

type GetRequest struct {
	Name string
}

type GetResponse struct {
	CertificateAuthority *api.CertificateAuthority
}

func makeGetEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRequest)
		ca, err := s.Get(ctx, req.Name)
		if err != nil {
			return nil, err
		}
		return GetResponse{
			CertificateAuthority: ca,
		}, nil
	}
}

type ListRequest struct{}

type ListResponse struct {
	CertificateAuthorities []*api.CertificateAuthority
}

func makeListEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		cas, err := s.List(ctx)
		if err != nil {
			return nil, err
		}
		return ListResponse{
			CertificateAuthorities: cas,
		}, nil
	}
}

type UpdateRequest struct {
	CertificateAuthority *api.CertificateAuthority
}

type UpdateResponse struct {
	CertificateAuthority *api.CertificateAuthority
}

func makeUpdateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateRequest)
		ca, err := s.Update(ctx, req.CertificateAuthority)
		if err != nil {
			return nil, err
		}
		return UpdateResponse{
			CertificateAuthority: ca,
		}, nil
	}
}
