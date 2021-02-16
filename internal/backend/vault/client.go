package vault

import (
	"context"
	"github.com/hashicorp/vault/api"
	"github.com/opentracing/opentracing-go"
)

var clientValue = struct{}{}

func GetClient(ctx context.Context) *Client {
	return ctx.Value(clientValue).(*Client)
}

func SetClient(ctx context.Context, client *Client) context.Context {
	return context.WithValue(ctx, clientValue, client)
}

type ClientConfig struct {
	CAFile string `mapstructure:"ca-file"`
	Token  string `validate:"required"`
	URL    string `validate:"required,url"`
}

type Client struct {
	c *api.Client
}

func New(cfg ClientConfig) (*Client, error) {
	conf := api.DefaultConfig()
	if cfg.CAFile != "" {
		conf.ConfigureTLS(&api.TLSConfig{CAPath: cfg.CAFile})
	}
	// conf.ConfigureTLS(&api.TLSConfig{TLSServerName: "vault"})

	c, err := api.NewClient(conf)
	if err != nil {
		return nil, err
	}

	if cfg.URL != "" {
		c.SetAddress(cfg.URL)
	}
	if cfg.Token != "" {
		c.SetToken(cfg.Token)
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
	span, _ := opentracing.StartSpanFromContext(ctx, "vault")
	defer span.Finish()
	span.SetTag("operation", operation)

	return c.c.Logical().Write(path, data)
}
