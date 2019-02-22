package endpoint

import (
	"context"
	"crypto/x509"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"

	"github.com/go-kit/kit/tracing/opentracing"
	stdopentracing "github.com/opentracing/opentracing-go"
	"powerssl.io/internal/pkg/util/middleware"
	"powerssl.io/pkg/controller/acme"
	"powerssl.io/pkg/controller/api"
)

type Endpoints struct {
	GetCreateAccountRequestEndpoint  endpoint.Endpoint
	SetCreateAccountResponseEndpoint endpoint.Endpoint

	GetDeactivateAccountRequestEndpoint  endpoint.Endpoint
	SetDeactivateAccountResponseEndpoint endpoint.Endpoint

	GetRekeyAccountRequestEndpoint  endpoint.Endpoint
	SetRekeyAccountResponseEndpoint endpoint.Endpoint

	GetUpdateAccountRequestEndpoint  endpoint.Endpoint
	SetUpdateAccountResponseEndpoint endpoint.Endpoint

	GetCreateOrderRequestEndpoint  endpoint.Endpoint
	SetCreateOrderResponseEndpoint endpoint.Endpoint

	GetFinalizeOrderRequestEndpoint  endpoint.Endpoint
	SetFinalizeOrderResponseEndpoint endpoint.Endpoint

	GetGetOrderRequestEndpoint  endpoint.Endpoint
	SetGetOrderResponseEndpoint endpoint.Endpoint

	GetCreateAuthorizationRequestEndpoint  endpoint.Endpoint
	SetCreateAuthorizationResponseEndpoint endpoint.Endpoint

	GetDeactivateAuthorizationRequestEndpoint  endpoint.Endpoint
	SetDeactivateAuthorizationResponseEndpoint endpoint.Endpoint

	GetGetAuthorizationRequestEndpoint  endpoint.Endpoint
	SetGetAuthorizationResponseEndpoint endpoint.Endpoint

	GetGetChallengeRequestEndpoint  endpoint.Endpoint
	SetGetChallengeResponseEndpoint endpoint.Endpoint

	GetValidateChallengeRequestEndpoint  endpoint.Endpoint
	SetValidateChallengeResponseEndpoint endpoint.Endpoint

	GetGetCertificateRequestEndpoint  endpoint.Endpoint
	SetGetCertificateResponseEndpoint endpoint.Endpoint

	GetRevokeCertificateRequestEndpoint  endpoint.Endpoint
	SetRevokeCertificateResponseEndpoint endpoint.Endpoint
}

