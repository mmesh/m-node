//go:build darwin
// +build darwin

package cmd

func ConsoleInit() error {
	return nil
}

func defaultConfigFile() string {
	return "/opt/mmesh/etc/mmesh-node.yml"
}

/*
func logFile() string {
	return "/opt/mmesh/var/log/mmesh-node.log"
}
*/
