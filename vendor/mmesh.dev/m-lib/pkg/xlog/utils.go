package xlog

import "strings"

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
