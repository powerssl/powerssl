package transport // import "powerssl.dev/common/transport"

import (
	"github.com/google/wire"
	"google.golang.org/grpc"
)

var Provider = wire.NewSet(
	New,
)

func NoDialOptions() []grpc.DialOption {
	return nil
}
