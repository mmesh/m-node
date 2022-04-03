package xlog

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/mgutz/ansi"
	"mmesh.dev/m-lib/pkg/utils/colors"
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

var logPrefixes = map[LogLevel]string{
	TRACE: "trace",
	DEBUG: "debug",
	INFO:  " info",
	WARN:  " warn",
	ERROR: "error",
	ALERT: "alert",
}

func (ll LogLevel) String() string {
	return logPrefixes[ll]
}

var logColorFuncs = map[LogLevel]func(string) string{
	TRACE: ansi.ColorFunc("magenta+bh"),
	DEBUG: ansi.ColorFunc("blue+b"),
	INFO:  ansi.ColorFunc("blue+bh"),
	WARN:  ansi.ColorFunc("yellow+b"),
	ERROR: ansi.ColorFunc("red+bh"),
	ALERT: ansi.ColorFunc("white+bh:red"),
}

type Priority string

const (
	LOW    Priority = "LOW"
	MEDIUM Priority = "MEDIUM"
	HIGH   Priority = "HIGH"
)

var logPriorities = map[LogLevel]Priority{
	TRACE: LOW,
	DEBUG: LOW,
	INFO:  LOW,
	WARN:  MEDIUM,
	ERROR: HIGH,
	ALERT: HIGH,
}

type LoggerSpec struct {
	logLevel LogLevel
	hostID   string

	ansiColor bool

	stdLog          map[LogLevel]*log.Logger
	stdLogFile      map[LogLevel]*log.Logger
	sumologicLogger *sumologicLogger
	slackLogger     *slackLogger
}

var l = &LoggerSpec{
	logLevel: INFO,
}

func Logger() *LoggerSpec {
	return l
}

func (l *LoggerSpec) SetLogLevel(level LogLevel) *LoggerSpec {
	l.logLevel = level
	return l
}

func (l *LoggerSpec) SetHostID(hostID string) *LoggerSpec {
	l.hostID = hostID
	return l
}

func (l *LoggerSpec) SetANSIColor(enabled bool) *LoggerSpec {
	l.ansiColor = enabled
	return l
}

func (l *LoggerSpec) SetStdLogger() *LoggerSpec {
	l.stdLog = map[LogLevel]*log.Logger{
		TRACE: log.New(os.Stdout, "["+logPrefixes[TRACE]+"]\t", log.Ldate|log.Ltime),
		DEBUG: log.New(os.Stdout, "["+logPrefixes[DEBUG]+"]\t", log.Ldate|log.Ltime),
		INFO:  log.New(os.Stdout, "["+logPrefixes[INFO]+"]\t", log.Ldate|log.Ltime),
		WARN:  log.New(os.Stdout, "["+logPrefixes[WARN]+"]\t", log.Ldate|log.Ltime),
		ERROR: log.New(os.Stdout, "["+logPrefixes[ERROR]+"]\t", log.Ldate|log.Ltime),
		ALERT: log.New(os.Stdout, "["+logPrefixes[ALERT]+"]\t", log.Ldate|log.Ltime),
	}
	return l
}

func (l *LoggerSpec) SetLogFile(logfile string) *LoggerSpec {
	f, err := os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		fmt.Println("Unable to open log file:", err)
		os.Exit(1)
	}

	l.stdLogFile = map[LogLevel]*log.Logger{
		TRACE: log.New(f, "["+logPrefixes[TRACE]+"]\t", log.Ldate|log.Ltime),
		DEBUG: log.New(f, "["+logPrefixes[DEBUG]+"]\t", log.Ldate|log.Ltime),
		INFO:  log.New(f, "["+logPrefixes[INFO]+"]\t", log.Ldate|log.Ltime),
		WARN:  log.New(f, "["+logPrefixes[WARN]+"]\t", log.Ldate|log.Ltime),
		ERROR: log.New(f, "["+logPrefixes[ERROR]+"]\t", log.Ldate|log.Ltime),
		ALERT: log.New(f, "["+logPrefixes[ALERT]+"]\t", log.Ldate|log.Ltime),
	}

	return l
}

func (l *LoggerSpec) Close() {
	if l.sumologicLogger != nil {
		l.sumologicLogger.endCh <- struct{}{}
	}
}

func (l *LoggerSpec) logLevelPrefix(level LogLevel) string {
	prefix := "[" + logPrefixes[level] + "]"

	if l.ansiColor {
		return logColorFuncs[level](prefix)
	}

	return prefix
}

func (l *LoggerSpec) logPrefix(level LogLevel, timestamp time.Time) string {
	// hostID := "[" + colors.White(l.hostID) + "]"
	// return l.logLevelPrefix(level) + " " + timestamp + " " + hostID

	if l.ansiColor {
		return l.logLevelPrefix(level) + " " + colors.Black(timestamp.Format(TIME_FORMAT))
	}

	return l.logLevelPrefix(level) + " " + timestamp.Format(TIME_FORMAT)
}

func (l *LoggerSpec) severity(level LogLevel) string {
	return strings.ToUpper(strings.TrimSpace(logPrefixes[level]))
}

func (l *LoggerSpec) priority(level LogLevel) Priority {
	return logPriorities[level]
}

func (l *LoggerSpec) log(level LogLevel, args ...interface{}) {
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

		if l.sumologicLogger != nil {
			if level >= l.sumologicLogger.logLevel {
				if err := l.sumologicLog(level, timestamp, fmt.Sprint(args...)); err != nil {
					sumologicErr := fmt.Sprintf("Unable to post log msg to SumoLogic: %v", err)
					fmt.Println(l.logPrefix(level, timestamp), sumologicErr)
				}
			}
		}

		if l.slackLogger != nil {
			if level >= l.slackLogger.logLevel {
				if err := l.slackLog(level, timestamp, fmt.Sprint(args...)); err != nil {
					slackErr := fmt.Sprintf("Unable to post log msg to Slack: %v", err)
					fmt.Println(l.logPrefix(level, timestamp), slackErr)
				}
			}
		}
	}
}

func (l *LoggerSpec) logf(level LogLevel, format string, args ...interface{}) {
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
