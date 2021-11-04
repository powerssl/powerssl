package runner // import "powerssl.dev/common/runner"

import (
	"context"

	"golang.org/x/sync/errgroup"

	"powerssl.dev/common/interrupthandler"
)

func Run(f func(ctx context.Context) ([]func() error, func(), error)) error {
	ctx := context.Background()
	runner, ctx := New(ctx)
	fn, cleanup, err := f(ctx)
	if err != nil {
		return err
	}
	defer cleanup()
	x, fn := fn[0], fn[:1]
	return runner.Run(x, fn...)
}

type Runner struct {
	*errgroup.Group
}

func New(ctx context.Context) (*Runner, context.Context) {
	g, ctx := errgroup.WithContext(ctx)
	return &Runner{
		Group: g,
	}, ctx
}

func (r *Runner) Run(f func() error, fn ...func() error) error {
	r.Go(f)
	for _, fi := range fn {
		r.Go(fi)
	}
	if err := r.Wait(); err != nil {
		switch err.(type) {
		case interrupthandler.InterruptError:
		default:
			return err
		}
	}
	return nil
}
