package telemetry

type Config struct {
	Component string       `flag:"-"`
	Meter     MeterConfig  `flag:"meter"`
	Tracer    TracerConfig `flag:"tracer"`
}

type MeterConfig struct {
	Addr     string `flag:"addr;;;metrics addr"`
	Exporter string `flag:"exporter;;prometheus;metrics exporter"`
}

type TracerConfig struct {
	Disabled bool `flag:"disabled;;;disable tracer"`
}