func NewEndpoints(svc acme.Service, logger log.Logger, tracer stdopentracing.Tracer, duration metrics.Histogram) Endpoints {
	var getCreateAccountRequestEndpoint endpoint.Endpoint
	{
		getCreateAccountRequestEndpoint = makeGetCreateAccountRequestEndpoint(svc)
		getCreateAccountRequestEndpoint = opentracing.TraceServer(tracer, "GetCreateAccountRequest")(getCreateAccountRequestEndpoint)
		getCreateAccountRequestEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "GetCreateAccountRequest"))(getCreateAccountRequestEndpoint)
		getCreateAccountRequestEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "GetCreateAccountRequest"))(getCreateAccountRequestEndpoint)
	}

	var setCreateAccountResponseEndpoint endpoint.Endpoint
	{
		setCreateAccountResponseEndpoint = makeSetCreateAccountResponseEndpoint(svc)
		setCreateAccountResponseEndpoint = opentracing.TraceServer(tracer, "SetCreateAccountResponse")(setCreateAccountResponseEndpoint)
		setCreateAccountResponseEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "SetCreateAccountResponse"))(setCreateAccountResponseEndpoint)
		setCreateAccountResponseEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "SetCreateAccountResponse"))(setCreateAccountResponseEndpoint)
	}

	var getDeactivateAccountRequestEndpoint endpoint.Endpoint
	{
		getDeactivateAccountRequestEndpoint = makeGetDeactivateAccountRequestEndpoint(svc)
		getDeactivateAccountRequestEndpoint = opentracing.TraceServer(tracer, "GetDeactivateAccountRequest")(getDeactivateAccountRequestEndpoint)
		getDeactivateAccountRequestEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "GetDeactivateAccountRequest"))(getDeactivateAccountRequestEndpoint)
		getDeactivateAccountRequestEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "GetDeactivateAccountRequest"))(getDeactivateAccountRequestEndpoint)
	}

	var setDeactivateAccountResponseEndpoint endpoint.Endpoint
	{
		setDeactivateAccountResponseEndpoint = makeSetDeactivateAccountResponseEndpoint(svc)
		setDeactivateAccountResponseEndpoint = opentracing.TraceServer(tracer, "SetDeactivateAccountResponse")(setDeactivateAccountResponseEndpoint)
		setDeactivateAccountResponseEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "SetDeactivateAccountResponse"))(setDeactivateAccountResponseEndpoint)
		setDeactivateAccountResponseEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "SetDeactivateAccountResponse"))(setDeactivateAccountResponseEndpoint)
	}

	var getRekeyAccountRequestEndpoint endpoint.Endpoint
	{
		getRekeyAccountRequestEndpoint = makeGetRekeyAccountRequestEndpoint(svc)
		getRekeyAccountRequestEndpoint = opentracing.TraceServer(tracer, "GetRekeyAccountRequest")(getRekeyAccountRequestEndpoint)
		getRekeyAccountRequestEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "GetRekeyAccountRequest"))(getRekeyAccountRequestEndpoint)
		getRekeyAccountRequestEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "GetRekeyAccountRequest"))(getRekeyAccountRequestEndpoint)
	}

	var setRekeyAccountResponseEndpoint endpoint.Endpoint
	{
		setRekeyAccountResponseEndpoint = makeSetRekeyAccountResponseEndpoint(svc)
		setRekeyAccountResponseEndpoint = opentracing.TraceServer(tracer, "SetRekeyAccountResponse")(setRekeyAccountResponseEndpoint)
		setRekeyAccountResponseEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "SetRekeyAccountResponse"))(setRekeyAccountResponseEndpoint)
		setRekeyAccountResponseEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "SetRekeyAccountResponse"))(setRekeyAccountResponseEndpoint)
	}

	var getUpdateAccountRequestEndpoint endpoint.Endpoint
	{
		getUpdateAccountRequestEndpoint = makeGetUpdateAccountRequestEndpoint(svc)
		getUpdateAccountRequestEndpoint = opentracing.TraceServer(tracer, "GetUpdateAccountRequest")(getUpdateAccountRequestEndpoint)
		getUpdateAccountRequestEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "GetUpdateAccountRequest"))(getUpdateAccountRequestEndpoint)
		getUpdateAccountRequestEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "GetUpdateAccountRequest"))(getUpdateAccountRequestEndpoint)
	}

	var setUpdateAccountResponseEndpoint endpoint.Endpoint
	{
		setUpdateAccountResponseEndpoint = makeSetUpdateAccountResponseEndpoint(svc)
		setUpdateAccountResponseEndpoint = opentracing.TraceServer(tracer, "SetUpdateAccountResponse")(setUpdateAccountResponseEndpoint)
		setUpdateAccountResponseEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "SetUpdateAccountResponse"))(setUpdateAccountResponseEndpoint)
		setUpdateAccountResponseEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "SetUpdateAccountResponse"))(setUpdateAccountResponseEndpoint)
	}

	var getCreateOrderRequestEndpoint endpoint.Endpoint
	{
		getCreateOrderRequestEndpoint = makeGetCreateOrderRequestEndpoint(svc)
		getCreateOrderRequestEndpoint = opentracing.TraceServer(tracer, "GetCreateOrderRequest")(getCreateOrderRequestEndpoint)
		getCreateOrderRequestEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "GetCreateOrderRequest"))(getCreateOrderRequestEndpoint)
		getCreateOrderRequestEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "GetCreateOrderRequest"))(getCreateOrderRequestEndpoint)
	}

	var setCreateOrderResponseEndpoint endpoint.Endpoint
	{
		setCreateOrderResponseEndpoint = makeSetCreateOrderResponseEndpoint(svc)
		setCreateOrderResponseEndpoint = opentracing.TraceServer(tracer, "SetCreateOrderResponse")(setCreateOrderResponseEndpoint)
		setCreateOrderResponseEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "SetCreateOrderResponse"))(setCreateOrderResponseEndpoint)
		setCreateOrderResponseEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "SetCreateOrderResponse"))(setCreateOrderResponseEndpoint)
	}

	var getFinalizeOrderRequestEndpoint endpoint.Endpoint
	{
		getFinalizeOrderRequestEndpoint = makeGetFinalizeOrderRequestEndpoint(svc)
		getFinalizeOrderRequestEndpoint = opentracing.TraceServer(tracer, "GetFinalizeOrderRequest")(getFinalizeOrderRequestEndpoint)
		getFinalizeOrderRequestEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "GetFinalizeOrderRequest"))(getFinalizeOrderRequestEndpoint)
		getFinalizeOrderRequestEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "GetFinalizeOrderRequest"))(getFinalizeOrderRequestEndpoint)
	}

	var setFinalizeOrderResponseEndpoint endpoint.Endpoint
	{
		setFinalizeOrderResponseEndpoint = makeSetFinalizeOrderResponseEndpoint(svc)
		setFinalizeOrderResponseEndpoint = opentracing.TraceServer(tracer, "SetFinalizeOrderResponse")(setFinalizeOrderResponseEndpoint)
		setFinalizeOrderResponseEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "SetFinalizeOrderResponse"))(setFinalizeOrderResponseEndpoint)
		setFinalizeOrderResponseEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "SetFinalizeOrderResponse"))(setFinalizeOrderResponseEndpoint)
	}

	var getGetOrderRequestEndpoint endpoint.Endpoint
	{
		getGetOrderRequestEndpoint = makeGetGetOrderRequestEndpoint(svc)
		getGetOrderRequestEndpoint = opentracing.TraceServer(tracer, "GetGetOrderRequest")(getGetOrderRequestEndpoint)
		getGetOrderRequestEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "GetGetOrderRequest"))(getGetOrderRequestEndpoint)
		getGetOrderRequestEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "GetGetOrderRequest"))(getGetOrderRequestEndpoint)
	}

	var setGetOrderResponseEndpoint endpoint.Endpoint
	{
		setGetOrderResponseEndpoint = makeSetGetOrderResponseEndpoint(svc)
		setGetOrderResponseEndpoint = opentracing.TraceServer(tracer, "SetGetOrderResponse")(setGetOrderResponseEndpoint)
		setGetOrderResponseEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "SetGetOrderResponse"))(setGetOrderResponseEndpoint)
		setGetOrderResponseEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "SetGetOrderResponse"))(setGetOrderResponseEndpoint)
	}

	var getCreateAuthorizationRequestEndpoint endpoint.Endpoint
	{
		getCreateAuthorizationRequestEndpoint = makeGetCreateAuthorizationRequestEndpoint(svc)
		getCreateAuthorizationRequestEndpoint = opentracing.TraceServer(tracer, "GetCreateAuthorizationRequest")(getCreateAuthorizationRequestEndpoint)
		getCreateAuthorizationRequestEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "GetCreateAuthorizationRequest"))(getCreateAuthorizationRequestEndpoint)
		getCreateAuthorizationRequestEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "GetCreateAuthorizationRequest"))(getCreateAuthorizationRequestEndpoint)
	}

	var setCreateAuthorizationResponseEndpoint endpoint.Endpoint
	{
		setCreateAuthorizationResponseEndpoint = makeSetCreateAuthorizationResponseEndpoint(svc)
		setCreateAuthorizationResponseEndpoint = opentracing.TraceServer(tracer, "SetCreateAuthorizationResponse")(setCreateAuthorizationResponseEndpoint)
		setCreateAuthorizationResponseEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "SetCreateAuthorizationResponse"))(setCreateAuthorizationResponseEndpoint)
		setCreateAuthorizationResponseEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "SetCreateAuthorizationResponse"))(setCreateAuthorizationResponseEndpoint)
	}

	var getDeactivateAuthorizationRequestEndpoint endpoint.Endpoint
	{
		getDeactivateAuthorizationRequestEndpoint = makeGetDeactivateAuthorizationRequestEndpoint(svc)
		getDeactivateAuthorizationRequestEndpoint = opentracing.TraceServer(tracer, "GetDeactivateAuthorizationRequest")(getDeactivateAuthorizationRequestEndpoint)
		getDeactivateAuthorizationRequestEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "GetDeactivateAuthorizationRequest"))(getDeactivateAuthorizationRequestEndpoint)
		getDeactivateAuthorizationRequestEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "GetDeactivateAuthorizationRequest"))(getDeactivateAuthorizationRequestEndpoint)
	}

	var setDeactivateAuthorizationResponseEndpoint endpoint.Endpoint
	{
		setDeactivateAuthorizationResponseEndpoint = makeSetDeactivateAuthorizationResponseEndpoint(svc)
		setDeactivateAuthorizationResponseEndpoint = opentracing.TraceServer(tracer, "SetDeactivateAuthorizationResponse")(setDeactivateAuthorizationResponseEndpoint)
		setDeactivateAuthorizationResponseEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "SetDeactivateAuthorizationResponse"))(setDeactivateAuthorizationResponseEndpoint)
		setDeactivateAuthorizationResponseEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "SetDeactivateAuthorizationResponse"))(setDeactivateAuthorizationResponseEndpoint)
	}

	var getGetAuthorizationRequestEndpoint endpoint.Endpoint
	{
		getGetAuthorizationRequestEndpoint = makeGetGetAuthorizationRequestEndpoint(svc)
		getGetAuthorizationRequestEndpoint = opentracing.TraceServer(tracer, "GetGetAuthorizationRequest")(getGetAuthorizationRequestEndpoint)
		getGetAuthorizationRequestEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "GetGetAuthorizationRequest"))(getGetAuthorizationRequestEndpoint)
		getGetAuthorizationRequestEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "GetGetAuthorizationRequest"))(getGetAuthorizationRequestEndpoint)
	}

	var setGetAuthorizationResponseEndpoint endpoint.Endpoint
	{
		setGetAuthorizationResponseEndpoint = makeSetGetAuthorizationResponseEndpoint(svc)
		setGetAuthorizationResponseEndpoint = opentracing.TraceServer(tracer, "SetGetAuthorizationResponse")(setGetAuthorizationResponseEndpoint)
		setGetAuthorizationResponseEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "SetGetAuthorizationResponse"))(setGetAuthorizationResponseEndpoint)
		setGetAuthorizationResponseEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "SetGetAuthorizationResponse"))(setGetAuthorizationResponseEndpoint)
	}

	var getGetChallengeRequestEndpoint endpoint.Endpoint
	{
		getGetChallengeRequestEndpoint = makeGetGetChallengeRequestEndpoint(svc)
		getGetChallengeRequestEndpoint = opentracing.TraceServer(tracer, "GetGetChallengeRequest")(getGetChallengeRequestEndpoint)
		getGetChallengeRequestEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "GetGetChallengeRequest"))(getGetChallengeRequestEndpoint)
		getGetChallengeRequestEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "GetGetChallengeRequest"))(getGetChallengeRequestEndpoint)
	}

	var setGetChallengeResponseEndpoint endpoint.Endpoint
	{
		setGetChallengeResponseEndpoint = makeSetGetChallengeResponseEndpoint(svc)
		setGetChallengeResponseEndpoint = opentracing.TraceServer(tracer, "SetGetChallengeResponse")(setGetChallengeResponseEndpoint)
		setGetChallengeResponseEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "SetGetChallengeResponse"))(setGetChallengeResponseEndpoint)
		setGetChallengeResponseEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "SetGetChallengeResponse"))(setGetChallengeResponseEndpoint)
	}

	var getValidateChallengeRequestEndpoint endpoint.Endpoint
	{
		getValidateChallengeRequestEndpoint = makeGetValidateChallengeRequestEndpoint(svc)
		getValidateChallengeRequestEndpoint = opentracing.TraceServer(tracer, "GetValidateChallengeRequest")(getValidateChallengeRequestEndpoint)
		getValidateChallengeRequestEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "GetValidateChallengeRequest"))(getValidateChallengeRequestEndpoint)
		getValidateChallengeRequestEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "GetValidateChallengeRequest"))(getValidateChallengeRequestEndpoint)
	}

	var setValidateChallengeResponseEndpoint endpoint.Endpoint
	{
		setValidateChallengeResponseEndpoint = makeSetValidateChallengeResponseEndpoint(svc)
		setValidateChallengeResponseEndpoint = opentracing.TraceServer(tracer, "SetValidateChallengeResponse")(setValidateChallengeResponseEndpoint)
		setValidateChallengeResponseEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "SetValidateChallengeResponse"))(setValidateChallengeResponseEndpoint)
		setValidateChallengeResponseEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "SetValidateChallengeResponse"))(setValidateChallengeResponseEndpoint)
	}

	var getGetCertificateRequestEndpoint endpoint.Endpoint
	{
		getGetCertificateRequestEndpoint = makeGetGetCertificateRequestEndpoint(svc)
		getGetCertificateRequestEndpoint = opentracing.TraceServer(tracer, "GetGetCertificateRequest")(getGetCertificateRequestEndpoint)
		getGetCertificateRequestEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "GetGetCertificateRequest"))(getGetCertificateRequestEndpoint)
		getGetCertificateRequestEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "GetGetCertificateRequest"))(getGetCertificateRequestEndpoint)
	}

	var setGetCertificateResponseEndpoint endpoint.Endpoint
	{
		setGetCertificateResponseEndpoint = makeSetGetCertificateResponseEndpoint(svc)
		setGetCertificateResponseEndpoint = opentracing.TraceServer(tracer, "SetGetCertificateResponse")(setGetCertificateResponseEndpoint)
		setGetCertificateResponseEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "SetGetCertificateResponse"))(setGetCertificateResponseEndpoint)
		setGetCertificateResponseEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "SetGetCertificateResponse"))(setGetCertificateResponseEndpoint)
	}

	var getRevokeCertificateRequestEndpoint endpoint.Endpoint
	{
		getRevokeCertificateRequestEndpoint = makeGetRevokeCertificateRequestEndpoint(svc)
		getRevokeCertificateRequestEndpoint = opentracing.TraceServer(tracer, "GetRevokeCertificateRequest")(getRevokeCertificateRequestEndpoint)
		getRevokeCertificateRequestEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "GetRevokeCertificateRequest"))(getRevokeCertificateRequestEndpoint)
		getRevokeCertificateRequestEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "GetRevokeCertificateRequest"))(getRevokeCertificateRequestEndpoint)
	}

	var setRevokeCertificateResponseEndpoint endpoint.Endpoint
	{
		setRevokeCertificateResponseEndpoint = makeSetRevokeCertificateResponseEndpoint(svc)
		setRevokeCertificateResponseEndpoint = opentracing.TraceServer(tracer, "SetRevokeCertificateResponse")(setRevokeCertificateResponseEndpoint)
		setRevokeCertificateResponseEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "SetRevokeCertificateResponse"))(setRevokeCertificateResponseEndpoint)
		setRevokeCertificateResponseEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "SetRevokeCertificateResponse"))(setRevokeCertificateResponseEndpoint)
	}

	return Endpoints{
		GetCreateAccountRequestEndpoint:  getCreateAccountRequestEndpoint,
		SetCreateAccountResponseEndpoint: setCreateAccountResponseEndpoint,

		GetDeactivateAccountRequestEndpoint:  getDeactivateAccountRequestEndpoint,
		SetDeactivateAccountResponseEndpoint: setDeactivateAccountResponseEndpoint,

		GetRekeyAccountRequestEndpoint:  getRekeyAccountRequestEndpoint,
		SetRekeyAccountResponseEndpoint: setRekeyAccountResponseEndpoint,

		GetUpdateAccountRequestEndpoint:  getUpdateAccountRequestEndpoint,
		SetUpdateAccountResponseEndpoint: setUpdateAccountResponseEndpoint,

		GetCreateOrderRequestEndpoint:  getCreateOrderRequestEndpoint,
		SetCreateOrderResponseEndpoint: setCreateOrderResponseEndpoint,

		GetFinalizeOrderRequestEndpoint:  getFinalizeOrderRequestEndpoint,
		SetFinalizeOrderResponseEndpoint: setFinalizeOrderResponseEndpoint,

		GetGetOrderRequestEndpoint:  getGetOrderRequestEndpoint,
		SetGetOrderResponseEndpoint: setGetOrderResponseEndpoint,

		GetCreateAuthorizationRequestEndpoint:  getCreateAuthorizationRequestEndpoint,
		SetCreateAuthorizationResponseEndpoint: setCreateAuthorizationResponseEndpoint,

		GetDeactivateAuthorizationRequestEndpoint:  getDeactivateAuthorizationRequestEndpoint,
		SetDeactivateAuthorizationResponseEndpoint: setDeactivateAuthorizationResponseEndpoint,

		GetGetAuthorizationRequestEndpoint:  getGetAuthorizationRequestEndpoint,
		SetGetAuthorizationResponseEndpoint: setGetAuthorizationResponseEndpoint,

		GetGetChallengeRequestEndpoint:  getGetChallengeRequestEndpoint,
		SetGetChallengeResponseEndpoint: setGetChallengeResponseEndpoint,

		GetValidateChallengeRequestEndpoint:  getValidateChallengeRequestEndpoint,
		SetValidateChallengeResponseEndpoint: setValidateChallengeResponseEndpoint,

		GetGetCertificateRequestEndpoint:  getGetCertificateRequestEndpoint,
		SetGetCertificateResponseEndpoint: setGetCertificateResponseEndpoint,

		GetRevokeCertificateRequestEndpoint:  getRevokeCertificateRequestEndpoint,
		SetRevokeCertificateResponseEndpoint: setRevokeCertificateResponseEndpoint,
	}
}

