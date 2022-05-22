package router

import (
	"fmt"
	"net"
	"strings"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"golang.org/x/net/ipv4"
	"golang.org/x/net/ipv6"
	"mmesh.dev/m-lib/pkg/ipnet"
	"mmesh.dev/m-lib/pkg/xlog"
)

type ipPacket struct {
	af        ipnet.AddressFamily
	srcIP     net.IP
	dstIP     net.IP
	dstAddr   string
	proto     layers.IPProtocol
	srcPort   uint16
	dstPort   uint16
	icmp4Type layers.ICMPv4TypeCode
	icmp6Type layers.ICMPv6TypeCode
	data      []byte
	plen      int
}

func (r *router) packetFilter(packet []byte) (*ipPacket, bool) {
	// decode packet
	ipPkt, err := decodeIPPacket(packet)
	if err != nil {
		return nil, true // drop the pkt
	}

	// drop non-unicast
	if !net.ParseIP(ipPkt.dstAddr).IsGlobalUnicast() {
		return nil, true // drop the pkt
	}
	if !ipPkt.srcIP.IsGlobalUnicast() {
		return nil, true // drop the pkt
	}

	// drop icmpv6 redirects
	if ipPkt.dstAddr == r.ipv6 {
		return nil, true // drop the pkt
	}

	// decode protocol layers
	if err := ipPkt.decodeIPLayers(); err != nil {
		return ipPkt, true // drop the pkt
	}

	// analysing decoded ip packet

	r.rib.RLock()
	defer r.rib.RUnlock()

	// get route
	ipDst := r.getIPDstFromRIB(ipPkt.dstAddr)
	if len(ipDst) == 0 {
		return ipPkt, true // drop the pkt
	}

	p, ok := r.rib.global.Policy[r.vrfID]
	if !ok {
		return ipPkt, true // no policy, drop the pkt
	}

	// if p == nil {
	// 	return ipPkt, true // drop the pkt
	// }

	for _, f := range p.NetworkFilters {
		var matchSrcIP, matchDstIP, matchProto, matchPort bool

		_, srcIPNet, err := net.ParseCIDR(f.SrcIPNet)
		if err != nil {
			return ipPkt, true // drop the pkt
		}
		if srcIPNet.Contains(ipPkt.srcIP) {
			matchSrcIP = true
		}

		_, dstIPNet, err := net.ParseCIDR(f.DstIPNet)
		if err != nil {
			return ipPkt, true // drop the pkt
		}
		if dstIPNet.Contains(ipPkt.dstIP) {
			matchDstIP = true
		}

		if strings.ToUpper(f.Proto) == "ANY" || ipnet.IPProtocol(f.Proto) == ipPkt.proto {
			matchProto = true
		}
		if f.DstPort == 0 || f.DstPort == int32(ipPkt.dstPort) {
			matchPort = true
		}

		if matchSrcIP && matchDstIP && matchProto && matchPort {
			switch strings.ToUpper(f.Policy) {
			case ipnet.Policy_ACCEPT:
				return ipPkt, false // accept the pkt
			case ipnet.Policy_DROP:
				return ipPkt, true // drop the pkt
			}
		}
	}

	switch strings.ToUpper(p.DefaultPolicy) {
	case ipnet.Policy_ACCEPT:
		return ipPkt, false // accept the pkt
	case ipnet.Policy_DROP:
		return ipPkt, true // drop the pkt
	}

	xlog.Warnf("Discarding packet from srcIP %s to dstIP %s\n", ipPkt.srcIP.String(), ipPkt.dstIP.String())

	return ipPkt, true // drop the pkt
}

func decodeIPPacket(pkt []byte) (*ipPacket, error) {
	header, err := ipv4.ParseHeader(pkt)
	if err != nil {
		xlog.Warnf("Unable to parse ip header: %v", err)
		return nil, err
	}

	switch ipnet.AddressFamily(header.Version) {
	case ipnet.AddressFamilyIPv4:
		// fmt.Println(header.String())

		return &ipPacket{
			af:      ipnet.AddressFamilyIPv4,
			srcIP:   header.Src,
			dstIP:   header.Dst,
			dstAddr: header.Dst.String(),
			data:    pkt,
			plen:    len(pkt),
		}, nil
	case ipnet.AddressFamilyIPv6:
		header, err := ipv6.ParseHeader(pkt)
		if err != nil {
			xlog.Warnf("Unable to parse ipv6 header: %v", err)
			return nil, err
		}

		// fmt.Println(header.String())

		dstAddr, err := ipnet.GetIPv6Endpoint(header.Dst.String())
		if err != nil {
			dstAddr = header.Dst.String()
			// xlog.Warnf("Unable to get valid pkt dst addr: %v", err)
			// return nil, err
		}

		return &ipPacket{
			af:      ipnet.AddressFamilyIPv6,
			srcIP:   header.Src,
			dstIP:   header.Dst,
			dstAddr: dstAddr,
			data:    pkt,
			plen:    len(pkt),
		}, nil
	}

	return nil, fmt.Errorf("unable to parse packet")
}

func (ipPkt *ipPacket) decodeIPLayers() error {
	var p gopacket.Packet

	switch ipPkt.af {
	case ipnet.AddressFamilyIPv4:
		p = gopacket.NewPacket(ipPkt.data, layers.LayerTypeIPv4, gopacket.Default)
	case ipnet.AddressFamilyIPv6:
		p = gopacket.NewPacket(ipPkt.data, layers.LayerTypeIPv6, gopacket.Default)
	default:
		return fmt.Errorf("unable to parse packet")
	}

	// Get the TCP layer from this packet
	if tcpLayer := p.Layer(layers.LayerTypeTCP); tcpLayer != nil {
		tcp, _ := tcpLayer.(*layers.TCP)
		ipPkt.proto = layers.IPProtocolTCP
		ipPkt.srcPort = uint16(tcp.SrcPort)
		ipPkt.dstPort = uint16(tcp.DstPort)
	}

	// Get the UDP layer from this packet
	if udpLayer := p.Layer(layers.LayerTypeUDP); udpLayer != nil {
		udp, _ := udpLayer.(*layers.UDP)
		ipPkt.proto = layers.IPProtocolUDP
		ipPkt.srcPort = uint16(udp.SrcPort)
		ipPkt.dstPort = uint16(udp.DstPort)
	}

	// Get the ICMPv4 layer from this packet
	if icmp4Layer := p.Layer(layers.LayerTypeICMPv4); icmp4Layer != nil {
		icmp4, _ := icmp4Layer.(*layers.ICMPv4)
		ipPkt.proto = layers.IPProtocolICMPv4
		ipPkt.icmp4Type = icmp4.TypeCode
	}

	// Get the ICMPv6 layer from this packet
	if icmp6Layer := p.Layer(layers.LayerTypeICMPv6); icmp6Layer != nil {
		icmp6, _ := icmp6Layer.(*layers.ICMPv6)
		ipPkt.proto = layers.IPProtocolICMPv6
		ipPkt.icmp6Type = icmp6.TypeCode
	}

	return nil
}
