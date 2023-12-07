//go:build linux
// +build linux

package config

import (
	"mmesh.dev/m-lib/pkg/xlog"
)

func defaultInterfaceName() string {
	return "mmesh0"
}

func setLogger(level xlog.LogLevel, hostID string) {
	xlog.Logger().SetLogLevel(level).SetHostID(hostID).SetANSIColor(true)
}
