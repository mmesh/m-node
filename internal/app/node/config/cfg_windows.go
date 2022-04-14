//go:build windows
// +build windows

package config

import (
	"os"

	"mmesh.dev/m-lib/pkg/xlog"
)

func defaultInterfaceName() string {
	return "mmesh"
}

func setLogger(level xlog.LogLevel, hostID string) {
	if len(os.Args) == 2 {
		if os.Args[1] == "service-start" {
			xlog.Logger().SetLogLevel(level).SetHostID(hostID).SetStdLogger()
			return
		}
	}

	xlog.Logger().SetLogLevel(level).SetHostID(hostID).SetANSIColor(true)
}