func (e Endpoints) GetCreateAccountRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, bool, []string, error) {
	resp, err := e.GetCreateAccountRequestEndpoint(ctx, GetCreateAccountRequestRequest{
		Activity: activity,
	})
	if err != nil {
		return nil, "", false, nil, err
	}
	response := resp.(GetCreateAccountRequestResponse)
	return response.Activity, response.DirectoryURL, response.TermsOfServiceAgreed, response.Contacts, nil
}

type GetCreateAccountRequestRequest struct {
	Activity *api.Activity
}

type GetCreateAccountRequestResponse struct {
	Activity             *api.Activity
	DirectoryURL         string
	TermsOfServiceAgreed bool
	Contacts             []string
}

func makeGetCreateAccountRequestEndpoint(s acme.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetCreateAccountRequestRequest)
		activity, directoryURL, termsOfServiceAgreed, contacts, err := s.GetCreateAccountRequest(ctx, req.Activity)
		if err != nil {
			return nil, err
		}
		return GetCreateAccountRequestResponse{
			Activity:             activity,
			DirectoryURL:         directoryURL,
			TermsOfServiceAgreed: termsOfServiceAgreed,
			Contacts:             contacts,
		}, nil
	}
}

func (e Endpoints) SetCreateAccountResponse(ctx context.Context, activity *api.Activity, account *api.Account, erro *api.Error) error {
	_, err := e.SetCreateAccountResponseEndpoint(ctx, SetCreateAccountResponseRequest{
		Activity: activity,
		Account:  account,
		Error:    erro,
	})
	return err
}

