package model

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"

	"powerssl.dev/powerssl/internal/pkg/uid"
	"powerssl.dev/powerssl/pkg/apiserver/api"
)

type Certificate struct {
	ID        string `pg:",pk"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `pg:",soft_delete"`

	DisplayName     string
	Title           string
	Description     string
	DNSNames        string
	KeyAlgorithm    string
	KeySize         int32
	DigestAlgorithm string
	AutoRenew       bool
}

func (c *Certificate) Name() string {
	return fmt.Sprintf("certificates/%s", c.ID)
}

var _ pg.BeforeInsertHook = (*Certificate)(nil)

func (certificate *Certificate) BeforeInsert(ctx context.Context) (context.Context, error) {
	certificate.ID = uid.New()
	return ctx, nil
}

func (c *Certificate) ToAPI() *api.Certificate {
	return &api.Certificate{
		Name: c.Name(),

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

func FindCertificateByName(name string, db *pg.DB) (*Certificate, error) {
	s := strings.Split(name, "/")
	if len(s) != 2 {
		return nil, status.Error(codes.InvalidArgument, "malformed name")
	}

	certificate := &Certificate{}
	if err := db.Model(certificate).Where("id = ?", s[1]).Limit(1).Select(); err != nil {
		if err == pg.ErrNoRows {
			return nil, status.Error(codes.NotFound, "not found")
		}
		return nil, err
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
