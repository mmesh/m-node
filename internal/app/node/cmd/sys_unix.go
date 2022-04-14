//go:build !windows
// +build !windows

package cmd

func ConsoleInit() error {
	return nil
}

func defaultConfigFile() string {
	return "/etc/mmesh/mmesh-node.yml"
}
