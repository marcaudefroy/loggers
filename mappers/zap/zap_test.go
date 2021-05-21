package zap

import (
	"bytes"
	"regexp"
	"testing"

	"github.com/marcaudefroy/loggers"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func newBufferedZapLog() (loggers.Contextual, *bytes.Buffer) {
	var b []byte
	var bb = bytes.NewBuffer(b)

	writer := zapcore.Lock(zapcore.AddSync(bb))
	enabler := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return true
	})
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	core := zapcore.NewCore(consoleEncoder, writer, enabler)
	l := zap.New(core).Sugar()
	return New(l), bb
}

func TestZapInterface(t *testing.T) {
	var _ loggers.Contextual = NewDefaultLogger()
}

func TestZapLevelOutput(t *testing.T) {
	l, b := newBufferedZapLog()
	l.Info("This is a test")

	expectedMatch := "(?i)info.*This is a test"
	actual := b.String()
	if ok, _ := regexp.Match(expectedMatch, []byte(actual)); !ok {
		t.Errorf("Log output mismatch %s (actual) != %s (expected)", actual, expectedMatch)
	}
}

func TestZapLevelfOutput(t *testing.T) {
	l, b := newBufferedZapLog()
	l.Errorf("This is %s test", "a")

	expectedMatch := "(?i)erro.*This is a test"
	actual := b.String()
	if ok, _ := regexp.Match(expectedMatch, []byte(actual)); !ok {
		t.Errorf("Log output mismatch %s (actual) != %s (expected)", actual, expectedMatch)
	}
}

func TestZapLevellnOutput(t *testing.T) {
	l, b := newBufferedZapLog()
	l.Debugln("This is a test.", "So is this.")

	expectedMatch := "(?i)debu.*This is a test. So is this."
	actual := b.String()
	if ok, _ := regexp.Match(expectedMatch, []byte(actual)); !ok {
		t.Errorf("Log output mismatch %s (actual) != %s (expected)", actual, expectedMatch)
	}
}

func TestZapWithFieldsOutput(t *testing.T) {
	l, b := newBufferedZapLog()
	l.WithFields("test", true).Warn("This is a message.")

	expectedMatch := "(?i)warn.*This is a message.*test.*true"
	actual := b.String()
	if ok, _ := regexp.Match(expectedMatch, []byte(actual)); !ok {
		t.Errorf("Log output mismatch %s (actual) != %s (expected)", actual, expectedMatch)
	}
}

func TestZapChainedWithFieldsOutput(t *testing.T) {
	l, b := newBufferedZapLog()
	l.WithFields("test", true).WithFields("test2", false).Warn("This is a message.")

	expectedMatch := "(?i)warn.*This is a message.*test.*true.*test2.*false"
	actual := b.String()
	if ok, _ := regexp.Match(expectedMatch, []byte(actual)); !ok {
		t.Errorf("Log output mismatch %s (actual) != %s (expected)", actual, expectedMatch)
	}
}

func TestZapWithFieldsfOutput(t *testing.T) {
	l, b := newBufferedZapLog()
	l.WithFields("test", true, "Error", "serious").Errorf("This is a %s.", "message")

	expectedMatch := "(?i)erro.*This is a message.*test.*true.*Error.*serious"
	actual := b.String()
	if ok, _ := regexp.Match(expectedMatch, []byte(actual)); !ok {
		t.Errorf("Log output mismatch %s (actual) != %s (expected)", actual, expectedMatch)
	}
}

func TestZapFieldsfOutput(t *testing.T) {
	l := NewDefaultLogger()
	l = l.WithFields("test", true, "Error", "serious")
	nl := l.WithField("foo", "bar")

	lFields := l.Fields()
	nlFields := nl.Fields()

	if len(lFields) != 4 {
		t.Errorf("Log fields must have %d elements, it have %d", 4, len(lFields))
	}
	if len(nlFields) != 6 {
		t.Errorf("Log fields must have %d elements, it have %d", 4, len(nlFields))
	}
}
