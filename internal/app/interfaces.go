package app

//go:generate mockgen -source interfaces.go -package mocks -destination mocks/interfaces_mock.go

type UserAccess interface {
}

type Logger interface {
	SetField(name, value string)
	WithFields(fields map[string]any)
	Debugf(format string, args ...any)
	Infof(format string, args ...any)
	Warnf(format string, args ...any)
	Errorf(format string, args ...any)
	Panicf(format string, args ...any)
	Debug(args ...any)
	Info(args ...any)
	Warn(args ...any)
	Error(args ...any)
	Panic(args ...any)
}
