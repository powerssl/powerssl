package db

import "github.com/jinzhu/gorm"

type Certificate struct {
	gorm.Model

	CommonName          string
	EncryptionAlgorithm string
	KeySize             int
	SignatureAlgorithm  string
	AutoRenew           bool
}
