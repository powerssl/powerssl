package powerutil

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ghodss/yaml"
	"github.com/hashicorp/vault/api"

	"powerssl.dev/powerssl/internal/app/powerutil/policy"
	"powerssl.dev/powerssl/internal/pkg/pki"
	"powerssl.dev/powerssl/internal/pkg/vault"
)

const (
	PkiPath     = "pki"
	TransitPath = "transit"
)

var mounts = map[string]*api.MountInput{
	PkiPath: {
		Type:        "pki",
		Description: "PowerSSL PKI",
		Config: api.MountConfigInput{
			MaxLeaseTTL: "8760h",
		},
	},
	TransitPath: {
		Type:        "transit",
		Description: "PowerSSL Transit",
	},
}

var roles = map[string]map[string]interface{}{
	fmt.Sprintf("%s/roles/powerssl-apiserver", PkiPath): {
		"allowed_domains":    "apiserver",
		"allow_localhost":    true, // TODO: Disable in production
		"allow_bare_domains": true,
		"max_ttl":            "24h",
	},
	fmt.Sprintf("%s/roles/powerssl-controller", PkiPath): {
		"allowed_domains":    "controller",
		"allow_localhost":    true, // TODO: Disable in production
		"allow_bare_domains": true,
		"max_ttl":            "24h",
	},
	fmt.Sprintf("%s/roles/powerssl-signer", PkiPath): {
		"allowed_domains":    "signer",
		"allow_localhost":    true, // TODO: Disable in production
		"allow_bare_domains": true,
		"max_ttl":            "24h",
	},
	fmt.Sprintf("%s/roles/powerssl-worker", PkiPath): {
		"allowed_domains":    "worker",
		"allow_localhost":    true, // TODO: Disable in production
		"allow_bare_domains": true,
		"max_ttl":            "24h",
	},
}

var tokens = map[string][]string{
	"powerssl-apiserver":  {"powerssl-apiserver"},
	"powerssl-controller": {"powerssl-controller"},
	"powerssl-signer":     {"powerssl-signer"},
	"powerssl-worker":     {"powerssl-worker"},
}

func RunVault(addr, ca, caKey string) error {
	var keys []string
	var rootToken string
	{
		c, err := vault.New(vault.ClientConfig{
			CAFile: ca,
			URL:    addr,
		})
		if err != nil {
			return err
		}

		var recoveryKeys []string
		keys, recoveryKeys, rootToken, err = vaultInit(c)
		if err != nil {
			return err
		}

		secret, err := yaml.Marshal(map[string]interface{}{
			"keys":         keys,
			"recoveryKeys": recoveryKeys,
			"rootToken":    rootToken,
		})
		if err != nil {
			return err
		}
		if err := os.MkdirAll("local/vault", 0755); err != nil {
			return err
		}
		if err := ioutil.WriteFile("local/vault/secret.yaml", secret, 0644); err != nil {
			return err
		}
	}

	c, err := vault.New(vault.ClientConfig{
		CAFile: ca,
		Token:  rootToken,
		URL:    addr,
	})
	if err != nil {
		return err
	}

	if err := vaultUnseal(c, keys); err != nil {
		return err
	}

	if err := vaultMount(c); err != nil {
		return err
	}

	if err := vaultInitPKI(c, addr, ca, caKey); err != nil {
		return err
	}

	if err := vaultWritePKIRoles(c); err != nil {
		return err
	}

	if err := vaultPutPolicies(c); err != nil {
		return err
	}

	if err := vaultCreateTokens(c); err != nil {
		return err
	}

	return nil
}

