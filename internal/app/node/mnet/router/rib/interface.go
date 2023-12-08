package rib

import (
	"net/netip"
	"sync"

	"mmesh.dev/m-api-go/grpc/network/nac"
	"mmesh.dev/m-api-go/grpc/network/routing"
	"mmesh.dev/m-api-go/grpc/resources/topology"
)

type Interface interface {
	Initialize(mq <-chan []byte, n *topology.Node, rd *nac.RoutingDomain, np *topology.Policy)
	Close()
	RouteEventQueue() <-chan *RouteEvent
	RelayConnQueue() <-chan *routing.NetHop
	RouterConnQueue() <-chan *routing.NetHop
	RoutingDomain() *nac.RoutingDomain
	GetPolicy(subnetID string) *topology.Policy
	DNSQuery(dnsName string) (ipv4, ipv6 string)
	CheckIPDst(addr *netip.Addr) error
	GetNetHop(addr *netip.Addr) (*routing.NetHop, error)
	GetRelayMAddrs(nh *routing.NetHop) []string
	GetRouterMAddrs(nh *routing.NetHop) []string

	wrkr(msgQueue <-chan []byte)
	cleanup()
	cleanupNetHops(nhm map[string]*routing.NetHop)
	cleanupRoutingTable(rt map[string]*routing.RoutingEntry)
	decoder(msg string) error
	processor(d *routing.RIBData)
	setRoutingScope(scope nac.RoutingScope)
	setRouter(router *routing.NetHop)
	setRelay(relay *routing.NetHop)
	setRoutingTable(rt map[string]*routing.RoutingEntry)
	setPolicy(npm map[string]*topology.Policy)
}
type ribData struct {
	rxQueue chan *routing.RIBData
	closeCh chan struct{}
	rib     *routing.RIB
	sync.RWMutex
}

func New() Interface {
	return &ribData{
		rxQueue: make(chan *routing.RIBData, 256),
		closeCh: make(chan struct{}, 1),
		rib: &routing.RIB{
			AccountID:     "",
			TenantID:      "",
			NetID:         "",
			RoutingDomain: nil,
			Routers:       make(map[string]*routing.NetHop, 0),
			Relays:        make(map[string]*routing.NetHop, 0),
			RoutingTable:  make(map[string]*routing.RoutingEntry, 0),
			Policy:        make(map[string]*topology.Policy, 0),
		},
	}
}

func (r *ribData) Initialize(mq <-chan []byte, n *topology.Node, rd *nac.RoutingDomain, np *topology.Policy) {
	r.Lock()

	r.rib.AccountID = n.AccountID
	r.rib.TenantID = n.TenantID
	r.rib.NetID = n.Cfg.NetID
	r.rib.RoutingDomain = rd
	r.rib.Policy[n.Cfg.SubnetID] = np

	r.Unlock()

	r.setRoutingScope(rd.Scope)

	go r.wrkr(mq)
}

func (r *ribData) Close() {
	r.closeCh <- struct{}{}
}

func (r *ribData) RouteEventQueue() <-chan *RouteEvent {
	return routeEventQueue
}

func (r *ribData) RelayConnQueue() <-chan *routing.NetHop {
	return relayConnQueue
}

func (r *ribData) RouterConnQueue() <-chan *routing.NetHop {
	return routerConnQueue
}

func (r *ribData) RoutingDomain() *nac.RoutingDomain {
	r.RLock()
	defer r.RUnlock()

	return r.rib.RoutingDomain
}

func (r *ribData) GetPolicy(subnetID string) *topology.Policy {
	r.RLock()
	defer r.RUnlock()

	if p, ok := r.rib.Policy[subnetID]; ok {
		return p
	}

	return nil
}

/*
func (r *ribData) RoutingScope() nac.RoutingScope {
	r.RLock()
	defer r.RUnlock()

	if r.rib.RoutingDomain == nil {
		return nac.RoutingScope_SUBNET
	}

	return r.rib.RoutingDomain.Scope
}

func (r *ribData) Routers() map[string]*routing.NetHop {
	r.RLock()
	defer r.RUnlock()

	return r.rib.Routers
}

func (r *ribData) Relays() map[string]*routing.NetHop {
	r.RLock()
	defer r.RUnlock()

	return r.rib.Relays
}

func (r *ribData) RoutingTable() map[string]*routing.RoutingEntry {
	r.RLock()
	defer r.RUnlock()

	return r.rib.RoutingTable
}

func (r *ribData) Policy() map[string]*topology.Policy {
	r.RLock()
	defer r.RUnlock()

	return r.rib.Policy
}
*/
