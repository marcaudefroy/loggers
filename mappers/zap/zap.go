package zap

import (
	"fmt"

	"github.com/marcaudefroy/loggers"
	"go.uber.org/zap"
)

// Logger is an Contextual logger wrapper over Zap's logger.
type Logger struct {
	zap    *zap.SugaredLogger
	fields []zap.Field
}

// New returns a Contextual Logger for Zap's logger.
func New(sugar *zap.SugaredLogger) loggers.Contextual {
	var l Logger
	l.zap = sugar
	l.Info("Now using Zap logger package (via loggers/mappers/zap).")
	// TODO: Handle kill pour sync
	return &l
}

func NewDefaultLogger() loggers.Contextual {
	var l Logger
	logger, _ := zap.NewDevelopment()
	l.zap = logger.Sugar()
	l.Info("Now using Zap logger package (via loggers/mappers/zap).")
	return &l
}

func (l Logger) Fatal(args ...interface{}) {
	l.zap.Fatal(args...)
}

func (l Logger) Fatalf(format string, args ...interface{}) {
	l.zap.Fatalf(format, args...)
}

func (l Logger) Fatalln(args ...interface{}) {
	l.zap.Fatal(sprintlnn(args...))
}

func (l Logger) Panic(args ...interface{}) {
	l.zap.Panic(args...)
}

func (l Logger) Panicf(format string, args ...interface{}) {
	l.zap.Panicf(format, args...)
}

func (l Logger) Panicln(args ...interface{}) {
	l.zap.Panic(sprintlnn(args...))
}

func (l Logger) Print(args ...interface{}) {
	// Zap does not support print without level.
	// We use info by default
	l.Info(args...)
}

func (l Logger) Printf(format string, args ...interface{}) {
	// Zap does not support print without level.
	// We use info by default
	l.Infof(format, args...)
}

func (l Logger) Println(args ...interface{}) {
	// Zap does not support print without level.
	// We use info by default
	l.Info(args...)
}

func (l Logger) Debug(args ...interface{}) {
	l.zap.Debug(args...)
}

func (l Logger) Debugf(format string, args ...interface{}) {
	l.zap.Debugf(format, args...)
}

func (l Logger) Debugln(args ...interface{}) {
	l.zap.Debug(sprintlnn(args...))
}

func (l Logger) Error(args ...interface{}) {
	l.zap.Error(args...)
}

func (l Logger) Errorf(format string, args ...interface{}) {
	l.zap.Errorf(format, args...)
}

func (l Logger) Errorln(args ...interface{}) {
	l.zap.Error(sprintlnn(args...))
}

func (l Logger) Info(args ...interface{}) {
	l.zap.Info(args...)
}

func (l Logger) Infof(format string, args ...interface{}) {
	l.zap.Infof(format, args...)
}

func (l Logger) Infoln(args ...interface{}) {
	l.zap.Info(sprintlnn(args...))
}

func (l Logger) Warn(args ...interface{}) {
	l.zap.Warn(args...)
}

func (l Logger) Warnf(format string, args ...interface{}) {
	l.zap.Warnf(format, args...)
}

func (l Logger) Warnln(args ...interface{}) {
	l.zap.Warn(sprintlnn(args...))
}

func (l Logger) WithField(key string, value interface{}) loggers.Contextual {
	var newLogger Logger
	field := zap.Any(key, value)
	newLogger.fields = append(l.fields, field)
	newLogger.zap = l.zap.With(field)
	return &newLogger
}

func (l Logger) WithFields(fields ...interface{}) loggers.Contextual {
	var newLogger Logger
	newFields := sliceToZapFields(fields...)
	newLogger.fields = append(l.fields, newFields...)
	newLogger.zap = l.zap.With(fields...)
	return &newLogger
}

func (l Logger) Fields() []interface{} {
	return fieldsToSlice(l.fields)
}

func sprintlnn(args ...interface{}) string {
	msg := fmt.Sprintln(args...)
	return msg[:len(msg)-1]
}

func fieldsToSlice(fields []zap.Field) []interface{} {
	res := make([]interface{}, 0, len(fields)*2)
	for _, field := range fields {
		res = append(res, field.Key, field.Interface)
	}
	return res
}

func sliceToZapFields(fields ...interface{}) []zap.Field {
	f := make([]zap.Field, 0, len(fields)/2)
	var key, value interface{}
	for i := 0; i+1 < len(fields); i = i + 2 {
		key = fields[i]
		value = fields[i+1]
		if s, ok := key.(string); ok {
			f = append(f, zap.Any(s, value))
		} else if s, ok := key.(fmt.Stringer); ok {
			f = append(f, zap.Any(s.String(), value))
		}
	}
	return f
}
