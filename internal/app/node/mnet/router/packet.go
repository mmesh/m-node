package router

import (
	"fmt"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"

	// "github.com/sanity-io/litter"
	"golang.org/x/net/ipv4"
	"golang.org/x/net/ipv6"
	"mmesh.dev/m-api-go/grpc/resources/topology"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/ipnet"
	"mmesh.dev/m-lib/pkg/xlog"
)

type ipHeader struct {
	af        ipnet.AddressFamily
	srcIP     net.IP
	dstIP     net.IP
	dstAddr   string
	proto     layers.IPProtocol
	srcPort   uint16
	dstPort   uint16
	icmp4Type layers.ICMPv4TypeCode
	icmp6Type layers.ICMPv6TypeCode
}

func parseHeader(pkt []byte) (*ipHeader, error) {
	header, err := ipv4.ParseHeader(pkt)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function ipv4.ParseHeader()", errors.Trace())
	}

	switch ipnet.AddressFamily(header.Version) {
	case ipnet.AddressFamilyIPv4:
		// fmt.Println("::::: IPv4 Packet HEADER")
		// fmt.Println(header.String())

		return &ipHeader{
			af:      ipnet.AddressFamilyIPv4,
			srcIP:   header.Src,
			dstIP:   header.Dst,
			dstAddr: header.Dst.String(),
		}, nil
	case ipnet.AddressFamilyIPv6:
		header, err := ipv6.ParseHeader(pkt)
		if err != nil {
			return nil, errors.Wrapf(err, "[%v] function ipv6.ParseHeader()", errors.Trace())
		}

		// fmt.Println("::::: IPv6 Packet HEADER")
		// fmt.Println(header.String())

		dstAddr, err := ipnet.GetIPv6Endpoint(header.Dst.String())
		if err != nil {
			dstAddr = header.Dst.String()
			// xlog.Warnf("Unable to get valid pkt dst addr: %v", err)
			// return nil, err
		}

		return &ipHeader{
			af:      ipnet.AddressFamilyIPv6,
			srcIP:   header.Src,
			dstIP:   header.Dst,
			dstAddr: dstAddr,
		}, nil
	}

	return nil, fmt.Errorf("unknown IP address family")
}

func (ipHdr *ipHeader) parseProtocol(pkt []byte) error {
	var p gopacket.Packet

	switch ipHdr.af {
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
		ipHdr.proto = layers.IPProtocolTCP
		ipHdr.srcPort = uint16(tcp.SrcPort)
		ipHdr.dstPort = uint16(tcp.DstPort)
	}

	// Get the UDP layer from this packet
	if udpLayer := p.Layer(layers.LayerTypeUDP); udpLayer != nil {
		udp, _ := udpLayer.(*layers.UDP)
		ipHdr.proto = layers.IPProtocolUDP
		ipHdr.srcPort = uint16(udp.SrcPort)
		ipHdr.dstPort = uint16(udp.DstPort)
	}

	// Get the ICMPv4 layer from this packet
	if icmp4Layer := p.Layer(layers.LayerTypeICMPv4); icmp4Layer != nil {
		icmp4, _ := icmp4Layer.(*layers.ICMPv4)
		ipHdr.proto = layers.IPProtocolICMPv4
		ipHdr.icmp4Type = icmp4.TypeCode
	}

	// Get the ICMPv6 layer from this packet
	if icmp6Layer := p.Layer(layers.LayerTypeICMPv6); icmp6Layer != nil {
		icmp6, _ := icmp6Layer.(*layers.ICMPv6)
		ipHdr.proto = layers.IPProtocolICMPv6
		ipHdr.icmp6Type = icmp6.TypeCode
	}

	return nil
}

func (ipHdr *ipHeader) isValidPacket(localIPv6 string) bool {
	// drop non-unicast
	if !net.ParseIP(ipHdr.dstAddr).IsGlobalUnicast() {
		return false // drop the pkt
	}
	if !ipHdr.srcIP.IsGlobalUnicast() {
		return false // drop the pkt
	}

	// drop icmpv6 redirects
	if ipHdr.dstAddr == localIPv6 {
		return false // drop the pkt
	}

	return true
}

