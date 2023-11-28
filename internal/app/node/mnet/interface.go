package mnet

import (
	"sync"

	"mmesh.dev/m-api-go/grpc/network/routing"
	"mmesh.dev/m-api-go/grpc/resources/topology"
	"mmesh.dev/m-node/internal/app/node/mnet/connection"
	"mmesh.dev/m-node/internal/app/node/mnet/router"
)

type LocalNodeInterface interface {
	Connection() connection.Interface
	Router() router.Interface
	AddNetworkEndpoint(endpointID, dnsName, reqIPv4 string) (string, error)
	RemoveNetworkEndpoint(endpointID string) error
	GetNodeLSA() *routing.LSA
	NodeReq() *topology.NodeReq
	Node() *topology.Node
	DNSPort() int
	IsK8sGwEnabled() bool
	Close()
}

type localNode struct {
	node        *topology.Node
	endpoints   *endpointsMap
	connection  connection.Interface
	router      router.Interface
	initialized bool
}

type endpointsMap struct {
	endpt map[string]*topology.Endpoint
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

func (ln *localNode) DNSPort() int {
	return int(ln.node.Agent.DNSPort)
}

func (ln *localNode) IsK8sGwEnabled() bool {
	return ln.node.Type == topology.NodeType_K8S_GATEWAY
}

func (ln *localNode) Close() {
	ln.Connection().Close()

	if ln.Node().Cfg.DisableNetworking || ln.Router() == nil {
		return
	}

	ln.Router().Disconnect()
}
