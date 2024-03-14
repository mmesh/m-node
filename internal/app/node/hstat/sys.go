package hstat

import (
	"fmt"
	"runtime"

	"mmesh.dev/m-api-go/grpc/resources/topology"
)

func uptimeStr(uptime uint64) string {
	var s string

	days := uptime / (60 * 60 * 24)

	if days == 1 {
		s = fmt.Sprintf("%d day", days)
	} else {
		s = fmt.Sprintf("%d days", days)
	}

	minutes := uptime / 60
	hours := minutes / 60
	hours %= 24
	minutes %= 60

	return fmt.Sprintf("%s, %d hours, %02d minutes", s, hours, minutes)
}

func getOSType() topology.OSType {
	switch runtime.GOOS {
	case "linux":
		return topology.OSType_LINUX
	case "darwin":
		return topology.OSType_DARWIN
	case "windows":
		return topology.OSType_WINDOWS
	}

	return topology.OSType_UNKNOWN_OS
}
