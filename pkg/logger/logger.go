package logger

import (
	"github.com/sirupsen/logrus"
)

// logger wraps logrus.Logger so that it can logger messages sharing a common set of fields.
type logger struct {
	logger *logrus.Logger
	entry  *logrus.Entry
	fields logrus.Fields
}

// NewLogger creates a logger object with the specified logrus.Logger and the fields that should be added to every message.
func NewLogger(fields logrus.Fields, hooks []logrus.Hook) *logger {
	logrusLogger := logrus.New()

	for _, hook := range hooks {
		logrusLogger.AddHook(hook)
	}

	return &logger{
		logger: logrusLogger,
		fields: fields,
	}
}

func (l *logger) SetField(name, value string) {
	l.fields[name] = value
}

func (l *logger) Debugf(format string, args ...interface{}) {
	l.tagged().Debugf(format, args...)
}

func (l *logger) Infof(format string, args ...interface{}) {
	l.tagged().Infof(format, args...)
}

func (l *logger) Warnf(format string, args ...interface{}) {
	l.tagged().Warnf(format, args...)
}

func (l *logger) Errorf(format string, args ...interface{}) {
	l.tagged().Errorf(format, args...)
}

func (l *logger) Panicf(format string, args ...interface{}) {
	l.tagged().Panicf(format, args...)
}

func (l *logger) Debug(args ...interface{}) {
	l.tagged().Debug(args...)
}

func (l *logger) Info(args ...interface{}) {
	l.tagged().Info(args...)
}

func (l *logger) Warn(args ...interface{}) {
	l.tagged().Warn(args...)
}

func (l *logger) Error(args ...interface{}) {
	l.tagged().Error(args...)
}

func (l *logger) Panic(args ...interface{}) {
	l.tagged().Panic(args...)
}

func (l *logger) WithFields(fields map[string]interface{}) {
	l.entry = l.logger.WithFields(fields)
}

func (l *logger) tagged() (entry *logrus.Entry) {
	if l.entry == nil {
		return l.logger.WithFields(l.fields)
	}

	entry = l.entry.WithFields(l.fields)
	l.entry = nil
	return
}
