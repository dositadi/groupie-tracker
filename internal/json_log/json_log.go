package jsonlog

import (
	"encoding/json"
	"io"
	"runtime/debug"
	"sync"
	"time"
)

type Level int8

const (
	LevelInfo Level = iota
	LevelError
	LevelFatal
	LevelOff
)

func (l Level) String() string {
	switch l {
	case LevelInfo:
		return "INFO"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	default:
		return ""
	}
}

type Logger struct {
	minLevel Level
	out      io.Writer
	mu       *sync.Mutex
	color    Color
}

// Constructor to create a new logger
func New(out io.Writer, minLevel Level) *Logger {
	return &Logger{
		out:      out,
		minLevel: minLevel,
		color:    Color{},
		mu:       &sync.Mutex{},
	}
}

// This takes charge of printing the info to the terminal
func (l *Logger) print(level Level, message string, properties map[string]string) (int, error) {
	// If level is less than the severity level then return 0 and do nothing.
	if level < l.minLevel {
		return 0, nil
	}

	aux := struct {
		Level      string            `json:"level"`
		Time       string            `json:"time"`
		Message    string            `json:"message"`
		Properties map[string]string `json:"properties,omitempty"`
		Trace      string            `json:"trace,omitempty"`
	}{
		Level:      level.String(),
		Time:       time.Now().Format(time.RFC1123),
		Message:    message,
		Properties: properties,
	}

	if level >= LevelError {
		aux.Trace = string(debug.Stack())
	}

	var line []byte

	line, err := json.Marshal(aux)
	if err != nil {
		line = []byte(level.String() + l.color.Get("red") + ": unable to marshal the json" + l.color.Get("reset"))
	}

	// Lock before wirting to the io.writer and unlock after writing is done
	l.mu.Lock()
	defer l.mu.Unlock()

	return l.out.Write(append(line, '\n'))
}

// Implementing the writer function of the io interface to write log with no additional properties
func (l *Logger) Write(message []byte) (int, error) {
	return l.print(LevelError, string(message), nil)
}

// Method to log error
func (l *Logger) PrintError(message string, properties map[string]string) (int, error) {
	return l.print(LevelError, message, properties)
}

// Method to log Info
func (l *Logger) PrintInfo(message string, properties map[string]string) (int, error) {
	return l.print(LevelInfo, message, properties)
}

// Method to log fatal
func (l *Logger) PrintFatal(message string, properties map[string]string) (int, error) {
	return l.print(LevelFatal, message, properties)
}
