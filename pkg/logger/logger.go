package logger

import (
	"context"
	"errors"

	"github.com/ferza17/kafka-basic/consumer/config"
)

// A global variable so that log functions can be directly accessed
var log Logger

// Fields Type to pass when we want to call WithFields for structured logging
type Fields map[string]interface{}

const (
	// Debug has verbose message
	Debug = "debug"
	// Info is default log level
	Info = "info"
	// Warn is for logging messages about possible issues
	Warn = "warn"
	// Error is for logging errors
	Error = "error"
	// Fatal is for logging fatal messages. The sytem shutsdown after logging the message.
	Fatal = "fatal"
)

const (
	InstanceLogrusLogger int = iota
)

var (
	errInvalidLoggerInstance = errors.New("invalid logger instance")
)

// Logger is our contract for the logger
type Logger interface {
	Debugf(ctx context.Context, format string, args ...interface{})
	Infof(ctx context.Context, format string, args ...interface{})
	Warnf(ctx context.Context, format string, args ...interface{})
	Errorf(ctx context.Context, format string, args ...interface{})
	Fatalf(ctx context.Context, format string, args ...interface{})
	Panicf(ctx context.Context, format string, args ...interface{})
	WithFields(keyValues Fields) Logger
	WithField(key string, value interface{}) Logger
}

// Configuration stores the config for the logger
// For some loggers there can only be one level across writers, for such the level of Console is picked by default
type Configuration struct {
	EnableConsole     bool
	ConsoleJSONFormat bool
	ConsoleLevel      string
	EnableFile        bool
	FileJSONFormat    bool
	FileLevel         string
	FileLocation      string
	ServiceName       string
}

// NewLogger returns an instance of logger
func NewLogger() error {

	var (
		logLevel = Info
	)

	if config.Get().LogLevel != "" {

		switch config.Get().LogLevel {
		case Debug, "Debug", "DEBUG":
			logLevel = Debug
		case Info, "Info", "INFO":
			logLevel = Info
		case Warn, "Warn", "WARN":
			logLevel = Warn
		case Error, "Error", "ERROR":
			logLevel = Error
		case Fatal, "Fatal", "FATAL":
			logLevel = Fatal

		default:
			logLevel = Info
		}
	}

	logger, err := newLogrusLogger(Configuration{
		ServiceName:       "odoo-integration-service",
		EnableConsole:     true,
		ConsoleLevel:      logLevel,
		ConsoleJSONFormat: false,
		EnableFile:        false,
		FileLevel:         logLevel,
		FileJSONFormat:    false,
		FileLocation:      "odoo-integration-backend.log",
	})
	if err != nil {
		return err
	}
	log = logger
	return nil

}

func Debugf(ctx context.Context, format string, args ...interface{}) {
	log.Debugf(ctx, format, args...)
}

func Infof(ctx context.Context, format string, args ...interface{}) {
	log.Infof(ctx, format, args...)
}

func Warnf(ctx context.Context, format string, args ...interface{}) {
	log.Warnf(ctx, format, args...)
}

func Errorf(ctx context.Context, format string, args ...interface{}) {
	log.Errorf(ctx, format, args...)
}

func Fatalf(ctx context.Context, format string, args ...interface{}) {
	log.Fatalf(ctx, format, args...)
}

func Panicf(ctx context.Context, format string, args ...interface{}) {
	log.Panicf(ctx, format, args...)
}

func WithFields(keyValues Fields) Logger {
	return log.WithFields(keyValues)
}

func WithField(key string, value interface{}) Logger {
	return log.WithField(key, value)
}
