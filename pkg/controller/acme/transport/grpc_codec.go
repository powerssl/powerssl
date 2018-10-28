package transport // import "powerssl.io/pkg/controller/acme/transport"

import (
	"context"

	"github.com/gogo/protobuf/types"

	"powerssl.io/pkg/controller/acme/endpoint"
	"powerssl.io/pkg/controller/api"
	apiv1 "powerssl.io/pkg/controller/api/v1"
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

func decodeGRPCAccount(account *apiv1.Account) (*api.Account, error) {
	return &api.Account{
		Status:               api.AccountStatus(account.GetStatus()),
		Contacts:             account.GetContacts(),
		TermsOfServiceAgreed: account.GetTermsOfServiceAgreed(),
		URL:                  account.GetUrl(),
	}, nil
}

func encodeGRPCAccount(account *api.Account) (*apiv1.Account, error) {
	return &apiv1.Account{
		Status:               apiv1.Account_Status(account.Status),
		Contacts:             account.Contacts,
		TermsOfServiceAgreed: account.TermsOfServiceAgreed,
		Url:                  account.URL,
	}, nil
}

// func decodeGRPCChallenge(challenge *apiv1.Challenge) (*api.Challenge, error) {
// 	return &api.Challenge{
// 		Type:    api.ChallengeType(challenge.GetType()),
// 		Details: challenge.GetDetails(),
// 	}, nil
// }

// func encodeGRPCChallenge(challenge *api.Challenge) (*apiv1.Challenge, error) {
// 	return &apiv1.Challenge{
// 		Type:    apiv1.ChallengeType(challenge.Type),
// 		Details: challenge.Details,
// 	}, nil
// }

// func decodeGRPCChallenges(grpcChallenges []*apiv1.Challenge) ([]*api.Challenge, error) {
// 	certificates := make([]*api.Challenge, len(grpcChallenges))
// 	for i, grpcChallenge := range grpcChallenges {
// 		certificate, err := decodeGRPCChallenge(grpcChallenge)
// 		if err != nil {
// 			return nil, err
// 		}
// 		certificates[i] = certificate
// 	}
// 	return certificates, nil
// }

// func encodeGRPCChallenges(certificates []*api.Challenge) ([]*apiv1.Challenge, error) {
// 	grpcChallenges := make([]*apiv1.Challenge, len(certificates))
// 	for i, certificate := range certificates {
// 		grpcChallenge, err := encodeGRPCChallenge(certificate)
// 		if err != nil {
// 			return nil, err
// 		}
// 		grpcChallenges[i] = grpcChallenge
// 	}
// 	return grpcChallenges, nil
// }

////////////////////////
func decodeGRPCGetCreateAccountRequestRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.Activity)
	activity, err := DecodeGRPCActivity(req)
	if err != nil {
		return nil, err
	}
	return endpoint.GetCreateAccountRequestRequest{
		Activity: activity,
	}, nil
}

func decodeGRPCGetCreateAccountRequestResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*apiv1.GetCreateAccountRequestResponse)
	activity, err := DecodeGRPCActivity(reply.GetActivity())
	if err != nil {
		return nil, err
	}
	return endpoint.GetCreateAccountRequestResponse{
		Activity:             activity,
		DirectoryURL:         reply.GetDirectoryUrl(),
		TermsOfServiceAgreed: reply.GetTermsOfServiceAgreed(),
		Contacts:             reply.GetContacts(),
	}, nil
}

func encodeGRPCGetCreateAccountRequestResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.GetCreateAccountRequestResponse)
	activity, err := EncodeGRPCActivity(resp.Activity)
	if err != nil {
		return nil, err
	}
	return &apiv1.GetCreateAccountRequestResponse{
		Activity:             activity,
		DirectoryUrl:         resp.DirectoryURL,
		TermsOfServiceAgreed: resp.TermsOfServiceAgreed,
		Contacts:             resp.Contacts,
	}, nil
}

func encodeGRPCGetCreateAccountRequestRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.GetCreateAccountRequestRequest)
	activity, err := EncodeGRPCActivity(req.Activity)
	if err != nil {
		return nil, err
	}
	return activity, nil
}

func decodeGRPCSetCreateAccountResponseRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.SetCreateAccountResponseRequest)
	activity, err := DecodeGRPCActivity(req.GetActivity())
	if err != nil {
		return nil, err
	}
	account, err := decodeGRPCAccount(req.GetAccount())
	if err != nil {
		return nil, err
	}
	erro, err := decodeGRPCError(req.GetError())
	if err != nil {
		return nil, err
	}
	return endpoint.SetCreateAccountResponseRequest{
		Activity: activity,
		Account:  account,
		Error:    erro,
	}, nil
}

func decodeGRPCSetCreateAccountResponseResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	return endpoint.SetCreateAccountResponseResponse{}, nil
}

func encodeGRPCSetCreateAccountResponseResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &types.Empty{}, nil
}

func encodeGRPCSetCreateAccountResponseRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.SetCreateAccountResponseRequest)
	activity, err := EncodeGRPCActivity(req.Activity)
	if err != nil {
		return nil, err
	}
	account, err := encodeGRPCAccount(req.Account)
	if err != nil {
		return nil, err
	}
	erro, err := encodeGRPCError(req.Error)
	if err != nil {
		return nil, err
	}
	return &apiv1.SetCreateAccountResponseRequest{
		Activity: activity,
		Account:  account,
		Error:    erro,
	}, nil
}

func decodeGRPCGetDeactivateAccountRequestRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.Activity)
	activity, err := DecodeGRPCActivity(req)
	if err != nil {
		return nil, err
	}
	return endpoint.GetDeactivateAccountRequestRequest{
		Activity: activity,
	}, nil
}

func decodeGRPCGetDeactivateAccountRequestResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*apiv1.GetDeactivateAccountRequestResponse)
	activity, err := DecodeGRPCActivity(reply.GetActivity())
	if err != nil {
		return nil, err
	}
	return endpoint.GetDeactivateAccountRequestResponse{
		Activity: activity,
		// Domain:   reply.GetDomain(), // TODO
	}, nil
}

func encodeGRPCGetDeactivateAccountRequestResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.GetDeactivateAccountRequestResponse)
	activity, err := EncodeGRPCActivity(resp.Activity)
	if err != nil {
		return nil, err
	}
	return &apiv1.GetDeactivateAccountRequestResponse{
		Activity: activity,
		// Domain:   resp.Domain, // TODO
	}, nil
}

func encodeGRPCGetDeactivateAccountRequestRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.GetDeactivateAccountRequestRequest)
	activity, err := EncodeGRPCActivity(req.Activity)
	if err != nil {
		return nil, err
	}
	return activity, nil
}

func decodeGRPCSetDeactivateAccountResponseRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.SetDeactivateAccountResponseRequest)
	activity, err := DecodeGRPCActivity(req.GetActivity())
	if err != nil {
		return nil, err
	}
	// TODO
	// erro, err := decodeGRPCError(req.GetError())
	// if err != nil {
	//   return nil, err
	// }
	// TODO
	// challenges, err := decodeGRPCChallenges(req.GetChallenges())
	// if err != nil {
	//   return nil, err
	// }
	return endpoint.SetDeactivateAccountResponseRequest{
		Activity: activity,
		// Error:      erro, // TODO
		// Challenges: challenges, // TODO
	}, nil
}

func decodeGRPCSetDeactivateAccountResponseResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	return endpoint.SetDeactivateAccountResponseResponse{}, nil
}

func encodeGRPCSetDeactivateAccountResponseResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &types.Empty{}, nil
}

func encodeGRPCSetDeactivateAccountResponseRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.SetDeactivateAccountResponseRequest)
	activity, err := EncodeGRPCActivity(req.Activity)
	if err != nil {
		return nil, err
	}
	// TODO
	// erro, err := encodeGRPCError(req.Error)
	// if err != nil {
	//   return nil, err
	// }
	// TODO
	// challenges, err := encodeGRPCChallenges(req.Challenges)
	// if err != nil {
	//   return nil, err
	// }
	return &apiv1.SetDeactivateAccountResponseRequest{
		Activity: activity,
		// Error:      erro, // TODO
		// Challenges: challenges, // TODO
	}, nil
}

func decodeGRPCGetRekeyAccountRequestRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.Activity)
	activity, err := DecodeGRPCActivity(req)
	if err != nil {
		return nil, err
	}
	return endpoint.GetRekeyAccountRequestRequest{
		Activity: activity,
	}, nil
}

func decodeGRPCGetRekeyAccountRequestResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*apiv1.GetRekeyAccountRequestResponse)
	activity, err := DecodeGRPCActivity(reply.GetActivity())
	if err != nil {
		return nil, err
	}
	return endpoint.GetRekeyAccountRequestResponse{
		Activity: activity,
		// Domain:   reply.GetDomain(), // TODO
	}, nil
}

func encodeGRPCGetRekeyAccountRequestResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.GetRekeyAccountRequestResponse)
	activity, err := EncodeGRPCActivity(resp.Activity)
	if err != nil {
		return nil, err
	}
	return &apiv1.GetRekeyAccountRequestResponse{
		Activity: activity,
		// Domain:   resp.Domain, // TODO
	}, nil
}

func encodeGRPCGetRekeyAccountRequestRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.GetRekeyAccountRequestRequest)
	activity, err := EncodeGRPCActivity(req.Activity)
	if err != nil {
		return nil, err
	}
	return activity, nil
}

func decodeGRPCSetRekeyAccountResponseRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.SetRekeyAccountResponseRequest)
	activity, err := DecodeGRPCActivity(req.GetActivity())
	if err != nil {
		return nil, err
	}
	// TODO
	// erro, err := decodeGRPCError(req.GetError())
	// if err != nil {
	//   return nil, err
	// }
	// TODO
	// challenges, err := decodeGRPCChallenges(req.GetChallenges())
	// if err != nil {
	//   return nil, err
	// }
	return endpoint.SetRekeyAccountResponseRequest{
		Activity: activity,
		// Error:      erro, // TODO
		// Challenges: challenges, // TODO
	}, nil
}

func decodeGRPCSetRekeyAccountResponseResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	return endpoint.SetRekeyAccountResponseResponse{}, nil
}

func encodeGRPCSetRekeyAccountResponseResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &types.Empty{}, nil
}

func encodeGRPCSetRekeyAccountResponseRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.SetRekeyAccountResponseRequest)
	activity, err := EncodeGRPCActivity(req.Activity)
	if err != nil {
		return nil, err
	}
	// TODO
	// erro, err := encodeGRPCError(req.Error)
	// if err != nil {
	//   return nil, err
	// }
	// TODO
	// challenges, err := encodeGRPCChallenges(req.Challenges)
	// if err != nil {
	//   return nil, err
	// }
	return &apiv1.SetRekeyAccountResponseRequest{
		Activity: activity,
		// Error:      erro, // TODO
		// Challenges: challenges, // TODO
	}, nil
}

func decodeGRPCGetUpdateAccountRequestRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.Activity)
	activity, err := DecodeGRPCActivity(req)
	if err != nil {
		return nil, err
	}
	return endpoint.GetUpdateAccountRequestRequest{
		Activity: activity,
	}, nil
}

func decodeGRPCGetUpdateAccountRequestResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*apiv1.GetUpdateAccountRequestResponse)
	activity, err := DecodeGRPCActivity(reply.GetActivity())
	if err != nil {
		return nil, err
	}
	return endpoint.GetUpdateAccountRequestResponse{
		Activity: activity,
		// Domain:   reply.GetDomain(), // TODO
	}, nil
}

func encodeGRPCGetUpdateAccountRequestResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.GetUpdateAccountRequestResponse)
	activity, err := EncodeGRPCActivity(resp.Activity)
	if err != nil {
		return nil, err
	}
	return &apiv1.GetUpdateAccountRequestResponse{
		Activity: activity,
		// Domain:   resp.Domain, // TODO
	}, nil
}

func encodeGRPCGetUpdateAccountRequestRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.GetUpdateAccountRequestRequest)
	activity, err := EncodeGRPCActivity(req.Activity)
	if err != nil {
		return nil, err
	}
	return activity, nil
}

func decodeGRPCSetUpdateAccountResponseRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.SetUpdateAccountResponseRequest)
	activity, err := DecodeGRPCActivity(req.GetActivity())
	if err != nil {
		return nil, err
	}
	// TODO
	// erro, err := decodeGRPCError(req.GetError())
	// if err != nil {
	//   return nil, err
	// }
	// TODO
	// challenges, err := decodeGRPCChallenges(req.GetChallenges())
	// if err != nil {
	//   return nil, err
	// }
	return endpoint.SetUpdateAccountResponseRequest{
		Activity: activity,
		// Error:      erro, // TODO
		// Challenges: challenges, // TODO
	}, nil
}

func decodeGRPCSetUpdateAccountResponseResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	return endpoint.SetUpdateAccountResponseResponse{}, nil
}

func encodeGRPCSetUpdateAccountResponseResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &types.Empty{}, nil
}

func encodeGRPCSetUpdateAccountResponseRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.SetUpdateAccountResponseRequest)
	activity, err := EncodeGRPCActivity(req.Activity)
	if err != nil {
		return nil, err
	}
	// TODO
	// erro, err := encodeGRPCError(req.Error)
	// if err != nil {
	//   return nil, err
	// }
	// TODO
	// challenges, err := encodeGRPCChallenges(req.Challenges)
	// if err != nil {
	//   return nil, err
	// }
	return &apiv1.SetUpdateAccountResponseRequest{
		Activity: activity,
		// Error:      erro, // TODO
		// Challenges: challenges, // TODO
	}, nil
}

func decodeGRPCGetCreateOrderRequestRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.Activity)
	activity, err := DecodeGRPCActivity(req)
	if err != nil {
		return nil, err
	}
	return endpoint.GetCreateOrderRequestRequest{
		Activity: activity,
	}, nil
}

func decodeGRPCGetCreateOrderRequestResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*apiv1.GetCreateOrderRequestResponse)
	activity, err := DecodeGRPCActivity(reply.GetActivity())
	if err != nil {
		return nil, err
	}
	return endpoint.GetCreateOrderRequestResponse{
		Activity: activity,
		// Domain:   reply.GetDomain(), // TODO
	}, nil
}

func encodeGRPCGetCreateOrderRequestResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.GetCreateOrderRequestResponse)
	activity, err := EncodeGRPCActivity(resp.Activity)
	if err != nil {
		return nil, err
	}
	return &apiv1.GetCreateOrderRequestResponse{
		Activity: activity,
		// Domain:   resp.Domain, // TODO
	}, nil
}

func encodeGRPCGetCreateOrderRequestRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.GetCreateOrderRequestRequest)
	activity, err := EncodeGRPCActivity(req.Activity)
	if err != nil {
		return nil, err
	}
	return activity, nil
}

func decodeGRPCSetCreateOrderResponseRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.SetCreateOrderResponseRequest)
	activity, err := DecodeGRPCActivity(req.GetActivity())
	if err != nil {
		return nil, err
	}
	// TODO
	// erro, err := decodeGRPCError(req.GetError())
	// if err != nil {
	//   return nil, err
	// }
	// TODO
	// challenges, err := decodeGRPCChallenges(req.GetChallenges())
	// if err != nil {
	//   return nil, err
	// }
	return endpoint.SetCreateOrderResponseRequest{
		Activity: activity,
		// Error:      erro, // TODO
		// Challenges: challenges, // TODO
	}, nil
}

func decodeGRPCSetCreateOrderResponseResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	return endpoint.SetCreateOrderResponseResponse{}, nil
}

func encodeGRPCSetCreateOrderResponseResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &types.Empty{}, nil
}

func encodeGRPCSetCreateOrderResponseRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.SetCreateOrderResponseRequest)
	activity, err := EncodeGRPCActivity(req.Activity)
	if err != nil {
		return nil, err
	}
	// TODO
	// erro, err := encodeGRPCError(req.Error)
	// if err != nil {
	//   return nil, err
	// }
	// TODO
	// challenges, err := encodeGRPCChallenges(req.Challenges)
	// if err != nil {
	//   return nil, err
	// }
	return &apiv1.SetCreateOrderResponseRequest{
		Activity: activity,
		// Error:      erro, // TODO
		// Challenges: challenges, // TODO
	}, nil
}

func decodeGRPCGetFinalizeOrderRequestRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.Activity)
	activity, err := DecodeGRPCActivity(req)
	if err != nil {
		return nil, err
	}
	return endpoint.GetFinalizeOrderRequestRequest{
		Activity: activity,
	}, nil
}

func decodeGRPCGetFinalizeOrderRequestResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*apiv1.GetFinalizeOrderRequestResponse)
	activity, err := DecodeGRPCActivity(reply.GetActivity())
	if err != nil {
		return nil, err
	}
	return endpoint.GetFinalizeOrderRequestResponse{
		Activity: activity,
		// Domain:   reply.GetDomain(), // TODO
	}, nil
}

func encodeGRPCGetFinalizeOrderRequestResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.GetFinalizeOrderRequestResponse)
	activity, err := EncodeGRPCActivity(resp.Activity)
	if err != nil {
		return nil, err
	}
	return &apiv1.GetFinalizeOrderRequestResponse{
		Activity: activity,
		// Domain:   resp.Domain, // TODO
	}, nil
}

func encodeGRPCGetFinalizeOrderRequestRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.GetFinalizeOrderRequestRequest)
	activity, err := EncodeGRPCActivity(req.Activity)
	if err != nil {
		return nil, err
	}
	return activity, nil
}

func decodeGRPCSetFinalizeOrderResponseRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.SetFinalizeOrderResponseRequest)
	activity, err := DecodeGRPCActivity(req.GetActivity())
	if err != nil {
		return nil, err
	}
	// TODO
	// erro, err := decodeGRPCError(req.GetError())
	// if err != nil {
	//   return nil, err
	// }
	// TODO
	// challenges, err := decodeGRPCChallenges(req.GetChallenges())
	// if err != nil {
	//   return nil, err
	// }
	return endpoint.SetFinalizeOrderResponseRequest{
		Activity: activity,
		// Error:      erro, // TODO
		// Challenges: challenges, // TODO
	}, nil
}

func decodeGRPCSetFinalizeOrderResponseResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	return endpoint.SetFinalizeOrderResponseResponse{}, nil
}

func encodeGRPCSetFinalizeOrderResponseResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &types.Empty{}, nil
}

func encodeGRPCSetFinalizeOrderResponseRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.SetFinalizeOrderResponseRequest)
	activity, err := EncodeGRPCActivity(req.Activity)
	if err != nil {
		return nil, err
	}
	// TODO
	// erro, err := encodeGRPCError(req.Error)
	// if err != nil {
	//   return nil, err
	// }
	// TODO
	// challenges, err := encodeGRPCChallenges(req.Challenges)
	// if err != nil {
	//   return nil, err
	// }
	return &apiv1.SetFinalizeOrderResponseRequest{
		Activity: activity,
		// Error:      erro, // TODO
		// Challenges: challenges, // TODO
	}, nil
}

func decodeGRPCGetGetOrderRequestRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.Activity)
	activity, err := DecodeGRPCActivity(req)
	if err != nil {
		return nil, err
	}
	return endpoint.GetGetOrderRequestRequest{
		Activity: activity,
	}, nil
}

func decodeGRPCGetGetOrderRequestResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*apiv1.GetGetOrderRequestResponse)
	activity, err := DecodeGRPCActivity(reply.GetActivity())
	if err != nil {
		return nil, err
	}
	return endpoint.GetGetOrderRequestResponse{
		Activity: activity,
		// Domain:   reply.GetDomain(), // TODO
	}, nil
}

func encodeGRPCGetGetOrderRequestResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.GetGetOrderRequestResponse)
	activity, err := EncodeGRPCActivity(resp.Activity)
	if err != nil {
		return nil, err
	}
	return &apiv1.GetGetOrderRequestResponse{
		Activity: activity,
		// Domain:   resp.Domain, // TODO
	}, nil
}

func encodeGRPCGetGetOrderRequestRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.GetGetOrderRequestRequest)
	activity, err := EncodeGRPCActivity(req.Activity)
	if err != nil {
		return nil, err
	}
	return activity, nil
}

func decodeGRPCSetGetOrderResponseRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.SetGetOrderResponseRequest)
	activity, err := DecodeGRPCActivity(req.GetActivity())
	if err != nil {
		return nil, err
	}
	// TODO
	// erro, err := decodeGRPCError(req.GetError())
	// if err != nil {
	//   return nil, err
	// }
	// TODO
	// challenges, err := decodeGRPCChallenges(req.GetChallenges())
	// if err != nil {
	//   return nil, err
	// }
	return endpoint.SetGetOrderResponseRequest{
		Activity: activity,
		// Error:      erro, // TODO
		// Challenges: challenges, // TODO
	}, nil
}

func decodeGRPCSetGetOrderResponseResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	return endpoint.SetGetOrderResponseResponse{}, nil
}

func encodeGRPCSetGetOrderResponseResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &types.Empty{}, nil
}

func encodeGRPCSetGetOrderResponseRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.SetGetOrderResponseRequest)
	activity, err := EncodeGRPCActivity(req.Activity)
	if err != nil {
		return nil, err
	}
	// TODO
	// erro, err := encodeGRPCError(req.Error)
	// if err != nil {
	//   return nil, err
	// }
	// TODO
	// challenges, err := encodeGRPCChallenges(req.Challenges)
	// if err != nil {
	//   return nil, err
	// }
	return &apiv1.SetGetOrderResponseRequest{
		Activity: activity,
		// Error:      erro, // TODO
		// Challenges: challenges, // TODO
	}, nil
}

func decodeGRPCGetCreateAuthorizationRequestRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.Activity)
	activity, err := DecodeGRPCActivity(req)
	if err != nil {
		return nil, err
	}
	return endpoint.GetCreateAuthorizationRequestRequest{
		Activity: activity,
	}, nil
}

func decodeGRPCGetCreateAuthorizationRequestResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*apiv1.GetCreateAuthorizationRequestResponse)
	activity, err := DecodeGRPCActivity(reply.GetActivity())
	if err != nil {
		return nil, err
	}
	return endpoint.GetCreateAuthorizationRequestResponse{
		Activity: activity,
		// Domain:   reply.GetDomain(), // TODO
	}, nil
}

func encodeGRPCGetCreateAuthorizationRequestResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.GetCreateAuthorizationRequestResponse)
	activity, err := EncodeGRPCActivity(resp.Activity)
	if err != nil {
		return nil, err
	}
	return &apiv1.GetCreateAuthorizationRequestResponse{
		Activity: activity,
		// Domain:   resp.Domain, // TODO
	}, nil
}

func encodeGRPCGetCreateAuthorizationRequestRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.GetCreateAuthorizationRequestRequest)
	activity, err := EncodeGRPCActivity(req.Activity)
	if err != nil {
		return nil, err
	}
	return activity, nil
}

func decodeGRPCSetCreateAuthorizationResponseRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.SetCreateAuthorizationResponseRequest)
	activity, err := DecodeGRPCActivity(req.GetActivity())
	if err != nil {
		return nil, err
	}
	// TODO
	// erro, err := decodeGRPCError(req.GetError())
	// if err != nil {
	//   return nil, err
	// }
	// TODO
	// challenges, err := decodeGRPCChallenges(req.GetChallenges())
	// if err != nil {
	//   return nil, err
	// }
	return endpoint.SetCreateAuthorizationResponseRequest{
		Activity: activity,
		// Error:      erro, // TODO
		// Challenges: challenges, // TODO
	}, nil
}

func decodeGRPCSetCreateAuthorizationResponseResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	return endpoint.SetCreateAuthorizationResponseResponse{}, nil
}

func encodeGRPCSetCreateAuthorizationResponseResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &types.Empty{}, nil
}

func encodeGRPCSetCreateAuthorizationResponseRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.SetCreateAuthorizationResponseRequest)
	activity, err := EncodeGRPCActivity(req.Activity)
	if err != nil {
		return nil, err
	}
	// TODO
	// erro, err := encodeGRPCError(req.Error)
	// if err != nil {
	//   return nil, err
	// }
	// TODO
	// challenges, err := encodeGRPCChallenges(req.Challenges)
	// if err != nil {
	//   return nil, err
	// }
	return &apiv1.SetCreateAuthorizationResponseRequest{
		Activity: activity,
		// Error:      erro, // TODO
		// Challenges: challenges, // TODO
	}, nil
}

func decodeGRPCGetDeactivateAuthorizationRequestRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.Activity)
	activity, err := DecodeGRPCActivity(req)
	if err != nil {
		return nil, err
	}
	return endpoint.GetDeactivateAuthorizationRequestRequest{
		Activity: activity,
	}, nil
}

