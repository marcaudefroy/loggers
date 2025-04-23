package log

import (
	"github.com/marcaudefroy/loggers"
	"github.com/marcaudefroy/loggers/mappers/stdlib"
)

// Logger is an Contextual logger interface.
var Logger loggers.Contextual

func init() {
	Logger = stdlib.NewDefaultLogger()
}

// Debug should be used when logging exessive debug info.
func Debug(v ...any) {
	Logger.Debug(v...)
}

// Debugf works the same as Debug but supports formatting.
func Debugf(format string, v ...any) {
	Logger.Debugf(format, v...)
}

// Debugln works the same as Debug but supports formatting.
func Debugln(v ...any) {
	Logger.Debugln(v...)
}

// Info is a general function to log something.
func Info(v ...any) {
	Logger.Info(v...)
}

// Infof works the same as Info but supports formatting.
func Infof(format string, v ...any) {
	Logger.Infof(format, v...)
}

// Infoln works the same as Info but supports formatting.
func Infoln(v ...any) {
	Logger.Infoln(v...)
}

// Warn is useful for alerting about something wrong.
func Warn(v ...any) {
	Logger.Warn(v...)
}

// Warnf works the same as Warn but supports formatting.
func Warnf(format string, v ...any) {
	Logger.Warnf(format, v...)
}

// Warnln works the same as Warn but prints each value on a line.
func Warnln(v ...any) {
	Logger.Warnln(v...)
}

// Error should be used only if real error occures.
func Error(v ...any) {
	Logger.Error(v...)
}

// Errorf works the same as Error but supports formatting.
func Errorf(format string, v ...any) {
	Logger.Errorf(format, v...)
}

// Errorln works the same as Error but prints each value on a line.
func Errorln(v ...any) {
	Logger.Errorln(v...)
}

// Fatal should be only used when it's not possible to continue program execution.
func Fatal(v ...any) {
	Logger.Fatal(v...)
}

// Fatalf works the same as Fatal but supports formatting.
func Fatalf(format string, v ...any) {
	Logger.Fatalf(format, v...)
}

// Fatalln works the same as Fatal but prints each value on a line.
func Fatalln(v ...any) {
	Logger.Fatalln(v...)
}

// Panic should be used only if real panic is desired.
func Panic(v ...any) {
	Logger.Panic(v...)
}

// Panicf works the same as Panic but supports formatting.
func Panicf(format string, v ...any) {
	Logger.Panicf(format, v...)
}

// Panicln works the same as Panic but prints each value on a line.
func Panicln(v ...any) {
	Logger.Panicln(v...)
}

// Print should be used for information messages.
func Print(v ...any) {
	Logger.Print(v...)
}

// Printf works the same as Print but supports formatting.
func Printf(format string, v ...any) {
	Logger.Printf(format, v...)
}

// Println works the same as Print but prints each value on a line.
func Println(v ...any) {
	Logger.Println(v...)
}

// WithField adds the key value as parameter to log.
func WithField(key string, value any) loggers.Contextual {
	return Logger.WithField(key, value)
}

// WithFields adds the fields as a list of key/value parameters to log. Even number expected.
func WithFields(fields ...any) loggers.Contextual {
	return Logger.WithFields(fields...)
}

func GetUnderlying[T any]() T {
	return Logger.GetUnderlying().(T)
}
