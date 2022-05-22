package router

import (
	"fmt"
	"net"
	"strings"

	"mmesh.dev/m-api-go/grpc/network/mmnp/routing"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/xlog"
)

func (r *router) UpdateRoutingTable(newRT *routing.RoutingTable) {
	r.rib.Lock()
	defer r.rib.Unlock()

	r.rib.global = newRT

	if r.localForwarding {
		go r.routesImport()
	}

	go r.connectRelays()
}

func (r *router) routesImport() {
	if !r.localForwarding || r.networkInterface == nil {
		return
	}

	r.routes.Lock()
	defer r.routes.Unlock()

	if err := r.updateLocalRoutes(); err != nil {
		xlog.Alertf("Unable to check configured routes: %v", err)
		return
	}

	oldRoutes := make(map[cidrIPDst]struct{})

	for ipDst, valid := range r.routes.local {
		if valid {
			oldRoutes[ipDst] = struct{}{}
		}
	}

	r.rib.RLock()
	defer r.rib.RUnlock()

	for route, re := range r.rib.global.RE {
		ipDst := cidrIPDst(route)
		if re.VRFID == r.vrfID || r.rib.global.Scope == routing.Scope_NETWORK {
			if re.Type == routing.RouteType_DYNAMIC {
				// only STATIC or CONNECTED routes are processed
				continue
			}

			if _, ok := oldRoutes[ipDst]; ok {
				delete(oldRoutes, ipDst)
			} else {
				if re.Type == routing.RouteType_STATIC {
					if !r.isValidRouteImport(route) {
						continue
					}
				}
				if err := r.networkInterface.routeAdd(ipDst); err != nil {
					xlog.Errorf("Unable to add route %s: %v", r, err)
					continue
				}
				r.routes.local[ipDst] = true
			}
		}
	}

	for ipDst := range oldRoutes {
		if err := r.networkInterface.routeDel(ipDst); err != nil {
			xlog.Alertf("Unable to delete route %s: %v", r, err)
			continue
		}
		delete(r.routes.local, ipDst)
	}
}

func (r *router) isValidRouteImport(route string) bool {
	for _, exportedRoute := range r.routes.exported {
		if exportedRoute == route {
			return false
		}
	}

	if strings.HasPrefix(route, "0.0.0.0") {
		return false
	}

	if !isValidRoute(route) {
		return false
	}

	for _, r := range r.routes.imported {
		if strings.ToLower(r) == "any" || r == route {
			return true
		}
	}

	return false
}

func isValidRoute(route string) bool {
	// check if route is directly connected
	connectedRoutes, err := getConnectedRoutes()
	if err != nil {
		xlog.Alertf("Unable to get connected routes: %v", errors.Cause(err))
		return false
	}

	for _, ipNet := range connectedRoutes {
		xlog.Tracef("Checking connected route %s", ipNet.String())

		if ipNet.String() == route {
			return false
		}
	}

	return true
}

func getConnectedRoutes() ([]*net.IPNet, error) {
	connectedRoutes := make([]*net.IPNet, 0)

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function net.InterfaceAddrs()", errors.Trace())
	}

	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok {
			connectedRoutes = append(connectedRoutes, ipNet)
		}
	}

	return connectedRoutes, nil
}

func (r *router) getNetHop(addr string) (*routing.NetHop, *int32, error) {
	r.rib.RLock()
	defer r.rib.RUnlock()

	if r.rib == nil {
		r.rtrQueue <- struct{}{}
		return nil, nil, fmt.Errorf("no routing entry for %s", addr)
	}

	// get route ipDst
	ipDst := r.getIPDstFromRIB(addr)
	if len(ipDst) == 0 {
		r.rtrQueue <- struct{}{}
		return nil, nil, fmt.Errorf("no routing entry for %s", addr)
	}

	var nh *routing.NetHop
	var p int32
	for prio, nHop := range r.rib.global.RE[ipDst].Gw {
		if nHop.Healthy {
			if prio < p || p == 0 {
				p = prio
				nh = nHop
			}
		}
	}

	if nh == nil {
		return nil, nil, fmt.Errorf("no routing entry for %s", addr)
	}

	return nh, &p, nil
}

func (r *router) setNetHopUnhealthy(ipDst string, prio int32) {
	r.rib.Lock()
	defer r.rib.Unlock()

	if re, ok := r.rib.global.RE[ipDst]; ok {
		if nHop, ok := re.Gw[prio]; ok {
			nHop.Healthy = false
		}
	}
}

func (r *router) getIPDstFromRIB(addr string) string {
	// try first internal routes
	ipv4Dst := addr + "/32"
	if re, ok := r.rib.global.RE[ipv4Dst]; ok {
		if re.VRFID == r.vrfID || r.rib.global.Scope == routing.Scope_NETWORK {
			return ipv4Dst
		}
	}
	ipv6Dst := addr + "/128"
	if re, ok := r.rib.global.RE[ipv6Dst]; ok {
		if re.VRFID == r.vrfID || r.rib.global.Scope == routing.Scope_NETWORK {
			return ipv6Dst
		}
	}

	// then static routes
	for ipDst, re := range r.rib.global.RE {
		_, netCIDR, err := net.ParseCIDR(ipDst)
		if err != nil {
			xlog.Alertf("Detected invalid route %s (please check your configs): %v", ipDst, err)
			continue
		}

		if netCIDR.Contains(net.ParseIP(addr)) {
			if re.VRFID == r.vrfID || r.rib.global.Scope == routing.Scope_NETWORK {
				return ipDst
			}
		}
	}

	return ""
}