func (r *router) packetFilter(ipHdr *ipHeader, pkt []byte, egress bool) bool {
	// // parse ip header
	// ipHdr, err := parseHeader(pkt)
	// if err != nil {
	// 	xlog.Warnf("Unable to parse IP header: %v", errors.Cause(err))
	// 	return nil, true // drop the pkt
	// }

	// // drop non-unicast
	// if !net.ParseIP(ipHdr.dstAddr).IsGlobalUnicast() {
	// 	return nil, true // drop the pkt
	// }
	// if !ipHdr.srcIP.IsGlobalUnicast() {
	// 	return nil, true // drop the pkt
	// }

	// // drop icmpv6 redirects
	// if ipHdr.dstAddr == r.ipv6 {
	// 	return nil, true // drop the pkt
	// }

	// // parse ip protocol
	// if err := ipHdr.parseProtocol(pkt); err != nil {
	// 	xlog.Warnf("Unable to parse IP protocol: %v", errors.Cause(err))
	// 	return true // drop the pkt
	// }

	// check decoded ip packet

	// check ipDst rib entry
	if err := r.RIB().CheckIPDst(ipHdr.dstAddr); err != nil {
		xlog.Warnf("Unable to check IP dstAddr: %v", err)
		return true // drop the pkt
	}

	p := r.RIB().GetPolicy(r.subnetID)
	if p == nil {
		return true // no policy, drop the pkt
	}

	for _, f := range p.NetworkFilters {
		var matchSrcIP, matchDstIP, matchProto, matchPort bool

		// litter.Dump(f)
		// litter.Dump(ipHdr)

		// srcIP
		_, srcIPNet, err := net.ParseCIDR(f.SrcIPNet)
		if err != nil {
			return true // drop the pkt
		}

		if egress {
			// outgoing pkt
			if srcIPNet.Contains(ipHdr.srcIP) {
				// fmt.Printf("*** MATCHED SrcIP (%s): %s\n", srcIPNet.String(), ipHdr.srcIP.String())
				matchSrcIP = true
			}
		} else { // ingress
			// response pkt
			if srcIPNet.Contains(ipHdr.dstIP) {
				// fmt.Printf("*** [Response] MATCHED DstIP (%s): %s\n", srcIPNet.String(), ipHdr.dstIP.String())
				matchSrcIP = true
			}
		}

		// dstIP
		_, dstIPNet, err := net.ParseCIDR(f.DstIPNet)
		if err != nil {
			return true // drop the pkt
		}

		if egress {
			// outgoing pkt
			if dstIPNet.Contains(ipHdr.dstIP) {
				// fmt.Printf("*** MATCHED DstIP (%s): %s\n", dstIPNet.String(), ipHdr.dstIP.String())
				matchDstIP = true
			}
		} else { // ingress
			// response pkt
			if dstIPNet.Contains(ipHdr.srcIP) {
				// fmt.Printf("*** [Response] MATCHED SrcIP (%s): %s\n", dstIPNet.String(), ipHdr.srcIP.String())
				matchDstIP = true
			}
		}

		// proto
		if f.Proto == topology.Protocol_ANY || ipnet.IPProtocol(f.Proto.String()) == ipHdr.proto {
			// fmt.Printf("*** MATCHED Proto (%s): %s\n", f.Proto.String(), ipHdr.proto.String())
			matchProto = true
		}

		// dstPort
		if egress {
			// outgoing pkt
			if f.DstPort == 0 || f.DstPort == uint32(ipHdr.dstPort) {
				// fmt.Printf("*** MATCHED DstPort (%d): %d\n", f.DstPort, ipHdr.dstPort)
				matchPort = true
			}
		} else { // ingress
			// response pkt
			if f.DstPort == 0 || f.DstPort == uint32(ipHdr.srcPort) {
				// fmt.Printf("*** [Response] MATCHED SrcPort (%d): %d\n", f.DstPort, ipHdr.srcPort)
				matchPort = true
			}
		}

		if matchSrcIP && matchDstIP && matchProto && matchPort {
			// fmt.Printf("*** MATCHED Policy: %s\n", f.Policy.String())

			switch f.Policy {
			case topology.SecurityPolicy_ACCEPT:
				return false // accept the pkt
			case topology.SecurityPolicy_DROP:
				return true // drop the pkt
			}
		}
	}

	// fmt.Printf("+++ DEFAULT Policy: %s\n", p.DefaultPolicy.String())

	switch p.DefaultPolicy {
	case topology.SecurityPolicy_ACCEPT:
		return false // accept the pkt
	case topology.SecurityPolicy_DROP:
		return true // drop the pkt
	}

	xlog.Warnf("Discarding packet from srcIP %s to dstIP %s\n", ipHdr.srcIP.String(), ipHdr.dstIP.String())

	return true // drop the pkt
}
