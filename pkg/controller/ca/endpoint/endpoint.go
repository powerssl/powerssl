package endpoint // import "powerssl.io/pkg/controller/ca/endpoint"

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"

	"powerssl.io/pkg/controller/api"
	service "powerssl.io/pkg/controller/ca/service"
	resource "powerssl.io/pkg/resource"
)

type Endpoints struct {
	GetAuthorizeDomainRequestEndpoint  endpoint.Endpoint
	SetAuthorizeDomainResponseEndpoint endpoint.Endpoint

	GetRequestCertificateRequestEndpoint  endpoint.Endpoint
	SetRequestCertificateResponseEndpoint endpoint.Endpoint

	GetRevokeCertificateRequestEndpoint  endpoint.Endpoint
	SetRevokeCertificateResponseEndpoint endpoint.Endpoint

	GetVerifyDomainRequestEndpoint  endpoint.Endpoint
	SetVerifyDomainResponseEndpoint endpoint.Endpoint
}

func NewEndpoints(svc service.Service, logger log.Logger, duration metrics.Histogram) Endpoints {
	var getAuthorizeDomainRequestEndpoint endpoint.Endpoint
	{
		getAuthorizeDomainRequestEndpoint = makeGetAuthorizeDomainRequestEndpoint(svc)
		getAuthorizeDomainRequestEndpoint = resource.LoggingMiddleware(log.With(logger, "method", "GetAuthorizeDomainRequest"))(getAuthorizeDomainRequestEndpoint)
		getAuthorizeDomainRequestEndpoint = resource.InstrumentingMiddleware(duration.With("method", "GetAuthorizeDomainRequest"))(getAuthorizeDomainRequestEndpoint)
	}

	var setAuthorizeDomainResponseEndpoint endpoint.Endpoint
	{
		setAuthorizeDomainResponseEndpoint = makeSetAuthorizeDomainResponseEndpoint(svc)
		setAuthorizeDomainResponseEndpoint = resource.LoggingMiddleware(log.With(logger, "method", "SetAuthorizeDomainRequest"))(setAuthorizeDomainResponseEndpoint)
		setAuthorizeDomainResponseEndpoint = resource.InstrumentingMiddleware(duration.With("method", "SetAuthorizeDomainRequest"))(setAuthorizeDomainResponseEndpoint)
	}

	var getRequestCertificateRequestEndpoint endpoint.Endpoint
	{
		getRequestCertificateRequestEndpoint = makeGetRequestCertificateRequestEndpoint(svc)
		getRequestCertificateRequestEndpoint = resource.LoggingMiddleware(log.With(logger, "method", "GetRequestCertificateRequest"))(getRequestCertificateRequestEndpoint)
		getRequestCertificateRequestEndpoint = resource.InstrumentingMiddleware(duration.With("method", "GetRequestCertificateRequest"))(getRequestCertificateRequestEndpoint)
	}

	var setRequestCertificateResponseEndpoint endpoint.Endpoint
	{
		setRequestCertificateResponseEndpoint = makeSetRequestCertificateResponseEndpoint(svc)
		setRequestCertificateResponseEndpoint = resource.LoggingMiddleware(log.With(logger, "method", "SetRequestCertificateRequest"))(setRequestCertificateResponseEndpoint)
		setRequestCertificateResponseEndpoint = resource.InstrumentingMiddleware(duration.With("method", "SetRequestCertificateRequest"))(setRequestCertificateResponseEndpoint)
	}

	var getRevokeCertificateRequestEndpoint endpoint.Endpoint
	{
		getRevokeCertificateRequestEndpoint = makeGetRevokeCertificateRequestEndpoint(svc)
		getRevokeCertificateRequestEndpoint = resource.LoggingMiddleware(log.With(logger, "method", "GetRevokeCertificateRequest"))(getRevokeCertificateRequestEndpoint)
		getRevokeCertificateRequestEndpoint = resource.InstrumentingMiddleware(duration.With("method", "GetRevokeCertificateRequest"))(getRevokeCertificateRequestEndpoint)
	}

	var setRevokeCertificateResponseEndpoint endpoint.Endpoint
	{
		setRevokeCertificateResponseEndpoint = makeSetRevokeCertificateResponseEndpoint(svc)
		setRevokeCertificateResponseEndpoint = resource.LoggingMiddleware(log.With(logger, "method", "SetRevokeCertificateRequest"))(setRevokeCertificateResponseEndpoint)
		setRevokeCertificateResponseEndpoint = resource.InstrumentingMiddleware(duration.With("method", "SetRevokeCertificateRequest"))(setRevokeCertificateResponseEndpoint)
	}

	var getVerifyDomainRequestEndpoint endpoint.Endpoint
	{
		getVerifyDomainRequestEndpoint = makeGetVerifyDomainRequestEndpoint(svc)
		getVerifyDomainRequestEndpoint = resource.LoggingMiddleware(log.With(logger, "method", "GetVerifyDomainRequest"))(getVerifyDomainRequestEndpoint)
		getVerifyDomainRequestEndpoint = resource.InstrumentingMiddleware(duration.With("method", "GetVerifyDomainRequest"))(getVerifyDomainRequestEndpoint)
	}

	var setVerifyDomainResponseEndpoint endpoint.Endpoint
	{
		setVerifyDomainResponseEndpoint = makeSetVerifyDomainResponseEndpoint(svc)
		setVerifyDomainResponseEndpoint = resource.LoggingMiddleware(log.With(logger, "method", "SetVerifyDomainRequest"))(setVerifyDomainResponseEndpoint)
		setVerifyDomainResponseEndpoint = resource.InstrumentingMiddleware(duration.With("method", "SetVerifyDomainRequest"))(setVerifyDomainResponseEndpoint)
	}

	return Endpoints{
		GetAuthorizeDomainRequestEndpoint:  getAuthorizeDomainRequestEndpoint,
		SetAuthorizeDomainResponseEndpoint: setAuthorizeDomainResponseEndpoint,

		GetRequestCertificateRequestEndpoint:  getRequestCertificateRequestEndpoint,
		SetRequestCertificateResponseEndpoint: setRequestCertificateResponseEndpoint,

		GetRevokeCertificateRequestEndpoint:  getRevokeCertificateRequestEndpoint,
		SetRevokeCertificateResponseEndpoint: setRevokeCertificateResponseEndpoint,

		GetVerifyDomainRequestEndpoint:  getVerifyDomainRequestEndpoint,
		SetVerifyDomainResponseEndpoint: setVerifyDomainResponseEndpoint,
	}
}

func (e Endpoints) GetAuthorizeDomainRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, error) {
	resp, err := e.GetAuthorizeDomainRequestEndpoint(ctx, GetAuthorizeDomainRequestRequest{
		Activity: activity,
	})
	if err != nil {
		return nil, "", err
	}
	response := resp.(GetAuthorizeDomainRequestResponse)
	return response.Activity, response.Domain, nil
}

type GetAuthorizeDomainRequestRequest struct {
	Activity *api.Activity
}

type GetAuthorizeDomainRequestResponse struct {
	Activity *api.Activity
	Domain   string
}

func makeGetAuthorizeDomainRequestEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAuthorizeDomainRequestRequest)
		activity, domain, err := s.GetAuthorizeDomainRequest(ctx, req.Activity)
		if err != nil {
			return nil, err
		}
		return GetAuthorizeDomainRequestResponse{
			Activity: activity,
			Domain:   domain,
		}, nil
	}
}

func (e Endpoints) SetAuthorizeDomainResponse(ctx context.Context, activity *api.Activity, erro *api.Error, challenges []*api.Challenge) error {
	_, err := e.SetAuthorizeDomainResponseEndpoint(ctx, SetAuthorizeDomainResponseRequest{
		Activity:   activity,
		Error:      erro,
		Challenges: challenges,
	})
	if err != nil {
		return err
	}
	return nil
}

type SetAuthorizeDomainResponseRequest struct {
	Activity   *api.Activity
	Error      *api.Error
	Challenges []*api.Challenge
}

type SetAuthorizeDomainResponseResponse struct{}

func makeSetAuthorizeDomainResponseEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SetAuthorizeDomainResponseRequest)
		err := s.SetAuthorizeDomainResponse(ctx, req.Activity, req.Error, req.Challenges)
		if err != nil {
			return nil, err
		}
		return SetAuthorizeDomainResponseResponse{}, nil
	}
}

//////////////////////////////////

func (e Endpoints) GetRequestCertificateRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, error) {
	resp, err := e.GetRequestCertificateRequestEndpoint(ctx, GetRequestCertificateRequestRequest{
		Activity: activity,
	})
	if err != nil {
		return nil, "", err
	}
	response := resp.(GetRequestCertificateRequestResponse)
	return response.Activity, response.CertificateSigningRequest, nil
}

type GetRequestCertificateRequestRequest struct {
	Activity *api.Activity
}

type GetRequestCertificateRequestResponse struct {
	Activity                  *api.Activity
	CertificateSigningRequest string
}

func makeGetRequestCertificateRequestEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRequestCertificateRequestRequest)
		activity, certificateSigningRequest, err := s.GetRequestCertificateRequest(ctx, req.Activity)
		if err != nil {
			return nil, err
		}
		return GetRequestCertificateRequestResponse{
			Activity:                  activity,
			CertificateSigningRequest: certificateSigningRequest,
		}, nil
	}
}

