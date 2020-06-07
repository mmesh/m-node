package netp2p

import (
	"bufio"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/pkg/errors"
	"mmesh.dev/mmesh/internal/app/node/metrics"
	"x6a.dev/pkg/xlog"
)

const BUFFER_SIZE = 9216 // Jumbo Frame Size

func handleStream(s network.Stream) {
	xlog.Info("[+] New stream connected")

	// Create a buffer stream for non blocking read and write.
	rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))

	// Create a thread to read data from new buffered stream.
	go mma.readStream(rw, "", false)

	// stream 's' will stay open until you close it (or the other side closes it).
	mma.linkStatus.Connections++
}

func (mma *mmAgent) readStream(rw *bufio.ReadWriter, maddr string, relayTunnel bool) {
	buf := make([]byte, BUFFER_SIZE)
	for {
		n, err := rw.Read(buf)
		if err != nil {
			xlog.Tracef("Unable to read from tunnel: %v", err)
			break
		}

		if err := mma.writeInterface(rw, buf, n); err != nil {
			xlog.Warnf("Unable to write packet to interface: %v", err)
			continue
		}
	}
	mma.linkStatus.Connections--
	if relayTunnel {
		metrics.DecrRelayConns(maddr)
	}
	xlog.Info("[!] Stream terminated")
}

func (mma *mmAgent) writeInterface(rw *bufio.ReadWriter, pkt []byte, plen int) error {
	ipPkt, err := decodeIPPacket(pkt[:plen])
	if err != nil {
		// xlog.Tracef("Unable to decode ip packet: %v", err)
		return err
	}

	mma.setTunnel(ipPkt.srcIP.String(), rw)

	// proxy64 forwarding
	if ipPkt.proxy64Forward() {
		return nil
	}

	// write to TUN interface
	if iface != nil {
		if _, err = iface.Write(pkt[:plen]); err != nil {
			return err
		}

		// update metrics
		go metrics.UpdateNetworkMetric(ipPkt.srcIP.String(), 0, uint64(plen), false)
	}

	return nil
}

func (mma *mmAgent) readInterface() {
	var queue = make(chan *ipPacket, 16*BUFFER_SIZE)

	// read interface
	go func() {
		pkt := make([]byte, BUFFER_SIZE)
		for {
			plen, err := iface.Read(pkt)
			if err != nil {
				xlog.Tracef("Unable to read from interface: %v", err)
				mma.closeInterface <- struct{}{}
				break
			}

			// packet filtering / firewalling
			ipPkt, dropped := mma.packetFilter(pkt[:plen])
			if dropped {
				// packet dropped by policy
				go func() {
					if ipPkt == nil {
						return
					}
					//xlog.Tracef("Discarding packet from srcIP %s to dstIP %s", ipPkt.srcIP.String(), ipPkt.dstIP.String())

					// udpate metrics with a new drop
					go metrics.UpdateNetworkMetric(ipPkt.dstIP.String(), 0, 0, true)
				}()
				continue
			}

			// proxy64 forwarding
			if ipPkt.proxy64Forward() {
				continue
			}

			queue <- ipPkt
		}
	}()

	// write stream
	go func() {
		var rw *bufio.ReadWriter

		for {
			ipPkt := <-queue

			rw = mma.getTunnel(ipPkt.dstAddr)
			if rw == nil {
				if err := mma.newTunnel(ipPkt); err != nil {
					xlog.Error(errors.Cause(err))
					continue
				}
				rw = mma.getTunnel(ipPkt.dstAddr)
				mma.linkStatus.Connections++
			}

			// real send
			if _, err := rw.Write(ipPkt.data[:ipPkt.plen]); err != nil {
				//xlog.Tracef("rw.Write(packet.data[:packet.len]): %v", err)
				mma.deleteTunnel(ipPkt.dstAddr)
				continue
			}

			if err := rw.Flush(); err != nil {
				//xlog.Tracef("rw.Flush(): %v", err)
				mma.deleteTunnel(ipPkt.dstAddr)
				continue
			}

			// udpate metrics
			go metrics.UpdateNetworkMetric(ipPkt.dstIP.String(), uint64(ipPkt.plen), 0, false)
		}
	}()

	<-mma.closeInterface
}
