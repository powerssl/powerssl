package acme

import (
	"context"
	"fmt"

	"github.com/eggsampler/acme"
	"github.com/opentracing/opentracing-go"
)

func NewClient(ctx context.Context, directoryURL string) (*acme.Client, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "NewClient")
	defer span.Finish()
	client, err := acme.NewClient(directoryURL)
	if err != nil {
		return nil, fmt.Errorf("Error connecting to acme directory: %v", err)
	}
	return &client, nil
}
