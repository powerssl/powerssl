package db

import "github.com/jinzhu/gorm"

type CertificateAuthority struct {
	gorm.Model

	Vendor string
}
