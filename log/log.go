package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var AppLog *zap.Logger
var err error

func init() {
	// TODO:出力先をconfで設定する。dockerだとファイル出力がうまくいかないので調査する
	logConfig := zap.Config{
		OutputPaths: []string{"stdout"},
		Level:       zap.NewAtomicLevelAt(zapcore.InfoLevel),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			CallerKey:    "caller",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	if AppLog, err = logConfig.Build(); err != nil {
		panic(err)
	}
}