type SetCreateAccountResponseRequest struct {
	Activity *api.Activity
	Account  *api.Account
	Error    *api.Error
}

type SetCreateAccountResponseResponse struct{}

func makeSetCreateAccountResponseEndpoint(s acme.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SetCreateAccountResponseRequest)
		err := s.SetCreateAccountResponse(ctx, req.Activity, req.Account, req.Error)
		if err != nil {
			return nil, err
		}
		return SetCreateAccountResponseResponse{}, nil
	}
}

func (e Endpoints) GetDeactivateAccountRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, error) {
	resp, err := e.GetDeactivateAccountRequestEndpoint(ctx, GetDeactivateAccountRequestRequest{
		Activity: activity,
	})
	if err != nil {
		return nil, "", err
	}
	response := resp.(GetDeactivateAccountRequestResponse)
	return response.Activity, response.AccountURL, nil
}

type GetDeactivateAccountRequestRequest struct {
	Activity *api.Activity
}

type GetDeactivateAccountRequestResponse struct {
	Activity   *api.Activity
	AccountURL string
}

func makeGetDeactivateAccountRequestEndpoint(s acme.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetDeactivateAccountRequestRequest)
		activity, accountURL, err := s.GetDeactivateAccountRequest(ctx, req.Activity)
		if err != nil {
			return nil, err
		}
		return GetDeactivateAccountRequestResponse{
			Activity:   activity,
			AccountURL: accountURL,
		}, nil
	}
}

