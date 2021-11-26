package netp2p

import (
	"fmt"
	"net"
	"strings"
	"sync"

	"mmesh.dev/m-api-go/grpc/network/mmnp/routing"
	"x6a.dev/pkg/errors"
	"x6a.dev/pkg/xlog"
)

var rtrq = make(chan struct{}, 64)

type routeMap struct {
	routes map[string]bool
	sync.RWMutex
}

var configuredRoutes *routeMap

func newRouteMap() *routeMap {
	return &routeMap{
		routes: make(map[string]bool),
	}
}

func GetRTRQueue() chan struct{} {
	return rtrq
}

func UpdateRoutingTable(newRT *routing.RoutingTable) {
	mma.rt.Lock()
	defer mma.rt.Unlock()

	mma.rt.RT = newRT

	go mma.routesImport()

	go mma.connectRelays()
}

func (mma *mmAgent) routesImport() {
	if configuredRoutes == nil {
		configuredRoutes = newRouteMap()
	}

	configuredRoutes.Lock()
	defer configuredRoutes.Unlock()

	if err := configuredRoutes.update(); err != nil {
		xlog.Alertf("Unable to check configured routes: %v", err)
		return
	}

	oldRoutes := make(map[string]struct{})

	for r, valid := range configuredRoutes.routes {
		if valid {
			oldRoutes[r] = struct{}{}
		}
	}

	mma.rt.RLock()
	defer mma.rt.RUnlock()

	for r, re := range mma.rt.RT.RE {
		if re.VRFID == mma.vrfID || mma.rt.RT.Scope == routing.Scope_NETWORK {
			if re.Type == routing.RouteType_DYNAMIC {
				// only STATIC or CONNECTED routes are processed
				continue
			}

			if _, ok := oldRoutes[r]; ok {
				delete(oldRoutes, r)
			} else {
				if re.Type == routing.RouteType_STATIC {
					if !mma.validRouteImport(r) {
						continue
					}
				}
				if err := routeAdd(r); err != nil {
					xlog.Errorf("Unable to add route %s: %v", r, err)
					continue
				}
				configuredRoutes.routes[r] = true
			}
		}
	}

	for r := range oldRoutes {
		if err := routeDel(r); err != nil {
			xlog.Alertf("Unable to delete route %s: %v", r, err)
			continue
		}
		delete(configuredRoutes.routes, r)
	}
}

func (mma *mmAgent) validRouteImport(route string) bool {
	for _, r := range mma.agent.Routes.Export {
		if r == route {
			return false
		}
	}

	if strings.HasPrefix(route, "0.0.0.0") {
		return false
	}

	if !validRoute(route) {
		return false
	}

	for _, r := range mma.agent.Routes.Import {
		if strings.ToLower(r) == "any" || r == route {
			return true
		}
	}

	return false
}

func validRoute(route string) bool {
	// check if route is directly connected
	connectedRoutes, err := getConnectedRoutes()
	if err != nil {
		xlog.Alertf("Unable to get connected routes: %v", errors.Cause(err))
		return false
	}

	for _, r := range connectedRoutes {
		xlog.Tracef("Checking connected route %s", r.String())

		if r.String() == route {
			return false
		}
	}

	return true
}

func getConnectedRoutes() ([]*net.IPNet, error) {
	var connectedRoutes []*net.IPNet

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

func (mma *mmAgent) connectRelays() {
	mma.rt.RLock()
	defer mma.rt.RUnlock()

	if mma.rt != nil {
		for _, r := range mma.rt.RT.Relays {
			if r.VRFID == mma.vrfID || mma.rt.RT.Scope == routing.Scope_NETWORK {
				if strings.Contains(r.MAddr, ":") {
					// relays are ipv4 only
					continue
				}
				peerInfo, err := getPeerInfo(r.MAddr)
				if err != nil {
					xlog.Errorf("Unable to get peer info: %v", err)
					continue
				}
				go mma.connectPeer(peerInfo)
			}
		}
	}
}

func (mma *mmAgent) getNetHop(addr string) (*routing.NetHop, *int32, error) {
	mma.rt.RLock()
	defer mma.rt.RUnlock()

	if mma.rt == nil {
		rtrq <- struct{}{}
		return nil, nil, fmt.Errorf("no routing entry found for addr %v", addr)
	}

	// get route
	r := mma.getRoute(addr)
	if len(r) == 0 {
		rtrq <- struct{}{}
		return nil, nil, fmt.Errorf("no routing entry found for addr %v", addr)
	}

	var nh *routing.NetHop
	var p int32
	for prio, nHop := range mma.rt.RT.RE[r].Gw {
		if nHop.Healthy {
			if prio < p || p == 0 {
				p = prio
				nh = nHop
			}
		}
	}

	if nh == nil {
		return nil, nil, fmt.Errorf("no routing entry available for addr %v", addr)
	}

	return nh, &p, nil
}

func (mma *mmAgent) setNxHopUnhealthy(addr string, prio int32) {
	mma.rt.Lock()
	defer mma.rt.Unlock()

	if re, ok := mma.rt.RT.RE[addr]; ok {
		if nxh, ok := re.Gw[prio]; ok {
			nxh.Healthy = false
		}
	}
}

func (mma *mmAgent) getRoute(addr string) string {
	// try first internal routes
	addrv4 := addr + "/32"
	if re, ok := mma.rt.RT.RE[addrv4]; ok {
		if re.VRFID == mma.vrfID || mma.rt.RT.Scope == routing.Scope_NETWORK {
			return addrv4
		}
	}
	addrv6 := addr + "/128"
	if re, ok := mma.rt.RT.RE[addrv6]; ok {
		if re.VRFID == mma.vrfID || mma.rt.RT.Scope == routing.Scope_NETWORK {
			return addrv6
		}
	}

	// then static routes
	for route, re := range mma.rt.RT.RE {
		_, netCIDR, err := net.ParseCIDR(route)
		if err != nil {
			xlog.Alertf("Detected invalid route %s (please check your configs): %v", route, err)
			continue
		}

		if netCIDR.Contains(net.ParseIP(addr)) {
			if re.VRFID == mma.vrfID || mma.rt.RT.Scope == routing.Scope_NETWORK {
				return route
			}
		}
	}

	return ""
}