func decodeGRPCGetDeactivateAuthorizationRequestResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*apiv1.GetDeactivateAuthorizationRequestResponse)
	activity, err := DecodeGRPCActivity(reply.GetActivity())
	if err != nil {
		return nil, err
	}
	return endpoint.GetDeactivateAuthorizationRequestResponse{
		Activity: activity,
		// Domain:   reply.GetDomain(), // TODO
	}, nil
}

func encodeGRPCGetDeactivateAuthorizationRequestResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.GetDeactivateAuthorizationRequestResponse)
	activity, err := EncodeGRPCActivity(resp.Activity)
	if err != nil {
		return nil, err
	}
	return &apiv1.GetDeactivateAuthorizationRequestResponse{
		Activity: activity,
		// Domain:   resp.Domain, // TODO
	}, nil
}

func encodeGRPCGetDeactivateAuthorizationRequestRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.GetDeactivateAuthorizationRequestRequest)
	activity, err := EncodeGRPCActivity(req.Activity)
	if err != nil {
		return nil, err
	}
	return activity, nil
}

func decodeGRPCSetDeactivateAuthorizationResponseRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.SetDeactivateAuthorizationResponseRequest)
	activity, err := DecodeGRPCActivity(req.GetActivity())
	if err != nil {
		return nil, err
	}
	// TODO
	// erro, err := decodeGRPCError(req.GetError())
	// if err != nil {
	//   return nil, err
	// }
	// TODO
	// challenges, err := decodeGRPCChallenges(req.GetChallenges())
	// if err != nil {
	//   return nil, err
	// }
	return endpoint.SetDeactivateAuthorizationResponseRequest{
		Activity: activity,
		// Error:      erro, // TODO
		// Challenges: challenges, // TODO
	}, nil
}

func decodeGRPCSetDeactivateAuthorizationResponseResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	return endpoint.SetDeactivateAuthorizationResponseResponse{}, nil
}

func encodeGRPCSetDeactivateAuthorizationResponseResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &types.Empty{}, nil
}

func encodeGRPCSetDeactivateAuthorizationResponseRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.SetDeactivateAuthorizationResponseRequest)
	activity, err := EncodeGRPCActivity(req.Activity)
	if err != nil {
		return nil, err
	}
	// TODO
	// erro, err := encodeGRPCError(req.Error)
	// if err != nil {
	//   return nil, err
	// }
	// TODO
	// challenges, err := encodeGRPCChallenges(req.Challenges)
	// if err != nil {
	//   return nil, err
	// }
	return &apiv1.SetDeactivateAuthorizationResponseRequest{
		Activity: activity,
		// Error:      erro, // TODO
		// Challenges: challenges, // TODO
	}, nil
}

func decodeGRPCGetGetAuthorizationRequestRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.Activity)
	activity, err := DecodeGRPCActivity(req)
	if err != nil {
		return nil, err
	}
	return endpoint.GetGetAuthorizationRequestRequest{
		Activity: activity,
	}, nil
}

func decodeGRPCGetGetAuthorizationRequestResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*apiv1.GetGetAuthorizationRequestResponse)
	activity, err := DecodeGRPCActivity(reply.GetActivity())
	if err != nil {
		return nil, err
	}
	return endpoint.GetGetAuthorizationRequestResponse{
		Activity: activity,
		// Domain:   reply.GetDomain(), // TODO
	}, nil
}

func encodeGRPCGetGetAuthorizationRequestResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.GetGetAuthorizationRequestResponse)
	activity, err := EncodeGRPCActivity(resp.Activity)
	if err != nil {
		return nil, err
	}
	return &apiv1.GetGetAuthorizationRequestResponse{
		Activity: activity,
		// Domain:   resp.Domain, // TODO
	}, nil
}

func encodeGRPCGetGetAuthorizationRequestRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.GetGetAuthorizationRequestRequest)
	activity, err := EncodeGRPCActivity(req.Activity)
	if err != nil {
		return nil, err
	}
	return activity, nil
}

func decodeGRPCSetGetAuthorizationResponseRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.SetGetAuthorizationResponseRequest)
	activity, err := DecodeGRPCActivity(req.GetActivity())
	if err != nil {
		return nil, err
	}
	// TODO
	// erro, err := decodeGRPCError(req.GetError())
	// if err != nil {
	//   return nil, err
	// }
	// TODO
	// challenges, err := decodeGRPCChallenges(req.GetChallenges())
	// if err != nil {
	//   return nil, err
	// }
	return endpoint.SetGetAuthorizationResponseRequest{
		Activity: activity,
		// Error:      erro, // TODO
		// Challenges: challenges, // TODO
	}, nil
}

func decodeGRPCSetGetAuthorizationResponseResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	return endpoint.SetGetAuthorizationResponseResponse{}, nil
}

func encodeGRPCSetGetAuthorizationResponseResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &types.Empty{}, nil
}

