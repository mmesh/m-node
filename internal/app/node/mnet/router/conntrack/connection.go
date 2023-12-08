package conntrack

import (
	"fmt"
	"net/netip"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"golang.org/x/net/ipv4"
	"golang.org/x/net/ipv6"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/ipnet"
)

type Connection struct {
	AF      ipnet.AddressFamily
	SrcIP   netip.Addr
	DstIP   netip.Addr
	DstAddr netip.Addr
	Proto   layers.IPProtocol
	SrcPort uint16
	DstPort uint16
	// ICMPv4ID       uint16
	ICMPv4TypeCode layers.ICMPv4TypeCode
	// ICMPv6ID       uint16
	ICMPv6TypeCode layers.ICMPv6TypeCode
	GREKey         uint32
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

	return true
}

func (conn *Connection) ParseProtocol(pkt []byte) error {
	var p gopacket.Packet

	switch conn.AF {
	case ipnet.AddressFamilyIPv4:
		p = gopacket.NewPacket(pkt, layers.LayerTypeIPv4, gopacket.Default)
	case ipnet.AddressFamilyIPv6:
		p = gopacket.NewPacket(pkt, layers.LayerTypeIPv6, gopacket.Default)
	default:
		return fmt.Errorf("unknown IP address family")
	}

	// Get the TCP layer from this packet
	if tcpLayer := p.Layer(layers.LayerTypeTCP); tcpLayer != nil {
		tcp, _ := tcpLayer.(*layers.TCP)
		conn.Proto = layers.IPProtocolTCP
		conn.SrcPort = uint16(tcp.SrcPort)
		conn.DstPort = uint16(tcp.DstPort)
		return nil
	}

	// Get the UDP layer from this packet
	if udpLayer := p.Layer(layers.LayerTypeUDP); udpLayer != nil {
		udp, _ := udpLayer.(*layers.UDP)
		conn.Proto = layers.IPProtocolUDP
		conn.SrcPort = uint16(udp.SrcPort)
		conn.DstPort = uint16(udp.DstPort)
		return nil
	}

	// Get the ICMPv4 layer from this packet
	if icmp4Layer := p.Layer(layers.LayerTypeICMPv4); icmp4Layer != nil {
		icmp4, _ := icmp4Layer.(*layers.ICMPv4)
		conn.Proto = layers.IPProtocolICMPv4
		conn.ICMPv4TypeCode = icmp4.TypeCode
		return nil
	}

	// Get the ICMPv6 layer from this packet
	if icmp6Layer := p.Layer(layers.LayerTypeICMPv6); icmp6Layer != nil {
		icmp6, _ := icmp6Layer.(*layers.ICMPv6)
		conn.Proto = layers.IPProtocolICMPv6
		conn.ICMPv6TypeCode = icmp6.TypeCode
		return nil
	}

	// Get the GRE layer from this packet
	if greLayer := p.Layer(layers.LayerTypeGRE); greLayer != nil {
		gre, _ := greLayer.(*layers.GRE)
		conn.Proto = layers.IPProtocolGRE
		conn.GREKey = gre.Key
		return nil
	}

	return nil
}

func (conn *Connection) outbound() Connection {
	return Connection{
		AF:      conn.AF,
		SrcIP:   conn.SrcIP,
		DstIP:   conn.DstIP,
		DstAddr: conn.DstAddr,
		Proto:   conn.Proto,
		SrcPort: conn.SrcPort,
		DstPort: conn.DstPort,
		// ICMPv4ID:       conn.ICMPv4ID,
		// ICMPv4TypeCode: conn.ICMPv4TypeCode,
		// ICMPv6ID:       conn.ICMPv6ID,
		// ICMPv6TypeCode: conn.ICMPv6TypeCode,
		GREKey: conn.GREKey,
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
		// ICMPv4ID:       conn.ICMPv4ID,
		// ICMPv4TypeCode: conn.ICMPv4TypeCode,
		// ICMPv6ID:       conn.ICMPv6ID,
		// ICMPv6TypeCode: conn.ICMPv6TypeCode,
		GREKey: conn.GREKey,
	}
}
