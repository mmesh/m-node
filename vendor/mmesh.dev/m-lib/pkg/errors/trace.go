package errors

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func Trace() string {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	return fmt.Sprintf("%s:%d | %s", filepath.Base(file), line, f.Name())
}

/*
func Trace2() string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return fmt.Sprintf("%s:%d/%s", filepath.Base(frame.File), frame.Line, frame.Function)
}
*/
