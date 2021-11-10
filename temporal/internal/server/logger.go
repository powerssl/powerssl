package server

import (
	"go.temporal.io/server/common/log/tag"
	"go.uber.org/zap"
)

type logger struct {
	*zap.SugaredLogger
}

func newLogger(sugaredLogger *zap.SugaredLogger) *logger {
	return &logger{
		SugaredLogger: sugaredLogger,
	}
}

func (l logger) Debug(msg string, tags ...tag.Tag) {
	l.Debugw(msg, keysAndValues(tags...)...)
}

func (l logger) Info(msg string, tags ...tag.Tag) {
	l.Infow(msg, keysAndValues(tags...)...)
}

func (l logger) Warn(msg string, tags ...tag.Tag) {
	l.Warnw(msg, keysAndValues(tags...)...)
}

func (l logger) Error(msg string, tags ...tag.Tag) {
	l.Errorw(msg, keysAndValues(tags...)...)
}

func (l logger) Fatal(msg string, tags ...tag.Tag) {
	l.Fatalw(msg, keysAndValues(tags...)...)
}

func keysAndValues(tags ...tag.Tag) []interface{} {
	var keysAndValues []interface{}
	for _, t := range tags {
		keysAndValues = append(keysAndValues, t.Key(), t.Value())
	}
	return keysAndValues
}
