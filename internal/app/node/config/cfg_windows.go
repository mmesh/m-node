//go:build windows
// +build windows

package config

import (
	"fmt"
	"os"

	"mmesh.dev/m-lib/pkg/version"
	"mmesh.dev/m-lib/pkg/xlog"
)

func defaultInterfaceName() string {
	return "mmesh"
}

func logFile() string {
	programFiles := os.Getenv("ProgramFiles")

	if len(programFiles) == 0 {
		programFiles = `C:\Program Files`
	}

	return fmt.Sprintf(`%s\mmesh\mmesh-node.log`, programFiles)
}

func setLogger(level xlog.LogLevel, hostID string) {
	xlog.Logger().SetLogLevel(level).SetHostID(hostID).SetANSIColor(false)

	xlog.Logger().SetLogFile(logFile())

	xlog.Logger().SetWindowsLogger(&xlog.EventLogOptions{
		Level:  level,
		Source: version.NODE_NAME,
	})
}
