package vault

import (
	"context"
	"fmt"

	"github.com/hashicorp/vault/api"
)

func (c *Client) CreateTransitKey(ctx context.Context, name string) error {
	_, err := c.logicalWrite(ctx,
		"CreateTransitKey",
		fmt.Sprintf("/transit/keys/%s", name),
		map[string]interface{}{"type": "rsa-4096"},
	)
	return err
}

func (c *Client) SignTransitData(ctx context.Context, name, input string) (*api.Secret, error) {
	return c.logicalWrite(ctx,
		"SignTransitData",
		fmt.Sprintf("/transit/sign/%s/sha2-256", name),
		map[string]interface{}{"input": input},
	)
}
