package mnet

import (
	"sync"

	"mmesh.dev/m-api-go/grpc/resources/network"
	"mmesh.dev/m-node/internal/app/node/mnet/connection"
	"mmesh.dev/m-node/internal/app/node/mnet/router"
)

type LocalNodeInterface interface {
	Connection() connection.Interface
	Router() router.Interface
	AddNetworkEndpoint(endpointID, dnsName, reqIPv4 string) (string, error)
	RemoveNetworkEndpoint(endpointID string) error
	NetworkNodeWithoutEndpoints() *network.Node
	NetworkNode() *network.Node
	IsK8sGwEnabled() bool
	Close()
}

type localNode struct {
	accountID  string
	tenantID   string
	netID      string
	vrfID      string
	nodeID     string
	agent      *network.Agent
	endpoints  *endpointsMap
	replicaSet bool
	connection connection.Interface
	router     router.Interface
}

type endpointsMap struct {
	endpt map[string]*network.NetworkEndpoint
	sync.RWMutex
}

var localnode *localNode

func LocalNode() LocalNodeInterface {
	return localnode
}

func (ln *localNode) Connection() connection.Interface {
	if ln == nil {
		return nil
	}

	return ln.connection
}

func (ln *localNode) Router() router.Interface {
	if ln == nil {
		return nil
	}

	return ln.router
}

func (ln *localNode) IsK8sGwEnabled() bool {
	return ln.agent.Options.KubernetesGw
}

func (ln *localNode) Close() {
	ln.Connection().Close()

	ln.Router().Disconnect()
}
