package vault

import (
	"crypto"
	"crypto/rsa"
	"fmt"
	"strconv"

	"github.com/hashicorp/vault/api"
	"github.com/kenshaw/pemutil"
)

func PrivateKeyFromKeyToken(keyToken string) (_ crypto.PrivateKey, err error) {
	conf := api.DefaultConfig()
	conf.Address = "https://localhost:8200"
	if err := conf.ConfigureTLS(&api.TLSConfig{
		CAPath: "/Users/dspangenberg/src/powerssl.io/powerssl/local/certs/ca.pem",
	}); err != nil {
		return nil, err
	}
	var c *api.Client
	if c, err = api.NewClient(conf); err != nil {
		return nil, err
	}
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

func SignerFromKeyToken(keyToken string) (_ crypto.Signer, err error) {
	var privateKey crypto.PrivateKey
	if privateKey, err = PrivateKeyFromKeyToken(keyToken); err != nil {
		return nil, err
	}
	var signer crypto.Signer
	var ok bool
	if signer, ok = privateKey.(*rsa.PrivateKey); !ok {
		return nil, fmt.Errorf("private key not an RSA private key")
	}
	return signer, nil
}