func (e Endpoints) SetDeactivateAccountResponse(ctx context.Context, activity *api.Activity, account *api.Account, erro *api.Error) error {
	_, err := e.SetDeactivateAccountResponseEndpoint(ctx, SetDeactivateAccountResponseRequest{
		Activity: activity,
		Account:  account,
		Error:    erro,
	})
	return err
}

type SetDeactivateAccountResponseRequest struct {
	Activity *api.Activity
	Account  *api.Account
	Error    *api.Error
}

type SetDeactivateAccountResponseResponse struct{}

func makeSetDeactivateAccountResponseEndpoint(s acme.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SetDeactivateAccountResponseRequest)
		err := s.SetDeactivateAccountResponse(ctx, req.Activity, req.Account, req.Error)
		if err != nil {
			return nil, err
		}
		return SetDeactivateAccountResponseResponse{}, nil
	}
}

func (e Endpoints) GetRekeyAccountRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, error) {
	resp, err := e.GetRekeyAccountRequestEndpoint(ctx, GetRekeyAccountRequestRequest{
		Activity: activity,
	})
	if err != nil {
		return nil, "", "", err
	}
	response := resp.(GetRekeyAccountRequestResponse)
	return response.Activity, response.AccountURL, response.DirectoryURL, nil
}

type GetRekeyAccountRequestRequest struct {
	Activity *api.Activity
}

type GetRekeyAccountRequestResponse struct {
	Activity     *api.Activity
	AccountURL   string
	DirectoryURL string
}

func makeGetRekeyAccountRequestEndpoint(s acme.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRekeyAccountRequestRequest)
		activity, accountURL, directoryURL, err := s.GetRekeyAccountRequest(ctx, req.Activity)
		if err != nil {
			return nil, err
		}
		return GetRekeyAccountRequestResponse{
			Activity:     activity,
			AccountURL:   accountURL,
			DirectoryURL: directoryURL,
		}, nil
	}
}

func (e Endpoints) SetRekeyAccountResponse(ctx context.Context, activity *api.Activity, account *api.Account, erro *api.Error) error {
	_, err := e.SetRekeyAccountResponseEndpoint(ctx, SetRekeyAccountResponseRequest{
		Activity: activity,
		Account:  account,
		Error:    erro,
	})
	return err
}

type SetRekeyAccountResponseRequest struct {
	Activity *api.Activity
	Account  *api.Account
	Error    *api.Error
}

type SetRekeyAccountResponseResponse struct{}

func makeSetRekeyAccountResponseEndpoint(s acme.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SetRekeyAccountResponseRequest)
		err := s.SetRekeyAccountResponse(ctx, req.Activity, req.Account, req.Error)
		if err != nil {
			return nil, err
		}
		return SetRekeyAccountResponseResponse{}, nil
	}
}

func (e Endpoints) GetUpdateAccountRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, []string, error) {
	resp, err := e.GetUpdateAccountRequestEndpoint(ctx, GetUpdateAccountRequestRequest{
		Activity: activity,
	})
	if err != nil {
		return nil, "", nil, err
	}
	response := resp.(GetUpdateAccountRequestResponse)
	return response.Activity, response.AccountURL, response.Contacts, nil
}

type GetUpdateAccountRequestRequest struct {
	Activity *api.Activity
}

type GetUpdateAccountRequestResponse struct {
	Activity   *api.Activity
	AccountURL string
	Contacts   []string
}

func makeGetUpdateAccountRequestEndpoint(s acme.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUpdateAccountRequestRequest)
		activity, accountURL, contacts, err := s.GetUpdateAccountRequest(ctx, req.Activity)
		if err != nil {
			return nil, err
		}
		return GetUpdateAccountRequestResponse{
			Activity:   activity,
			AccountURL: accountURL,
			Contacts:   contacts,
		}, nil
	}
}

func (e Endpoints) SetUpdateAccountResponse(ctx context.Context, activity *api.Activity, account *api.Account, erro *api.Error) error {
	_, err := e.SetUpdateAccountResponseEndpoint(ctx, SetUpdateAccountResponseRequest{
		Activity: activity,
		Account:  account,
		Error:    erro,
	})
	return err
}

type SetUpdateAccountResponseRequest struct {
	Activity *api.Activity
	Account  *api.Account
	Error    *api.Error
}

type SetUpdateAccountResponseResponse struct{}

func makeSetUpdateAccountResponseEndpoint(s acme.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SetUpdateAccountResponseRequest)
		err := s.SetUpdateAccountResponse(ctx, req.Activity, req.Account, req.Error)
		if err != nil {
			return nil, err
		}
		return SetUpdateAccountResponseResponse{}, nil
	}
}

func (e Endpoints) GetCreateOrderRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, []*api.Identifier, string, string, error) {
	resp, err := e.GetCreateOrderRequestEndpoint(ctx, GetCreateOrderRequestRequest{
		Activity: activity,
	})
	if err != nil {
		return nil, "", "", nil, "", "", err
	}
	response := resp.(GetCreateOrderRequestResponse)
	return response.Activity, response.DirectoryURL, response.AccountURL, response.Identifiers, response.NotBefore, response.NotAfter, nil
}

type GetCreateOrderRequestRequest struct {
	Activity *api.Activity
}

type GetCreateOrderRequestResponse struct {
	Activity     *api.Activity
	DirectoryURL string
	AccountURL   string
	Identifiers  []*api.Identifier
	NotBefore    string
	NotAfter     string
}

func makeGetCreateOrderRequestEndpoint(s acme.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetCreateOrderRequestRequest)
		activity, directoryURL, accountURL, identifiers, notBefore, notAfter, err := s.GetCreateOrderRequest(ctx, req.Activity)
		if err != nil {
			return nil, err
		}
		return GetCreateOrderRequestResponse{
			Activity:     activity,
			DirectoryURL: directoryURL,
			AccountURL:   accountURL,
			Identifiers:  identifiers,
			NotBefore:    notBefore,
			NotAfter:     notAfter,
		}, nil
	}
}

