package connection

import (
	"time"

	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/resources/iam/auth"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/grpc/client"
	"mmesh.dev/m-lib/pkg/xlog"
)

func newConnection() *connection {
	c := &connection{}
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

	for c.nxnc == nil || err != nil {
		c.nxnc, c.grpcClientConn, err = client.NewNetworkAPIClient(endpoint, authKey, authSecret)
		if err != nil {
			xlog.Errorf("Unable to connect to controller %s: %v", endpoint, errors.Cause(err))

			connectionFailed = true
			fc.setUnhealthy(endpoint)

			endpoint = fc.endpoint()
			xlog.Infof("Reconnecting to controller %s...", endpoint)

			time.Sleep(time.Second)
			continue
		}

		if err := fc.update(c.nxnc); err != nil {
			xlog.Errorf("Unable to get federation controllers: %v", errors.Cause(err))
		} else if !connectionFailed {
			// get the least crowded federation controller endpoint
			e := fc.endpoint()
			if endpoint != e {
				xlog.Infof("Found less loaded controller %s, reconnecting...", e)

				endpoint = e

				if err := c.grpcClientConn.Close(); err != nil {
					xlog.Errorf("Unable to close gRPC network connection: %v", err)
				}
				c.nxnc = nil
				continue
			}
		}

		if err = c.newSession(); err != nil {
			xlog.Errorf("Unable to create a network session: %v", errors.Cause(err))
			time.Sleep(5 * time.Second)
			continue
		}
	}

	xlog.Info("Node CONNECTED :-)")

	return c
}
