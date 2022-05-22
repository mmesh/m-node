package maddr

import "strings"

func IsRelay(maddr string) bool {
	return strings.Contains(maddr, "/p2p-circuit/")
}
