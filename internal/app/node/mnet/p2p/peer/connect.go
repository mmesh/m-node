package peer

import (
	"context"
	"fmt"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	// swarm "github.com/libp2p/go-libp2p-swarm"
	"github.com/libp2p/go-libp2p/p2p/net/swarm"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/xlog"
)

func Connect(p2pHost host.Host, hop *NetHop) (*peer.AddrInfo, error) {
	pm := newPeerMapFromNetHop(hop)

	for _, peerInfo := range pm.peer {
		if peerInfo.ID == p2pHost.ID() {
			continue
		}

		// peerID := peerInfo.ID.Pretty()[:maxIDChars]

		// xlog.Debugf("Trying connection to peer %s", peerID)
		if err := connect(p2pHost, peerInfo); err != nil {
			// xlog.Warnf("Unable to connect peer %s: %v", peerID, err)
			continue
		} else {
			// xlog.Debugf("CONNECTED to peer %s", peerID)
			return peerInfo, nil
		}
	}

	return nil, fmt.Errorf("unable to connect to peer")
}

func connect(p2pHost host.Host, peerInfo *peer.AddrInfo) error {
	p2pHost.Network().(*swarm.Swarm).Backoff().Clear(peerInfo.ID)
	if err := p2pHost.Connect(context.TODO(), *peerInfo); err != nil {
		xlog.Trace(err)
		return errors.Wrapf(err, "[%v] function p2pHost.Connect()", errors.Trace())
	}
	p2pHost.Network().(*swarm.Swarm).Backoff().Clear(peerInfo.ID)

	return nil
}
