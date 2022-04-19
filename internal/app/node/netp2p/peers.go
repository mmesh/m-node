package netp2p

import (
	"context"

	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/xlog"

	"github.com/libp2p/go-libp2p-core/peer"
	swarm "github.com/libp2p/go-libp2p-swarm"
	// "github.com/libp2p/go-libp2p/p2p/protocol/circuitv2/client"
	"github.com/multiformats/go-multiaddr"
)

func getPeerInfo(maddr string) (*peer.AddrInfo, error) {
	peerAddr, err := multiaddr.NewMultiaddr(maddr)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function multiaddr.NewMultiaddr()", errors.Trace())
	}

	peerInfo, err := peer.AddrInfoFromP2pAddr(peerAddr)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function peer.AddrInfoFromP2pAddr()", errors.Trace())
	}

	return peerInfo, nil
}

/*
func relayReservation(peerInfo *peer.AddrInfo) error {
	if peerInfo.ID.Pretty() == mma.p2pHost.ID().Pretty() {
		return nil
	}

	if _, err := client.Reserve(context.TODO(), mma.p2pHost, *peerInfo); err != nil {
		return errors.Wrapf(err, "[%v] function client.Reserve()", errors.Trace())
	}

	xlog.Debugf("Got RESERVATION from relay %s", peerInfo.ID.Pretty())

	return nil
}
*/

func (mma *mmAgent) connectPeer(peerInfo *peer.AddrInfo) error {
	if peerInfo.ID.Pretty() == mma.p2pHost.ID().Pretty() {
		return nil
	}

	ctx := context.TODO()

	mma.p2pHost.Network().(*swarm.Swarm).Backoff().Clear(peerInfo.ID)
	if err := mma.p2pHost.Connect(ctx, *peerInfo); err != nil {
		xlog.Debugf("Connection FAILED with peer %s", peerInfo.ID.Pretty())
		xlog.Trace(err)
		return errors.Wrapf(err, "[%v] function p2pHost.Connect()", errors.Trace())
	}
	mma.p2pHost.Network().(*swarm.Swarm).Backoff().Clear(peerInfo.ID)

	xlog.Debugf("Connection ESTABLISHED with peer %s", peerInfo.ID.Pretty())

	return nil
}
