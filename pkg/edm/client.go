package edm

import (
	"context"

	"github.com/coreos/etcd/clientv3"
	"gopkg.in/go-playground/validator.v9"
)

type Client struct {
	Config *Config          `validate:"required"`
	ctx    context.Context  `validate:"required"`
	Etcd   *clientv3.Client `validate:"required"`
}

func (c *Client) Close() {
	c.Etcd.Close()
}

func New(ctx context.Context, config *Config) (*Client, error) {
	err := config.Validate()
	if err != nil {
		return nil, err
	}
}
