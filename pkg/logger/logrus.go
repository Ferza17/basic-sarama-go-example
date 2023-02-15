package logger

import (
	"context"
	"fmt"
	"io"
	"os"
	"runtime"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

type logrusLogEntry struct {
	entry       *logrus.Entry
	serviceName string
}

type logrusLogger struct {
	logger      *logrus.Logger
	serviceName string
}

func fileInfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "<???>"
		line = 1
	}
	return fmt.Sprintf("%s:%d", file, line)
}

func getFormatter(isJSON bool) logrus.Formatter {
	if isJSON {
		return &logrus.JSONFormatter{}
	}
	return &logrus.TextFormatter{
		FullTimestamp:          true,
		DisableLevelTruncation: true,
	}
}

func newLogrusLogger(config Configuration) (Logger, error) {
	logLevel := config.ConsoleLevel
	if logLevel == "" {
		logLevel = config.FileLevel
	}

	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return nil, err
	}

	stdOutHandler := os.Stdout
	fileHandler := &lumberjack.Logger{
		Filename: config.FileLocation,
		MaxSize:  100,
		Compress: true,
		MaxAge:   28,
	}

	lLogger := &logrus.Logger{
		Out:          stdOutHandler,
		Formatter:    getFormatter(config.FileJSONFormat),
		Hooks:        make(logrus.LevelHooks),
		Level:        level,
		ReportCaller: false,
	}

	if config.EnableConsole && config.EnableFile {
		lLogger.SetOutput(io.MultiWriter(stdOutHandler, fileHandler))
	} else {
		if config.EnableFile {
			lLogger.SetOutput(fileHandler)
			lLogger.SetFormatter(getFormatter(config.FileJSONFormat))
		}
	}

	return &logrusLogger{
		serviceName: config.ServiceName,
		logger:      lLogger,
	}, nil
}

var skipperCaler = 3

func (l *logrusLogger) Debugf(ctx context.Context, format string, args ...interface{}) {
	entry := l.logger.WithField("service", l.serviceName)
	entry.Data["file"] = fileInfo(skipperCaler)
	entry.Debugf(format, args...)
}

func (l *logrusLogger) Infof(ctx context.Context, format string, args ...interface{}) {
	entry := l.logger.WithField("service", l.serviceName)
	entry.Data["file"] = fileInfo(skipperCaler)
	entry.Infof(format, args...)
}

func (l *logrusLogger) Warnf(ctx context.Context, format string, args ...interface{}) {
	entry := l.logger.WithField("service", l.serviceName)
	entry.Data["file"] = fileInfo(skipperCaler)
	entry.Warnf(format, args...)
}

func (l *logrusLogger) Errorf(ctx context.Context, format string, args ...interface{}) {
	entry := l.logger.WithField("service", l.serviceName)
	entry.Data["file"] = fileInfo(skipperCaler)
	entry.Errorf(format, args...)
}

func (l *logrusLogger) Fatalf(ctx context.Context, format string, args ...interface{}) {
	entry := l.logger.WithField("service", l.serviceName)
	entry.Data["file"] = fileInfo(skipperCaler)
	entry.Fatalf(format, args...)
}

func (l *logrusLogger) Panicf(ctx context.Context, format string, args ...interface{}) {
	entry := l.logger.WithField("service", l.serviceName)
	entry.Data["file"] = fileInfo(skipperCaler)
	entry.Fatalf(format, args...)
}

func (l *logrusLogger) WithFields(fields Fields) Logger {
	return &logrusLogEntry{
		entry:       l.logger.WithFields(convertToLogrusFields(fields)),
		serviceName: l.serviceName,
	}
}

func (l *logrusLogger) WithField(key string, value interface{}) Logger {
	return &logrusLogEntry{
		entry:       l.logger.WithFields(convertToLogrusFields(map[string]interface{}{key: value})),
		serviceName: l.serviceName,
	}
}

func (l *logrusLogEntry) Debugf(ctx context.Context, format string, args ...interface{}) {
	l.entry.WithField("service", l.serviceName).Debugf(format, args...)
}

func (l *logrusLogEntry) Infof(ctx context.Context, format string, args ...interface{}) {
	l.entry.WithField("service", l.serviceName).Infof(format, args...)
}

func (l *logrusLogEntry) Warnf(ctx context.Context, format string, args ...interface{}) {
	l.entry.WithField("service", l.serviceName).Warnf(format, args...)
}

func (l *logrusLogEntry) Errorf(ctx context.Context, format string, args ...interface{}) {
	l.entry.WithField("service", l.serviceName).Errorf(format, args...)
}

func (l *logrusLogEntry) Fatalf(ctx context.Context, format string, args ...interface{}) {
	l.entry.WithField("service", l.serviceName).Fatalf(format, args...)
}

func (l *logrusLogEntry) Panicf(ctx context.Context, format string, args ...interface{}) {
	l.entry.WithField("service", l.serviceName).Fatalf(format, args...)
}

func (l *logrusLogEntry) WithField(key string, value interface{}) Logger {
	return &logrusLogEntry{
		entry:       l.entry.WithField("service", l.serviceName).WithField(key, value),
		serviceName: l.serviceName,
	}
}

func (l *logrusLogEntry) WithFields(fields Fields) Logger {
	return &logrusLogEntry{
		entry:       l.entry.WithField("service", l.serviceName).WithFields(convertToLogrusFields(fields)),
		serviceName: l.serviceName,
	}
}

func convertToLogrusFields(fields Fields) logrus.Fields {
	logrusFields := map[string]interface{}{}
	for index, val := range fields {
		logrusFields[index] = val
	}
	return logrusFields
}
