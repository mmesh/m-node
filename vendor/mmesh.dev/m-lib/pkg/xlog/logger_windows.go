//go:build windows
// +build windows

package xlog

import (
	"fmt"
	"log"
	"os"
	"time"
)

type LoggerSpec struct {
	logLevel LogLevel
	hostID   string

	ansiColor bool

	stdLog           map[LogLevel]*log.Logger
	stdLogFile       *os.File
	stdLogFileLogger map[LogLevel]*log.Logger
	windowsLogger    *windowsLogger
	sumologicLogger  *sumologicLogger
	slackLogger      *slackLogger
}

func (l *LoggerSpec) Close() {
	if l.stdLogFile != nil {
		l.stdLogFile.Close()
	}

	if l.windowsLogger != nil {
		l.windowsLogger.close()
	}

	if l.sumologicLogger != nil {
		l.sumologicLogger.endCh <- struct{}{}
	}
}

func (l *LoggerSpec) writeLog(level LogLevel, tm time.Time, msg ...any) {
	all := append([]interface{}{l.logPrefix(level, tm)}, msg...)

	fmt.Println(all...)

	if l.windowsLogger != nil {
		if level >= l.windowsLogger.logLevel {
			for _, m := range msg {
				str := fmt.Sprintf("%v", m)
				if err := l.windowsLogger.writeLog(level, str); err != nil {
					eventLogErr := fmt.Sprintf("Unable to write msg to windows event log: %v", err)
					fmt.Println(l.logPrefix(level, tm), eventLogErr)
				}
			}
		}
	}
}
