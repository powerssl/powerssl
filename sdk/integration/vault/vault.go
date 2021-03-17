package vault

import (
	"crypto"
	"crypto/rsa"
	"fmt"
	"strconv"

	"github.com/hashicorp/vault/api"
	"github.com/kenshaw/pemutil"
)

type Config struct {
	Address       string
	CAFile        string
	ClientCert    string
	ClientKey     string
	Insecure      bool
	TLSServerName string
}

type Client struct {
	c *api.Client
}

func NewClient(cfg Config) (_ *Client, err error) {
	conf := api.DefaultConfig()
	conf.Address = cfg.Address
	if err = conf.ConfigureTLS(&api.TLSConfig{
		CACert:        cfg.CAFile,
		ClientCert:    cfg.ClientCert,
		ClientKey:     cfg.ClientKey,
		Insecure:      cfg.Insecure,
		TLSServerName: cfg.TLSServerName,
	}); err != nil {
		return nil, err
	}
	var c *api.Client
	if c, err = api.NewClient(conf); err != nil {
		return nil, err
	}
	return &Client{
		c: c,
	}, nil
}

func (c *Client) Logical() *api.Logical {
	return c.c.Logical()
}

func (c *Client) PrivateKeyFromKeyToken(keyToken string) (_ crypto.PrivateKey, err error) {
	var secret *api.Secret
	if secret, err = c.Logical().Unwrap(keyToken); err != nil {
		return nil, err
	}
	keysInterface, ok := secret.Data["keys"]
	if !ok {
		return nil, fmt.Errorf("key \"keys\" not found")
	}
	keys, ok := keysInterface.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("\"keys\" not an object")
	}
	var key string
	var version int
	for versionString, keyInterface := range keys {
		var versionInt int
		if versionInt, err = strconv.Atoi(versionString); err != nil {
			return nil, err
		}
		if versionInt <= version {
			continue
		}
		version = versionInt
		key, ok = keyInterface.(string)
		if !ok {
			return nil, fmt.Errorf("key not a string")
		}
	}
	var store pemutil.Store
	if store, err = pemutil.DecodeBytes([]byte(key)); err != nil {
		return nil, err
	}
	var privateKey crypto.PrivateKey
	if privateKey, ok = store.PrivateKey(); !ok {
		return nil, fmt.Errorf("PEM does not contain an private key")
	}
	return privateKey, nil
}

func (c *Client) SignerFromKeyToken(keyToken string) (_ crypto.Signer, err error) {
	var privateKey crypto.PrivateKey
	if privateKey, err = c.PrivateKeyFromKeyToken(keyToken); err != nil {
		return nil, err
	}
	var signer crypto.Signer
	var ok bool
	if signer, ok = privateKey.(*rsa.PrivateKey); !ok {
		return nil, fmt.Errorf("private key not an RSA private key")
	}
	return signer, nil
}
