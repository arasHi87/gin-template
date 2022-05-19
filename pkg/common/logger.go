package common

import (
	"os"
	"time"

	"github.com/arashi87/gin-template/pkg/setting"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func InitLogger() {
	Logger = logrus.New()

	// basic env
	logPath := setting.CONFIG.LogPath
	logExpire := setting.CONFIG.LogExpire
	logRotate := setting.CONFIG.LogRotate

	// logger basic setting
	Logger.Out = os.Stdout
	Logger.SetLevel(logrus.InfoLevel)
	Logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// setting rotatelogs
	writer, _ := rotatelogs.New(
		logPath+"%Y%m%d.log",
		rotatelogs.WithLinkName(logPath),
		rotatelogs.WithMaxAge(time.Duration(logExpire)*time.Hour),
		rotatelogs.WithRotationTime(time.Duration(logRotate)*time.Hour),
	)

	// lfshook write map
	writeMap := lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}
	Logger.AddHook(lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}))
}
