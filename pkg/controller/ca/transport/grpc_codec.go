package transport // import "powerssl.io/pkg/controller/workflow/transport"

import (
	"context"

	"github.com/gogo/protobuf/types"

	"powerssl.io/pkg/controller/api"
	apiv1 "powerssl.io/pkg/controller/api/v1"
	"powerssl.io/pkg/controller/ca/endpoint"
)

func DecodeGRPCActivity(activity *apiv1.Activity) (*api.Activity, error) {
	return &api.Activity{
		Token: activity.GetToken(),
		// TODO
		Name: api.ActivityName(activity.GetName()),
		// Workflow
		Signature: activity.GetSignature(),
	}, nil
}

func EncodeGRPCActivity(activity *api.Activity) (*apiv1.Activity, error) {
	if activity == nil {
		activity = &api.Activity{}
	}
	return &apiv1.Activity{
		Token: activity.Token,
		// TODO
		Name: apiv1.Activity_Name(activity.Name),
		// Name
		// Workflow
		Signature: activity.Signature,
	}, nil
}

func decodeGRPCError(erro *apiv1.Error) (*api.Error, error) {
	return &api.Error{
		Message: erro.GetMessage(),
	}, nil
}

func encodeGRPCError(erro *api.Error) (*apiv1.Error, error) {
	if erro == nil {
		erro = &api.Error{}
	}
	return &apiv1.Error{
		Message: erro.Message,
	}, nil
}

func decodeGRPCChallenge(challenge *apiv1.Challenge) (*api.Challenge, error) {
	return &api.Challenge{
		Type:    api.ChallengeType(challenge.GetType()),
		Details: challenge.GetDetails(),
	}, nil
}

func encodeGRPCChallenge(challenge *api.Challenge) (*apiv1.Challenge, error) {
	return &apiv1.Challenge{
		Type:    apiv1.ChallengeType(challenge.Type),
		Details: challenge.Details,
	}, nil
}

func decodeGRPCChallenges(grpcChallenges []*apiv1.Challenge) ([]*api.Challenge, error) {
	certificates := make([]*api.Challenge, len(grpcChallenges))
	for i, grpcChallenge := range grpcChallenges {
		certificate, err := decodeGRPCChallenge(grpcChallenge)
		if err != nil {
			return nil, err
		}
		certificates[i] = certificate
	}
	return certificates, nil
}

func encodeGRPCChallenges(certificates []*api.Challenge) ([]*apiv1.Challenge, error) {
	grpcChallenges := make([]*apiv1.Challenge, len(certificates))
	for i, certificate := range certificates {
		grpcChallenge, err := encodeGRPCChallenge(certificate)
		if err != nil {
			return nil, err
		}
		grpcChallenges[i] = grpcChallenge
	}
	return grpcChallenges, nil
}

////////////////////////

func decodeGRPCGetAuthorizeDomainRequestRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.Activity)
	activity, err := DecodeGRPCActivity(req)
	if err != nil {
		return nil, err
	}
	return endpoint.GetAuthorizeDomainRequestRequest{
		Activity: activity,
	}, nil
}

func decodeGRPCGetAuthorizeDomainRequestResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*apiv1.GetAuthorizeDomainRequestResponse)
	activity, err := DecodeGRPCActivity(reply.GetActivity())
	if err != nil {
		return nil, err
	}
	return endpoint.GetAuthorizeDomainRequestResponse{
		Activity: activity,
		Domain:   reply.GetDomain(),
	}, nil
}

func encodeGRPCGetAuthorizeDomainRequestResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.GetAuthorizeDomainRequestResponse)
	activity, err := EncodeGRPCActivity(resp.Activity)
	if err != nil {
		return nil, err
	}
	return &apiv1.GetAuthorizeDomainRequestResponse{
		Activity: activity,
		Domain:   resp.Domain,
	}, nil
}

func encodeGRPCGetAuthorizeDomainRequestRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.GetAuthorizeDomainRequestRequest)
	activity, err := EncodeGRPCActivity(req.Activity)
	if err != nil {
		return nil, err
	}
	return activity, nil
}

