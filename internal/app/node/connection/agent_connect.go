package connection

import (
	"os"

	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/network/mmnp/natProbe"
	"mmesh.dev/m-api-go/grpc/rpc"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/xlog"
	"mmesh.dev/m-node/internal/app/node/netp2p"
)

// agentConnect connects the agent to mmesh network
func agentConnect(nxnc rpc.NetworkAPIClient) error {
	// hiddenNode := viper.GetBool("agent.hidden")
	cfgExternalIPv4 := viper.GetString("agent.externalIPv4") // could be ""

	if netp2p.AgentConfigured() {
		// netp2p.UpdateNetworkClient(nxnc)

		if err := netp2p.InitMMAgent(cfgExternalIPv4, nxnc); err != nil {
			xlog.Alertf("Unable to initialize agent: %v", errors.Cause(err))
			os.Exit(1)
		}

		return nil
	}

	natp := &natProbe.NATProbe{
		Port:         int32(viper.GetInt("agent.port")),
		ExternalIPv4: cfgExternalIPv4,
	}

	natp, err := agentNATProbe(nxnc, natp)
	if err != nil {
		xlog.Alertf("Unable to initialize agent: %v", errors.Cause(err))
		os.Exit(1)
	}

	// configured externalIPv4 overrides controller's detected externalIPv4
	if len(cfgExternalIPv4) > 0 {
		natp.ExternalIPv4 = cfgExternalIPv4
	}

	// if hiddenNode {
	// 	natp.ExternalIPv4 = ""
	// }

	if err := netp2p.InitMMAgent(natp.ExternalIPv4, nxnc); err != nil {
		xlog.Alertf("Unable to initialize agent: %v", errors.Cause(err))
		os.Exit(1)
	}

	return nil
}
