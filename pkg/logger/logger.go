package logger

import "go.uber.org/zap"

// NewDevZapLogger provides zap development logger
func NewDevZapLogger(level string) (*zap.Logger, error) {
	configZap := zap.NewDevelopmentConfig()
	configZap.OutputPaths = []string{"stdout", "../medhublogs"}
	configZap.Development = true
	configZap.DisableStacktrace = true
	configZap.Encoding = "console"
	configZap.ErrorOutputPaths = []string{"stderr"}
	configZap.EncoderConfig = zap.NewDevelopmentEncoderConfig()

	switch level {
	case "debug":
		configZap.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		configZap.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		configZap.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		configZap.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	case "dpanic":
		configZap.Level = zap.NewAtomicLevelAt(zap.DPanicLevel)
	case "panic":
		configZap.Level = zap.NewAtomicLevelAt(zap.PanicLevel)
	case "fatal":
		configZap.Level = zap.NewAtomicLevelAt(zap.FatalLevel)
	default:
		configZap.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}

	return configZap.Build()
}
