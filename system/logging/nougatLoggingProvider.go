package logging

import (
	"inspiranesia/system/logging/zap"
)

type NougatLoggingProvider interface {
	Warn(args ...interface{})
	Info(args ...interface{})
	Error(args ...interface{})
	Panic(args ...interface{})
	Debug(args ...interface{})
	Fatal(args ...interface{})
	Fatalf(template string, args ...interface{})
	Sync() error
}

func ProvideLogging() NougatLoggingProvider {
	return zap.New()
}
