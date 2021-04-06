package tracing // import "powerssl.dev/common/tracing"

import (
	"bytes"
	"io"
	"io/ioutil"

	"github.com/opentracing/opentracing-go"

	"powerssl.dev/common/log"
)

func NewNoopTracer(_ string, _ log.Logger) (opentracing.Tracer, io.Closer, error) {
	return opentracing.NoopTracer{}, ioutil.NopCloser(bytes.NewBuffer([]byte{})), nil
}
