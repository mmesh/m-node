//go:build darwin
// +build darwin

package config

import "mmesh.dev/m-lib/pkg/xlog"

func defaultInterfaceName() string {
	return "utun2"
}

func setLogger(level xlog.LogLevel, hostID string) {
	xlog.Logger().SetLogLevel(level).SetHostID(hostID).SetANSIColor(true)
}
