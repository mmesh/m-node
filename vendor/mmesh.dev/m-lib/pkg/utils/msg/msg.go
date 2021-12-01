package msg

import (
	"fmt"

	"github.com/mgutz/ansi"
)

const (
	TRACE = iota
	DEBUG
	INFO
	OK
	FAIL
	WARN
	ERROR
	ALERT
)

var msgPrefixes = map[int]string{
	TRACE: "trace",
	DEBUG: "debug",
	INFO:  " info",
	OK:    " o k ",
	FAIL:  " fail",
	WARN:  " warn",
	ERROR: "error",
	ALERT: "alert",
}

var msgColorFuncs = map[int]func(string) string{
	TRACE: ansi.ColorFunc("white+bh:black"),
	DEBUG: ansi.ColorFunc("white+bh:cyan"),
	INFO:  ansi.ColorFunc("white+bh:blue"),
	OK:    ansi.ColorFunc("white+bh:green"),
	FAIL:  ansi.ColorFunc("white+bh:red"),
	WARN:  ansi.ColorFunc("white+bh:yellow"),
	ERROR: ansi.ColorFunc("white+bh:red"),
	ALERT: ansi.ColorFunc("white+bh:magenta"),
}

func msgLevelPrefix(level int) string {
	prefix := "[" + msgPrefixes[level] + "]"

	return msgColorFuncs[level](prefix)
}

func msg(level int, args ...interface{}) {
	all := append([]interface{}{msgLevelPrefix(level)}, args...)
	fmt.Println()
	fmt.Println(all...)
}

func msgf(level int, format string, args ...interface{}) {
	fmt.Println()
	fmt.Println(msgLevelPrefix(level), fmt.Sprintf(format, args...))
}

func Trace(args ...interface{}) {
	msg(TRACE, args...)
}

func Debug(args ...interface{}) {
	msg(DEBUG, args...)
}

func Info(args ...interface{}) {
	msg(INFO, args...)
}

func Ok(args ...interface{}) {
	msg(OK, args...)
}

func Fail(args ...interface{}) {
	msg(FAIL, args...)
}

func Warn(args ...interface{}) {
	msg(WARN, args...)
}

func Error(args ...interface{}) {
	msg(ERROR, args...)
}

func Alert(args ...interface{}) {
	msg(ALERT, args...)
}

func Tracef(format string, args ...interface{}) {
	msgf(TRACE, format, args...)
}

func Debugf(format string, args ...interface{}) {
	msgf(DEBUG, format, args...)
}

func Infof(format string, args ...interface{}) {
	msgf(INFO, format, args...)
}

func Okf(format string, args ...interface{}) {
	msgf(OK, format, args...)
}

func Failf(format string, args ...interface{}) {
	msgf(FAIL, format, args...)
}

func Warnf(format string, args ...interface{}) {
	msgf(WARN, format, args...)
}

func Errorf(format string, args ...interface{}) {
	msgf(ERROR, format, args...)
}

func Alertf(format string, args ...interface{}) {
	msgf(ALERT, format, args...)
}
