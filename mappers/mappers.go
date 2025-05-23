package mappers

import "github.com/marcaudefroy/loggers"

type (
	// Level indicates a specific log level.
	Level byte

	// LevelMapper interfaces allows a logger to map to any Advanced Logger.
	LevelMapper interface {
		LevelPrint(Level, ...any)
		LevelPrintf(Level, string, ...any)
		LevelPrintln(Level, ...any)
	}

	// ContextualMapper interfaces allows a logger to map to any Contextual Logger.
	ContextualMapper interface {
		LevelMapper
		WithField(key string, value any) loggers.Contextual
		WithFields(fields ...any) loggers.Contextual
	}
)

const (
	// LevelDebug is a log Level.
	LevelDebug Level = iota
	// LevelInfo is a log Level.
	LevelInfo
	// LevelWarn is a log Level.
	LevelWarn
	// LevelError is a log Level.
	LevelError
	// LevelFatal is a log Level.
	LevelFatal
	// LevelPanic is a log Level.
	LevelPanic
)

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG "
	case LevelInfo:
		return "INFO  "
	case LevelWarn:
		return "WARN  "
	case LevelError:
		return "ERROR "
	case LevelFatal:
		return "FATAL "
	case LevelPanic:
		return "PANIC "
	default:
		panic("Missing case statement in Level String.")
	}
}
