package config

import (
	"os"

	"github.com/spf13/viper"
	"mmesh.dev/mmesh/internal/pkg/logging"
	"mmesh.dev/mmesh/internal/pkg/mmid"
	"mmesh.dev/mmesh/internal/pkg/version"
	"x6a.dev/pkg/msg"
	"x6a.dev/pkg/xlog"
)

func Init() {
	hostID, err := os.Hostname()
	if err != nil {
		msg.Error(err)
		os.Exit(1)
	}

	accountID := viper.GetString("node.account")
	if len(accountID) == 0 {
		msg.Error("Invalid configuration: missing node.account")
		os.Exit(1)
	}
	tenantID := viper.GetString("node.tenant")
	if len(tenantID) == 0 {
		msg.Error("Invalid configuration: missing node.tenant")
		os.Exit(1)
	}
	netID := viper.GetString("node.network")
	if len(netID) == 0 {
		msg.Error("Invalid configuration: missing node.network")
		os.Exit(1)
	}
	vrfID := viper.GetString("node.subnet")
	if len(vrfID) == 0 {
		msg.Error("Invalid configuration: missing node.subnet")
		os.Exit(1)
	}
	nodeID := viper.GetString("node.id")
	if len(nodeID) == 0 {
		msg.Error("Invalid configuration: missing node.id")
		os.Exit(1)
	}

	mmID := mmid.NewMMNodeID(accountID, tenantID, netID, vrfID, nodeID)

	viper.Set("mm.id", mmID.String())

	viper.Set("mm.app", version.AGENT_NAME)

	viper.Set("host.id", hostID)

	setDefaults(nodeID)

	// logger config
	logging.LogLevel = xlog.GetLogLevel(viper.GetString("loglevel"))
	if logging.LogLevel == -1 {
		logging.LogLevel = xlog.INFO
	}

	logging.Interactive = false

	logLevel := logging.LogLevel

	if viper.GetBool("verbose") {
		logLevel = xlog.TRACE
	}

	xlog.SetLogger(logLevel, hostID)

	xlog.Infof("Initialized %s %s", version.AGENT_NAME, version.GetVersion())
}

func setDefaults(nodeID string) {
	dnsName := viper.GetString("endpoint.dnsName")
	if len(dnsName) == 0 {
		viper.Set("endpoint.dnsName", nodeID)
	}

	ipv4 := viper.GetString("endpoint.ipv4")
	if len(ipv4) == 0 {
		viper.Set("endpoint.ipv4", "auto")
	}

	port := int32(viper.GetInt("agent.port"))
	if port == 0 {
		viper.Set("agent.port", int(57775))
	}

	dnsPort := viper.GetInt("agent.dns.port")
	if dnsPort == 0 {
		viper.Set("agent.dns.port", int(53))
	}

	ifaceName := viper.GetString("agent.iface")
	if len(ifaceName) == 0 {
		viper.Set("agent.iface", "mmesh0")
	}
}
