package router

import (
	"bufio"
	"context"

	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/xlog"
	"mmesh.dev/m-node/internal/app/node/metrics"
	"mmesh.dev/m-node/internal/app/node/mnet/p2p"
	"mmesh.dev/m-node/internal/app/node/mnet/p2p/peer"
)

func (r *router) connectTunnel(peerHop *peer.Hop) (*bufio.ReadWriter, *peer.Connection, error) {
	pc, err := peer.ConnectHop(r.p2pHost, peerHop)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "[%v] function peer.ConnectHop()", errors.Trace())
	}

	s, err := r.p2pHost.NewStream(context.TODO(), pc.PeerInfo.ID, p2p.ProtocolID)
	if err != nil {
		xlog.Errorf("Unable to create a stream with peer %s: %v", pc.PeerInfo.ID.Pretty(), err)
		return nil, nil, errors.Wrapf(err, "[%v] function r.p2pHost.NewStream()", errors.Trace())
	}

	// Create a buffered stream so that read and writes are non blocking.
	return bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s)), pc, nil
}

func (r *router) newTunnel(ipPkt *ipPacket) error {
	// xlog.Infof("Requested new tunnel to %s", ipPkt.dstAddr)

	nh, prio, err := r.getNetHop(ipPkt.dstAddr)
	if err != nil {
		return errors.Wrapf(err, "[%v] function r.getNetHop()", errors.Trace())
	}

	peerHop := &peer.Hop{
		ID:     nh.Agent.AgentID,
		MAddrs: nh.Agent.MAddrs,
	}

	rw, pc, err := r.connectTunnel(peerHop)
	if err != nil {
		r.setNetHopUnhealthy(ipPkt.dstAddr, *prio)
		return errors.Wrapf(err, "[%v] function r.connectTunnel()", errors.Trace())
	}

	if !r.setTunnel(ipPkt.dstAddr, rw) {
		return nil
	}

	// xlog.Infof("Tunnel connected to %s via %s [DIRECT]", ipPkt.dstAddr, nextHop)

	if pc.DirectConnection {
		xlog.Infof("%s tunnel connected to %s [DIRECT]", pc.Proto.String(), ipPkt.dstAddr)
	} else {
		xlog.Infof("%s tunnel connected to %s [INDIRECT]", pc.Proto.String(), ipPkt.dstAddr)
		go metrics.IncrRelayConns(pc.PeerInfo.ID.Pretty())
	}

	// Create a thread to read data from new buffered stream.
	go r.readStream(rw, pc)

	return nil
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