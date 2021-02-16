package acme

import (
	"context"
	"fmt"

	acme "github.com/eggsampler/acme/v2"
	"github.com/opentracing/opentracing-go"
)

func NewClient(ctx context.Context, directoryURL string) (*acme.Client, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "NewClient")
	defer span.Finish()
	client, err := acme.NewClient(directoryURL)
	if err != nil {
		return nil, fmt.Errorf("error connecting to acme directory: %v", err)
	}
	return &client, nil
}