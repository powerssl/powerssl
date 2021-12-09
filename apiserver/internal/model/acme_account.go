package model

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	apiv1 "powerssl.dev/api/apiserver/v1"
	"powerssl.dev/apiserver/internal/repository"
)

// TODO: Doesn't get called
// func (a *ACMEAccount) Validate(db *pg.DB) (map[string]error, bool) {
//	 var errors map[string]error
//
//	 if !a.TermsOfServiceAgreed {
//		 errors["TermsOfServiceAgreed"] = status.Error(codes.InvalidArgument, "terms of service need to be agreed")
//	 }
//
//	 if a.ACMEServer == nil {
//		 errors["TermsOfServiceAgreed"] = status.Error(codes.NotFound, "ACME server not found")
//	 }
//
//	 return errors, len(errors) == 0
// }

var acmeAccountMessageType *apiv1.ACMEAccount
var acmeAccountUpdateMaskSanitizer = NewUpdateMaskSanitizer(acmeAccountMessageType).
	Allowed("display_name", "title", "description", "contacts").
	Internal("terms_of_service_agreed", "contacts", "account_url")

func AcmeAccountUpdateParams(ctx context.Context, id uuid.UUID, fm *fieldmaskpb.FieldMask, acmeAccount *apiv1.ACMEAccount) (updateParams repository.UpdateACMEAccountParams, err error) {
	updateParams.ID = id
	err = setUpdateParams(acmeAccountUpdateMaskSanitizer.Sanitize(ctx, fm), acmeAccount, &updateParams)
	return updateParams, err
}

func ParseAcmeAccountID(name string) (id uuid.UUID, err error) {
	n := strings.Split(name, "/")
	if len(n) != 4 || n[0] != "acmeServers" || n[2] == "acmeAccounts" {
		return id, fmt.Errorf("acme account name format mismatch")
	}
	return uuid.Parse(n[3])
}

type AcmeAccount struct {
	repository.AcmeAccount
}

func NewAcmeAccount(acmeAccount repository.AcmeAccount) AcmeAccount {
	return AcmeAccount{
		AcmeAccount: acmeAccount,
	}
}

func (a AcmeAccount) Encode() *apiv1.ACMEAccount {
	var contacts []string
	if a.Contacts != "" {
		contacts = strings.Split(a.Contacts, ",")
	}
	return &apiv1.ACMEAccount{
		Name:                 a.Name(),
		CreateTime:           timestamppb.New(a.CreatedAt),
		UpdateTime:           timestamppb.New(a.UpdatedAt),
		DisplayName:          a.DisplayName,
		Title:                a.Title,
		Description:          a.Description,
		Labels:               map[string]string{"not": "implemented"},
		TermsOfServiceAgreed: a.TermsOfServiceAgreed,
		Contacts:             contacts,
		AccountUrl:           a.AccountUrl,
	}
}

func (a AcmeAccount) Name() string {
	return fmt.Sprintf("acmeServers/%s/acmeAccounts/%s", a.AcmeServerID, a.ID)
}

type AcmeAccounts []AcmeAccount

func NewAcmeAccounts(acmeAccounts []repository.AcmeAccount) AcmeAccounts {
	as := make(AcmeAccounts, len(acmeAccounts))
	for i, acmeAccount := range acmeAccounts {
		as[i] = NewAcmeAccount(acmeAccount)
	}
	return as
}

func (a AcmeAccounts) Encode() []*apiv1.ACMEAccount {
	acmeAccounts := make([]*apiv1.ACMEAccount, len(a))
	for i, account := range a {
		acmeAccounts[i] = account.Encode()
	}
	return acmeAccounts
}
