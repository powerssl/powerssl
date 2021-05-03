package vault // import "powerssl.dev/backend/vault"

import (
	"context"

	"github.com/hashicorp/vault/api"
	"github.com/opentracing/opentracing-go"

	"powerssl.dev/backend/ctxkey"
)

var clientValue = ctxkey.New("dev.powerssl.backend.vault")

func GetClient(ctx context.Context) *Client {
	return ctx.Value(clientValue).(*Client)
}

func SetClient(ctx context.Context, client *Client) context.Context {
	return context.WithValue(ctx, clientValue, client)
}

type ClientConfig struct {
	AppRole struct {
		RoleID   string `mapstructure:"role-id"`
		SecretID string `mapstructure:"secret-id"`
	} `mapstructure:",squash"`
	CAFile string `mapstructure:"ca-file"`
	Token  string
	URL    string `validate:"url"`
}

type Client struct {
	c   *api.Client
	cfg ClientConfig
}

func New(cfg *ClientConfig) (*Client, error) {
	conf := api.DefaultConfig()
	if cfg.CAFile != "" {
		if err := conf.ConfigureTLS(&api.TLSConfig{
			CAPath: cfg.CAFile,
		}); err != nil {
			return nil, err
		}
	}
	if cfg.URL != "" {
		conf.Address = cfg.URL
	}

	c, err := api.NewClient(conf)
	if err != nil {
		return nil, err
	}
	if cfg.Token == "" && cfg.AppRole.RoleID != "" && cfg.AppRole.SecretID != "" {
		var resp *api.Secret
		if resp, err = c.Logical().Write("auth/approle/login", map[string]interface{}{
			"role_id":   cfg.AppRole.RoleID,
			"secret_id": cfg.AppRole.SecretID,
		}); err != nil {
			return nil, err
		}
		cfg.Token = resp.Auth.ClientToken
	}
	if cfg.Token != "" {
		c.SetToken(cfg.Token)
	}

	return &Client{
		c:   c,
		cfg: *cfg,
	}, nil
}

func (c *Client) Auth() *api.Auth {
	return c.c.Auth()
}

func (c *Client) Clone() (*api.Client, error) {
	vault, err := c.c.Clone()
	if err != nil {
		return nil, err
	}
	if c.cfg.Token != "" {
		vault.SetToken(c.cfg.Token)
	}
	return vault, nil
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

	return c.Logical().Write(path, data)
}

func LogicalRead(ctx context.Context, vault *api.Client, operation, path string) (*api.Secret, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "vault")
	defer span.Finish()
	span.SetTag("operation", operation)

	return vault.Logical().Read(path)
}
