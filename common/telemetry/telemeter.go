package telemetry

import (
	"context"

	"powerssl.dev/common/log"
)

type Telemeter struct {
	cfg    Config
	log    log.Logger
	Meter  *Meter
	Tracer *Tracer
}

func New(cfg Config, logger log.Logger) (*Telemeter, error) {
	tracer := NewTracer(cfg.Tracer, logger)
	meter, err := NewMeter(cfg.Meter, logger)
	if err != nil {
		return nil, err
	}
	return &Telemeter{
		cfg:    cfg,
		log:    logger,
		Tracer: tracer,
		Meter:  meter,
	}, nil
}

func (t *Telemeter) Cleanup(_ context.Context) error {
	return nil
}

func (t *Telemeter) F(ctx context.Context) F {
	return func() error {
		return t.Meter.Serve(ctx)
	}
}
