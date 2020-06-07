// Copyright (C) 2019 x6a
//
// pkg is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// pkg is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with pkg. If not, see <http://www.gnu.org/licenses/>.

package xlog

import (
	"fmt"
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

const (
	logOptionSlack = iota
	logOptionFile
	logOptionSyslog
)

type slackLoggerCfg struct {
	webhook  string
	user     string
	icon     string
	logLevel LogLevel
	channels map[LogLevel]string
	colors   map[LogLevel]string
}

type LogOption struct {
	key   int
	value interface{}
}

type logger struct {
	logLevel LogLevel
	hostID   string

	slackLogger *slackLoggerCfg
	outputFile  string
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

type SlackOption struct {
	Level        LogLevel
	Webhook      string
	User         string
	Icon         string
	TraceChannel string
	DebugChannel string
	InfoChannel  string
	WarnChannel  string
	ErrorChannel string
	AlertChannel string
}

func WithSlack(opt *SlackOption) *LogOption {
	return &LogOption{
		key: logOptionSlack,
		value: &slackLoggerCfg{
			webhook:  opt.Webhook,
			user:     opt.User,
			icon:     opt.Icon,
			logLevel: opt.Level,
			channels: map[LogLevel]string{
				TRACE: opt.TraceChannel,
				DEBUG: opt.DebugChannel,
				INFO:  opt.InfoChannel,
				WARN:  opt.WarnChannel,
				ERROR: opt.ErrorChannel,
				ALERT: opt.AlertChannel,
			},
			colors: map[LogLevel]string{
				TRACE: "#ff77ff",
				DEBUG: "#444999",
				INFO:  "#009999",
				WARN:  "#fff000",
				ERROR: "#ff4444",
				ALERT: "#990000",
			},
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
		case logOptionSlack:
			l.slackLogger = opt.value.(*slackLoggerCfg)
		case logOptionFile:
			l.outputFile = opt.value.(string)
		}
	}
}

func (l *logger) logLevelPrefix(level LogLevel) string {
	prefix := "[" + logPrefixes[level] + "]"

	return logColorFuncs[level](prefix)
}

func (l *logger) logPrefix(level LogLevel, timestamp time.Time) string {
	//hostID := "[" + colors.White(l.hostID) + "]"

	// return l.logLevelPrefix(level) + " " + timestamp + " " + hostID
	return l.logLevelPrefix(level) + " " + colors.Black(timestamp.Format(TIME_FORMAT))
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

		all := append([]interface{}{l.logPrefix(level, timestamp)}, args...)
		fmt.Println(all...)

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

		fmt.Println(l.logPrefix(level, timestamp), fmt.Sprintf(format, args...))

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
