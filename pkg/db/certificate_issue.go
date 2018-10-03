package db

import "github.com/jinzhu/gorm"

type CertificateIssue struct {
	gorm.Model

	CommonName          string
	EncryptionAlgorithm string
	KeySize             int
	SignatureAlgorithm  string
}
