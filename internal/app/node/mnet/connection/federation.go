package connection

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/resources/controller"
	"mmesh.dev/m-api-go/grpc/resources/network"
	"mmesh.dev/m-api-go/grpc/rpc"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/mmid"
)

type federationConnection struct {
	node               *network.Node
	controllerEndpoint string
	controllers        map[string]*controller.Controller // map[controllerID]*controller.Controller
	healthy            map[string]bool                   // map[endpoint]bool
	sync.RWMutex
}

var fc *federationConnection

func newFederationConnection() *federationConnection {
	mmID := viper.GetString("mm.id")

	return &federationConnection{
		node:               mmid.MMNodeID(mmID).Node(),
		controllerEndpoint: viper.GetString("controller.endpoint"),
		controllers:        make(map[string]*controller.Controller),
		healthy:            make(map[string]bool),
	}
}

func (f *federationConnection) update(nxnc rpc.NetworkAPIClient) error {
	f.Lock()
	defer f.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fe, err := nxnc.FederationEndpoints(ctx, f.node)
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

		if connections == 0 || c.Status.RoutedNodes < connections {
			connections = c.Status.RoutedNodes
			f.controllerEndpoint = e
		}
	}

	return f.controllerEndpoint
}

func (f *federationConnection) setUnhealthy(endpoint string) {
	f.Lock()
	defer f.Unlock()

	f.healthy[endpoint] = false
}

func FederationUpdate(nxnc rpc.NetworkAPIClient) error {
	if fc == nil {
		fc = newFederationConnection()
	}

	return fc.update(nxnc)
}
