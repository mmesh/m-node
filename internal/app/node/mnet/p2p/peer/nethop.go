package peer

import (
	"github.com/libp2p/go-libp2p-core/peer"
	"mmesh.dev/m-node/internal/app/node/mnet/maddr"
	"mmesh.dev/m-node/internal/app/node/mnet/p2p/transport"
)

// const maxIDChars int = 8

type NetHop struct {
	PeerMAddrs   []string
	RelayMAddrs  []string
	RouterMAddrs []string
}

type peerMap struct {
	peer   map[peer.ID]*peer.AddrInfo
	relay  map[peer.ID]*peer.AddrInfo
	router map[peer.ID]*peer.AddrInfo
}

func newPeerMapFromNetHop(hop *NetHop) *peerMap {
	return &peerMap{
		peer:   getPeerGroup(hop.PeerMAddrs),
		relay:  getPeerGroup(hop.RelayMAddrs),
		router: getPeerGroup(hop.RouterMAddrs),
	}
}

func getPeerGroup(maddrs []string) map[peer.ID]*peer.AddrInfo {
	peerGroup := make(map[peer.ID]*peer.AddrInfo, 0)

	for _, ma := range maddrs {
		proto := maddr.GetTransport(ma)

		if proto == transport.Invalid {
			continue
		}

		pi, err := getPeerInfo(ma)
		if err != nil {
			// xlog.Errorf("Unable to parse multiaddr %s: %v", ma, err)
			continue
		}

		if _, ok := peerGroup[pi.ID]; !ok {
			peerGroup[pi.ID] = pi
		} else {
			peerGroup[pi.ID].Addrs = append(peerGroup[pi.ID].Addrs, pi.Addrs...)
		}
	}

	return peerGroup
}
