package powerutil

import (
	"fmt"
	"path/filepath"

	"github.com/hashicorp/vault/api"

	"powerssl.io/internal/app/powerutil/policy"
	"powerssl.io/internal/pkg/vault"
)

const (
	PKI_PATH     = "pki"
	TRANSIT_PATH = "transit"
)

var mounts = map[string]*api.MountInput{
	PKI_PATH: {
		Type:        "pki",
		Description: "PowerSSL PKI",
		Config: api.MountConfigInput{
			MaxLeaseTTL: "8760h",
		},
	},
	TRANSIT_PATH: {
		Type:        "transit",
		Description: "PowerSSL Transit",
	},
}

var roles = map[string]map[string]interface{}{
	fmt.Sprintf("%s/roles/powerssl-apiserver", PKI_PATH): map[string]interface{}{
		"allowed_domains":    "apiserver",
		"allow_localhost":    false,
		"allow_bare_domains": true,
		"max_ttl":            "24h",
	},
	fmt.Sprintf("%s/roles/powerssl-controller", PKI_PATH): map[string]interface{}{
		"allowed_domains":    "controller",
		"allow_localhost":    false,
		"allow_bare_domains": true,
		"max_ttl":            "24h",
	},
	fmt.Sprintf("%s/roles/powerssl-signer", PKI_PATH): map[string]interface{}{
		"allowed_domains":    "signer",
		"allow_localhost":    false,
		"allow_bare_domains": true,
		"max_ttl":            "24h",
	},
}

func RunVault(addr, caFile string) error {
	var keys []string
	var rootToken string
	{
		c, err := vault.New(addr, "", caFile)
		if err != nil {
			return err
		}

		keys, _, rootToken, err = vaultInit(c)
		if err != nil {
			return err
		}
	}

	c, err := vault.New(addr, rootToken, caFile)
	if err != nil {
		return err
	}

	if err := vaultUnseal(c, keys); err != nil {
		return err
	}

	if err := vaultMount(c); err != nil {
		return err
	}

	if err := vaultWritePKIRoles(c); err != nil {
		return err
	}

	if err := vaultPutPolicies(c); err != nil {
		return err
	}

	return nil
}

func vaultInit(c *vault.Client) ([]string, []string, string, error) {
	var SecretShares int
	SecretShares = 1
	var SecretThreshold int
	SecretThreshold = 1
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
		if sealStatus.Sealed == false {
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