func vaultInit(c *vault.Client) ([]string, []string, string, error) {
	var SecretShares = 1
	var SecretThreshold = 1
	var StoredShares int
	var PGPKeys []string
	var RecoveryShares int
	var RecoveryThreshold int
	var RecoveryPGPKeys []string
	var RootTokenPGPKey string

	var keys, recoveryKeys []string
	var rootToken string

	initialized, err := c.Sys().InitStatus()
	if err != nil {
		return keys, recoveryKeys, rootToken, err
	}

	if !initialized {
		response, err := c.Sys().Init(&api.InitRequest{
			SecretShares:      SecretShares,
			SecretThreshold:   SecretThreshold,
			StoredShares:      StoredShares,
			PGPKeys:           PGPKeys,
			RecoveryShares:    RecoveryShares,
			RecoveryThreshold: RecoveryThreshold,
			RecoveryPGPKeys:   RecoveryPGPKeys,
			RootTokenPGPKey:   RootTokenPGPKey,
		})
		if err != nil {
			return keys, recoveryKeys, rootToken, err
		}

		keys = response.Keys
		recoveryKeys = response.RecoveryKeys
		rootToken = response.RootToken

		fmt.Printf("Keys: %s\n", response.Keys)
		fmt.Printf("KeysB64: %s\n", response.KeysB64)
		fmt.Printf("RecoveryKeys: %s\n", response.RecoveryKeys)
		fmt.Printf("RecoveryKeysB64: %s\n", response.RecoveryKeysB64)
		fmt.Printf("RootToken: %s\n", response.RootToken)
	}
	return keys, recoveryKeys, rootToken, nil
}

func vaultUnseal(c *vault.Client, keys []string) error {
	for _, key := range keys {
		sealStatus, err := c.Sys().Unseal(key)
		if err != nil {
			return err
		}
		if !sealStatus.Sealed {
			break
		}
	}
	return nil
}

func vaultPutPolicies(c *vault.Client) error {
	for _, assetFile := range policy.AssetNames() {
		byt, err := policy.Asset(assetFile)
		if err != nil {
			return err
		}
		name := assetFile[0 : len(assetFile)-len(filepath.Ext(assetFile))]
		if err := c.Sys().PutPolicy(name, string(byt)); err != nil {
			return err
		}
	}
	return nil
}

func vaultWritePKIRoles(c *vault.Client) error {
	for path, data := range roles {
		if _, err := c.Logical().Write(path, data); err != nil {
			return err
		}
	}
	return nil
}

func vaultMount(c *vault.Client) error {
	for name, data := range mounts {
		if err := c.Sys().Mount(name, data); err != nil {
			return err
		}
	}
	return nil
}

func vaultCreateTokens(c *vault.Client) error {
	for id, policies := range tokens {
		opts := &api.TokenCreateRequest{
			ID:       id,
			Policies: policies,
		}
		if _, err := c.Auth().Token().Create(opts); err != nil {
			return err
		}
	}
	return nil
}

func vaultInitPKI(c *vault.Client, addr, ca, caKey string) error {
	var cert, csr string

	{
		path := fmt.Sprintf("%s/intermediate/generate/internal", PkiPath)
		data := map[string]interface{}{
			"common_name": "PowerSSL Intermediate Authority",
			"ttl":         "8760h",
		}
		secret, err := c.Logical().Write(path, data)
		if err != nil {
			return err
		}
		csr = secret.Data["csr"].(string)
	}

	{
		pem, err := pki.Sign(ca, caKey, csr)
		if err != nil {
			return err
		}
		if err := ioutil.WriteFile("local/certs/intermediate.pem", pem, 0644); err != nil {
			return err
		}
		cert = string(pem)
	}

	{
		path := fmt.Sprintf("%s/intermediate/set-signed", PkiPath)
		data := map[string]interface{}{
			"certificate": cert,
		}
		_, err := c.Logical().Write(path, data)
		if err != nil {
			return err
		}
	}

	{
		path := fmt.Sprintf("%s/config/urls", PkiPath)
		data := map[string]interface{}{
			"crl_distribution_points": fmt.Sprintf("%s/v1/pki/crl", addr),
			"issuing_certificates":    fmt.Sprintf("%s/v1/pki/ca", addr),
		}
		_, err := c.Logical().Write(path, data)
		if err != nil {
			return err
		}
	}

	return nil
}
