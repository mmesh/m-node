package rib

import (
	"fmt"

	"mmesh.dev/m-api-go/grpc/network/routing"
)

func (r *ribData) GetRelayMAddrs(nh *routing.NetHop) []string {
	r.RLock()
	defer r.RUnlock()

	return getRGroupMAddrs(nh.P2PHostID, r.rib.Relays)
}

func (r *ribData) GetRouterMAddrs(nh *routing.NetHop) []string {
	r.RLock()
	defer r.RUnlock()

	return getRGroupMAddrs(nh.P2PHostID, r.rib.Routers)
}

func getRGroupMAddrs(p2pHostID string, rgroup map[string]*routing.NetHop) []string {
	rMAddrs := make([]string, 0)

	for rP2PHostID, r := range rgroup {
		if p2pHostID == rP2PHostID {
			continue
		}

		// if strings.Contains(r.MAddr, p2pHostID) {
		// 	continue
		// }

		for _, rmaddr := range r.MAddrs {
			rMAddrs = append(rMAddrs,
				fmt.Sprintf("%s/p2p-circuit/p2p/%s", rmaddr, p2pHostID))
		}
	}

	return rMAddrs
}
