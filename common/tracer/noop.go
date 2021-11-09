package tracer // import "powerssl.dev/common/tracer"

import (
	"bytes"
	"io"
	"io/ioutil"

	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

func NewNoopTracer(_ Config, _ *zap.SugaredLogger) (opentracing.Tracer, io.Closer, error) {
	return opentracing.NoopTracer{}, ioutil.NopCloser(bytes.NewBuffer([]byte{})), nil
}
