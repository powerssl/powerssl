package tracer // import "powerssl.dev/common/tracer"

type Config struct {
	Component      string `flag:"-"`
	Implementation string `flag:"implementation;;jaeger;tracer implementation"`
}
