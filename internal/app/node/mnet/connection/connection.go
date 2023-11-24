package connection

import (
	"time"

	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/grpc/client"
	"mmesh.dev/m-lib/pkg/xlog"
)

func (c *connection) new() {
	for {
		if err := c.networkAdmissionRequest(); err != nil {
			xlog.Errorf("Unable to connect to network controller: %v", errors.Cause(err))
			xlog.Info("Retrying in 5s...")
			time.Sleep(5 * time.Second)
			continue
		}
		break
	}

	if fc == nil {
		fc = c.newFederationConnection()
	} else {
		fc.node = c.node
		fc.controllerEndpoint = c.defaultControllerEndpoint
	}

	endpoint := fc.endpoint()
	connectionFailed := false
	conns := 0
	c.nxnc = nil

	var err error

	for c.nxnc == nil || err != nil {
		c.nxnc, c.grpcClientConn, err = client.NewNetworkAPIClient(endpoint, c.authKey, c.authSecret)
		if err != nil {
			xlog.Errorf("Unable to connect to controller %s: %v", endpoint, errors.Cause(err))

			connectionFailed = true
			conns = 0
			fc.setUnhealthy(endpoint)

			endpoint = fc.endpoint()
			xlog.Infof("Reconnecting to controller %s...", endpoint)

			time.Sleep(time.Second)
			continue
		}

		if err := fc.update(c.nxnc); err != nil {
			xlog.Errorf("Unable to get federation controllers: %v", errors.Cause(err))
		} else if !connectionFailed && conns < 2 {
			conns++

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

		if !c.initialized {
			if err = c.newSession(); err != nil {
				xlog.Errorf("Unable to create a network session: %v", errors.Cause(err))
				time.Sleep(5 * time.Second)
				continue
			}
		}

		// test
		// c.node.Cfg.DisableNetworking = true // test
		// test

		if c.node.Cfg.KubernetesGw {
			c.node.Cfg.DisableNetworking = false
		}

		if !c.node.Cfg.DisableNetworking {
			if err = c.newRoutingClient(fc.host()); err != nil {
				xlog.Errorf("Unable to create a routing session: %v", errors.Cause(err))
				time.Sleep(1 * time.Second)
				continue
			}
		}
	}

	c.initialized = true

	xlog.Info("Node CONNECTED :-)")
}
