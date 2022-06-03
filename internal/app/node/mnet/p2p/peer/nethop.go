package peer

import (
	"github.com/libp2p/go-libp2p-core/peer"
	"mmesh.dev/m-node/internal/app/node/mnet/maddr"
	"mmesh.dev/m-node/internal/app/node/mnet/p2p/transport"
)

// const maxIDChars int = 8

type NetHop struct {
	// ID     string
	MAddrs []string
}

type peerMap struct {
	peer map[peer.ID]*peer.AddrInfo
}

func newPeerMapFromNetHop(hop *NetHop) *peerMap {
	pm := &peerMap{
		peer: make(map[peer.ID]*peer.AddrInfo, 0),
	}

	for _, ma := range hop.MAddrs {
		proto := maddr.GetTransport(ma)

		if proto == transport.Invalid {
			continue
		}

		pi, err := getPeerInfo(ma)
		if err != nil {
			// xlog.Errorf("Unable to parse multiaddr %s: %v", ma, err)
			continue
		}

		if _, ok := pm.peer[pi.ID]; !ok {
			pm.peer[pi.ID] = pi
		} else {
			pm.peer[pi.ID].Addrs = append(pm.peer[pi.ID].Addrs, pi.Addrs...)
		}
	}

	return pm
}
