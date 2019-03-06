package vault

import (
	"context"

	"github.com/hashicorp/vault/api"
	"github.com/opentracing/opentracing-go"
)

type Client struct {
	c *api.Client
}

func New(address, token, caFile string) (*Client, error) {
	conf := api.DefaultConfig()
	if caFile != "" {
		conf.ConfigureTLS(&api.TLSConfig{CAPath: caFile})
	}
	// conf.ConfigureTLS(&api.TLSConfig{TLSServerName: "vault"})

	c, err := api.NewClient(conf)
	if err != nil {
		return nil, err
	}

	if address != "" {
		c.SetAddress(address)
	}
	if token != "" {
		c.SetToken(token)
	}

	return &Client{c: c}, nil
}

func (c *Client) Auth() *api.Auth {
	return c.c.Auth()
}

func (c *Client) Logical() *api.Logical {
	return c.c.Logical()
}

func (c *Client) Sys() *api.Sys {
	return c.c.Sys()
}

func (c *Client) logicalWrite(ctx context.Context, operation, path string, data map[string]interface{}) (*api.Secret, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "vault")
	defer span.Finish()
	span.SetTag("operation", operation)

	return c.c.Logical().Write(path, data)
}
