package server

import (
	"go.temporal.io/server/common/log/tag"

	"powerssl.dev/common/log"
)

type temporalLogger struct {
	log.Logger
}

func (l *temporalLogger) Debug(msg string, tags ...tag.Tag) {
	l.Debugw(msg, keysAndValues(tags...)...)
}

func (l *temporalLogger) Info(msg string, tags ...tag.Tag) {
	l.Infow(msg, keysAndValues(tags...)...)
}

func (l *temporalLogger) Warn(msg string, tags ...tag.Tag) {
	l.Warnw(msg, keysAndValues(tags...)...)
}

func (l *temporalLogger) Error(msg string, tags ...tag.Tag) {
	l.Errorw(msg, keysAndValues(tags...)...)
}

func (l *temporalLogger) Fatal(msg string, tags ...tag.Tag) {
	l.Fatalw(msg, keysAndValues(tags...)...)
}

func keysAndValues(tags ...tag.Tag) []interface{} {
	var keysAndValues []interface{}
	for _, t := range tags {
		keysAndValues = append(keysAndValues, t.Key(), t.Value())
	}
	return keysAndValues
}
