package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// Interface contract for logger -.
type Interface interface {
	Debug(message interface{}, args ...interface{})
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message interface{}, args ...interface{})
	Fatal(args ...interface{})
}

// Logger -.
type Logger struct {
	logger *zap.SugaredLogger
}

var _ Interface = (*Logger)(nil)

// New -.
func New(level string) *Logger {

	logger, _ := zap.NewDevelopment()
	sugaredLogger := logger.Sugar().WithOptions(zap.AddStacktrace(zapcore.InvalidLevel))

	return &Logger{
		logger: sugaredLogger,
	}
}

// Debug -.
func (l *Logger) Debug(message interface{}, args ...interface{}) {
	l.logger.Debug(message, args)
}

// Info -.
func (l *Logger) Info(message string, args ...interface{}) {
	l.logger.Info(message, args)
}

// Warn -.
func (l *Logger) Warn(message string, args ...interface{}) {
	l.logger.Warn(message, args)
}

// Error -.
func (l *Logger) Error(message interface{}, args ...interface{}) {
	l.logger.Error(message, args)
}

// Fatal -.
func (l *Logger) Fatal(args ...interface{}) {
	l.logger.Fatal(args)
	os.Exit(1)
}
