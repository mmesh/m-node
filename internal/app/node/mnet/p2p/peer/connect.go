package peer

import (
	"context"
	"fmt"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/net/swarm"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/xlog"
	"mmesh.dev/m-node/internal/app/node/mnet/p2p"
	"mmesh.dev/m-node/internal/app/node/mnet/p2p/conn"
)

func NewStream(p2pHost host.Host, hop *NetHop) (network.Stream, error) {
	pm := newPeerAddrInfoMapFromNetHop(hop)

	// fmt.Println("----- pm - start -----")
	// pm.show()
	// fmt.Println("----- pm - end -----")

	// try direct/relayed connection
	peerInfo := connectPeerGroup(p2pHost, pm.peer)
	if peerInfo == nil {
		return nil, fmt.Errorf("unable to connect to peer")
	}

	conns := p2pHost.Network().ConnsToPeer(peerInfo.ID)

	xlog.Infof("Peer %s CONNECTED (%d conns)",
		peerInfo.ID.ShortString(), len(conns))

	transientConnection := false

	for _, c := range conns {
		if c.Stat().Transient {
			transientConnection = true
		}
		conn.Log(c)
	}

	ctx := context.TODO() // context for direct connection
	if transientConnection {
		ctx = network.WithUseTransient(ctx, "mmesh") // context for relayed connection
	}

	s, err := p2pHost.NewStream(ctx, peerInfo.ID, p2p.ProtocolID)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function p2pHost.NewStream()", errors.Trace())
	}

	return s, nil
}

func connectPeerGroup(p2pHost host.Host, peers map[peer.ID]*peer.AddrInfo) *peer.AddrInfo {
	for _, peerInfo := range peers {
		if peerInfo.ID == p2pHost.ID() {
			continue
		}

		if err := connect(p2pHost, peerInfo); err != nil {
			continue
		}

		return peerInfo
	}

	return nil
}

func connect(p2pHost host.Host, peerInfo *peer.AddrInfo) error {
	p2pHost.Network().(*swarm.Swarm).Backoff().Clear(peerInfo.ID)
	if err := p2pHost.Connect(context.TODO(), *peerInfo); err != nil {
		xlog.Tracef("Unable to connect to peer %s: %v", peerInfo.ID.ShortString(), err)
		return errors.Wrapf(err, "[%v] function p2pHost.Connect()", errors.Trace())
	}
	p2pHost.Network().(*swarm.Swarm).Backoff().Clear(peerInfo.ID)

	return nil
}