func (e Endpoints) SetCreateOrderResponse(ctx context.Context, activity *api.Activity, order *api.Order, erro *api.Error) error {
	_, err := e.SetCreateOrderResponseEndpoint(ctx, SetCreateOrderResponseRequest{
		Activity: activity,
		Order:    order,
		Error:    erro,
	})
	return err
}

type SetCreateOrderResponseRequest struct {
	Activity *api.Activity
	Order    *api.Order
	Error    *api.Error
}

type SetCreateOrderResponseResponse struct{}

func makeSetCreateOrderResponseEndpoint(s acme.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SetCreateOrderResponseRequest)
		err := s.SetCreateOrderResponse(ctx, req.Activity, req.Order, req.Error)
		if err != nil {
			return nil, err
		}
		return SetCreateOrderResponseResponse{}, nil
	}
}

func (e Endpoints) GetFinalizeOrderRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, *x509.CertificateRequest, error) {
	resp, err := e.GetFinalizeOrderRequestEndpoint(ctx, GetFinalizeOrderRequestRequest{
		Activity: activity,
	})
	if err != nil {
		return nil, "", "", nil, err
	}
	response := resp.(GetFinalizeOrderRequestResponse)
	return response.Activity, response.AccountURL, response.OrderURL, response.CertificateSigningRequest, nil
}

type GetFinalizeOrderRequestRequest struct {
	Activity *api.Activity
}

type GetFinalizeOrderRequestResponse struct {
	Activity                  *api.Activity
	AccountURL                string
	OrderURL                  string
	CertificateSigningRequest *x509.CertificateRequest
}

func makeGetFinalizeOrderRequestEndpoint(s acme.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetFinalizeOrderRequestRequest)
		activity, accountURL, orderURL, certificateSigningRequest, err := s.GetFinalizeOrderRequest(ctx, req.Activity)
		if err != nil {
			return nil, err
		}
		return GetFinalizeOrderRequestResponse{
			Activity:                  activity,
			AccountURL:                accountURL,
			OrderURL:                  orderURL,
			CertificateSigningRequest: certificateSigningRequest,
		}, nil
	}
}

func (e Endpoints) SetFinalizeOrderResponse(ctx context.Context, activity *api.Activity, order *api.Order, erro *api.Error) error {
	_, err := e.SetFinalizeOrderResponseEndpoint(ctx, SetFinalizeOrderResponseRequest{
		Activity: activity,
		Order:    order,
		Error:    erro,
	})
	return err
}

type SetFinalizeOrderResponseRequest struct {
	Activity *api.Activity
	Order    *api.Order
	Error    *api.Error
}

type SetFinalizeOrderResponseResponse struct{}

func makeSetFinalizeOrderResponseEndpoint(s acme.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SetFinalizeOrderResponseRequest)
		err := s.SetFinalizeOrderResponse(ctx, req.Activity, req.Order, req.Error)
		if err != nil {
			return nil, err
		}
		return SetFinalizeOrderResponseResponse{}, nil
	}
}

func (e Endpoints) GetGetOrderRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, error) {
	resp, err := e.GetGetOrderRequestEndpoint(ctx, GetGetOrderRequestRequest{
		Activity: activity,
	})
	if err != nil {
		return nil, "", "", err
	}
	response := resp.(GetGetOrderRequestResponse)
	return response.Activity, response.AccountURL, response.OrderURL, nil
}

type GetGetOrderRequestRequest struct {
	Activity *api.Activity
}

type GetGetOrderRequestResponse struct {
	Activity   *api.Activity
	AccountURL string
	OrderURL   string
}

func makeGetGetOrderRequestEndpoint(s acme.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetGetOrderRequestRequest)
		activity, accountURL, orderURL, err := s.GetGetOrderRequest(ctx, req.Activity)
		if err != nil {
			return nil, err
		}
		return GetGetOrderRequestResponse{
			Activity:   activity,
			AccountURL: accountURL,
			OrderURL:   orderURL,
		}, nil
	}
}

func (e Endpoints) SetGetOrderResponse(ctx context.Context, activity *api.Activity, order *api.Order, erro *api.Error) error {
	_, err := e.SetGetOrderResponseEndpoint(ctx, SetGetOrderResponseRequest{
		Activity: activity,
		Order:    order,
		Error:    erro,
	})
	return err
}

type SetGetOrderResponseRequest struct {
	Activity *api.Activity
	Order    *api.Order
	Error    *api.Error
}

type SetGetOrderResponseResponse struct{}

func makeSetGetOrderResponseEndpoint(s acme.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SetGetOrderResponseRequest)
		err := s.SetGetOrderResponse(ctx, req.Activity, req.Order, req.Error)
		if err != nil {
			return nil, err
		}
		return SetGetOrderResponseResponse{}, nil
	}
}

func (e Endpoints) GetCreateAuthorizationRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, *api.Identifier, error) {
	resp, err := e.GetCreateAuthorizationRequestEndpoint(ctx, GetCreateAuthorizationRequestRequest{
		Activity: activity,
	})
	if err != nil {
		return nil, "", "", nil, err
	}
	response := resp.(GetCreateAuthorizationRequestResponse)
	return response.Activity, response.DirectoryURL, response.AccountURL, response.Identifier, nil
}

type GetCreateAuthorizationRequestRequest struct {
	Activity *api.Activity
}

type GetCreateAuthorizationRequestResponse struct {
	Activity     *api.Activity
	DirectoryURL string
	AccountURL   string
	Identifier   *api.Identifier
}

func makeGetCreateAuthorizationRequestEndpoint(s acme.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetCreateAuthorizationRequestRequest)
		activity, directoryURL, accountURL, identifier, err := s.GetCreateAuthorizationRequest(ctx, req.Activity)
		if err != nil {
			return nil, err
		}
		return GetCreateAuthorizationRequestResponse{
			Activity:     activity,
			DirectoryURL: directoryURL,
			AccountURL:   accountURL,
			Identifier:   identifier,
		}, nil
	}
}

func (e Endpoints) SetCreateAuthorizationResponse(ctx context.Context, activity *api.Activity, authorization *api.Authorization, erro *api.Error) error {
	_, err := e.SetCreateAuthorizationResponseEndpoint(ctx, SetCreateAuthorizationResponseRequest{
		Activity:      activity,
		Authorization: authorization,
		Error:         erro,
	})
	return err
}

