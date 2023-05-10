package connection

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"mmesh.dev/m-api-go/grpc/resources/controller"
	"mmesh.dev/m-api-go/grpc/resources/topology"
	"mmesh.dev/m-api-go/grpc/rpc"
	"mmesh.dev/m-lib/pkg/errors"
)

type federationConnection struct {
	node               *topology.Node
	controllerHost     string
	controllerEndpoint string
	controllers        map[string]*controller.Controller // map[controllerID]*controller.Controller
	healthy            map[string]bool                   // map[endpoint]bool
	sync.RWMutex
}

var fc *federationConnection

func (c *connection) newFederationConnection() *federationConnection {
	return &federationConnection{
		node:               c.node,
		controllerHost:     strings.Split(c.defaultControllerEndpoint, ":")[0],
		controllerEndpoint: c.defaultControllerEndpoint,
		controllers:        make(map[string]*controller.Controller),
		healthy:            make(map[string]bool),
	}
}

func (f *federationConnection) update(nxnc rpc.NetworkAPIClient) error {
	f.Lock()
	defer f.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	nr := &topology.NodeReq{
		AccountID: f.node.AccountID,
		TenantID:  f.node.TenantID,
		NetID:     f.node.NetID,
		SubnetID:  f.node.SubnetID,
		NodeID:    f.node.NodeID,
	}

	fe, err := nxnc.FederationEndpoints(ctx, nr)
	if err != nil {
		f.healthy[f.controllerEndpoint] = false
		return errors.Wrapf(err, "[%v] function nxnc.FederationEndpoints()", errors.Trace())
	}

	f.healthy[f.controllerEndpoint] = true
	f.controllers = fe.Controllers

	return nil
}

func (f *federationConnection) endpoint() string {
	f.Lock()
	defer f.Unlock()

	if len(f.controllers) == 0 {
		return f.controllerEndpoint
	}

	currentControllerEndpoint := f.controllerEndpoint

	var connections int32

	for _, c := range f.controllers {
		e := fmt.Sprintf("%s:%d", c.Host, c.Port)

		// current controller is not healthy
		if currentControllerEndpoint == e && !f.healthy[e] {
			continue
		}

		tm := time.Unix(c.Status.LastUpdated, 0)
		if time.Since(tm) > 420*time.Second {
			f.healthy[e] = false
			continue
		} else {
			f.healthy[e] = true
		}

		if connections == 0 || c.Status.Connections < connections {
			connections = c.Status.Connections
			f.controllerEndpoint = e
			f.controllerHost = c.Host
		}
	}

	return f.controllerEndpoint
}

func (f *federationConnection) host() string {
	return f.controllerHost
}

func (f *federationConnection) setUnhealthy(endpoint string) {
	f.Lock()
	defer f.Unlock()

	f.healthy[endpoint] = false
}

func FederationUpdate(nxnc rpc.NetworkAPIClient) error {
	if fc == nil {
		return nil
	}

	return fc.update(nxnc)
}
