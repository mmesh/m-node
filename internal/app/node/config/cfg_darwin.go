//go:build darwin
// +build darwin

package config

func defaultInterfaceName() string {
	return "utun2"
}
