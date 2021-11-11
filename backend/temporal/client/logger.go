package client

import "go.uber.org/zap"

type logger struct {
	*zap.SugaredLogger
}

func newLogger(sugaredLogger *zap.SugaredLogger) *logger {
	return &logger{
		SugaredLogger: sugaredLogger,
	}
}

func (l logger) Debug(msg string, keysAndValues ...interface{}) {
	l.Debugw(msg, keysAndValues...)
}

func (l logger) Info(msg string, keysAndValues ...interface{}) {
	l.Infow(msg, keysAndValues...)
}

func (l logger) Warn(msg string, keysAndValues ...interface{}) {
	l.Warnw(msg, keysAndValues...)
}

func (l logger) Error(msg string, keysAndValues ...interface{}) {
	l.Errorw(msg, keysAndValues...)
}