type SetCreateAuthorizationResponseRequest struct {
	Activity      *api.Activity
	Authorization *api.Authorization
	Error         *api.Error
}

type SetCreateAuthorizationResponseResponse struct{}

func makeSetCreateAuthorizationResponseEndpoint(s acme.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SetCreateAuthorizationResponseRequest)
		err := s.SetCreateAuthorizationResponse(ctx, req.Activity, req.Authorization, req.Error)
		if err != nil {
			return nil, err
		}
		return SetCreateAuthorizationResponseResponse{}, nil
	}
}

func (e Endpoints) GetDeactivateAuthorizationRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, error) {
	resp, err := e.GetDeactivateAuthorizationRequestEndpoint(ctx, GetDeactivateAuthorizationRequestRequest{
		Activity: activity,
	})
	if err != nil {
		return nil, "", "", err
	}
	response := resp.(GetDeactivateAuthorizationRequestResponse)
	return response.Activity, response.AccountURL, response.AuthorizationURL, nil
}

type GetDeactivateAuthorizationRequestRequest struct {
	Activity *api.Activity
}

type GetDeactivateAuthorizationRequestResponse struct {
	Activity         *api.Activity
	AccountURL       string
	AuthorizationURL string
}

func makeGetDeactivateAuthorizationRequestEndpoint(s acme.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetDeactivateAuthorizationRequestRequest)
		activity, accountURL, authorizationURL, err := s.GetDeactivateAuthorizationRequest(ctx, req.Activity)
		if err != nil {
			return nil, err
		}
		return GetDeactivateAuthorizationRequestResponse{
			Activity:         activity,
			AccountURL:       accountURL,
			AuthorizationURL: authorizationURL,
		}, nil
	}
}

func (e Endpoints) SetDeactivateAuthorizationResponse(ctx context.Context, activity *api.Activity, authorization *api.Authorization, erro *api.Error) error {
	_, err := e.SetDeactivateAuthorizationResponseEndpoint(ctx, SetDeactivateAuthorizationResponseRequest{
		Activity:      activity,
		Authorization: authorization,
		Error:         erro,
	})
	return err
}

type SetDeactivateAuthorizationResponseRequest struct {
	Activity      *api.Activity
	Authorization *api.Authorization
	Error         *api.Error
}

type SetDeactivateAuthorizationResponseResponse struct{}

func makeSetDeactivateAuthorizationResponseEndpoint(s acme.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SetDeactivateAuthorizationResponseRequest)
		err := s.SetDeactivateAuthorizationResponse(ctx, req.Activity, req.Authorization, req.Error)
		if err != nil {
			return nil, err
		}
		return SetDeactivateAuthorizationResponseResponse{}, nil
	}
}

func (e Endpoints) GetGetAuthorizationRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, error) {
	resp, err := e.GetGetAuthorizationRequestEndpoint(ctx, GetGetAuthorizationRequestRequest{
		Activity: activity,
	})
	if err != nil {
		return nil, "", "", err
	}
	response := resp.(GetGetAuthorizationRequestResponse)
	return response.Activity, response.AccountURL, response.AuthorizationURL, nil
}

type GetGetAuthorizationRequestRequest struct {
	Activity *api.Activity
}

type GetGetAuthorizationRequestResponse struct {
	Activity         *api.Activity
	AccountURL       string
	AuthorizationURL string
}

func makeGetGetAuthorizationRequestEndpoint(s acme.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetGetAuthorizationRequestRequest)
		activity, accountURL, authorizationURL, err := s.GetGetAuthorizationRequest(ctx, req.Activity)
		if err != nil {
			return nil, err
		}
		return GetGetAuthorizationRequestResponse{
			Activity:         activity,
			AccountURL:       accountURL,
			AuthorizationURL: authorizationURL,
		}, nil
	}
}

func (e Endpoints) SetGetAuthorizationResponse(ctx context.Context, activity *api.Activity, authorization *api.Authorization, erro *api.Error) error {
	_, err := e.SetGetAuthorizationResponseEndpoint(ctx, SetGetAuthorizationResponseRequest{
		Activity:      activity,
		Authorization: authorization,
		Error:         erro,
	})
	return err
}

type SetGetAuthorizationResponseRequest struct {
	Activity      *api.Activity
	Authorization *api.Authorization
	Error         *api.Error
}

type SetGetAuthorizationResponseResponse struct{}

func makeSetGetAuthorizationResponseEndpoint(s acme.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SetGetAuthorizationResponseRequest)
		err := s.SetGetAuthorizationResponse(ctx, req.Activity, req.Authorization, req.Error)
		if err != nil {
			return nil, err
		}
		return SetGetAuthorizationResponseResponse{}, nil
	}
}

func (e Endpoints) GetGetChallengeRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, error) {
	resp, err := e.GetGetChallengeRequestEndpoint(ctx, GetGetChallengeRequestRequest{
		Activity: activity,
	})
	if err != nil {
		return nil, "", "", err
	}
	response := resp.(GetGetChallengeRequestResponse)
	return response.Activity, response.AccountURL, response.ChallengeURL, nil
}

type GetGetChallengeRequestRequest struct {
	Activity *api.Activity
}

type GetGetChallengeRequestResponse struct {
	Activity     *api.Activity
	AccountURL   string
	ChallengeURL string
}

func makeGetGetChallengeRequestEndpoint(s acme.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetGetChallengeRequestRequest)
		activity, accountURL, challengeURL, err := s.GetGetChallengeRequest(ctx, req.Activity)
		if err != nil {
			return nil, err
		}
		return GetGetChallengeRequestResponse{
			Activity:     activity,
			AccountURL:   accountURL,
			ChallengeURL: challengeURL,
		}, nil
	}
}

func (e Endpoints) SetGetChallengeResponse(ctx context.Context, activity *api.Activity, challenge *api.Challenge, erro *api.Error) error {
	_, err := e.SetGetChallengeResponseEndpoint(ctx, SetGetChallengeResponseRequest{
		Activity:  activity,
		Challenge: challenge,
		Error:     erro,
	})
	return err
}

type SetGetChallengeResponseRequest struct {
	Activity  *api.Activity
	Challenge *api.Challenge
	Error     *api.Error
}

type SetGetChallengeResponseResponse struct{}

