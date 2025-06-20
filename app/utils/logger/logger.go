package logger

import (
	"os"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger
func InitLogger(env string) {
	Log = logrus.New()

	if env == "production" {
		Log.SetFormatter(&logrus.JSONFormatter{})
		Log.SetLevel(logrus.WarnLevel)
		Log.SetOutput(&lumberjack.Logger{
			Filename:   "./logs/app.log",
			MaxSize:    100,
			MaxBackups: 7,
			MaxAge:     30,
			Compress:   true,
		})
	} else {
		Log.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
		Log.SetLevel(logrus.DebugLevel)
		Log.SetOutput(os.Stdout)
	}
}
