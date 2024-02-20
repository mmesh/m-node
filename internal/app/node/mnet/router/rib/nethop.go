package rib

import (
	"fmt"
	"net/netip"

	"mmesh.dev/m-api-go/grpc/network/nac"
	"mmesh.dev/m-api-go/grpc/network/routing"
	"mmesh.dev/m-lib/pkg/xlog"
)

func (r *ribData) CheckIPDst(addr *netip.Addr) error {
	r.RLock()
	defer r.RUnlock()

	// get route ipDst
	ipDst := getIPDstFromRIB(addr, r.rib)
	if len(ipDst) == 0 {
		return fmt.Errorf("no routing entry for %s", addr)
	}

	return nil
}

func (r *ribData) GetNetHop(addr *netip.Addr) (*routing.NetHop, error) {
	r.RLock()
	defer r.RUnlock()

	// get route ipDst
	ipDst := getIPDstFromRIB(addr, r.rib)
	if len(ipDst) == 0 {
		return nil, fmt.Errorf("no routing entry for %s", addr)
	}

	re, ok := r.rib.RoutingTable[ipDst]
	if !ok {
		return nil, fmt.Errorf("no routing entry for %s", addr)
	}

	var netHop *routing.NetHop
	var prio int32

	for _, nh := range re.Gw {
		if nh.Priority < prio || prio == 0 {
			prio = nh.Priority
			netHop = nh
		}
	}

	if netHop == nil {
		return nil, fmt.Errorf("no routing entry for %s", addr)
	}

	return netHop, nil
}

func getIPDstFromRIB(addr *netip.Addr, r *routing.RIB) string {
	// try first internal routes
	ipv4Dst := addr.String() + "/32"
	if re, ok := r.RoutingTable[ipv4Dst]; ok {
		if re.SubnetID == r.RoutingDomain.SubnetID || r.RoutingDomain.Scope == nac.RoutingScope_NETWORK {
			return ipv4Dst
		}
	}
	ipv6Dst := addr.String() + "/128"
	if re, ok := r.RoutingTable[ipv6Dst]; ok {
		if len(re.SubnetID) == 0 && re.Type == routing.RouteType_PROXY {
			return ipv6Dst
		}

		if re.SubnetID == r.RoutingDomain.SubnetID || r.RoutingDomain.Scope == nac.RoutingScope_NETWORK {
			return ipv6Dst
		}
	}

	// then static routes
	for ipDst, re := range r.RoutingTable {
		netCIDR, err := netip.ParsePrefix(ipDst)
		if err != nil {
			xlog.Alertf("Detected invalid route %s (please check your configs): %v", ipDst, err)
			continue
		}

		if netCIDR.Contains(*addr) {
			if re.SubnetID == r.RoutingDomain.SubnetID || r.RoutingDomain.Scope == nac.RoutingScope_NETWORK {
				return ipDst
			}
		}
	}

	return ""
}
