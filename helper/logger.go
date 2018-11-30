package helper

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// production environments.
func productionEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

// productionConfig is a reasonable production logging configuration.
// Logging is enabled at InfoLevel and above.
func productionConfig() zap.Config {
	return zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "json",
		EncoderConfig:    productionEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

// Logger ...
var Logger *zap.SugaredLogger

func init() {
	pConfig := productionConfig()
	loggerr, err := pConfig.Build()

	if err != nil {
		panic(err)
	}

	Logger = loggerr.Sugar()

}
