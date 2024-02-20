package router

import (
	"mmesh.dev/m-api-go/grpc/network/routing"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-node/internal/app/node/mnet/p2p/peer"
)

func (r *router) relayConnect(nh *routing.NetHop) error {
	return r.rconnect(nh)
}

func (r *router) routerConnect(nh *routing.NetHop) error {
	return r.rconnect(nh)
}

func (r *router) proxyConnect(nh *routing.NetHop) error {
	if nh.P2PHostID == r.p2pHost.ID().String() {
		return nil
	}

	peerHop := &peer.NetHop{
		PeerMAddrs:   nh.MAddrs,
		RelayMAddrs:  nil,
		RouterMAddrs: nil,
	}

	if err := peer.ProxyConnect(r.p2pHost, peerHop); err != nil {
		return errors.Wrapf(err, "[%v] function peer.ProxyConnect()", errors.Trace())
	}

	return nil
}

func (r *router) rconnect(nh *routing.NetHop) error {
	if nh.P2PHostID == r.p2pHost.ID().String() {
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
