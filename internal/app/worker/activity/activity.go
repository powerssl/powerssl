package activity

import (
	"context"

	"powerssl.dev/powerssl/pkg/apiserver/api"
)

func CreateACMEAccount(ctx context.Context, directoryURL string, termsOfServiceAgreed bool, contacts []string) error {
	return nil
}

func UpdateAccount(ctx context.Context, name string, updateMask []string, acmeAccount *api.ACMEAccount) error {
	return nil
}
