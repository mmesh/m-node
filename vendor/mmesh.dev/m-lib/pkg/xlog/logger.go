package xlog

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/mgutz/ansi"
	"x6a.dev/pkg/colors"
)

const TIME_FORMAT = "2006-01-02 15:04:05.000"

type LogLevel int

const (
	TRACE LogLevel = iota
	DEBUG
	INFO
	WARN
	ERROR
	ALERT
)

type Priority string

const (
	LOW    Priority = "LOW"
	MEDIUM Priority = "MEDIUM"
	HIGH   Priority = "HIGH"
)

type logOption int

const (
	logOptionOutputStdout logOption = iota
	logOptionOutputStderr
	logOptionOutputStdLogger
	logOptionOutputFile
	logOptionOutputSyslog
	logOptionOutputSlack
	logOptionANSIColor
)

type LogOption struct {
	key   logOption
	value interface{}
}

var logPrefixes = map[LogLevel]string{
	TRACE: "trace",
	DEBUG: "debug",
	INFO:  " info",
	WARN:  " warn",
	ERROR: "error",
	ALERT: "alert",
}

var logPriorities = map[LogLevel]Priority{
	TRACE: LOW,
	DEBUG: LOW,
	INFO:  LOW,
	WARN:  MEDIUM,
	ERROR: HIGH,
	ALERT: HIGH,
}

var logColorFuncs = map[LogLevel]func(string) string{
	TRACE: ansi.ColorFunc("magenta+bh"),
	DEBUG: ansi.ColorFunc("blue+b"),
	INFO:  ansi.ColorFunc("blue+bh"),
	WARN:  ansi.ColorFunc("yellow+b"),
	ERROR: ansi.ColorFunc("red+bh"),
	ALERT: ansi.ColorFunc("white+bh:red"),
}

type logger struct {
	logLevel LogLevel
	hostID   string

	ansiColor bool

	stdLog      map[LogLevel]*log.Logger
	stdLogFile  map[LogLevel]*log.Logger
	slackLogger *slackLoggerCfg
}

var l = &logger{
	logLevel: INFO,
}

func SetLogger(level LogLevel, hostID string, logOpts ...*LogOption) {
	logger := &logger{
		logLevel: level,
		hostID:   hostID,
	}
	logger.setOptions(logOpts...)

	l = logger
}

func WithANSIColor(enabled bool) *LogOption {
	return &LogOption{
		key:   logOptionANSIColor,
		value: enabled,
	}
}

func WithStdLogger() *LogOption {
	return &LogOption{
		key: logOptionOutputStdLogger,
		value: map[LogLevel]*log.Logger{
			TRACE: log.New(os.Stdout, "["+logPrefixes[TRACE]+"]\t", log.Ldate|log.Ltime),
			DEBUG: log.New(os.Stdout, "["+logPrefixes[DEBUG]+"]\t", log.Ldate|log.Ltime),
			INFO:  log.New(os.Stdout, "["+logPrefixes[INFO]+"]\t", log.Ldate|log.Ltime),
			WARN:  log.New(os.Stdout, "["+logPrefixes[WARN]+"]\t", log.Ldate|log.Ltime),
			ERROR: log.New(os.Stdout, "["+logPrefixes[ERROR]+"]\t", log.Ldate|log.Ltime),
			ALERT: log.New(os.Stdout, "["+logPrefixes[ALERT]+"]\t", log.Ldate|log.Ltime),
		},
	}
}

func WithLogFile(logfile string) *LogOption {
	f, err := os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		fmt.Println("Unable to open log file:", err)
		os.Exit(1)
	}

	return &LogOption{
		key: logOptionOutputFile,
		value: map[LogLevel]*log.Logger{
			TRACE: log.New(f, "["+logPrefixes[TRACE]+"]\t", log.Ldate|log.Ltime),
			DEBUG: log.New(f, "["+logPrefixes[DEBUG]+"]\t", log.Ldate|log.Ltime),
			INFO:  log.New(f, "["+logPrefixes[INFO]+"]\t", log.Ldate|log.Ltime),
			WARN:  log.New(f, "["+logPrefixes[WARN]+"]\t", log.Ldate|log.Ltime),
			ERROR: log.New(f, "["+logPrefixes[ERROR]+"]\t", log.Ldate|log.Ltime),
			ALERT: log.New(f, "["+logPrefixes[ALERT]+"]\t", log.Ldate|log.Ltime),
		},
	}
}

