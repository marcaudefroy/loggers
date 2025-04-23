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
	fields []any
}

func NewLogger(l *slog.Logger) loggers.Contextual {
	nl := &Logger{
		logger: l,
		fields: []any{},
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

func (l *Logger) WithField(key string, value any) loggers.Contextual {
	return l.WithFields(key, value)
}

func (l *Logger) WithFields(fields ...any) loggers.Contextual {
	nl := l.logger.With(fields...)
	newFields := append([]any{}, l.fields...)
	newFields = append(newFields, fields...)

	nL := &Logger{
		logger: nl,
		fields: newFields,
	}
	mp := mappers.NewContextualMap(nL)
	return mp
}

func (l *Logger) Fields() []any {
	return l.fields
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
	log(msg, l.toAttrs(args...)...)
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

func (l *Logger) toAttrs(args ...any) []any {
	var attrs []any
	allArgs := append(l.fields, args...)
	for i := 0; i+1 < len(allArgs); i += 2 {
		key, ok := allArgs[i].(string)
		if !ok {
			continue
		}
		attrs = append(attrs, slog.Any(key, allArgs[i+1]))
	}
	return attrs
}
