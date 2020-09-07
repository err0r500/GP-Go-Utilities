package logger

import (
	"github.com/newrelic/go-agent/v3/integrations/logcontext/nrlogrusplugin"
	"github.com/newrelic/go-agent/v3/integrations/nrlogrus"
	"github.com/newrelic/go-agent/v3/newrelic"
	log "github.com/sirupsen/logrus"
)

func StandardLogger() newrelic.Logger {
	return nrlogrus.StandardLogger()
}

func SetLevel(level log.Level) {
	log.SetLevel(level)
}

func SetFormatter(formatter log.Formatter) {
	log.SetFormatter(formatter)
}

func SetNewRelicFormatter() {
	log.SetFormatter(nrlogrusplugin.ContextFormatter{})
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

func Tracef(format string, args ...interface{}) {
	log.Tracef(format, args...)
}

func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

func Printf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

func Warningf(format string, args ...interface{}) {
	log.Warningf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	log.Panicf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}
