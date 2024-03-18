//go:build darwin
// +build darwin

package kvstore

func dbDir() string {
	return "/opt/mmesh/var/lib/db"
}
