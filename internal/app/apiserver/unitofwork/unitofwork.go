package unitofwork

import (
	"context"

	"github.com/freerware/work/v4/unit"
)

var unityOfWork = struct{}{}

func GetUnit(ctx context.Context) unit.Unit {
	return ctx.Value(unityOfWork).(unit.Unit)
}

func SetUnit(ctx context.Context, unit unit.Unit) (context.Context, error) {
	return context.WithValue(ctx, unityOfWork, unit), nil
}
