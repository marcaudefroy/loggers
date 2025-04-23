package loggers

// Standard is the interface used by Go's standard library's log package.
type Standard interface {
	GetUnderlying() any
	Fatal(args ...any)
	Fatalf(format string, args ...any)
	Fatalln(args ...any)

	Panic(args ...any)
	Panicf(format string, args ...any)
	Panicln(args ...any)

	Print(args ...any)
	Printf(format string, args ...any)
	Println(args ...any)
}

// Advanced is an interface with commonly used log level methods.
type Advanced interface {
	Standard

	Debug(args ...any)
	Debugf(format string, args ...any)
	Debugln(args ...any)

	Error(args ...any)
	Errorf(format string, args ...any)
	Errorln(args ...any)

	Info(args ...any)
	Infof(format string, args ...any)
	Infoln(args ...any)

	Warn(args ...any)
	Warnf(format string, args ...any)
	Warnln(args ...any)
}

// Contextual is an interface that allows context addition to a log statement before
// calling the final print (message/level) method.
type Contextual interface {
	Advanced

	WithField(key string, value any) Contextual
	WithFields(fields ...any) Contextual
}
