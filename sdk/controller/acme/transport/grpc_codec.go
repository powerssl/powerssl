package transport // import "powerssl.dev/sdk/controller/acme/transport"

import (
	"context"

	apiv1 "powerssl.dev/api/controller/v1"

	"powerssl.dev/sdk/controller/acme/endpoint"
	"powerssl.dev/sdk/controller/api"
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

func encodeGRPCError(erro *api.Error) (*apiv1.Error, error) {
	if erro == nil {
		erro = &api.Error{}
	}
	return &apiv1.Error{
		Message: erro.Message,
	}, nil
}

func encodeGRPCAccount(account *api.Account) (*apiv1.Account, error) {
	if account == nil { // TODO
		account = new(api.Account)
	}
	return &apiv1.Account{
		Status:               apiv1.Account_Status(account.Status),
		Contacts:             account.Contacts,
		TermsOfServiceAgreed: account.TermsOfServiceAgreed,
		Url:                  account.URL,
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
		Contacts:             reply.GetContacts(),
		DirectoryURL:         reply.GetDirectoryUrl(),
		KeyToken:             reply.GetKeyToken(),
		TermsOfServiceAgreed: reply.GetTermsOfServiceAgreed(),
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

func decodeGRPCSetCreateAccountResponseResponse(_ context.Context, _ interface{}) (interface{}, error) {
	return endpoint.SetCreateAccountResponseResponse{}, nil
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

func encodeGRPCGetDeactivateAccountRequestRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.GetDeactivateAccountRequestRequest)
	activity, err := EncodeGRPCActivity(req.Activity)
	if err != nil {
		return nil, err
	}
	return activity, nil
}

func decodeGRPCSetDeactivateAccountResponseResponse(_ context.Context, _ interface{}) (interface{}, error) {
	return endpoint.SetDeactivateAccountResponseResponse{}, nil
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

func encodeGRPCGetRekeyAccountRequestRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.GetRekeyAccountRequestRequest)
	activity, err := EncodeGRPCActivity(req.Activity)
	if err != nil {
		return nil, err
	}
	return activity, nil
}

func decodeGRPCSetRekeyAccountResponseResponse(_ context.Context, _ interface{}) (interface{}, error) {
	return endpoint.SetRekeyAccountResponseResponse{}, nil
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

func encodeGRPCGetUpdateAccountRequestRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.GetUpdateAccountRequestRequest)
	activity, err := EncodeGRPCActivity(req.Activity)
	if err != nil {
		return nil, err
	}
	return activity, nil
}

func decodeGRPCSetUpdateAccountResponseResponse(_ context.Context, _ interface{}) (interface{}, error) {
	return endpoint.SetUpdateAccountResponseResponse{}, nil
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

func encodeGRPCGetCreateOrderRequestRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.GetCreateOrderRequestRequest)
	activity, err := EncodeGRPCActivity(req.Activity)
	if err != nil {
		return nil, err
	}
	return activity, nil
}

func decodeGRPCSetCreateOrderResponseResponse(_ context.Context, _ interface{}) (interface{}, error) {
	return endpoint.SetCreateOrderResponseResponse{}, nil
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

func encodeGRPCGetFinalizeOrderRequestRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.GetFinalizeOrderRequestRequest)
	activity, err := EncodeGRPCActivity(req.Activity)
	if err != nil {
		return nil, err
	}
	return activity, nil
}

func decodeGRPCSetFinalizeOrderResponseResponse(_ context.Context, _ interface{}) (interface{}, error) {
	return endpoint.SetFinalizeOrderResponseResponse{}, nil
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

func encodeGRPCGetGetOrderRequestRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.GetGetOrderRequestRequest)
	activity, err := EncodeGRPCActivity(req.Activity)
	if err != nil {
		return nil, err
	}
	return activity, nil
}

func decodeGRPCSetGetOrderResponseResponse(_ context.Context, _ interface{}) (interface{}, error) {
	return endpoint.SetGetOrderResponseResponse{}, nil
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

func encodeGRPCGetCreateAuthorizationRequestRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.GetCreateAuthorizationRequestRequest)
	activity, err := EncodeGRPCActivity(req.Activity)
	if err != nil {
		return nil, err
	}
	return activity, nil
}

func decodeGRPCSetCreateAuthorizationResponseResponse(_ context.Context, _ interface{}) (interface{}, error) {
	return endpoint.SetCreateAuthorizationResponseResponse{}, nil
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

func encodeGRPCGetDeactivateAuthorizationRequestRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.GetDeactivateAuthorizationRequestRequest)
	activity, err := EncodeGRPCActivity(req.Activity)
	if err != nil {
		return nil, err
	}
	return activity, nil
}

func decodeGRPCSetDeactivateAuthorizationResponseResponse(_ context.Context, _ interface{}) (interface{}, error) {
	return endpoint.SetDeactivateAuthorizationResponseResponse{}, nil
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

func encodeGRPCGetGetAuthorizationRequestRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.GetGetAuthorizationRequestRequest)
	activity, err := EncodeGRPCActivity(req.Activity)
	if err != nil {
		return nil, err
	}
	return activity, nil
}

func decodeGRPCSetGetAuthorizationResponseResponse(_ context.Context, _ interface{}) (interface{}, error) {
	return endpoint.SetGetAuthorizationResponseResponse{}, nil
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

func encodeGRPCGetGetChallengeRequestRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.GetGetChallengeRequestRequest)
	activity, err := EncodeGRPCActivity(req.Activity)
	if err != nil {
		return nil, err
	}
	return activity, nil
}

func decodeGRPCSetGetChallengeResponseResponse(_ context.Context, _ interface{}) (interface{}, error) {
	return endpoint.SetGetChallengeResponseResponse{}, nil
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

func encodeGRPCGetValidateChallengeRequestRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.GetValidateChallengeRequestRequest)
	activity, err := EncodeGRPCActivity(req.Activity)
	if err != nil {
		return nil, err
	}
	return activity, nil
}

func decodeGRPCSetValidateChallengeResponseResponse(_ context.Context, _ interface{}) (interface{}, error) {
	return endpoint.SetValidateChallengeResponseResponse{}, nil
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

func encodeGRPCGetGetCertificateRequestRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.GetGetCertificateRequestRequest)
	activity, err := EncodeGRPCActivity(req.Activity)
	if err != nil {
		return nil, err
	}
	return activity, nil
}

func decodeGRPCSetGetCertificateResponseResponse(_ context.Context, _ interface{}) (interface{}, error) {
	return endpoint.SetGetCertificateResponseResponse{}, nil
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

func encodeGRPCGetRevokeCertificateRequestRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.GetRevokeCertificateRequestRequest)
	activity, err := EncodeGRPCActivity(req.Activity)
	if err != nil {
		return nil, err
	}
	return activity, nil
}

func decodeGRPCSetRevokeCertificateResponseResponse(_ context.Context, _ interface{}) (interface{}, error) {
	return endpoint.SetRevokeCertificateResponseResponse{}, nil
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
