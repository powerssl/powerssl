package certificate

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gogo/status"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc/codes"

	"powerssl.io/pkg/apiserver/api"
)

type Certificate struct {
	gorm.Model

	DisplayName     string
	Title           string
	Description     string
	DNSNames        string
	KeyAlgorithm    string
	KeySize         int32
	DigestAlgorithm string
	AutoRenew       bool
}

func (c *Certificate) ToAPI() *api.Certificate {
	return &api.Certificate{
		Name: fmt.Sprintf("certificates/%d", c.ID),

		CreateTime:  c.CreatedAt,
		UpdateTime:  c.UpdatedAt,
		DisplayName: c.DNSNames,
		Title:       c.DNSNames,
		Description: c.Description,
		Labels:      map[string]string{"not": "implemented"},

		Dnsnames:        strings.Split(c.DNSNames, ","),
		KeyAlgorithm:    c.KeyAlgorithm,
		KeySize:         c.KeySize,
		DigestAlgorithm: c.DigestAlgorithm,
		AutoRenew:       c.AutoRenew,
	}
}

type Certificates []*Certificate

func (c Certificates) ToAPI() []*api.Certificate {
	certs := make([]*api.Certificate, len(c))
	for i, cert := range c {
		certs[i] = cert.ToAPI()
	}
	return certs
}

func FindCertificateByName(name string, db *gorm.DB) (*Certificate, error) {
	s := strings.Split(name, "/")
	if len(s) != 2 {
		return nil, status.Error(codes.InvalidArgument, "malformed name")
	}
	id, err := strconv.Atoi(s[1])
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "malformed name")
	}

	certificate := &Certificate{}
	if db.Where("id = ?", id).First(&certificate).RecordNotFound() {
		return nil, status.Error(codes.NotFound, "not found")
	}
	return certificate, nil
}

func NewCertificateFromAPI(certificate *api.Certificate) *Certificate {
	return &Certificate{
		DNSNames:        strings.Join(certificate.Dnsnames, ","),
		KeyAlgorithm:    certificate.KeyAlgorithm,
		KeySize:         certificate.KeySize,
		DigestAlgorithm: certificate.DigestAlgorithm,
		AutoRenew:       certificate.AutoRenew,
	}
}
