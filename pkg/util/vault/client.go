package vault

import (
	"github.com/hashicorp/vault/api"
)

type Client struct {
	c *api.Logical
}

func New(address, token, caFile string) (*Client, error) {
	conf := api.DefaultConfig()
	conf.ConfigureTLS(&api.TLSConfig{CAPath: caFile})
	c, err := api.NewClient(conf)
	if err != nil {
		return nil, err
	}
	c.SetAddress(address)
	c.SetToken(token)
	return &Client{c: c.Logical()}, nil
}
