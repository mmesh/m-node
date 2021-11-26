package netp2p

import (
	"bufio"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/pkg/errors"
	"mmesh.dev/m-node/internal/app/node/metrics"
	"x6a.dev/pkg/xlog"
)

// https://github.com/lucas-clemente/quic-go/wiki/UDP-Receive-Buffer-Size
// const BUFFER_SIZE = 2500000

// const BUFFER_SIZE = 9216 // Jumbo Frame Size

const BUFFER_SIZE = 1 << 16

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
	pkt := make([]byte, BUFFER_SIZE)
	for {
		plen, err := rw.Read(pkt)
		if err != nil {
			xlog.Tracef("Unable to read from tunnel: %v", err)
			break
		}

		if err := mma.writeInterface(rw, pkt, plen); err != nil {
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
		if _, err = ifaceWrite(pkt[:plen]); err != nil {
			return err
		}

		// update metrics
		go metrics.UpdateNetworkMetric(ipPkt.srcIP.String(), 0, uint64(plen), false)
	}

	return nil
}

func (mma *mmAgent) readInterface() {
	// read interface
	go func() {
		pkt := make([]byte, BUFFER_SIZE)
		for {
			plen, err := ifaceRead(pkt)
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

					// udpate metrics with a new drop
					go metrics.UpdateNetworkMetric(ipPkt.dstIP.String(), 0, 0, true)
				}()
				continue
			}

			// proxy64 forwarding
			if ipPkt.proxy64Forward() {
				continue
			}

			rw := mma.getTunnel(ipPkt.dstAddr)
			if rw == nil {
				go func(ipPkt ipPacket) {
					if err := mma.newTunnel(&ipPkt); err != nil {
						xlog.Error(errors.Cause(err))
						return
					}
					rw := mma.getTunnel(ipPkt.dstAddr)
					mma.linkStatus.Connections++

					writeStream(ipPkt, rw)
				}(*ipPkt)
			} else {
				writeStream(*ipPkt, rw)
			}
		}
	}()

	<-mma.closeInterface
}

func writeStream(ipPkt ipPacket, rw *bufio.ReadWriter) {
	// real send
	if _, err := rw.Write(ipPkt.data[:ipPkt.plen]); err != nil {
		// xlog.Errorf("rw.Write(packet.data[:packet.len]): %v", err)
		mma.deleteTunnel(ipPkt.dstAddr)
		return
	}

	if err := rw.Flush(); err != nil {
		// xlog.Errorf("rw.Flush(): %v", err)
		mma.deleteTunnel(ipPkt.dstAddr)
		return
	}

	// udpate metrics
	go metrics.UpdateNetworkMetric(ipPkt.dstIP.String(), uint64(ipPkt.plen), 0, false)
}
