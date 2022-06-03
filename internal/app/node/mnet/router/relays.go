package router

import (
	"mmesh.dev/m-api-go/grpc/network/mmnp/routing"
	"mmesh.dev/m-node/internal/app/node/mnet/p2p/peer"
)

func (r *router) connectRelays() {
	r.rib.RLock()
	defer r.rib.RUnlock()

	if r.rib.global.Relays == nil {
		return
	}

	for _, relay := range r.rib.global.Relays {
		if relay.VRFID == r.vrfID || r.rib.global.Scope == routing.Scope_NETWORK {
			if relay.P2PHostID == r.p2pHost.ID().Pretty() {
				continue
			}

			peerHop := &peer.NetHop{
				// ID:      relay.P2PHostID,
				MAddrs: relay.MAddrs,
			}

			// circuitv1 relays
			go peer.Connect(r.p2pHost, peerHop)

			/*
				// circuitv2 requires slot reservations in relays
				go func() {
					if err := relayReservation(peerInfo); err != nil {
						xlog.Errorf("Unable to get reservation at relay: %v", err)
						return
					}

					if err := peer.Connect(r.p2pHost, peerInfo); err != nil {
						xlog.Errorf("Unable to connect to relay: %v", err)
						return
					}
					xlog.Debugf("Connection ESTABLISHED with relay %s", peerInfo.ID.Pretty())
				}()
			*/
		}
	}
}
