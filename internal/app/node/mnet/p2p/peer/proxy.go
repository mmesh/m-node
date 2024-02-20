package peer

import (
	"fmt"

	"github.com/libp2p/go-libp2p/core/host"
)

func ProxyConnect(p2pHost host.Host, hop *NetHop) error {
	pm := newPeerAddrInfoMapFromNetHop(hop)

	for _, peerInfo := range pm.peer {
		if peerInfo.ID == p2pHost.ID() {
			continue
		}

		if err := connect(p2pHost, peerInfo); err != nil {
			continue
		}

		return nil
	}

	return fmt.Errorf("unable to connect to proxy")
}
