package domain

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Domain struct {
	gorm.Model

	DNSName string
	Name    string `gorm:"-"`
}

func (d Domain) ExternalName() string {
	return fmt.Sprintf("domains/%s", d.DNSName)
}
