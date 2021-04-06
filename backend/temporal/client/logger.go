package client // import "powerssl.dev/backend/temporal/client"

import "powerssl.dev/common/log"

type temporalLogger struct {
	log.Logger
}

func (l temporalLogger) Debug(msg string, keysAndValues ...interface{}) {
	l.Debugw(msg, keysAndValues...)
}

func (l temporalLogger) Info(msg string, keysAndValues ...interface{}) {
	l.Infow(msg, keysAndValues...)
}

func (l temporalLogger) Warn(msg string, keysAndValues ...interface{}) {
	l.Warnw(msg, keysAndValues...)
}

func (l temporalLogger) Error(msg string, keysAndValues ...interface{}) {
	l.Errorw(msg, keysAndValues...)
}