func decodeGRPCSetAuthorizeDomainResponseRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.SetAuthorizeDomainResponseRequest)
	activity, err := DecodeGRPCActivity(req.GetActivity())
	if err != nil {
		return nil, err
	}
	erro, err := decodeGRPCError(req.GetError())
	if err != nil {
		return nil, err
	}
	challenges, err := decodeGRPCChallenges(req.GetChallenges())
	if err != nil {
		return nil, err
	}
	return endpoint.SetAuthorizeDomainResponseRequest{
		Activity:   activity,
		Error:      erro,
		Challenges: challenges,
	}, nil
}

func decodeGRPCSetAuthorizeDomainResponseResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	return endpoint.SetAuthorizeDomainResponseResponse{}, nil
}

func encodeGRPCSetAuthorizeDomainResponseResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &types.Empty{}, nil
}

func encodeGRPCSetAuthorizeDomainResponseRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.SetAuthorizeDomainResponseRequest)
	activity, err := EncodeGRPCActivity(req.Activity)
	if err != nil {
		return nil, err
	}
	erro, err := encodeGRPCError(req.Error)
	if err != nil {
		return nil, err
	}
	challenges, err := encodeGRPCChallenges(req.Challenges)
	if err != nil {
		return nil, err
	}
	return &apiv1.SetAuthorizeDomainResponseRequest{
		Activity:   activity,
		Error:      erro,
		Challenges: challenges,
	}, nil
}

////////////////////////

func decodeGRPCGetRequestCertificateRequestRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.Activity)
	activity, err := DecodeGRPCActivity(req)
	if err != nil {
		return nil, err
	}
	return endpoint.GetRequestCertificateRequestRequest{
		Activity: activity,
	}, nil
}

func decodeGRPCGetRequestCertificateRequestResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*apiv1.GetRequestCertificateRequestResponse)
	activity, err := DecodeGRPCActivity(reply.GetActivity())
	if err != nil {
		return nil, err
	}
	return endpoint.GetRequestCertificateRequestResponse{
		Activity:                  activity,
		CertificateSigningRequest: reply.GetCertificateSigningRequest(),
	}, nil
}

func encodeGRPCGetRequestCertificateRequestResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.GetRequestCertificateRequestResponse)
	activity, err := EncodeGRPCActivity(resp.Activity)
	if err != nil {
		return nil, err
	}
	return &apiv1.GetRequestCertificateRequestResponse{
		Activity:                  activity,
		CertificateSigningRequest: resp.CertificateSigningRequest,
	}, nil
}

func encodeGRPCGetRequestCertificateRequestRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.GetRequestCertificateRequestRequest)
	activity, err := EncodeGRPCActivity(req.Activity)
	if err != nil {
		return nil, err
	}
	return activity, nil
}

func decodeGRPCSetRequestCertificateResponseRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.SetRequestCertificateResponseRequest)
	activity, err := DecodeGRPCActivity(req.GetActivity())
	if err != nil {
		return nil, err
	}
	erro, err := decodeGRPCError(req.GetError())
	if err != nil {
		return nil, err
	}
	return endpoint.SetRequestCertificateResponseRequest{
		Activity:    activity,
		Error:       erro,
		Certificate: req.GetCertificate(),
	}, nil
}

func decodeGRPCSetRequestCertificateResponseResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	return endpoint.SetRequestCertificateResponseResponse{}, nil
}

func encodeGRPCSetRequestCertificateResponseResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &types.Empty{}, nil
}

func encodeGRPCSetRequestCertificateResponseRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.SetRequestCertificateResponseRequest)
	activity, err := EncodeGRPCActivity(req.Activity)
	if err != nil {
		return nil, err
	}
	erro, err := encodeGRPCError(req.Error)
	if err != nil {
		return nil, err
	}
	return &apiv1.SetRequestCertificateResponseRequest{
		Activity:    activity,
		Error:       erro,
		Certificate: req.Certificate,
	}, nil
}

////////////////////////

func decodeGRPCGetRevokeCertificateRequestRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.Activity)
	activity, err := DecodeGRPCActivity(req)
	if err != nil {
		return nil, err
	}
	return endpoint.GetRevokeCertificateRequestRequest{
		Activity: activity,
	}, nil
}

func decodeGRPCGetRevokeCertificateRequestResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*apiv1.GetRevokeCertificateRequestResponse)
	activity, err := DecodeGRPCActivity(reply.GetActivity())
	if err != nil {
		return nil, err
	}
	return endpoint.GetRevokeCertificateRequestResponse{
		Activity:    activity,
		Certificate: reply.GetCertificate(),
	}, nil
}

