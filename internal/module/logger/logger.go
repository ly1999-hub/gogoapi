package logger

import (
	"go.elastic.co/apm/module/apmzap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	LogData map[string]interface{}
)

var (
	zapLogger *zap.Logger
	err       error
	appName   = "gogoAPI"
)

func Init(server string) {
	cfg := zap.Config{
		Encoding:      "json",
		Level:         zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths:   []string{"stdout"},
		InitialFields: map[string]interface{}{"server": server, "capture": appName},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "message",
			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,
			TimeKey:     "time",
			EncodeTime:  zapcore.ISO8601TimeEncoder,
		},
	}
	zapLogger, err := cfg.Build(zap.WrapCore((&apmzap.Core{}).WrapCore))
	if err != nil {
		panic(err)
	}
	println(zapLogger)
}
