//go:build darwin
// +build darwin

package utils

func Ifconfig(sargs string) error {
	return execCommand("ifconfig", sargs)
}

func Route(sargs string) error {
	return execCommand("route", sargs)
}
