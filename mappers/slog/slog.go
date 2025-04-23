package slog

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/marcaudefroy/loggers"
	"github.com/marcaudefroy/loggers/mappers"
)

type Logger struct {
	logger *slog.Logger
}

func NewLogger(l *slog.Logger) loggers.Contextual {
	nl := &Logger{
		logger: l,
	}
	mp := mappers.NewContextualMap(nl)
	return mp
}

func NewDefaultLogger() loggers.Contextual {
	handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	return NewLogger(slog.New(handler))
}

func (l *Logger) GetUnderlying() any {
	return l.logger
}

func (l *Logger) WithField(key string, value any) loggers.Contextual {
	return l.WithFields(key, value)
}

func (l *Logger) WithFields(fields ...any) loggers.Contextual {
	nl := l.logger.With(fields...)
	nL := &Logger{
		logger: nl,
	}
	mp := mappers.NewContextualMap(nL)
	return mp
}

// LevelPrint is a Mapper method
func (l *Logger) LevelPrint(lev mappers.Level, i ...any) {
	var log func(msg string, args ...any)
	switch lev {
	case mappers.LevelDebug:
		log = l.logger.Debug
	case mappers.LevelInfo:
		log = l.logger.Info
	case mappers.LevelWarn:
		log = l.logger.Warn
	case mappers.LevelError:
		log = l.logger.Error
	default:
		log = l.logger.Info
	}
	msg, args := l.extractMsgAndAttrs(i...)
	log(msg, args...)
}

// LevelPrintf is a Mapper method
func (l *Logger) LevelPrintf(lev mappers.Level, format string, i ...any) {
	l.LevelPrint(lev, fmt.Sprintf(format, i...))
}

// LevelPrintln is a Mapper method
func (l *Logger) LevelPrintln(lev mappers.Level, i ...any) {
	l.LevelPrint(lev, i...)
}

func (l *Logger) extractMsgAndAttrs(args ...any) (string, []any) {
	var msg string

	if len(args) > 0 {
		if m, ok := args[0].(string); ok {
			msg = m
			args = args[1:]
		}
	}
	return msg, args
}
