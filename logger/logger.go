package logger

import (
	"context"

	"github.com/newrelic/go-agent/v3/integrations/logcontext/nrlogrusplugin"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
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

func TraceCf(c context.Context, format string, args ...interface{}) {
	if txn := nrgin.Transaction(c); txn != nil {
		ctx := newrelic.NewContext(context.Background(), txn)
		log.WithContext(ctx).Tracef(format, args...)
	} else {
		log.Tracef(format, args...)
	}
}

func DebugCf(c context.Context, format string, args ...interface{}) {
	if txn := nrgin.Transaction(c); txn != nil {
		ctx := newrelic.NewContext(context.Background(), txn)
		log.WithContext(ctx).Debugf(format, args...)
	} else {
		log.Debugf(format, args...)
	}
}

func PrintCf(c context.Context, format string, args ...interface{}) {
	if txn := nrgin.Transaction(c); txn != nil {
		ctx := newrelic.NewContext(context.Background(), txn)
		log.WithContext(ctx).Printf(format, args...)
	} else {
		log.Printf(format, args...)
	}
}

func InfoCf(c context.Context, format string, args ...interface{}) {
	if txn := nrgin.Transaction(c); txn != nil {
		ctx := newrelic.NewContext(context.Background(), txn)
		log.WithContext(ctx).Infof(format, args...)
	} else {
		log.Infof(format, args...)
	}
}

func WarnCf(c context.Context, format string, args ...interface{}) {
	if txn := nrgin.Transaction(c); txn != nil {
		ctx := newrelic.NewContext(context.Background(), txn)
		log.WithContext(ctx).Warnf(format, args...)
	} else {
		log.Warnf(format, args...)
	}
}

func WarningCf(c context.Context, format string, args ...interface{}) {
	if txn := nrgin.Transaction(c); txn != nil {
		ctx := newrelic.NewContext(context.Background(), txn)
		log.WithContext(ctx).Warningf(format, args...)
	} else {
		log.Warningf(format, args...)
	}
}

func ErrorCf(c context.Context, format string, args ...interface{}) {
	if txn := nrgin.Transaction(c); txn != nil {
		ctx := newrelic.NewContext(context.Background(), txn)
		log.WithContext(ctx).Errorf(format, args...)
	} else {
		log.Errorf(format, args...)
	}
}

func PanicCf(c context.Context, format string, args ...interface{}) {
	if txn := nrgin.Transaction(c); txn != nil {
		ctx := newrelic.NewContext(context.Background(), txn)
		log.WithContext(ctx).Panicf(format, args...)
	} else {
		log.Panicf(format, args...)
	}
}

func FatalCf(c context.Context, format string, args ...interface{}) {
	if txn := nrgin.Transaction(c); txn != nil {
		ctx := newrelic.NewContext(context.Background(), txn)
		log.WithContext(ctx).Fatalf(format, args...)
	} else {
		log.Fatalf(format, args...)
	}
}
