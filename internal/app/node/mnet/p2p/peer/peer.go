package peer

import (
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multiaddr"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/xlog"
	"mmesh.dev/m-node/internal/app/node/mnet/maddr"
	"mmesh.dev/m-node/internal/app/node/mnet/p2p/transport"
)

func getPeerAddrs(maddrs []string) map[peer.ID]*peer.AddrInfo {
	peerInfo := make(map[peer.ID]*peer.AddrInfo, 0)

	for _, ma := range maddrs {
		proto := maddr.GetTransport(ma)

		if proto == transport.Invalid {
			continue
		}

		pi, err := getPeerAddrInfo(ma)
		if err != nil {
			xlog.Warnf("Unable to parse peer multiaddr %s: %v", ma, errors.Cause(err))
			continue
		}

		if _, ok := peerInfo[pi.ID]; !ok {
			peerInfo[pi.ID] = pi
		} else {
			peerInfo[pi.ID].Addrs = append(peerInfo[pi.ID].Addrs, pi.Addrs...)
		}
	}

	return peerInfo
}

func getPeerAddrInfo(maddr string) (*peer.AddrInfo, error) {
	peerAddr, err := multiaddr.NewMultiaddr(maddr)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function multiaddr.NewMultiaddr()", errors.Trace())
	}

	peerAddrInfo, err := peer.AddrInfoFromP2pAddr(peerAddr)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function peer.AddrInfoFromP2pAddr()", errors.Trace())
	}

	return peerAddrInfo, nil
}
