package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	apiv1 "powerssl.dev/api/apiserver/v1"
	"powerssl.dev/backend/auth"
)

var acmeAccountMessageType *apiv1.ACMEAccount
var acmeAccountUpdateMaskSanitizer = NewUpdateMaskSanitizer(acmeAccountMessageType).
	Allowed("display_name", "title", "description", "contacts").
	Internal("terms_of_service_agreed", "contacts", "account_url")

func (a *AcmeAccount) Name() string {
	return fmt.Sprintf("acmeServers/%s/acmeAccounts/%s", a.AcmeServerID, a.ID)
}

func (a *AcmeAccount) ToAPI() *apiv1.ACMEAccount {
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

func (q *Queries) CreateACMEAccountFromAPI(ctx context.Context, parent string, acmeAccount *apiv1.ACMEAccount) (AcmeAccount, error) {
	s := strings.Split(parent, "/")
	if len(s) != 2 || s[0] != "acmeServers" {
		return AcmeAccount{}, status.Error(codes.InvalidArgument, "malformed parent")
	}
	acmeServerID, err := uuid.Parse(s[1])
	if err != nil {
		return AcmeAccount{}, err
	}
	return q.CreateACMEAccount(ctx, CreateACMEAccountParams{
		AcmeServerID:         acmeServerID,
		DisplayName:          acmeAccount.DisplayName,
		Title:                acmeAccount.Title,
		Description:          acmeAccount.Description,
		TermsOfServiceAgreed: acmeAccount.TermsOfServiceAgreed,
		Contacts:             strings.Join(acmeAccount.Contacts, ","),
	})
}

func (q *Queries) UpdateACMEAccountWithMask(ctx context.Context, id uuid.UUID, fm *fieldmaskpb.FieldMask, acmeAccount *apiv1.ACMEAccount) (AcmeAccount, error) {
	fm = acmeAccountUpdateMaskSanitizer.Sanitize(fm, auth.IsInternal(ctx))
	updateParams := UpdateACMEAccountParams{ID: id}
	if err := setUpdateParams(fm, acmeAccount, &updateParams); err != nil {
		return AcmeAccount{}, err
	}
	return q.UpdateACMEAccount(ctx, updateParams)
}

type AcmeAccounts []AcmeAccount

func (a AcmeAccounts) ToAPI() []*apiv1.ACMEAccount {
	acmeAccounts := make([]*apiv1.ACMEAccount, len(a))
	for i, account := range a {
		acmeAccounts[i] = account.ToAPI()
	}
	return acmeAccounts
}

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