func encodeGRPCGetRevokeCertificateRequestResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.GetRevokeCertificateRequestResponse)
	activity, err := EncodeGRPCActivity(resp.Activity)
	if err != nil {
		return nil, err
	}
	return &apiv1.GetRevokeCertificateRequestResponse{
		Activity:    activity,
		Certificate: resp.Certificate,
	}, nil
}

func encodeGRPCGetRevokeCertificateRequestRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.GetRevokeCertificateRequestRequest)
	activity, err := EncodeGRPCActivity(req.Activity)
	if err != nil {
		return nil, err
	}
	return activity, nil
}

func decodeGRPCSetRevokeCertificateResponseRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.SetRevokeCertificateResponseRequest)
	activity, err := DecodeGRPCActivity(req.GetActivity())
	if err != nil {
		return nil, err
	}
	erro, err := decodeGRPCError(req.GetError())
	if err != nil {
		return nil, err
	}
	return endpoint.SetRevokeCertificateResponseRequest{
		Activity: activity,
		Error:    erro,
	}, nil
}

func decodeGRPCSetRevokeCertificateResponseResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	return endpoint.SetRevokeCertificateResponseResponse{}, nil
}

func encodeGRPCSetRevokeCertificateResponseResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &types.Empty{}, nil
}

func encodeGRPCSetRevokeCertificateResponseRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.SetRevokeCertificateResponseRequest)
	activity, err := EncodeGRPCActivity(req.Activity)
	if err != nil {
		return nil, err
	}
	erro, err := encodeGRPCError(req.Error)
	if err != nil {
		return nil, err
	}
	return &apiv1.SetRevokeCertificateResponseRequest{
		Activity: activity,
		Error:    erro,
	}, nil
}

////////////////////////

func decodeGRPCGetVerifyDomainRequestRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.Activity)
	activity, err := DecodeGRPCActivity(req)
	if err != nil {
		return nil, err
	}
	return endpoint.GetVerifyDomainRequestRequest{
		Activity: activity,
	}, nil
}

func decodeGRPCGetVerifyDomainRequestResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*apiv1.GetVerifyDomainRequestResponse)
	activity, err := DecodeGRPCActivity(reply.GetActivity())
	if err != nil {
		return nil, err
	}
	return endpoint.GetVerifyDomainRequestResponse{
		Activity:      activity,
		Domain:        reply.GetDomain(),
		ChallengeType: api.ChallengeType(reply.GetChallengeType()),
	}, nil
}

func encodeGRPCGetVerifyDomainRequestResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.GetVerifyDomainRequestResponse)
	activity, err := EncodeGRPCActivity(resp.Activity)
	if err != nil {
		return nil, err
	}
	return &apiv1.GetVerifyDomainRequestResponse{
		Activity:      activity,
		Domain:        resp.Domain,
		ChallengeType: apiv1.ChallengeType(resp.ChallengeType),
	}, nil
}

func encodeGRPCGetVerifyDomainRequestRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.GetVerifyDomainRequestRequest)
	activity, err := EncodeGRPCActivity(req.Activity)
	if err != nil {
		return nil, err
	}
	return activity, nil
}

func decodeGRPCSetVerifyDomainResponseRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.SetVerifyDomainResponseRequest)
	activity, err := DecodeGRPCActivity(req.GetActivity())
	if err != nil {
		return nil, err
	}
	erro, err := decodeGRPCError(req.GetError())
	if err != nil {
		return nil, err
	}
	return endpoint.SetVerifyDomainResponseRequest{
		Activity: activity,
		Error:    erro,
	}, nil
}

func decodeGRPCSetVerifyDomainResponseResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	return endpoint.SetVerifyDomainResponseResponse{}, nil
}

func encodeGRPCSetVerifyDomainResponseResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &types.Empty{}, nil
}

func encodeGRPCSetVerifyDomainResponseRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.SetVerifyDomainResponseRequest)
	activity, err := EncodeGRPCActivity(req.Activity)
	if err != nil {
		return nil, err
	}
	erro, err := encodeGRPCError(req.Error)
	if err != nil {
		return nil, err
	}
	return &apiv1.SetVerifyDomainResponseRequest{
		Activity: activity,
		Error:    erro,
	}, nil
}