func (e Endpoints) SetRequestCertificateResponse(ctx context.Context, activity *api.Activity, erro *api.Error, certificate string) error {
	_, err := e.SetRequestCertificateResponseEndpoint(ctx, SetRequestCertificateResponseRequest{
		Activity:    activity,
		Error:       erro,
		Certificate: certificate,
	})
	if err != nil {
		return err
	}
	return nil
}

type SetRequestCertificateResponseRequest struct {
	Activity    *api.Activity
	Error       *api.Error
	Certificate string
}

type SetRequestCertificateResponseResponse struct{}

func makeSetRequestCertificateResponseEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SetRequestCertificateResponseRequest)
		err := s.SetRequestCertificateResponse(ctx, req.Activity, req.Error, req.Certificate)
		if err != nil {
			return nil, err
		}
		return SetRequestCertificateResponseResponse{}, nil
	}
}

//////////////////////////////////

func (e Endpoints) GetRevokeCertificateRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, error) {
	resp, err := e.GetRevokeCertificateRequestEndpoint(ctx, GetRevokeCertificateRequestRequest{
		Activity: activity,
	})
	if err != nil {
		return nil, "", err
	}
	response := resp.(GetRevokeCertificateRequestResponse)
	return response.Activity, response.Certificate, nil
}

type GetRevokeCertificateRequestRequest struct {
	Activity *api.Activity
}

type GetRevokeCertificateRequestResponse struct {
	Activity    *api.Activity
	Certificate string
}

func makeGetRevokeCertificateRequestEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRevokeCertificateRequestRequest)
		activity, certificate, err := s.GetRevokeCertificateRequest(ctx, req.Activity)
		if err != nil {
			return nil, err
		}
		return GetRevokeCertificateRequestResponse{
			Activity:    activity,
			Certificate: certificate,
		}, nil
	}
}

func (e Endpoints) SetRevokeCertificateResponse(ctx context.Context, activity *api.Activity, erro *api.Error) error {
	_, err := e.SetRevokeCertificateResponseEndpoint(ctx, SetRevokeCertificateResponseRequest{
		Activity: activity,
		Error:    erro,
	})
	if err != nil {
		return err
	}
	return nil
}

type SetRevokeCertificateResponseRequest struct {
	Activity *api.Activity
	Error    *api.Error
}

type SetRevokeCertificateResponseResponse struct{}

func makeSetRevokeCertificateResponseEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SetRevokeCertificateResponseRequest)
		err := s.SetRevokeCertificateResponse(ctx, req.Activity, req.Error)
		if err != nil {
			return nil, err
		}
		return SetRevokeCertificateResponseResponse{}, nil
	}
}

//////////////////////////////////

func (e Endpoints) GetVerifyDomainRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, api.ChallengeType, error) {
	resp, err := e.GetVerifyDomainRequestEndpoint(ctx, GetVerifyDomainRequestRequest{
		Activity: activity,
	})
	if err != nil {
		return nil, "", 0, err
	}
	response := resp.(GetVerifyDomainRequestResponse)
	return response.Activity, response.Domain, response.ChallengeType, nil
}

type GetVerifyDomainRequestRequest struct {
	Activity *api.Activity
}

type GetVerifyDomainRequestResponse struct {
	Activity      *api.Activity
	Domain        string
	ChallengeType api.ChallengeType
}

func makeGetVerifyDomainRequestEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetVerifyDomainRequestRequest)
		activity, domain, challengeType, err := s.GetVerifyDomainRequest(ctx, req.Activity)
		if err != nil {
			return nil, err
		}
		return GetVerifyDomainRequestResponse{
			Activity:      activity,
			Domain:        domain,
			ChallengeType: challengeType,
		}, nil
	}
}

func (e Endpoints) SetVerifyDomainResponse(ctx context.Context, activity *api.Activity, erro *api.Error) error {
	_, err := e.SetVerifyDomainResponseEndpoint(ctx, SetVerifyDomainResponseRequest{
		Activity: activity,
		Error:    erro,
	})
	if err != nil {
		return err
	}
	return nil
}

type SetVerifyDomainResponseRequest struct {
	Activity *api.Activity
	Error    *api.Error
}

type SetVerifyDomainResponseResponse struct{}

func makeSetVerifyDomainResponseEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SetVerifyDomainResponseRequest)
		err := s.SetVerifyDomainResponse(ctx, req.Activity, req.Error)
		if err != nil {
			return nil, err
		}
		return SetVerifyDomainResponseResponse{}, nil
	}
}
