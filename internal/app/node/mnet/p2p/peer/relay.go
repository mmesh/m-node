package peer

import (
	"context"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/protocol/circuitv2/client"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/xlog"
)

func RConnect(p2pHost host.Host, hop *NetHop) error {
	pm := newPeerAddrInfoMapFromNetHop(hop)

	if err := rconnect(p2pHost, pm.peer); err != nil {
		return errors.Wrapf(err, "[%v] function rconnect()", errors.Trace())
	}

	return nil
}

func rconnect(p2pHost host.Host, peers map[peer.ID]*peer.AddrInfo) error {
	for _, peerInfo := range peers {
		if peerInfo.ID == p2pHost.ID() {
			continue
		}

		if err := connect(p2pHost, peerInfo); err != nil {
			continue
		}

		if err := getRelayReservation(p2pHost, peerInfo); err != nil {
			continue
		}

		return nil
	}

	return nil
}

func getRelayReservation(p2pHost host.Host, relayPeerInfo *peer.AddrInfo) error {
	if _, err := client.Reserve(context.TODO(), p2pHost, *relayPeerInfo); err != nil {
		xlog.Warnf("Unable to get reservation from relay peer %s: %v",
			relayPeerInfo.ID.ShortString(), err)
		return errors.Wrapf(err, "[%v] function client.Reserve()", errors.Trace())
	}

	return nil
}
