package stdlib

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/marcaudefroy/loggers"
	"github.com/marcaudefroy/loggers/mappers"
)

// goLog maps the standard log package logger to an Contextual log interface.
// However it mostly ignores any level info.
type goLog struct {
	logger *log.Logger
	fields []any
}

// NewDefaultLogger returns a Contextual logger using a log.Logger with stderr output.
func NewDefaultLogger() loggers.Contextual {
	var g goLog
	g.logger = log.New(os.Stderr, "", log.Ldate|log.Ltime)
	g.fields = []any{}

	a := mappers.NewContextualMap(&g)

	return a
}

// NewLogger creates a Contextual logger from a log.Logger.
func NewLogger(l *log.Logger) loggers.Contextual {
	var g goLog
	g.logger = l
	g.fields = []any{}
	a := mappers.NewContextualMap(&g)

	return a
}

func (l *goLog) GetUnderlying() any {
	return l.logger
}

// LevelPrint is a Mapper method
func (l *goLog) LevelPrint(lev mappers.Level, i ...any) {
	v := []any{lev}
	v = append(v, i...)
	l.logger.Print(v...)
}

// LevelPrintf is a Mapper method
func (l *goLog) LevelPrintf(lev mappers.Level, format string, i ...any) {
	f := "%s" + format
	v := []any{lev}
	v = append(v, i...)
	l.logger.Printf(f, v...)
}

// LevelPrintln is a Mapper method
func (l *goLog) LevelPrintln(lev mappers.Level, i ...any) {
	v := []any{lev}
	v = append(v, i...)
	l.logger.Println(v...)
}

// WithField returns an Contextual logger with a pre-set field.
func (l *goLog) WithField(key string, value any) loggers.Contextual {
	return l.WithFields(key, value)
}

// WithFields returns an Contextual logger with pre-set fields.
func (l *goLog) WithFields(fields ...any) loggers.Contextual {
	if l == nil {
		return nil
	}
	newL := *l
	if newL.fields == nil {
		newL.fields = []any{}
	}
	newL.fields = append(newL.fields, fields...)

	r := gologPostfixLogger{&newL}
	return mappers.NewContextualMap(&r)
}

type gologPostfixLogger struct {
	*goLog
}

func (r *gologPostfixLogger) GetUnderlying() any {
	return r.logger
}

func (r *gologPostfixLogger) postfixFromFields() string {
	if len(r.fields) > 1 {
		s := make([]string, 0, len(r.fields)/2)
		for i := 0; i+1 < len(r.fields); i = i + 2 {
			key := r.fields[i]
			value := r.fields[i+1]
			s = append(s, fmt.Sprint(key, "=", value))
		}
		return "[" + strings.Join(s, ", ") + "]"
	}
	return ""
}

func (r *gologPostfixLogger) LevelPrint(lev mappers.Level, i ...any) {
	i = append(i, " ", r.postfixFromFields())

	r.goLog.LevelPrint(lev, i...)
}

func (r *gologPostfixLogger) LevelPrintf(lev mappers.Level, format string, i ...any) {
	format = format + " %s"
	i = append(i, r.postfixFromFields())

	r.goLog.LevelPrintf(lev, format, i...)
}

func (r *gologPostfixLogger) LevelPrintln(lev mappers.Level, i ...any) {
	i = append(i, r.postfixFromFields())
	r.goLog.LevelPrintln(lev, i...)
}
