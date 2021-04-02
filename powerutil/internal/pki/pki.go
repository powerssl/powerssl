package pki

import cfssllog "github.com/cloudflare/cfssl/log"

func init() {
	cfssllog.Level = cfssllog.LevelError
}
