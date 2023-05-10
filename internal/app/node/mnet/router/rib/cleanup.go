package rib

import (
	"time"

	"mmesh.dev/m-api-go/grpc/network/routing"
	"mmesh.dev/m-lib/pkg/xlog"
)

const rtHoldTime = 180

func (r *ribData) cleanup() {
	r.cleanupNetHops(r.rib.Routers)
	r.cleanupNetHops(r.rib.Relays)
	r.cleanupRoutingTable(r.rib.RoutingTable)
}

func (r *ribData) cleanupNetHops(nhm map[string]*routing.NetHop) {
	r.Lock()
	defer r.Unlock()

	offlineNetHop := make(map[string]struct{})

	for p2pHostID, nh := range nhm {
		if netHopIsOnline(nh) {
			continue
		}
		offlineNetHop[p2pHostID] = struct{}{}
	}

	for p2pHostID := range offlineNetHop {
		delete(nhm, p2pHostID)
	}
}

func (r *ribData) cleanupRoutingTable(rt map[string]*routing.RoutingEntry) {
	r.Lock()
	defer r.Unlock()

	orphanAddr := make(map[string]struct{})

	for addr, re := range rt {
		offlineNetHop := make(map[string]struct{})

		for p2pHostID, nh := range re.Gw {
			if netHopIsOnline(nh) {
				continue
			}
			offlineNetHop[p2pHostID] = struct{}{}
		}

		for p2pHostID := range offlineNetHop {
			delete(re.Gw, p2pHostID)
		}

		if len(re.Gw) == 0 {
			// CONNECTED routes are not removed
			if re.Type != routing.RouteType_CONNECTED {
				xlog.Debugf("Removing orphan route to %s", addr)
				evtDeleteRoute(addr, re.Type)
				orphanAddr[addr] = struct{}{}
			}
		}
	}

	for addr := range orphanAddr {
		delete(rt, addr)
	}
}

func netHopIsOnline(nh *routing.NetHop) bool {
	tm := time.Unix(nh.LastSeen, 0)

	return time.Since(tm) <= rtHoldTime*time.Second
}
