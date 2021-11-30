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

func setLogger(level xlog.LogLevel, hostID string, logOpts ...*xlog.LogOption) {
	if len(os.Args) == 2 {
		if os.Args[1] == "service-start" {
			xlog.SetLogger(level, hostID, xlog.WithStdLogger())
			return
		}
	}

	xlog.SetLogger(level, hostID, xlog.WithANSIColor(true))
}
