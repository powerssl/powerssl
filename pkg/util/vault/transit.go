package vault

import (
	"context"
	"fmt"

	"github.com/hashicorp/vault/api"
	"github.com/opentracing/opentracing-go"
)

func (c *Client) CreateTransitKey(ctx context.Context, name string) error {
	_, err := c.write(ctx,
		"CreateTransitKey",
		fmt.Sprintf("/transit/keys/%s", name),
		map[string]interface{}{"type": "rsa-4096"},
	)
	return err
}

func (c *Client) SignTransitData(ctx context.Context, name, input string) (*api.Secret, error) {
	return c.write(ctx,
		"SignTransitData",
		fmt.Sprintf("/transit/sign/%s/sha2-256", name),
		map[string]interface{}{"input": input},
	)
}

func (c *Client) write(ctx context.Context, operation, path string, data map[string]interface{}) (*api.Secret, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "vault")
	defer span.Finish()
	span.SetTag("operation", operation)

	return c.c.Write(path, data)
}
