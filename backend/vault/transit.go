package vault // import "powerssl.dev/backend/vault"

import (
	"context"
	"fmt"
	"strings"
)

func (c *Client) CreateTransitKey(ctx context.Context, name string) error {
	_, err := c.logicalWrite(ctx,
		"CreateTransitKey",
		fmt.Sprintf("transit/keys/%s", sanitizeTransitKey(name)),
		map[string]interface{}{
			"exportable": true,
			"type":       "rsa-4096",
		},
	)
	return err
}

func (c *Client) ExportWrappedTransitKey(ctx context.Context, name string) (string, error) {
	vault, err := c.Clone()
	if err != nil {
		return "", err
	}
	vault.SetWrappingLookupFunc(func(operation, path string) string {
		return "5m"
	})
	secret, err := LogicalRead(ctx, vault, "ExportTransitKey",
		fmt.Sprintf("transit/export/signing-key/%s/latest", sanitizeTransitKey(name)))
	if err != nil {
		return "", err
	}
	if secret == nil {
		return "", fmt.Errorf("secret does not exist")
	}
	if secret.WrapInfo == nil {
		return "", fmt.Errorf("no wrapping info")
	}
	return secret.WrapInfo.Token, nil
}

func sanitizeTransitKey(name string) string {
	return strings.ReplaceAll(name, "/", "-")
}
