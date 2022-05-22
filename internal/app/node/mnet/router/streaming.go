package router

import (
	"bufio"

	"github.com/libp2p/go-libp2p-core/network"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/xlog"
	"mmesh.dev/m-node/internal/app/node/metrics"
	"mmesh.dev/m-node/internal/app/node/mnet/p2p/peer"
)

// https://github.com/lucas-clemente/quic-go/wiki/UDP-Receive-Buffer-Size
// const BUFFER_SIZE = 2500000

// const BUFFER_SIZE = 9216 // Jumbo Frame Size

const BUFFER_SIZE = 1 << 16

func (r *router) handleStream(s network.Stream) {
	xlog.Info("[+] New stream connected")

	// Create a buffer stream for non blocking read and write.
	rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))

	// Create a thread to read data from new buffered stream.
	go r.readStream(rw, nil)

	// stream 's' will stay open until you close it (or the other side closes it).
	r.linkStatus.Connections++
}

func (r *router) readStream(rw *bufio.ReadWriter, pc *peer.Connection) {
	pkt := make([]byte, BUFFER_SIZE)
	for {
		plen, err := rw.Read(pkt)
		if err != nil {
			xlog.Tracef("Unable to read from tunnel: %v", err)
			break
		}

		if err := r.writeInterface(rw, pkt, plen); err != nil {
			xlog.Warnf("Unable to write packet to interface: %v", err)
			continue
		}
	}
	r.linkStatus.Connections--

	if pc != nil {
		if !pc.DirectConnection {
			go metrics.DecrRelayConns(pc.PeerInfo.ID.Pretty())
		}
	}

	xlog.Info("[!] Stream terminated")
}

func (r *router) writeInterface(rw *bufio.ReadWriter, pkt []byte, plen int) error {
	ipPkt, err := decodeIPPacket(pkt[:plen])
	if err != nil {
		// xlog.Tracef("Unable to decode ip packet: %v", err)
		return err
	}

	r.setTunnel(ipPkt.srcIP.String(), rw)

	// proxy64 forwarding
	if r.proxy64Forward(ipPkt) {
		return nil
	}

	// write to TUN interface
	if r.networkInterface != nil {
		if _, err = r.networkInterface.write(pkt[:plen]); err != nil {
			return err
		}

		// update metrics
		go metrics.UpdateNetworkMetric(ipPkt.srcIP.String(), 0, uint64(plen), false)
	}

	return nil
}

func (r *router) readInterface() {
	// read interface
	go func() {
		pkt := make([]byte, BUFFER_SIZE)
		for {
			plen, err := r.networkInterface.read(pkt)
			if err != nil {
				xlog.Tracef("Unable to read from interface: %v", err)
				r.networkInterface.closeCh <- struct{}{}
				break
			}

			// packet filtering / firewalling
			ipPkt, dropped := r.packetFilter(pkt[:plen])
			if dropped {
				// packet dropped by policy
				go func() {
					if ipPkt == nil {
						return
					}

					// udpate metrics with a new drop
					go metrics.UpdateNetworkMetric(ipPkt.dstIP.String(), 0, 0, true)
				}()
				continue
			}

			// proxy64 forwarding
			if r.proxy64Forward(ipPkt) {
				continue
			}

			rw := r.getTunnel(ipPkt.dstAddr)
			if rw == nil {
				go func(ipPkt ipPacket) {
					if err := r.newTunnel(&ipPkt); err != nil {
						xlog.Warnf("Unable to get tunnel to %s: %v",
							ipPkt.dstIP.String(), errors.Cause(err))
						return
					}
					rw := r.getTunnel(ipPkt.dstAddr)
					r.linkStatus.Connections++

					r.writeStream(ipPkt, rw)
				}(*ipPkt)
			} else {
				r.writeStream(*ipPkt, rw)
			}
		}
	}()

	<-r.networkInterface.closeCh
}

func (r *router) writeStream(ipPkt ipPacket, rw *bufio.ReadWriter) {
	// real send
	if _, err := rw.Write(ipPkt.data[:ipPkt.plen]); err != nil {
		// xlog.Errorf("rw.Write(packet.data[:packet.len]): %v", err)
		r.deleteTunnel(ipPkt.dstAddr)
		return
	}

	if err := rw.Flush(); err != nil {
		// xlog.Errorf("rw.Flush(): %v", err)
		r.deleteTunnel(ipPkt.dstAddr)
		return
	}

	// udpate metrics
	go metrics.UpdateNetworkMetric(ipPkt.dstIP.String(), uint64(ipPkt.plen), 0, false)
}
