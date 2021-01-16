package connection

import (
	"time"

	"github.com/spf13/viper"
	nrpc "mmesh.dev/m-api-go/grpc/network/rpc"
	"mmesh.dev/m-api-go/grpc/resources/iam/auth"
	"mmesh.dev/m-lib/pkg/grpc/client"
	"x6a.dev/pkg/errors"
	"x6a.dev/pkg/xlog"
)

func AgentConnect() nrpc.NetworkAPIClient {
	var nxnc nrpc.NetworkAPIClient
	var err error

	authKey := &auth.AuthKey{
		Key: viper.GetString("node.authKey"),
	}
	authSecret := viper.GetString("node.authSecret")

	if fc == nil {
		fc = newFederationConnection()
	}
	endpoint := fc.endpoint()
	connectionFailed := false

	for nxnc == nil || err != nil {
		nxnc, err = client.NewNetworkAPIClient(endpoint, authKey, authSecret)
		if err != nil {
			xlog.Errorf("Unable to connect to controller %s: %v", endpoint, errors.Cause(err))

			connectionFailed = true
			fc.setUnhealthy(endpoint)

			endpoint = fc.endpoint()
			xlog.Infof("Reconnecting to controller %s...", endpoint)

			time.Sleep(time.Second)
			continue
		}

		if err := fc.update(nxnc); err != nil {
			xlog.Errorf("Unable to get federation controllers: %v", errors.Cause(err))
		} else if !connectionFailed {
			// get the least crowded federation controller endpoint
			e := fc.endpoint()
			if endpoint != e {
				xlog.Infof("Found better (less loaded) controller %s, reconnecting...", e)

				endpoint = e

				if err := client.NxClientConn.Close(); err != nil {
					xlog.Alertf("Unable to close gRPC network connection: %v", err)
				}
				nxnc = nil
				continue
			}
		}

		if err = agentConnect(nxnc); err != nil {
			xlog.Errorf("Unable to register network agent: %v", errors.Cause(err))
			time.Sleep(5 * time.Second)
			continue
		}
	}

	xlog.Info("Node CONNECTED :-)")

	return nxnc
}
