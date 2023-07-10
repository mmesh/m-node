package router

import (
	"bufio"

	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/xlog"
	"mmesh.dev/m-node/internal/app/node/mnet/p2p/peer"
)

func (r *router) connectTunnel(peerHop *peer.NetHop) (*bufio.ReadWriter, error) {
	s, err := peer.NewStream(r.p2pHost, peerHop)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function peer.NewStream()", errors.Trace())
	}

	// create a buffered stream so that read and writes are non blocking
	return bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s)), nil
}

func (r *router) newTunnel(ipPkt *ipPacket) (bool, error) {
	// xlog.Infof("Requested new tunnel to %s", ipPkt.dstAddr)

	if !r.setDial(ipPkt.dstAddr) {
		return false, nil
	}
	defer r.unsetDial(ipPkt.dstAddr)

	nh, err := r.RIB().GetNetHop(ipPkt.dstAddr)
	if err != nil {
		return false, errors.Wrapf(err, "[%v] function r.RIB().GetNetHop()", errors.Trace())
	}

	peerHop := &peer.NetHop{
		PeerMAddrs:   nh.MAddrs,
		RelayMAddrs:  r.RIB().GetRelayMAddrs(nh),
		RouterMAddrs: r.RIB().GetRouterMAddrs(nh),
	}

	rw, err := r.connectTunnel(peerHop)
	if err != nil {
		// r.RIB().SetNetHopUnhealthy(ipPkt.dstAddr, nh.P2PHostID)
		return false, errors.Wrapf(err, "[%v] function r.connectTunnel()", errors.Trace())
	}

	if !r.setTunnel(ipPkt.dstAddr, rw) {
		return true, nil
	}

	xlog.Infof("Tunnel connected to %s", ipPkt.dstAddr)

	// create a thread to read data from new buffered stream
	go r.readStream(rw)

	return true, nil
}

func (r *router) setTunnel(dstAddr string, rw *bufio.ReadWriter) bool {
	r.streams.Lock()
	defer r.streams.Unlock()

	if _, ok := r.streams.tunnel[dstAddr]; ok {
		return false
	}

	r.streams.tunnel[dstAddr] = rw

	return true
}

func (r *router) getTunnel(dstAddr string) *bufio.ReadWriter {
	r.streams.RLock()
	defer r.streams.RUnlock()

	if rw, ok := r.streams.tunnel[dstAddr]; ok {
		return rw
	}

	return nil
}

func (r *router) deleteTunnel(dstAddr string) {
	r.streams.Lock()
	defer r.streams.Unlock()

	if _, ok := r.streams.tunnel[dstAddr]; ok {
		delete(r.streams.tunnel, dstAddr)
		xlog.Infof("Deleted tunnel to %s", dstAddr)
	}
}
