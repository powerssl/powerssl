package log // import "powerssl.dev/common/log"

import (
	"go.uber.org/zap"
	zapadapter "logur.dev/adapter/zap"
)

type CertifyLogger interface {
	Trace(msg string, fields ...map[string]interface{})
	Debug(msg string, fields ...map[string]interface{})
	Info(msg string, fields ...map[string]interface{})
	Warn(msg string, fields ...map[string]interface{})
	Error(msg string, fields ...map[string]interface{})
}

type Logger interface {
	CertifyLogger() CertifyLogger
	DPanic(args ...interface{})
	DPanicf(template string, args ...interface{})
	DPanicw(msg string, keysAndValues ...interface{})
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
	Debugw(msg string, keysAndValues ...interface{})
	Desugar() *zap.Logger
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	Fatal(args ...interface{})
	Fatalf(template string, args ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Named(name string) Logger
	Panic(args ...interface{})
	Panicf(template string, args ...interface{})
	Panicw(msg string, keysAndValues ...interface{})
	Sync() error
	TemporalLogger() TemporalLogger
	Warn(args ...interface{})
	Warnf(template string, args ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
	With(args ...interface{}) Logger
}

type logger struct {
	*zap.SugaredLogger
}

func (l *logger) CertifyLogger() CertifyLogger {
	return zapadapter.New(l.Desugar())
}

func (l *logger) Named(name string) Logger {
	return &logger{SugaredLogger: l.SugaredLogger.Named(name)}
}

func (l *logger) TemporalLogger() TemporalLogger {
	return &temporalLogger{
		Logger: l,
	}
}

func (l *logger) With(args ...interface{}) Logger {
	return &logger{SugaredLogger: l.SugaredLogger.With(args...)}
}

func New(cfg Config) (Logger, error) {
	config := zapConfig(cfg)
	var l *zap.Logger
	var err error
	if l, err = config.Build(cfg.Options...); err != nil {
		return nil, err
	}
	zap.ReplaceGlobals(l)
	return &logger{SugaredLogger: l.Sugar()}, nil
}

func zapConfig(cfg Config) (config zap.Config) {
	if cfg.Production() {
		config = zap.NewProductionConfig()
	} else {
		config = zap.NewDevelopmentConfig()
	}
	config.OutputPaths = []string{"stdout"}
	return config
}
