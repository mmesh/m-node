package connection

import (
	"os"

	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/network/mmnp/natProbe"
	nrpc "mmesh.dev/m-api-go/grpc/network/rpc"
	"mmesh.dev/m-node/internal/app/node/netp2p"
	"x6a.dev/pkg/errors"
	"x6a.dev/pkg/xlog"
)

// agentConnect connects the agent to mmesh network
func agentConnect(nxnc nrpc.NetworkAPIClient) error {
	if netp2p.AgentConfigured() {
		netp2p.UpdateNetworkClient(nxnc)
	} else {
		var err error

		// hiddenNode := viper.GetBool("agent.hidden")
		cfgExternalIPv4 := viper.GetString("agent.externalIPv4") // could be ""

		natp := &natProbe.NATProbe{
			Port:         int32(viper.GetInt("agent.port")),
			ExternalIPv4: cfgExternalIPv4,
		}

		natp, err = agentNATProbe(nxnc, natp)
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

		if err = netp2p.InitMMAgent(natp.ExternalIPv4, nxnc); err != nil {
			xlog.Alertf("Unable to initialize agent: %v", errors.Cause(err))
			os.Exit(1)
		}
	}

	return nil
}
