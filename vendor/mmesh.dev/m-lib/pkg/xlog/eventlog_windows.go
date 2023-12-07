//go:build windows
// +build windows

package xlog

import (
	"fmt"
	"strings"

	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/svc/eventlog"
	"mmesh.dev/m-lib/pkg/errors"
)

type EventLogOptions struct {
	Level  LogLevel
	Source string
}

type windowsLogger struct {
	logLevel      LogLevel
	infoLogger    *eventLogWriter
	warningLogger *eventLogWriter
	errorLogger   *eventLogWriter
}

type eventLogWriter struct {
	level LogLevel
	src   string
	elog  *eventlog.Log
}

func (l *LoggerSpec) SetWindowsLogger(opt *EventLogOptions) *LoggerSpec {
	wl, err := setupWindowsLogger(opt)
	if err != nil {
		fmt.Printf("Unable to setup windows event logger: %v", err)
	}

	l.windowsLogger = wl

	return l
}

func (wl *windowsLogger) close() {
	if wl.infoLogger != nil {
		wl.infoLogger.close()
	}

	if wl.warningLogger != nil {
		wl.warningLogger.close()
	}

	if wl.errorLogger != nil {
		wl.errorLogger.close()
	}
}

func (wl *windowsLogger) writeLog(level LogLevel, msg string) error {
	switch level {
	case TRACE:
		return wl.infoLogger.writeLog(level, msg)
	case DEBUG:
		return wl.infoLogger.writeLog(level, msg)
	case INFO:
		return wl.infoLogger.writeLog(level, msg)
	case WARN:
		return wl.warningLogger.writeLog(level, msg)
	case ERROR:
		return wl.errorLogger.writeLog(level, msg)
	case ALERT:
		return wl.errorLogger.writeLog(level, msg)
	}

	return wl.infoLogger.writeLog(level, msg)
}

// writeLog sends a log message to the Event Log.
func (w *eventLogWriter) writeLog(level LogLevel, msg string) error {
	switch level {
	case TRACE:
		return w.elog.Info(1, msg)
	case DEBUG:
		return w.elog.Info(1, msg)
	case INFO:
		return w.elog.Info(1, msg)
	case WARN:
		return w.elog.Warning(3, msg)
	case ERROR:
		return w.elog.Error(2, msg)
	case ALERT:
		return w.elog.Error(2, msg)
	}

	return fmt.Errorf("unrecognized logLevel: %v", w.level)
}

func (w *eventLogWriter) close() error {
	return w.elog.Close()
}

func setupWindowsLogger(opt *EventLogOptions) (*windowsLogger, error) {
	infoL, err := newW(INFO, opt.Source)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function newW()", errors.Trace())
	}

	warningL, err := newW(WARN, opt.Source)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function newW()", errors.Trace())
	}

	errL, err := newW(ERROR, opt.Source)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function newW()", errors.Trace())
	}

	return &windowsLogger{
		logLevel:      opt.Level,
		infoLogger:    infoL,
		warningLogger: warningL,
		errorLogger:   errL,
	}, nil
}

func newW(level LogLevel, src string) (*eventLogWriter, error) {
	// Continue if we receive "registry key already exists" or if we get
	// ERROR_ACCESS_DENIED so that we can log without administrative permissions
	// for pre-existing eventlog sources.
	if err := eventlog.InstallAsEventCreate(src, eventlog.Info|eventlog.Warning|eventlog.Error); err != nil {
		if !strings.Contains(err.Error(), "registry key already exists") && err != windows.ERROR_ACCESS_DENIED {
			return nil, errors.Wrapf(err, "[%v] function eventlog.InstallAsEventCreate()", errors.Trace())
		}
	}
	elog, err := eventlog.Open(src)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function eventlog.Open()", errors.Trace())
	}
	return &eventLogWriter{
		level: level,
		src:   src,
		elog:  elog,
	}, nil
}