func makeSetGetChallengeResponseEndpoint(s acme.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SetGetChallengeResponseRequest)
		err := s.SetGetChallengeResponse(ctx, req.Activity, req.Challenge, req.Error)
		if err != nil {
			return nil, err
		}
		return SetGetChallengeResponseResponse{}, nil
	}
}

func (e Endpoints) GetValidateChallengeRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, error) {
	resp, err := e.GetValidateChallengeRequestEndpoint(ctx, GetValidateChallengeRequestRequest{
		Activity: activity,
	})
	if err != nil {
		return nil, "", "", err
	}
	response := resp.(GetValidateChallengeRequestResponse)
	return response.Activity, response.AccountURL, response.ChallengeURL, nil
}

type GetValidateChallengeRequestRequest struct {
	Activity *api.Activity
}

type GetValidateChallengeRequestResponse struct {
	Activity     *api.Activity
	AccountURL   string
	ChallengeURL string
}

func makeGetValidateChallengeRequestEndpoint(s acme.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetValidateChallengeRequestRequest)
		activity, accountURL, challengeURL, err := s.GetValidateChallengeRequest(ctx, req.Activity)
		if err != nil {
			return nil, err
		}
		return GetValidateChallengeRequestResponse{
			Activity:     activity,
			AccountURL:   accountURL,
			ChallengeURL: challengeURL,
		}, nil
	}
}

func (e Endpoints) SetValidateChallengeResponse(ctx context.Context, activity *api.Activity, challenge *api.Challenge, erro *api.Error) error {
	_, err := e.SetValidateChallengeResponseEndpoint(ctx, SetValidateChallengeResponseRequest{
		Activity:  activity,
		Challenge: challenge,
		Error:     erro,
	})
	return err
}

type SetValidateChallengeResponseRequest struct {
	Activity  *api.Activity
	Challenge *api.Challenge
	Error     *api.Error
}

type SetValidateChallengeResponseResponse struct{}

func makeSetValidateChallengeResponseEndpoint(s acme.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SetValidateChallengeResponseRequest)
		err := s.SetValidateChallengeResponse(ctx, req.Activity, req.Challenge, req.Error)
		if err != nil {
			return nil, err
		}
		return SetValidateChallengeResponseResponse{}, nil
	}
}

func (e Endpoints) GetGetCertificateRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, error) {
	resp, err := e.GetGetCertificateRequestEndpoint(ctx, GetGetCertificateRequestRequest{
		Activity: activity,
	})
	if err != nil {
		return nil, "", "", err
	}
	response := resp.(GetGetCertificateRequestResponse)
	return response.Activity, response.AccountURL, response.CertificateURL, nil
}

type GetGetCertificateRequestRequest struct {
	Activity *api.Activity
}

type GetGetCertificateRequestResponse struct {
	Activity       *api.Activity
	AccountURL     string
	CertificateURL string
}

func makeGetGetCertificateRequestEndpoint(s acme.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetGetCertificateRequestRequest)
		activity, accountURL, certificateURL, err := s.GetGetCertificateRequest(ctx, req.Activity)
		if err != nil {
			return nil, err
		}
		return GetGetCertificateRequestResponse{
			Activity:       activity,
			AccountURL:     accountURL,
			CertificateURL: certificateURL,
		}, nil
	}
}

func (e Endpoints) SetGetCertificateResponse(ctx context.Context, activity *api.Activity, certificates []*x509.Certificate, erro *api.Error) error {
	_, err := e.SetGetCertificateResponseEndpoint(ctx, SetGetCertificateResponseRequest{
		Activity:     activity,
		Certificates: certificates,
		Error:        erro,
	})
	return err
}

type SetGetCertificateResponseRequest struct {
	Activity     *api.Activity
	Certificates []*x509.Certificate
	Error        *api.Error
}

type SetGetCertificateResponseResponse struct{}

func makeSetGetCertificateResponseEndpoint(s acme.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SetGetCertificateResponseRequest)
		err := s.SetGetCertificateResponse(ctx, req.Activity, req.Certificates, req.Error)
		if err != nil {
			return nil, err
		}
		return SetGetCertificateResponseResponse{}, nil
	}
}

func (e Endpoints) GetRevokeCertificateRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, *x509.Certificate, *api.RevocationReason, error) {
	resp, err := e.GetRevokeCertificateRequestEndpoint(ctx, GetRevokeCertificateRequestRequest{
		Activity: activity,
	})
	if err != nil {
		return nil, "", "", nil, nil, err
	}
	response := resp.(GetRevokeCertificateRequestResponse)
	return response.Activity, response.DirectoryURL, response.AccountURL, response.Certificate, response.Reason, nil
}

type GetRevokeCertificateRequestRequest struct {
	Activity *api.Activity
}

type GetRevokeCertificateRequestResponse struct {
	Activity     *api.Activity
	DirectoryURL string
	AccountURL   string
	Certificate  *x509.Certificate
	Reason       *api.RevocationReason
}

func makeGetRevokeCertificateRequestEndpoint(s acme.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRevokeCertificateRequestRequest)
		activity, directoryURL, accountURL, certificate, reason, err := s.GetRevokeCertificateRequest(ctx, req.Activity)
		if err != nil {
			return nil, err
		}
		return GetRevokeCertificateRequestResponse{
			Activity:     activity,
			DirectoryURL: directoryURL,
			AccountURL:   accountURL,
			Certificate:  certificate,
			Reason:       reason,
		}, nil
	}
}

func (e Endpoints) SetRevokeCertificateResponse(ctx context.Context, activity *api.Activity, erro *api.Error) error {
	_, err := e.SetRevokeCertificateResponseEndpoint(ctx, SetRevokeCertificateResponseRequest{
		Activity: activity,
		Error:    erro,
	})
	return err
}

type SetRevokeCertificateResponseRequest struct {
	Activity *api.Activity
	Error    *api.Error
}

type SetRevokeCertificateResponseResponse struct{}

func makeSetRevokeCertificateResponseEndpoint(s acme.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SetRevokeCertificateResponseRequest)
		err := s.SetRevokeCertificateResponse(ctx, req.Activity, req.Error)
		if err != nil {
			return nil, err
		}
		return SetRevokeCertificateResponseResponse{}, nil
	}
}
