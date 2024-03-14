package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"mmesh.dev/m-lib/pkg/logging"
	"mmesh.dev/m-lib/pkg/utils/msg"
	"mmesh.dev/m-lib/pkg/version"
	"mmesh.dev/m-lib/pkg/xlog"
)

func Init() {
	hostID, err := os.Hostname()
	if err != nil {
		msg.Error(err)
		os.Exit(1)
	}

	nodeToken := viper.GetString("token")
	if len(nodeToken) == 0 {
		msg.Error("Authorization token not found")
		os.Exit(1)
	}

	viper.Set("mm.app", version.NODE_NAME)

	viper.Set("host.id", hostID)

	setDefaults()

	// logger config
	logging.LogLevel = xlog.GetLogLevel(viper.GetString("loglevel"))
	if logging.LogLevel == -1 {
		logging.LogLevel = xlog.INFO
	}

	logging.Interactive = false

	logLevel := logging.LogLevel

	xlog.Logger().SetLogLevel(logLevel).SetHostID(hostID).SetANSIColor(true)

	fmt.Print("[settings loaded]\n\n")
}

func setDefaults() {
	ifaceName := viper.GetString("iface")
	if len(ifaceName) == 0 {
		viper.Set("iface", defaultInterfaceName())
	}

	port := viper.GetInt("port")
	if port == 0 {
		viper.Set("port", int(57775))
	}

	dnsPort := viper.GetInt("dnsPort")
	if dnsPort == 0 {
		viper.Set("dnsPort", int(53535))
	}
}
