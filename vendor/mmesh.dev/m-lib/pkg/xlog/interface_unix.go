//go:build !windows
// +build !windows

package xlog

import (
	"fmt"
	"log"
	"os"
)

type LoggerSpec struct {
	logLevel LogLevel
	hostID   string

	ansiColor bool

	stdLog           map[LogLevel]*log.Logger
	stdLogFile       *os.File
	stdLogFileLogger map[LogLevel]*log.Logger
	sumologicLogger  *sumologicLogger
	slackLogger      *slackLogger
}

func (l *LoggerSpec) Close() {
	if l.stdLogFile != nil {
		l.stdLogFile.Close()
	}

	if l.sumologicLogger != nil {
		l.sumologicLogger.endCh <- struct{}{}
	}
}

func (l *LoggerSpec) writeLog(level LogLevel, msg ...any) {
	fmt.Println(msg...)
}
