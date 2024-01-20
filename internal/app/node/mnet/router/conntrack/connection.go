package conntrack

import (
	"fmt"
	"net/netip"

	"github.com/google/gopacket/layers"
	"golang.org/x/net/ipv4"
	"golang.org/x/net/ipv6"
	"mmesh.dev/m-api-go/grpc/resources/nstore/netdb"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/ipnet"
)

type Connection struct {
	AF        ipnet.AddressFamily
	SrcIP     netip.Addr
	DstIP     netip.Addr
	DstAddr   netip.Addr
	Proto     layers.IPProtocol
	SrcPort   uint16
	DstPort   uint16
	protoInfo *protoInfo
}

type protoInfo struct {
	tcp   *layers.TCP
	udp   *layers.UDP
	icmp4 *layers.ICMPv4
	icmp6 *layers.ICMPv6
	gre   *layers.GRE
	sctp  *layers.SCTP
}

func ParseHeader(pkt []byte) (*Connection, error) {
	header, err := ipv4.ParseHeader(pkt)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function ipv4.ParseHeader()", errors.Trace())
	}

	switch ipnet.AddressFamily(header.Version) {
	case ipnet.AddressFamilyIPv4:
		srcIP, err := netip.ParseAddr(header.Src.String())
		if err != nil {
			return nil, errors.Wrapf(err, "[%v] function netip.ParseAddr()", errors.Trace())
		}

		dstIP, err := netip.ParseAddr(header.Dst.String())
		if err != nil {
			return nil, errors.Wrapf(err, "[%v] function netip.ParseAddr()", errors.Trace())
		}

		return &Connection{
			AF:      ipnet.AddressFamilyIPv4,
			SrcIP:   srcIP,
			DstIP:   dstIP,
			DstAddr: dstIP,
		}, nil
	case ipnet.AddressFamilyIPv6:
		header, err := ipv6.ParseHeader(pkt)
		if err != nil {
			return nil, errors.Wrapf(err, "[%v] function ipv6.ParseHeader()", errors.Trace())
		}

		srcIP, err := netip.ParseAddr(header.Src.String())
		if err != nil {
			return nil, errors.Wrapf(err, "[%v] function netip.ParseAddr()", errors.Trace())
		}

		dstIP, err := netip.ParseAddr(header.Dst.String())
		if err != nil {
			return nil, errors.Wrapf(err, "[%v] function netip.ParseAddr()", errors.Trace())
		}

		dstAddr, err := ipnet.GetIPv6Endpoint(dstIP.String())
		if err != nil {
			dstAddr = dstIP.String()
			// xlog.Warnf("Unable to get valid pkt dst addr: %v", err)
			// return nil, err
		}

		dstNetIPAddr, err := netip.ParseAddr(dstAddr)
		if err != nil {
			return nil, errors.Wrapf(err, "[%v] function netip.ParseAddr()", errors.Trace())
		}

		return &Connection{
			AF:      ipnet.AddressFamilyIPv6,
			SrcIP:   srcIP,
			DstIP:   dstIP,
			DstAddr: dstNetIPAddr,
		}, nil
	}

	return nil, fmt.Errorf("unknown IP address family")
}

func (conn *Connection) IsValid(localIPv6 string) bool {
	// drop non-unicast
	if !conn.DstAddr.IsGlobalUnicast() {
		return false // drop the pkt
	}
	if !conn.SrcIP.IsGlobalUnicast() {
		return false // drop the pkt
	}

	// drop icmpv6 redirects
	if conn.DstAddr.String() == localIPv6 {
		return false // drop the pkt
	}

	return true // accept the pkt
}

func (conn *Connection) IsActive(pktlen int) bool {
	if conntrack == nil {
		return false
	}

	if !conntrack.isActiveConnection(conn, uint64(pktlen)) {
		return false
	}

	if conn.Proto == layers.IPProtocolICMPv4 || conn.Proto == layers.IPProtocolICMPv6 {
		if conn.invalidICMPTypeReply() {
			return false // only icmp echo reply is permitted, drop the pkt
		}
	}

	// store netflow
	if nfMap == nil {
		newNetflowMap()
	}
	nfMap.inboundConnection(conn, uint64(pktlen))

	return true
}

func (conn *Connection) direct() Connection {
	return Connection{
		AF:      conn.AF,
		SrcIP:   conn.SrcIP,
		DstIP:   conn.DstIP,
		DstAddr: conn.DstAddr,
		Proto:   conn.Proto,
		SrcPort: conn.SrcPort,
		DstPort: conn.DstPort,
	}
}

func (conn *Connection) reverse() Connection {
	return Connection{
		AF:      conn.AF,
		SrcIP:   conn.DstIP,
		DstIP:   conn.SrcIP,
		DstAddr: conn.SrcIP,
		Proto:   conn.Proto,
		SrcPort: conn.DstPort,
		DstPort: conn.SrcPort,
	}
}

func (conn *Connection) outbound() Connection {
	return conn.direct()
}

func (conn *Connection) flow() Connection {
	return conn.direct()
}

func (conn *Connection) GetAddressFamily() netdb.AddressFamily {
	switch conn.AF {
	case ipnet.AddressFamilyIPv4:
		return netdb.AddressFamily_IP4
	case ipnet.AddressFamilyIPv6:
		return netdb.AddressFamily_IP6
	}

	return netdb.AddressFamily_UNKNOWN_AF
}

func (conn *Connection) GetProtocol() netdb.Protocol {
	switch conn.Proto {
	case layers.IPProtocolTCP:
		return netdb.Protocol_TCP
	case layers.IPProtocolUDP:
		return netdb.Protocol_UDP
	case layers.IPProtocolICMPv4:
		return netdb.Protocol_ICMP4
	case layers.IPProtocolICMPv6:
		return netdb.Protocol_ICMP6
	case layers.IPProtocolGRE:
		return netdb.Protocol_GRE
	case layers.IPProtocolSCTP:
		return netdb.Protocol_SCTP
	}

	return netdb.Protocol_UNKNOWN_PROTO
}

func (conn *Connection) GetNetConnection() *netdb.Connection {
	return &netdb.Connection{
		AF:      conn.GetAddressFamily(),
		SrcIP:   conn.SrcIP.String(),
		DstIP:   conn.DstAddr.String(),
		Proto:   conn.GetProtocol(),
		SrcPort: uint32(conn.SrcPort),
		DstPort: uint32(conn.DstPort),
	}
}