func encodeGRPCSetGetAuthorizationResponseRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.SetGetAuthorizationResponseRequest)
	activity, err := EncodeGRPCActivity(req.Activity)
	if err != nil {
		return nil, err
	}
	// TODO
	// erro, err := encodeGRPCError(req.Error)
	// if err != nil {
	//   return nil, err
	// }
	// TODO
	// challenges, err := encodeGRPCChallenges(req.Challenges)
	// if err != nil {
	//   return nil, err
	// }
	return &apiv1.SetGetAuthorizationResponseRequest{
		Activity: activity,
		// Error:      erro, // TODO
		// Challenges: challenges, // TODO
	}, nil
}

func decodeGRPCGetGetChallengeRequestRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.Activity)
	activity, err := DecodeGRPCActivity(req)
	if err != nil {
		return nil, err
	}
	return endpoint.GetGetChallengeRequestRequest{
		Activity: activity,
	}, nil
}

func decodeGRPCGetGetChallengeRequestResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*apiv1.GetGetChallengeRequestResponse)
	activity, err := DecodeGRPCActivity(reply.GetActivity())
	if err != nil {
		return nil, err
	}
	return endpoint.GetGetChallengeRequestResponse{
		Activity: activity,
		// Domain:   reply.GetDomain(), // TODO
	}, nil
}

func encodeGRPCGetGetChallengeRequestResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.GetGetChallengeRequestResponse)
	activity, err := EncodeGRPCActivity(resp.Activity)
	if err != nil {
		return nil, err
	}
	return &apiv1.GetGetChallengeRequestResponse{
		Activity: activity,
		// Domain:   resp.Domain, // TODO
	}, nil
}

func encodeGRPCGetGetChallengeRequestRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.GetGetChallengeRequestRequest)
	activity, err := EncodeGRPCActivity(req.Activity)
	if err != nil {
		return nil, err
	}
	return activity, nil
}

func decodeGRPCSetGetChallengeResponseRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.SetGetChallengeResponseRequest)
	activity, err := DecodeGRPCActivity(req.GetActivity())
	if err != nil {
		return nil, err
	}
	// TODO
	// erro, err := decodeGRPCError(req.GetError())
	// if err != nil {
	//   return nil, err
	// }
	// TODO
	// challenges, err := decodeGRPCChallenges(req.GetChallenges())
	// if err != nil {
	//   return nil, err
	// }
	return endpoint.SetGetChallengeResponseRequest{
		Activity: activity,
		// Error:      erro, // TODO
		// Challenges: challenges, // TODO
	}, nil
}

func decodeGRPCSetGetChallengeResponseResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	return endpoint.SetGetChallengeResponseResponse{}, nil
}

func encodeGRPCSetGetChallengeResponseResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &types.Empty{}, nil
}

func encodeGRPCSetGetChallengeResponseRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.SetGetChallengeResponseRequest)
	activity, err := EncodeGRPCActivity(req.Activity)
	if err != nil {
		return nil, err
	}
	// TODO
	// erro, err := encodeGRPCError(req.Error)
	// if err != nil {
	//   return nil, err
	// }
	// TODO
	// challenges, err := encodeGRPCChallenges(req.Challenges)
	// if err != nil {
	//   return nil, err
	// }
	return &apiv1.SetGetChallengeResponseRequest{
		Activity: activity,
		// Error:      erro, // TODO
		// Challenges: challenges, // TODO
	}, nil
}

func decodeGRPCGetValidateChallengeRequestRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.Activity)
	activity, err := DecodeGRPCActivity(req)
	if err != nil {
		return nil, err
	}
	return endpoint.GetValidateChallengeRequestRequest{
		Activity: activity,
	}, nil
}

func decodeGRPCGetValidateChallengeRequestResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*apiv1.GetValidateChallengeRequestResponse)
	activity, err := DecodeGRPCActivity(reply.GetActivity())
	if err != nil {
		return nil, err
	}
	return endpoint.GetValidateChallengeRequestResponse{
		Activity: activity,
		// Domain:   reply.GetDomain(), // TODO
	}, nil
}

func encodeGRPCGetValidateChallengeRequestResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.GetValidateChallengeRequestResponse)
	activity, err := EncodeGRPCActivity(resp.Activity)
	if err != nil {
		return nil, err
	}
	return &apiv1.GetValidateChallengeRequestResponse{
		Activity: activity,
		// Domain:   resp.Domain, // TODO
	}, nil
}

func encodeGRPCGetValidateChallengeRequestRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.GetValidateChallengeRequestRequest)
	activity, err := EncodeGRPCActivity(req.Activity)
	if err != nil {
		return nil, err
	}
	return activity, nil
}

func decodeGRPCSetValidateChallengeResponseRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.SetValidateChallengeResponseRequest)
	activity, err := DecodeGRPCActivity(req.GetActivity())
	if err != nil {
		return nil, err
	}
	// TODO
	// erro, err := decodeGRPCError(req.GetError())
	// if err != nil {
	//   return nil, err
	// }
	// TODO
	// challenges, err := decodeGRPCChallenges(req.GetChallenges())
	// if err != nil {
	//   return nil, err
	// }
	return endpoint.SetValidateChallengeResponseRequest{
		Activity: activity,
		// Error:      erro, // TODO
		// Challenges: challenges, // TODO
	}, nil
}

func decodeGRPCSetValidateChallengeResponseResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	return endpoint.SetValidateChallengeResponseResponse{}, nil
}

func encodeGRPCSetValidateChallengeResponseResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &types.Empty{}, nil
}

func encodeGRPCSetValidateChallengeResponseRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.SetValidateChallengeResponseRequest)
	activity, err := EncodeGRPCActivity(req.Activity)
	if err != nil {
		return nil, err
	}
	// TODO
	// erro, err := encodeGRPCError(req.Error)
	// if err != nil {
	//   return nil, err
	// }
	// TODO
	// challenges, err := encodeGRPCChallenges(req.Challenges)
	// if err != nil {
	//   return nil, err
	// }
	return &apiv1.SetValidateChallengeResponseRequest{
		Activity: activity,
		// Error:      erro, // TODO
		// Challenges: challenges, // TODO
	}, nil
}

func decodeGRPCGetGetCertificateRequestRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.Activity)
	activity, err := DecodeGRPCActivity(req)
	if err != nil {
		return nil, err
	}
	return endpoint.GetGetCertificateRequestRequest{
		Activity: activity,
	}, nil
}

func decodeGRPCGetGetCertificateRequestResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*apiv1.GetGetCertificateRequestResponse)
	activity, err := DecodeGRPCActivity(reply.GetActivity())
	if err != nil {
		return nil, err
	}
	return endpoint.GetGetCertificateRequestResponse{
		Activity: activity,
		// Domain:   reply.GetDomain(), // TODO
	}, nil
}

func encodeGRPCGetGetCertificateRequestResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.GetGetCertificateRequestResponse)
	activity, err := EncodeGRPCActivity(resp.Activity)
	if err != nil {
		return nil, err
	}
	return &apiv1.GetGetCertificateRequestResponse{
		Activity: activity,
		// Domain:   resp.Domain, // TODO
	}, nil
}

func encodeGRPCGetGetCertificateRequestRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.GetGetCertificateRequestRequest)
	activity, err := EncodeGRPCActivity(req.Activity)
	if err != nil {
		return nil, err
	}
	return activity, nil
}

func decodeGRPCSetGetCertificateResponseRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.SetGetCertificateResponseRequest)
	activity, err := DecodeGRPCActivity(req.GetActivity())
	if err != nil {
		return nil, err
	}
	// TODO
	// erro, err := decodeGRPCError(req.GetError())
	// if err != nil {
	//   return nil, err
	// }
	// TODO
	// challenges, err := decodeGRPCChallenges(req.GetChallenges())
	// if err != nil {
	//   return nil, err
	// }
	return endpoint.SetGetCertificateResponseRequest{
		Activity: activity,
		// Error:      erro, // TODO
		// Challenges: challenges, // TODO
	}, nil
}

func decodeGRPCSetGetCertificateResponseResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	return endpoint.SetGetCertificateResponseResponse{}, nil
}

func encodeGRPCSetGetCertificateResponseResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &types.Empty{}, nil
}

func encodeGRPCSetGetCertificateResponseRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.SetGetCertificateResponseRequest)
	activity, err := EncodeGRPCActivity(req.Activity)
	if err != nil {
		return nil, err
	}
	// TODO
	// erro, err := encodeGRPCError(req.Error)
	// if err != nil {
	//   return nil, err
	// }
	// TODO
	// challenges, err := encodeGRPCChallenges(req.Challenges)
	// if err != nil {
	//   return nil, err
	// }
	return &apiv1.SetGetCertificateResponseRequest{
		Activity: activity,
		// Error:      erro, // TODO
		// Challenges: challenges, // TODO
	}, nil
}

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
		Activity: activity,
		// Domain:   reply.GetDomain(), // TODO
	}, nil
}

func encodeGRPCGetRevokeCertificateRequestResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.GetRevokeCertificateRequestResponse)
	activity, err := EncodeGRPCActivity(resp.Activity)
	if err != nil {
		return nil, err
	}
	return &apiv1.GetRevokeCertificateRequestResponse{
		Activity: activity,
		// Domain:   resp.Domain, // TODO
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
	// TODO
	// erro, err := decodeGRPCError(req.GetError())
	// if err != nil {
	//   return nil, err
	// }
	// TODO
	// challenges, err := decodeGRPCChallenges(req.GetChallenges())
	// if err != nil {
	//   return nil, err
	// }
	return endpoint.SetRevokeCertificateResponseRequest{
		Activity: activity,
		// Error:      erro, // TODO
		// Challenges: challenges, // TODO
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
	// TODO
	// erro, err := encodeGRPCError(req.Error)
	// if err != nil {
	//   return nil, err
	// }
	// TODO
	// challenges, err := encodeGRPCChallenges(req.Challenges)
	// if err != nil {
	//   return nil, err
	// }
	return &apiv1.SetRevokeCertificateResponseRequest{
		Activity: activity,
		// Error:      erro, // TODO
		// Challenges: challenges, // TODO
	}, nil
}