func GetLogLevel(loglevel string) LogLevel {
	if strings.Contains(strings.ToUpper(loglevel), "TRACE") {
		return TRACE
	}
	if strings.Contains(strings.ToUpper(loglevel), "DEBUG") {
		return DEBUG
	}
	if strings.Contains(strings.ToUpper(loglevel), "INFO") {
		return INFO
	}
	if strings.Contains(strings.ToUpper(loglevel), "WARN") {
		return WARN
	}
	if strings.Contains(strings.ToUpper(loglevel), "ERROR") {
		return ERROR
	}
	if strings.Contains(strings.ToUpper(loglevel), "ALERT") {
		return ALERT
	}

	return -1
}

func (ll LogLevel) String() string {
	return logPrefixes[ll]
}

func (l *logger) setOptions(logOpts ...*LogOption) {
	for _, opt := range logOpts {
		switch opt.key {
		case logOptionANSIColor:
			l.ansiColor = opt.value.(bool)
		case logOptionOutputStdLogger:
			l.stdLog = opt.value.(map[LogLevel]*log.Logger)
		case logOptionOutputFile:
			l.stdLogFile = opt.value.(map[LogLevel]*log.Logger)
		case logOptionOutputSlack:
			l.slackLogger = opt.value.(*slackLoggerCfg)
		}
	}
}

func (l *logger) logLevelPrefix(level LogLevel) string {
	prefix := "[" + logPrefixes[level] + "]"

	if l.ansiColor {
		return logColorFuncs[level](prefix)
	}

	return prefix
}

func (l *logger) logPrefix(level LogLevel, timestamp time.Time) string {
	//hostID := "[" + colors.White(l.hostID) + "]"
	// return l.logLevelPrefix(level) + " " + timestamp + " " + hostID

	if l.ansiColor {
		return l.logLevelPrefix(level) + " " + colors.Black(timestamp.Format(TIME_FORMAT))
	}

	return l.logLevelPrefix(level) + " " + timestamp.Format(TIME_FORMAT)
}

func (l *logger) severity(level LogLevel) string {
	return strings.ToUpper(strings.TrimSpace(logPrefixes[level]))
}

func (l *logger) priority(level LogLevel) Priority {
	return logPriorities[level]
}

func (l *logger) log(level LogLevel, args ...interface{}) {
	if level >= l.logLevel {
		timestamp := time.Now()

		if l.stdLogFile != nil {
			l.stdLogFile[level].Println(args...)
		} else if l.stdLog != nil {
			l.stdLog[level].Println(args...)
		} else {
			all := append([]interface{}{l.logPrefix(level, timestamp)}, args...)
			fmt.Println(all...)
		}

		if l.slackLogger != nil {
			if level >= l.slackLogger.logLevel {
				if err := l.slackLog(level, timestamp, fmt.Sprint(args...)); err != nil {
					slackErr := fmt.Errorf("Unable to post to Slack: %v", err)
					fmt.Println(l.logPrefix(level, timestamp), slackErr)
				}
			}
		}
	}
}

func (l *logger) logf(level LogLevel, format string, args ...interface{}) {
	if level >= l.logLevel {
		timestamp := time.Now()

		if l.stdLogFile != nil {
			l.stdLogFile[level].Println(fmt.Sprintf(format, args...))
		} else if l.stdLog != nil {
			l.stdLog[level].Println(fmt.Sprintf(format, args...))
		} else {
			fmt.Println(l.logPrefix(level, timestamp), fmt.Sprintf(format, args...))
		}

		if l.slackLogger != nil {
			if level >= l.slackLogger.logLevel {
				if err := l.slackLog(level, timestamp, fmt.Sprintf(format, args...)); err != nil {
					slackErr := fmt.Errorf("Unable to post to Slack: %v", err)
					fmt.Println(l.logPrefix(level, timestamp), slackErr)
				}
			}
		}
	}
}

func Trace(args ...interface{}) {
	l.log(TRACE, args...)
}

func Debug(args ...interface{}) {
	l.log(DEBUG, args...)
}

func Info(args ...interface{}) {
	l.log(INFO, args...)
}

func Warn(args ...interface{}) {
	l.log(WARN, args...)
}

func Error(args ...interface{}) {
	l.log(ERROR, args...)
}

func Alert(args ...interface{}) {
	l.log(ALERT, args...)
}

func Tracef(format string, args ...interface{}) {
	l.logf(TRACE, format, args...)
}

func Debugf(format string, args ...interface{}) {
	l.logf(DEBUG, format, args...)
}

func Infof(format string, args ...interface{}) {
	l.logf(INFO, format, args...)
}

func Warnf(format string, args ...interface{}) {
	l.logf(WARN, format, args...)
}

func Errorf(format string, args ...interface{}) {
	l.logf(ERROR, format, args...)
}

func Alertf(format string, args ...interface{}) {
	l.logf(ALERT, format, args...)
}
