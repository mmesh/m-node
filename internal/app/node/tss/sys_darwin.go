//go:build darwin
// +build darwin

package tss

func tssDir() string {
	return "/var/local/mmesh"
}
