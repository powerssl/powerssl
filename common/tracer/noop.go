package tracer // import "powerssl.dev/common/tracer"

import (
	"bytes"
	"io"
	"io/ioutil"

	"github.com/opentracing/opentracing-go"

	"powerssl.dev/common/log"
)

func NewNoopTracer(_ Config, _ log.Logger) (opentracing.Tracer, io.Closer, error) {
	return opentracing.NoopTracer{}, ioutil.NopCloser(bytes.NewBuffer([]byte{})), nil
}
