package rib

import (
	"fmt"

	"mmesh.dev/m-api-go/grpc/network/nac"
	"mmesh.dev/m-api-go/grpc/network/routing"
	"mmesh.dev/m-api-go/grpc/resources/topology"
	"mmesh.dev/m-lib/pkg/ipnet"
	"mmesh.dev/m-lib/pkg/xlog"
)

func (r *ribData) processor(d *routing.RIBData) {
	switch d.Type {
	case routing.RIBDataType_ROUTING_SCOPE:
		r.setRoutingScope(d.Scope)
	case routing.RIBDataType_ROUTER:
		r.setRouter(d.Router)
	case routing.RIBDataType_RELAY:
		r.setRelay(d.Relay)
	case routing.RIBDataType_ROUTING_TABLE:
		r.setRoutingTable(d.RoutingTable)
	case routing.RIBDataType_POLICY:
		r.setPolicy(d.Policy)
	}
}

func (r *ribData) setRoutingScope(scope nac.RoutingScope) {
	r.Lock()
	defer r.Unlock()

	r.rib.RoutingDomain.Scope = scope

	rd := r.rib.RoutingDomain

	// set connected routes
	switch rd.Scope {
	case nac.RoutingScope_SUBNET:
		delete(r.rib.RoutingTable, rd.NetworkCIDR)
		evtDeleteRoute(rd.NetworkCIDR, routing.RouteType_CONNECTED)

		r.rib.RoutingTable[rd.SubnetCIDR] = newRoutingEntry()
		r.rib.RoutingTable[rd.SubnetCIDR].SubnetID = rd.SubnetID
		r.rib.RoutingTable[rd.SubnetCIDR].AddressFamily = routing.AddressFamily_IP4
		r.rib.RoutingTable[rd.SubnetCIDR].Type = routing.RouteType_CONNECTED
		evtAddRoute(rd.SubnetCIDR, routing.RouteType_CONNECTED)
	case nac.RoutingScope_NETWORK:
		delete(r.rib.RoutingTable, rd.SubnetCIDR)
		evtDeleteRoute(rd.SubnetCIDR, routing.RouteType_CONNECTED)

		r.rib.RoutingTable[rd.NetworkCIDR] = newRoutingEntry()
		r.rib.RoutingTable[rd.NetworkCIDR].SubnetID = rd.SubnetID
		r.rib.RoutingTable[rd.NetworkCIDR].AddressFamily = routing.AddressFamily_IP4
		r.rib.RoutingTable[rd.NetworkCIDR].Type = routing.RouteType_CONNECTED
		evtAddRoute(rd.NetworkCIDR, routing.RouteType_CONNECTED)
	}

	ipnetv6 := fmt.Sprintf("%s:/32", ipnet.IPv6Prefix())
	if _, ok := r.rib.RoutingTable[ipnetv6]; !ok {
		r.rib.RoutingTable[ipnetv6] = newRoutingEntry()
		r.rib.RoutingTable[ipnetv6].SubnetID = rd.SubnetID
		r.rib.RoutingTable[ipnetv6].AddressFamily = routing.AddressFamily_IP6
		r.rib.RoutingTable[ipnetv6].Type = routing.RouteType_CONNECTED
		evtAddRoute(ipnetv6, routing.RouteType_CONNECTED)
	}
}

func (r *ribData) setRouter(router *routing.NetHop) {
	r.Lock()
	defer r.Unlock()

	r.rib.Routers[router.P2PHostID] = router

	if len(router.SubnetID) == 0 {
		evtRouter(router)
	}
}

func (r *ribData) setRelay(relay *routing.NetHop) {
	r.Lock()
	defer r.Unlock()

	r.rib.Relays[relay.P2PHostID] = relay

	if relay.SubnetID == r.rib.RoutingDomain.SubnetID || r.rib.RoutingDomain.Scope == nac.RoutingScope_NETWORK {
		evtRelay(relay)
	}
}

func (r *ribData) setRoutingTable(rt map[string]*routing.RoutingEntry) {
	r.Lock()
	defer r.Unlock()

	for addr, re := range rt {
		if _, ok := r.rib.RoutingTable[addr]; !ok {
			evtAddRoute(addr, re.Type)
		}

		r.rib.RoutingTable[addr] = re

		if re.Type == routing.RouteType_PROXY && len(r.appSvcs) > 0 {
			for _, nh := range re.Gw {
				evtProxy(nh)
			}
		}
	}
}

func (r *ribData) setPolicy(npm map[string]*topology.Policy) {
	r.Lock()
	defer r.Unlock()

	for subnetID, p := range npm {
		xlog.Infof("Updated network policy for %s", subnetID)
		r.rib.Policy[subnetID] = p
	}
}

func newRoutingEntry() *routing.RoutingEntry {
	return &routing.RoutingEntry{
		Gw: make(map[string]*routing.NetHop, 0), // map[p2pHostID]*routing.NetHop
	}
}
