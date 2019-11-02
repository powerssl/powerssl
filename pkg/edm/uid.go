package edm

import (
	"encoding/base64"

	"github.com/pborman/uuid"
)

func NewUID() string {
	return base64.URLEncoding.EncodeToString(uuid.NewRandom())[:22]
}
