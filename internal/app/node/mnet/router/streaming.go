package router

import (
	"bufio"
	"encoding/binary"
	"io"

	"github.com/libp2p/go-libp2p/core/network"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/ipnet"
	"mmesh.dev/m-lib/pkg/xlog"
	"mmesh.dev/m-node/internal/app/node/metrics"
	"mmesh.dev/m-node/internal/app/node/mnet/p2p/conn"
)

// const protocolV1 byte = 1

// https://github.com/lucas-clemente/quic-go/wiki/UDP-Receive-Buffer-Size
// const BUFFER_SIZE = 2500000

// const BUFFER_SIZE = 9216 // Jumbo Frame Size
// const BUFFER_SIZE = 1 << 16 // 65536
const BUFFER_SIZE = 1 << 17 // 128KB
// const BUFFER_SIZE = 1 << 20 // 1MB
// const BUFFER_SIZE = 1 << 22 // 4MB

func (r *router) handleStream(s network.Stream) {
	// xlog.Info("[+] New stream connected")
	conn.Log(s.Conn())

	// Create a buffer stream for non blocking read and write.
	rw := bufio.NewReadWriter(
		bufio.NewReaderSize(s, BUFFER_SIZE),
		bufio.NewWriterSize(s, BUFFER_SIZE),
	)
	// bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))

	// Create a thread to read data from new buffered stream.
	go r.readStream(rw)

	// stream 's' will stay open until you close it (or the other side closes it).
	r.connections++
}

func (r *router) readStream(rw *bufio.ReadWriter) {
	hdr := make([]byte, 4)
	for {
		// protocol version
		// protoVersion, err := rw.ReadByte()
		// if err != nil {
		// 	xlog.Warnf("Unable to read proto version from stream: %v", err)
		// 	break
		// }

		// hdr (packet length)
		if _, err := io.ReadFull(rw, hdr); err != nil {
			xlog.Warnf("Unable to read proto header from stream: %v", err)
			break
		}

		plen := binary.BigEndian.Uint32(hdr)

		// pkt
		pkt, err := io.ReadAll(io.LimitReader(rw, int64(plen)))
		if err != nil {
			xlog.Warnf("Unable to read packet from stream: %v", err)
			break
		}

		// if protoVersion == protocolV1 {
		// 	// protocolV1 stuff
		// }

		if err := r.writeInterface(rw, pkt); err != nil {
			xlog.Warnf("Unable to write packet to interface: %v", err)
			continue
		}
	}
	r.connections--

	xlog.Info("[!] Stream terminated")
}

func (r *router) writeInterface(rw *bufio.ReadWriter, pkt []byte) error {
	// parse ip header
	ipHdr, err := parseHeader(pkt)
	if err != nil {
		xlog.Warnf("Unable to parse IP header: %v", errors.Cause(err))
		return nil
	}

	// parse ip protocol
	if err := ipHdr.parseProtocol(pkt); err != nil {
		xlog.Warnf("Unable to parse IP protocol: %v", errors.Cause(err))
		return nil
	}

	// packet filtering (firewalling) -- ingress
	if r.packetFilter(ipHdr, pkt, false) {
		// packet dropped by policy
		go func() {
			if ipHdr == nil {
				return
			}

			// udpate metrics with a new drop
			go metrics.UpdateNetworkMetric(ipHdr.srcIP.String(), 0, 0, true)
		}()
		return nil
	}

	r.setTunnel(ipHdr.srcIP.String(), rw)

	// proxy64 forwarding
	if ipHdr.af == ipnet.AddressFamilyIPv6 {
		if r.proxy64Forward(ipHdr, pkt) {
			return nil
		}
	}

	// write to TUN interface
	if r.networkInterface != nil {
		if _, err := r.networkInterface.write(pkt); err != nil {
			return err
		}

		// update metrics
		go metrics.UpdateNetworkMetric(ipHdr.srcIP.String(), 0, uint64(len(pkt)), false)
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

			ipHdr, err := parseHeader(pkt[:plen])
			if err != nil {
				xlog.Warnf("Unable to parse IP header: %v", errors.Cause(err))
				continue
			}

			if !ipHdr.isValidPacket(r.ipv6) {
				continue
			}

			// proxy64 forwarding
			if ipHdr.af == ipnet.AddressFamilyIPv6 {
				if err := ipHdr.parseProtocol(pkt); err != nil {
					xlog.Warnf("Unable to parse IP protocol: %v", errors.Cause(err))
					continue
				}

				if r.proxy64Forward(ipHdr, pkt[:plen]) {
					continue
				}
			}

			rw := r.getTunnel(ipHdr.dstAddr)
			if rw == nil {
				go func(ipHdr ipHeader) {
					ok, err := r.newTunnel(&ipHdr)
					if err != nil {
						xlog.Warnf("Unable to get tunnel to %s: %v",
							ipHdr.dstIP.String(), errors.Cause(err))
						return
					}
					if !ok {
						return
					}
					rw := r.getTunnel(ipHdr.dstAddr)
					r.connections++

					r.writeStream(rw, pkt[:plen], ipHdr)
				}(*ipHdr)
			} else {
				r.writeStream(rw, pkt[:plen], *ipHdr)
			}
		}
	}()

	<-r.networkInterface.closeCh
}

func (r *router) writeStream(rw *bufio.ReadWriter, pkt []byte, ipHdr ipHeader) {
	// real send

	// protocol version
	// protoVersion := []byte{protocolV1}

	// if _, err := rw.Write(protoVersion); err != nil {
	// 	// xlog.Errorf("rw.Write(): %v", err)
	// 	r.deleteTunnel(ipHdr.dstAddr)
	// 	return
	// }

	// hdr (packet length)
	hdr := make([]byte, 4)
	binary.BigEndian.PutUint32(hdr, uint32(len(pkt)))

	if _, err := rw.Write(hdr); err != nil {
		// xlog.Errorf("rw.Write(): %v", err)
		r.deleteTunnel(ipHdr.dstAddr)
		return
	}

	// packet
	if _, err := rw.Write(pkt); err != nil {
		// xlog.Errorf("rw.Write(): %v", err)
		r.deleteTunnel(ipHdr.dstAddr)
		return
	}

	if err := rw.Flush(); err != nil {
		// xlog.Errorf("rw.Flush(): %v", err)
		r.deleteTunnel(ipHdr.dstAddr)
		return
	}

	// udpate metrics
	go metrics.UpdateNetworkMetric(ipHdr.dstIP.String(), uint64(4+len(pkt)), 0, false)
}
