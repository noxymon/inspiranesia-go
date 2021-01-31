package zap

import "go.uber.org/zap"

type ZapLogger struct{
	logger *zap.SugaredLogger
}

func New() *ZapLogger {
	log, _ := zap.NewProduction()
	logSugar := log.Sugar()

	return &ZapLogger{
		logger: logSugar,
	}
}

func (z ZapLogger) Warn(args ...interface{}) {
	z.logger.Warn(args)
}

func (z ZapLogger) Info(args ...interface{}) {
	z.logger.Info(args)
}

func (z ZapLogger) Error(args ...interface{}) {
	z.logger.Error(args)
}

func (z ZapLogger) Panic(args ...interface{}) {
	z.logger.Panic(args)
}

func (z ZapLogger) Debug(args ...interface{}) {
	z.logger.Debug(args)
}

func (z ZapLogger) Fatal(args ...interface{}) {
	z.logger.Fatal(args)
}

func (z ZapLogger) Fatalf(template string, args ...interface{}) {
	z.logger.Fatalf(template, args)
}

func (z ZapLogger) Sync() error {
	return z.logger.Sync()
}

