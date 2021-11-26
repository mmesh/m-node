package netp2p

import (
	"bufio"
	"context"
	"fmt"
	"sort"
	"strings"

	"mmesh.dev/m-api-go/grpc/network/mmnp/routing"
	"mmesh.dev/m-node/internal/app/node/metrics"
	"x6a.dev/pkg/errors"
	"x6a.dev/pkg/xlog"
)

func (mma *mmAgent) connectTunnel(maddr string) (*bufio.ReadWriter, error) {
	ctx := context.TODO()

	xlog.Debugf("Connecting to endpoint maddr %s", maddr)

	peerInfo, err := getPeerInfo(maddr)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function getPeerInfo(maddr)", errors.Trace())
	}

	if err := mma.connectPeer(peerInfo); err != nil {
		return nil, errors.Wrapf(err, "[%v] function mma.connectPeer(peerInfo)", errors.Trace())
	}

	s, err := mma.p2pHost.NewStream(ctx, peerInfo.ID, mmProtocolID)
	if err != nil {
		xlog.Errorf("Unable to create a stream with peer %s: %v", peerInfo.ID.Pretty(), err)
		return nil, errors.Wrapf(err, "[%v] function mma.p2pHost.NewStream(ctx, peerInfo.ID, mmProtocolID)", errors.Trace())
	}

	// Create a buffered stream so that read and writes are non blocking.
	return bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s)), nil
}

func (mma *mmAgent) newTunnel(ipPkt *ipPacket) error {
	xlog.Infof("Requested new tunnel to %s", ipPkt.dstAddr)

	nh, prio, err := mma.getNetHop(ipPkt.dstAddr)
	if err != nil {
		return errors.Wrapf(err, "[%v] function mma.getNetHop(addr)", errors.Trace())
	}

	// Try direct connection
	xlog.Infof("Trying direct connection to %s", ipPkt.dstAddr)
	for _, ma := range nh.Agent.MAddrs {
		nextHop := strings.Split(ma, "/")[2]
		if !strings.Contains(ma, "/p2p-circuit/") {
			if rw, err := mma.connectTunnel(ma); err != nil {
				continue
			} else {
				mma.setTunnel(ipPkt.dstAddr, rw)
				xlog.Infof("Tunnel connected to %s via %s [DIRECT]", ipPkt.dstAddr, nextHop)

				// Create a thread to read data from new buffered stream.
				go mma.readStream(rw, ma, false)
				return nil
			}
		}
	}
	// Try connection via relay
	xlog.Infof("Trying connection via relay to %s", ipPkt.dstAddr)
	if err := mma.newRTunnel(ipPkt.dstAddr, nh); err == nil {
		return nil
	}

	mma.setNxHopUnhealthy(ipPkt.dstAddr, *prio)

	return errors.Errorf("unable to establish a new tunnel to reach %v", ipPkt.dstIP.String())
}

func (mma *mmAgent) newRTunnel(dstAddr string, nh *routing.NetHop) error {
	sort.SliceStable(mma.rt.RT.Relays, func(i, j int) bool {
		return mma.rt.RT.Relays[i].Connections < mma.rt.RT.Relays[j].Connections
	})

	for _, r := range mma.rt.RT.Relays {
		if r.VRFID == mma.vrfID || mma.rt.RT.Scope == routing.Scope_NETWORK {
			for _, ma := range nh.Agent.MAddrs {
				nextHop := strings.Split(ma, "/")[2]
				if strings.HasPrefix(ma, r.MAddr+"/p2p-circuit/") {
					if rw, err := mma.connectTunnel(ma); err != nil {
						continue
					} else {
						mma.setTunnel(dstAddr, rw)
						xlog.Infof("Tunnel connected to %s via %s [RELAY]", dstAddr, nextHop)

						metrics.IncrRelayConns(r.MAddr)

						// Create a thread to read data from new buffered stream.
						go mma.readStream(rw, r.MAddr, true)
						return nil
					}
				}
			}
		}
	}

	return fmt.Errorf("no valid relay found")
}

func (mma *mmAgent) setTunnel(dstAddr string, rw *bufio.ReadWriter) {
	mma.streams.Lock()
	defer mma.streams.Unlock()

	if _, ok := mma.streams.tunnel[dstAddr]; ok {
		return
	}

	mma.streams.tunnel[dstAddr] = rw
}

func (mma *mmAgent) getTunnel(dstAddr string) *bufio.ReadWriter {
	mma.streams.RLock()
	defer mma.streams.RUnlock()

	if rw, ok := mma.streams.tunnel[dstAddr]; ok {
		return rw
	}

	return nil
}

func (mma *mmAgent) deleteTunnel(dstAddr string) {
	mma.streams.Lock()
	defer mma.streams.Unlock()

	if _, ok := mma.streams.tunnel[dstAddr]; ok {
		delete(mma.streams.tunnel, dstAddr)
		xlog.Infof("Deleted tunnel to %s", dstAddr)
	}
}
