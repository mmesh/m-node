package hstat

import (
	"fmt"
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
