//go:build darwin
// +build darwin

package kvstore

func dbDir() string {
	return "/var/local/mmesh/db"
}
