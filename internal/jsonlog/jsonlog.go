package jsonlog

import (
	"io"
	"runtime/debug"
	"strconv"

	"acad.learn2earn.ng/git/dositadi/groupie-tracker/internal/utils"
)

type Level int

func (l Level) String() string {
	return strconv.Itoa(int(l))
}

const (
	INFO Level = iota
	ERROR
	FATAL
	DEBUG
)

type Logger struct {
	out      io.Writer
	minLevel Level
}

func New(out io.Writer, minLevel Level) *Logger {
	return &Logger{
		out:      out,
		minLevel: minLevel,
	}
}

func (l *Logger) print(level Level, message string, properties map[string]string) (n int, err error) {
	if level > l.minLevel {
		return
	}

	data := struct {
		Level      string
		Message    string
		Properties map[string]string
		Trace      string
	}{
		Level:      level.String(),
		Message:    message,
		Properties: properties,
	}

	if level >= ERROR {
		data.Trace = string(debug.Stack())
	}

	log := utils.MarshalObject(data)

	return l.out.Write(append(log, '\n'))
}

func (l *Logger) PrintError(message string, properties map[string]string) {
	l.print(ERROR, message, properties)
}

func (l *Logger) PrintInfo(message string, properties map[string]string) {
	l.print(INFO, message, properties)
}

func (l *Logger) PrintFatal(message string, properties map[string]string) {
	l.print(FATAL, message, properties)
}

func (l *Logger) Write(p []byte) (n int, err error) {
	return l.print(ERROR, string(p), nil)
}
