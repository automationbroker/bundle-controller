package log

import (
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func NewLog() {
	logger.Level = logrus.DebugLevel
}

func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Printf(format string, args ...interface{}) {
	logger.Printf(format, args...)
}

func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

func Warningf(format string, args ...interface{}) {
	logger.Warningf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	logger.Panicf(format, args...)
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Print(args ...interface{}) {
	logger.Print(args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func Warning(args ...interface{}) {
	logger.Warning(args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

func Panic(args ...interface{}) {
	logger.Panic(args...)
}

func Debugln(args ...interface{}) {
	logger.Debugln(args...)
}

func Infoln(args ...interface{}) {
	logger.Infoln(args...)
}

func Println(args ...interface{}) {
	logger.Println(args...)
}

func Warnln(args ...interface{}) {
	logger.Warnln(args...)
}

func Warningln(args ...interface{}) {
	logger.Warningln(args...)
}

func Errorln(args ...interface{}) {
	logger.Errorln(args...)
}

func Fatalln(args ...interface{}) {
	logger.Fatalln(args...)
}

func Panicln(args ...interface{}) {
	logger.Panicln(args...)
}
