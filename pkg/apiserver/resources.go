package apiserver

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"github.com/jinzhu/gorm"

	resource "powerssl.io/pkg/resource"
	"powerssl.io/pkg/resource/generated/certificate"
	"powerssl.io/pkg/resource/generated/certificateauthority"
	"powerssl.io/pkg/resource/generated/certificateissue"
)

func makeResources(db *gorm.DB, logger log.Logger, duration metrics.Histogram) []resource.Resource {
	return []resource.Resource{
		certificate.New(db, logger, duration),
		certificateauthority.New(db, logger, duration),
		certificateissue.New(db, logger, duration),
	}
}
