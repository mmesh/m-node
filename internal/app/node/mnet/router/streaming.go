package router

import (
	"bufio"
	"encoding/binary"
	"io"

	"github.com/google/gopacket/layers"
	"github.com/libp2p/go-libp2p/core/network"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/ipnet"
	"mmesh.dev/m-lib/pkg/xlog"
	"mmesh.dev/m-node/internal/app/node/metrics"
	"mmesh.dev/m-node/internal/app/node/mnet/p2p/conn"
	"mmesh.dev/m-node/internal/app/node/mnet/router/conntrack"
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
	// parse ip connection
	conn, err := conntrack.ParseHeader(pkt)
	if err != nil {
		xlog.Warnf("Unable to parse IP header: %v", errors.Cause(err))
		return nil
	}

	// parse ip protocol
	if err := conn.ParseProtocol(pkt); err != nil {
		xlog.Warnf("Unable to parse IP protocol: %v", errors.Cause(err))
		return nil
	}

	// packet filtering (firewalling) -- ingress
	if r.packetFilter(conn, len(pkt), false) {
		// packet dropped by policy

		dropPkt := false

		// check conntrack table
		if !conntrack.Ctrl().IsValidConnection(conn, len(pkt)) {
			xlog.Warnf("[conntrack] Dropping %s connection from %s:%d to %s:%d",
				conn.Proto.String(),
				conn.SrcIP.String(),
				conn.SrcPort,
				conn.DstAddr.String(),
				conn.DstPort,
			)
			// packet dropped because it's not present in conntrack table
			dropPkt = true
		}

		// filter specific proto packet types
		switch conn.Proto {
		case layers.IPProtocolICMPv4:
			if conn.ICMPv4TypeCode.Type() != layers.ICMPv4TypeEchoReply {
				dropPkt = true // only icmp echo request is permitted, drop the pkt
			}
		case layers.IPProtocolICMPv6:
			if conn.ICMPv6TypeCode.Type() != layers.ICMPv6TypeEchoReply {
				dropPkt = true // only icmp echo request is permitted, drop the pkt
			}
		}

		if dropPkt {
			go func() {
				// udpate metrics with a new drop
				go metrics.UpdateNetworkMetric(conn.SrcIP.String(), 0, 0, true)
			}()
			return nil
		}
	}

	r.setTunnel(conn.SrcIP, rw)

	// proxy64 forwarding
	if conn.AF == ipnet.AddressFamilyIPv6 {
		if r.proxy64Forward(conn, pkt) {
			return nil
		}
	}

	// write to TUN interface
	if r.networkInterface != nil {
		if _, err := r.networkInterface.write(pkt); err != nil {
			return err
		}

		// update metrics
		go metrics.UpdateNetworkMetric(conn.SrcIP.String(), 0, uint64(len(pkt)), false)
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

			conn, err := conntrack.ParseHeader(pkt[:plen])
			if err != nil {
				xlog.Warnf("Unable to parse IP header: %v", errors.Cause(err))
				continue
			}

			if !conn.IsValid(r.ipv6) {
				continue
			}

			// parse ip protocol
			if err := conn.ParseProtocol(pkt); err != nil {
				xlog.Warnf("Unable to parse IP protocol: %v", errors.Cause(err))
				continue
			}

			conntrack.Ctrl().OutboundConnection(conn, plen)

			// proxy64 forwarding
			if conn.AF == ipnet.AddressFamilyIPv6 {
				if r.proxy64Forward(conn, pkt[:plen]) {
					continue
				}
			}

			rw := r.getTunnel(conn.DstAddr)
			if rw == nil {
				go func(conn conntrack.Connection) {
					ok, err := r.newTunnel(&conn)
					if err != nil {
						xlog.Warnf("Unable to get tunnel to %s: %v",
							conn.DstIP.String(), errors.Cause(err))
						return
					}
					if !ok {
						return
					}
					rw := r.getTunnel(conn.DstAddr)
					r.connections++

					r.writeStream(rw, pkt[:plen], conn)
				}(*conn)
			} else {
				r.writeStream(rw, pkt[:plen], *conn)
			}
		}
	}()

	<-r.networkInterface.closeCh
}

func (r *router) writeStream(rw *bufio.ReadWriter, pkt []byte, conn conntrack.Connection) {
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
		r.deleteTunnel(conn.DstAddr)
		return
	}

	// packet
	if _, err := rw.Write(pkt); err != nil {
		// xlog.Errorf("rw.Write(): %v", err)
		r.deleteTunnel(conn.DstAddr)
		return
	}

	if err := rw.Flush(); err != nil {
		// xlog.Errorf("rw.Flush(): %v", err)
		r.deleteTunnel(conn.DstAddr)
		return
	}

	// udpate metrics
	go metrics.UpdateNetworkMetric(conn.DstIP.String(), uint64(4+len(pkt)), 0, false)
}
