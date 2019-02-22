package engine

import (
	"context"
)

type Engine struct{}

func New() *Engine {
	return &Engine{}
}

func (e *Engine) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
