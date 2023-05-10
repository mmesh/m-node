package router

import (
	"mmesh.dev/m-api-go/grpc/network/routing"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-node/internal/app/node/mnet/p2p/peer"
)

func (r *router) rconnect(nh *routing.NetHop) error {
	if nh.P2PHostID == r.p2pHost.ID().Pretty() {
		return nil
	}

	peerHop := &peer.NetHop{
		PeerMAddrs:   nh.MAddrs,
		RelayMAddrs:  nil,
		RouterMAddrs: nil,
	}

	// circuitv1 relays
	if _, err := peer.Connect(r.p2pHost, peerHop); err != nil {
		return errors.Wrapf(err, "[%v] function peer.Connect()", errors.Trace())
	}

	/*
		// circuitv2 requires slot reservations in relays
		if err := relayReservation(peerInfo); err != nil {
			xlog.Errorf("Unable to get reservation at relay: %v", err)
			return
		}

		if err := peer.Connect(r.p2pHost, peerInfo); err != nil {
			xlog.Errorf("Unable to connect to relay: %v", err)
			return
		}
		xlog.Debugf("Connection ESTABLISHED with relay %s", peerInfo.ID.Pretty())

	*/

	return nil
}
