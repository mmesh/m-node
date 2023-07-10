package peer

import (
	"context"
	"fmt"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/protocol/circuitv2/client"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/xlog"
)

func RConnect(p2pHost host.Host, hop *NetHop) error {
	pm := newPeerAddrInfoMapFromNetHop(hop)

	for _, peerInfo := range pm.peer {
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

	return fmt.Errorf("unable to reserve a slot in relay")
}

func getRelayReservation(p2pHost host.Host, relayPeerInfo *peer.AddrInfo) error {
	if _, err := client.Reserve(context.TODO(), p2pHost, *relayPeerInfo); err != nil {
		xlog.Tracef("Unable to get reservation from relay peer %s: %v",
			relayPeerInfo.ID.ShortString(), err)
		return errors.Wrapf(err, "[%v] function client.Reserve()", errors.Trace())
	}

	return nil
}
