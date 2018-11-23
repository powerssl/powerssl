package uid

import (
	"encoding/base64"

	"github.com/pborman/uuid"
)

func New() string {
	return base64.URLEncoding.EncodeToString(uuid.NewRandom())[:22]
}
