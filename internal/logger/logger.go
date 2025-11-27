package logger

import "go.uber.org/zap"

var Log *zap.SugaredLogger

func Init(appEnv string) {
	var l *zap.Logger
	var err error

	if appEnv == "production" {
		l, err = zap.NewProduction()
	} else {
		l, err = zap.NewDevelopment()
	}
	if err != nil {
		panic(err)
	}
	Log = l.Sugar()
}
