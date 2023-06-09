package peer

import (
	"github.com/libp2p/go-libp2p/core/peer"
)

type NetHop struct {
	PeerMAddrs   []string
	RelayMAddrs  []string
	RouterMAddrs []string
}

type peerAddrInfoMap struct {
	peer map[peer.ID]*peer.AddrInfo
	// relay  map[peer.ID]*peer.AddrInfo
	// router map[peer.ID]*peer.AddrInfo
}

func newPeerAddrInfoMapFromNetHop(hop *NetHop) *peerAddrInfoMap {
	netMAddrs := make([]string, 0)
	netMAddrs = append(netMAddrs, hop.PeerMAddrs...)

	// netMAddrs = append(netMAddrs, hop.RelayMAddrs...)
	// netMAddrs = append(netMAddrs, hop.RouterMAddrs...)

	if len(hop.RelayMAddrs) > 0 {
		netMAddrs = append(netMAddrs, hop.RelayMAddrs...)
	} else {
		netMAddrs = append(netMAddrs, hop.RouterMAddrs...)
	}

	return &peerAddrInfoMap{
		peer: getPeerAddrs(netMAddrs),
		// peer:   getPeerAddrs(hop.PeerMAddrs),
		// relay:  getPeerAddrs(hop.RelayMAddrs),
		// router: getPeerAddrs(hop.RouterMAddrs),
	}
}

/*
func (pm *peerAddrInfoMap) show() {
	fmt.Println()
	fmt.Println("///////////////////////////////////////////////////////////////")
	fmt.Println()

	fmt.Println("-----------------")
	fmt.Println("peerMap - pm.peer")
	fmt.Println("-----------------")

	for _, pi := range pm.peer {
		showPeerAddrInfo(pi)
	}

	// fmt.Println("------------------")
	// fmt.Println("peerMap - pm.relay")
	// fmt.Println("------------------")

	// for _, pi := range pm.relay {
	// 	showPeerAddrInfo(pi)
	// }

	// fmt.Println("-------------------")
	// fmt.Println("peerMap - pm.router")
	// fmt.Println("-------------------")

	// for _, pi := range pm.router {
	// 	showPeerAddrInfo(pi)
	// }

	fmt.Println()
	fmt.Println("///////////////////////////////////////////////////////////////")
	fmt.Println()
}

func showPeerAddrInfo(pi *peer.AddrInfo) {
	fmt.Println("***************************************************************")
	fmt.Println()
	// fmt.Println("peer.AddrInfo")
	// fmt.Println("=============")
	// fmt.Println(pi.String())
	// fmt.Println()
	fmt.Println("--------------------------------------------")
	fmt.Printf("Peer ID: %s\n", pi.ID.Pretty())
	fmt.Println("--------------------------------------------")
	fmt.Println(" MultiAddrs:")
	for _, ma := range pi.Addrs {
		fmt.Printf("  - %s\n", ma.String())
	}
	fmt.Println()
	fmt.Println("***************************************************************")
}
*/
