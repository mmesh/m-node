package connection

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/spf13/viper"
	"mmesh.dev/api-go/grpc/network/resources/controller"
	"mmesh.dev/api-go/grpc/network/resources/network"
	"mmesh.dev/api-go/grpc/network/rpc"
	"mmesh.dev/mmesh/internal/pkg/mmid"
	"x6a.dev/pkg/errors"
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

func (f *federationConnection) update(nxnc rpc.NxNetworkClient) error {
	f.Lock()
	defer f.Unlock()

	fe, err := nxnc.FederationEndpoints(context.TODO(), f.node)
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

	var connections int32
	var controller *controller.Controller

	for _, c := range f.controllers {
		e := fmt.Sprintf("%s:%d", c.Host, c.Port)

		if !f.healthy[e] && f.controllerEndpoint == e {
			continue
		}

		tm := time.Unix(c.Status.LastUpdated, 0)
		if time.Since(tm) > 420*time.Second {
			f.healthy[e] = false
		} else {
			f.healthy[e] = true
		}

		if connections == 0 || c.Status.RoutedNodes < connections {
			connections = c.Status.RoutedNodes
			controller = c
		}
	}

	f.controllerEndpoint = fmt.Sprintf("%s:%d", controller.Host, controller.Port)

	return f.controllerEndpoint
}

func (f *federationConnection) setUnhealthy(endpoint string) {
	f.Lock()
	defer f.Unlock()

	f.healthy[endpoint] = false
}

func FederationUpdate(nxnc rpc.NxNetworkClient) error {
	if fc == nil {
		fc = newFederationConnection()
	}

	return fc.update(nxnc)
}
