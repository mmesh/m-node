package netp2p

import (
	"context"

	"x6a.dev/pkg/errors"
	"x6a.dev/pkg/xlog"

	"github.com/libp2p/go-libp2p-core/peer"
	swarm "github.com/libp2p/go-libp2p-swarm"
	"github.com/multiformats/go-multiaddr"
)

func (mma *mmAgent) connectPeer(peerInfo *peer.AddrInfo) error {
	if peerInfo.ID.Pretty() != mma.p2pHost.ID().Pretty() {
		ctx := context.Background()

		mma.p2pHost.Network().(*swarm.Swarm).Backoff().Clear(peerInfo.ID)
		if err := mma.p2pHost.Connect(ctx, *peerInfo); err != nil {
			xlog.Warnf("Connection FAILED with peer %s", peerInfo.ID.Pretty())
			xlog.Trace(err)
			return errors.Wrapf(err, "[%v] function nxHost.Connect(ctx, *peerInfo)", errors.Trace())
		}
		mma.p2pHost.Network().(*swarm.Swarm).Backoff().Clear(peerInfo.ID)

		xlog.Debugf("Connection ESTABLISHED with peer %v", peerInfo.ID.Pretty())
	}

	return nil
}

func getPeerInfo(maddr string) (*peer.AddrInfo, error) {
	peerAddr, err := multiaddr.NewMultiaddr(maddr)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function multiaddr.NewMultiaddr(maddr)", errors.Trace())
	}

	peerInfo, err := peer.AddrInfoFromP2pAddr(peerAddr)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function peer.AddrInfoFromP2pAddr(peerAddr)", errors.Trace())
	}

	return peerInfo, nil
}
