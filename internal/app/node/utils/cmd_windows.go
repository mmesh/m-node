//go:build windows
// +build windows

package utils

func Netsh(sargs string) error {
	return execCommand("netsh", sargs)
}
