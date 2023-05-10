//go:build linux
// +build linux

package tss

func tssDir() string {
	return "/var/lib/mmesh"
}
