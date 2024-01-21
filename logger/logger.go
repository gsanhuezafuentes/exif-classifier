package logger

import (
	"io"
	"log"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARNING
	ERROR
)

type Logger interface {
	SetLogLevel(level LogLevel)
	SetOutput(w io.Writer)
	Debug(message ...any)
	Info(message ...any)
	Warning(message ...any)
	Error(message ...any)
	Debugf(format string, args ...any)
	Infof(format string, args ...any)
	Warningf(format string, args ...any)
	Errorf(format string, args ...any)
}

type CustomLogger struct {
	*log.Logger
	logLevel LogLevel
}

var logger = New(io.Discard, "", log.LstdFlags, INFO)

func New(out io.Writer, prefix string, flag int, level LogLevel) *CustomLogger {
	return &CustomLogger{
		Logger:   log.New(out, prefix, flag),
		logLevel: level,
	}
}

func (r *CustomLogger) SetLogLevel(level LogLevel) {
	r.logLevel = level
}

func (r *CustomLogger) SetOutput(w io.Writer) {
	r.Logger.SetOutput(w)
}

func (r *CustomLogger) Debug(message ...any) {
	if r.logLevel <= DEBUG {
		r.Logger.Println("[DEBUG] ", message)
	}
}

func (r *CustomLogger) Info(message ...any) {
	if r.logLevel <= INFO {
		r.Logger.Println("[INFO] ", message)
	}
}

func (r *CustomLogger) Warning(message ...any) {
	if r.logLevel <= WARNING {
		r.Logger.Println("[WARNING] ", message)
	}
}

func (r *CustomLogger) Error(message ...any) {
	if r.logLevel <= ERROR {
		r.Logger.Println("[ERROR] ", message)
	}
}

func (r *CustomLogger) Debugf(format string, args ...any) {
	if r.logLevel <= DEBUG {
		r.Logger.Printf("[DEBUG] "+format, args...)
	}
}

func (r *CustomLogger) Infof(format string, args ...any) {
	if r.logLevel <= INFO {
		r.Logger.Printf("[INFO] "+format, args...)
	}
}

func (r *CustomLogger) Warningf(format string, args ...any) {
	if r.logLevel <= WARNING {
		r.Logger.Printf("[WARNING] "+format, args...)
	}
}

func (r *CustomLogger) Errorf(format string, args ...any) {
	if r.logLevel <= ERROR {
		r.Logger.Printf("[ERROR] "+format, args...)
	}
}

func GetLogger() *CustomLogger {
	return logger
}
