package endpoint // import "powerssl.dev/sdk/controller/acme/endpoint"

import (
	"context"
	"crypto/x509"
	"powerssl.dev/sdk/controller/acme"

	"github.com/go-kit/kit/endpoint"

	"powerssl.dev/sdk/controller/api"
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

func MakeGetCreateAccountRequestEndpoint(s acme.Service) endpoint.Endpoint {
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

func MakeSetCreateAccountResponseEndpoint(s acme.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SetCreateAccountResponseRequest)
		err := s.SetCreateAccountResponse(ctx, req.Activity, req.Account, req.Error)
		if err != nil {
			return nil, err
		}
		return SetCreateAccountResponseResponse{}, nil
	}
}

func MakeGetDeactivateAccountRequestEndpoint(s acme.Service) endpoint.Endpoint {
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

func MakeSetDeactivateAccountResponseEndpoint(s acme.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SetDeactivateAccountResponseRequest)
		err := s.SetDeactivateAccountResponse(ctx, req.Activity, req.Account, req.Error)
		if err != nil {
			return nil, err
		}
		return SetDeactivateAccountResponseResponse{}, nil
	}
}

func MakeGetRekeyAccountRequestEndpoint(s acme.Service) endpoint.Endpoint {
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

func MakeSetRekeyAccountResponseEndpoint(s acme.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SetRekeyAccountResponseRequest)
		err := s.SetRekeyAccountResponse(ctx, req.Activity, req.Account, req.Error)
		if err != nil {
			return nil, err
		}
		return SetRekeyAccountResponseResponse{}, nil
	}
}

func MakeGetUpdateAccountRequestEndpoint(s acme.Service) endpoint.Endpoint {
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

func MakeSetUpdateAccountResponseEndpoint(s acme.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SetUpdateAccountResponseRequest)
		err := s.SetUpdateAccountResponse(ctx, req.Activity, req.Account, req.Error)
		if err != nil {
			return nil, err
		}
		return SetUpdateAccountResponseResponse{}, nil
	}
}

func MakeGetCreateOrderRequestEndpoint(s acme.Service) endpoint.Endpoint {
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

func MakeSetCreateOrderResponseEndpoint(s acme.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SetCreateOrderResponseRequest)
		err := s.SetCreateOrderResponse(ctx, req.Activity, req.Order, req.Error)
		if err != nil {
			return nil, err
		}
		return SetCreateOrderResponseResponse{}, nil
	}
}

func MakeGetFinalizeOrderRequestEndpoint(s acme.Service) endpoint.Endpoint {
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

func MakeSetFinalizeOrderResponseEndpoint(s acme.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SetFinalizeOrderResponseRequest)
		err := s.SetFinalizeOrderResponse(ctx, req.Activity, req.Order, req.Error)
		if err != nil {
			return nil, err
		}
		return SetFinalizeOrderResponseResponse{}, nil
	}
}

func MakeGetGetOrderRequestEndpoint(s acme.Service) endpoint.Endpoint {
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

func MakeSetGetOrderResponseEndpoint(s acme.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SetGetOrderResponseRequest)
		err := s.SetGetOrderResponse(ctx, req.Activity, req.Order, req.Error)
		if err != nil {
			return nil, err
		}
		return SetGetOrderResponseResponse{}, nil
	}
}

func MakeGetCreateAuthorizationRequestEndpoint(s acme.Service) endpoint.Endpoint {
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

func MakeSetCreateAuthorizationResponseEndpoint(s acme.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SetCreateAuthorizationResponseRequest)
		err := s.SetCreateAuthorizationResponse(ctx, req.Activity, req.Authorization, req.Error)
		if err != nil {
			return nil, err
		}
		return SetCreateAuthorizationResponseResponse{}, nil
	}
}

func MakeGetDeactivateAuthorizationRequestEndpoint(s acme.Service) endpoint.Endpoint {
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

func MakeSetDeactivateAuthorizationResponseEndpoint(s acme.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SetDeactivateAuthorizationResponseRequest)
		err := s.SetDeactivateAuthorizationResponse(ctx, req.Activity, req.Authorization, req.Error)
		if err != nil {
			return nil, err
		}
		return SetDeactivateAuthorizationResponseResponse{}, nil
	}
}

func MakeGetGetAuthorizationRequestEndpoint(s acme.Service) endpoint.Endpoint {
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

func MakeSetGetAuthorizationResponseEndpoint(s acme.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SetGetAuthorizationResponseRequest)
		err := s.SetGetAuthorizationResponse(ctx, req.Activity, req.Authorization, req.Error)
		if err != nil {
			return nil, err
		}
		return SetGetAuthorizationResponseResponse{}, nil
	}
}

func MakeGetGetChallengeRequestEndpoint(s acme.Service) endpoint.Endpoint {
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

func MakeSetGetChallengeResponseEndpoint(s acme.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SetGetChallengeResponseRequest)
		err := s.SetGetChallengeResponse(ctx, req.Activity, req.Challenge, req.Error)
		if err != nil {
			return nil, err
		}
		return SetGetChallengeResponseResponse{}, nil
	}
}

func MakeGetValidateChallengeRequestEndpoint(s acme.Service) endpoint.Endpoint {
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

func MakeSetValidateChallengeResponseEndpoint(s acme.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SetValidateChallengeResponseRequest)
		err := s.SetValidateChallengeResponse(ctx, req.Activity, req.Challenge, req.Error)
		if err != nil {
			return nil, err
		}
		return SetValidateChallengeResponseResponse{}, nil
	}
}

func MakeGetGetCertificateRequestEndpoint(s acme.Service) endpoint.Endpoint {
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

func MakeSetGetCertificateResponseEndpoint(s acme.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SetGetCertificateResponseRequest)
		err := s.SetGetCertificateResponse(ctx, req.Activity, req.Certificates, req.Error)
		if err != nil {
			return nil, err
		}
		return SetGetCertificateResponseResponse{}, nil
	}
}

func MakeGetRevokeCertificateRequestEndpoint(s acme.Service) endpoint.Endpoint {
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

func MakeSetRevokeCertificateResponseEndpoint(s acme.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SetRevokeCertificateResponseRequest)
		err := s.SetRevokeCertificateResponse(ctx, req.Activity, req.Error)
		if err != nil {
			return nil, err
		}
		return SetRevokeCertificateResponseResponse{}, nil
	}
}
