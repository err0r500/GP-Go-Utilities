package logger

import (
	log "github.com/sirupsen/logrus"
)

func SetLevel(level log.Level) {
	log.SetLevel(level)
}

func Trace(args ...interface{}) {
	log.Trace(args)
}

func Debug(args ...interface{}) {
	log.Debug(args...)
}

func Info(args ...interface{}) {
	log.Info(args)
}

func Print(args ...interface{}) {
	log.Print(args)
}

func Warn(args ...interface{}) {
	log.Warn(args)
}

func Error(args ...interface{}) {
	log.Error(args)
}

func Fatal(args ...interface{}) {
	log.Fatal(args)
}

func Panic(args ...interface{}) {
	log.Panic(args)
}
