package router

import (
	"mmesh.dev/m-api-go/grpc/network/routing"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-node/internal/app/node/mnet/p2p/peer"
)

func (r *router) rconnect(nh *routing.NetHop) error {
	if nh.P2PHostID == r.p2pHost.ID().Pretty() {
		return nil
	}

	peerHop := &peer.NetHop{
		PeerMAddrs:   nh.MAddrs,
		RelayMAddrs:  nil,
		RouterMAddrs: nil,
	}

	if err := peer.RConnect(r.p2pHost, peerHop); err != nil {
		return errors.Wrapf(err, "[%v] function peer.RConnect()", errors.Trace())
	}

	return nil
}
