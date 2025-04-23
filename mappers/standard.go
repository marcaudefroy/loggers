package mappers

import (
	"errors"
	"fmt"
	"os"
)

type standardMap struct {
	LevelMapper
}

func (s *standardMap) GetUnderlying() any {
	return s
}

// Print should be used only if real error occures.
func (s *standardMap) Print(v ...any) {
	s.LevelPrint(LevelInfo, v...)
}

// Printf works the same as Print but supports formatting.
func (s *standardMap) Printf(format string, v ...any) {
	s.LevelPrintf(LevelInfo, format, v...)
}

// Println works the same as Print but supports formatting.
func (s *standardMap) Println(v ...any) {
	s.LevelPrintln(LevelInfo, v...)
}

// Fatal works the same as Error but it terminates the program right after logging.
// Fatal should be only used when it's not possible to continue program execution.
func (s *standardMap) Fatal(v ...any) {
	s.LevelPrint(LevelFatal, v...)
	os.Exit(1)
}

// Fatalf works the same as Fatal but supports formatting.
func (s *standardMap) Fatalf(format string, v ...any) {
	s.LevelPrintf(LevelFatal, format, v...)
	os.Exit(1)
}

// Fatalln works the same as Info but supports formatting.
func (s *standardMap) Fatalln(v ...any) {
	s.LevelPrintln(LevelFatal, v...)
	os.Exit(1)
}

// Panic works the same as Error but it terminates the program right after logging.
func (s *standardMap) Panic(v ...any) {
	s.LevelPrint(LevelPanic, v...)
	panic(errors.New(fmt.Sprint(v...)))
}

// Panicf works the same as Panic but supports formatting.
func (s *standardMap) Panicf(format string, v ...any) {
	s.LevelPrintf(LevelPanic, format, v...)
	panic(fmt.Errorf(format, v...))
}

// Panicln works the same as Panic but supports formatting.
func (s *standardMap) Panicln(v ...any) {
	s.LevelPrintln(LevelPanic, v...)
	panic(errors.New(fmt.Sprint(v...)))
}
