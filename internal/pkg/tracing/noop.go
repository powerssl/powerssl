package tracing

import (
	"bytes"
	"io"
	"io/ioutil"

	"github.com/go-kit/kit/log"
	"github.com/opentracing/opentracing-go"
)

func NewNoopTracer(serviceName string, logger log.Logger) (opentracing.Tracer, io.Closer, error) {
	return opentracing.NoopTracer{}, ioutil.NopCloser(bytes.NewBuffer([]byte{})), nil
}
