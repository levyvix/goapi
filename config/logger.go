package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func NewLogger(prefix string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, prefix, log.Ldate|log.Ltime|log.Lshortfile)
	return &Logger{
		debug:   log.New(writer, "DEBUG: ", logger.Flags()),
		info:    log.New(writer, "INFO: ", logger.Flags()),
		warning: log.New(writer, "WARNING: ", logger.Flags()),
		err:     log.New(writer, "ERROR: ", logger.Flags()),
		writer:  writer,
	}

}

// Debug

func (l *Logger) Debug(v ...any) {
	l.debug.Println(v...)
}

func (l *Logger) Debugf(format string, v ...any) {
	l.debug.Printf(format, v...)
}

// Info
func (l *Logger) Info(v ...any) {
	l.info.Println(v...)
}

func (l *Logger) Infof(format string, v ...any) {
	l.info.Printf(format, v...)
}

// Warning
func (l *Logger) Warning(v ...any) {
	l.warning.Println(v...)
}

func (l *Logger) Warningf(format string, v ...any) {
	l.warning.Printf(format, v...)
}

// Error
func (l *Logger) Error(v ...any) {
	l.err.Println(v...)
}

func (l *Logger) Errorf(format string, v ...any) {
	l.err.Printf(format, v...)
}
