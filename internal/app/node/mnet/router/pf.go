package router

import (
	"net/netip"

	"github.com/google/gopacket/layers"
	"mmesh.dev/m-api-go/grpc/resources/topology"
	"mmesh.dev/m-lib/pkg/ipnet"
	"mmesh.dev/m-lib/pkg/xlog"
	"mmesh.dev/m-node/internal/app/node/mnet/router/conntrack"
)

func (r *router) packetFilter(conn *conntrack.Connection, pktlen int, egress bool) bool {
	// check ipDst rib entry
	if err := r.RIB().CheckIPDst(&conn.DstAddr); err != nil {
		xlog.Warnf("Unable to check IP dstAddr: %v", err)
		return true // drop the pkt
	}

	p := r.RIB().GetPolicy(r.subnetID)
	if p == nil {
		return true // no policy, drop the pkt
	}

	// filter specific proto packet types
	switch conn.Proto {
	case layers.IPProtocolICMPv4:
		if conn.ICMPv4TypeCode.Type() != layers.ICMPv4TypeEchoRequest {
			return true // only icmp echo request is permitted, drop the pkt
		}
	case layers.IPProtocolICMPv6:
		if conn.ICMPv6TypeCode.Type() != layers.ICMPv6TypeEchoRequest {
			return true // only icmp echo request is permitted, drop the pkt
		}
	}

	// check configured network filters
	for _, f := range p.NetworkFilters {
		var matchSrcIP, matchDstIP, matchProto, matchPort bool

		// proto
		if f.Proto == topology.Protocol_ANY || ipnet.IPProtocol(f.Proto.String()) == conn.Proto {
			// fmt.Printf("*** MATCHED Proto (%s): %s\n", f.Proto.String(), conn.proto.String())
			matchProto = true
		} else {
			continue
		}

		// srcIP
		srcIPNet, err := netip.ParsePrefix(f.SrcIPNet)
		if err != nil {
			return true // drop the pkt
		}

		// dstIP
		dstIPNet, err := netip.ParsePrefix(f.DstIPNet)
		if err != nil {
			return true // drop the pkt
		}

		if egress {
			// outgoing pkt
			if srcIPNet.Contains(conn.SrcIP) &&
				dstIPNet.Contains(conn.DstIP) &&
				(f.DstPort == 0 || f.DstPort == uint32(conn.DstPort)) {
				// fmt.Printf("*** MATCHED SrcIP (%s): %s\n", srcIPNet.String(), conn.SrcIP.String())
				// fmt.Printf("*** MATCHED DstIP (%s): %s\n", dstIPNet.String(), conn.DstIP.String())
				// fmt.Printf("*** MATCHED DstPort (%d): %d\n", f.DstPort, conn.DstPort)
				matchSrcIP = true
				matchDstIP = true
				matchDstIP = true
			}
		} else {
			// incoming pkt
			if srcIPNet.Contains(conn.SrcIP) &&
				dstIPNet.Contains(conn.DstIP) &&
				(f.DstPort == 0 || f.DstPort == uint32(conn.DstPort)) {
				// fmt.Printf("*** MATCHED SrcIP (%s): %s\n", srcIPNet.String(), conn.SrcIP.String())
				// fmt.Printf("*** MATCHED DstIP (%s): %s\n", dstIPNet.String(), conn.DstIP.String())
				// fmt.Printf("*** MATCHED DstPort (%d): %d\n", f.DstPort, conn.DstPort)
				matchSrcIP = true
				matchDstIP = true
				matchPort = true
			} else {
				// response pkt
				// if srcIPNet.Contains(conn.DstIP) &&
				// 	dstIPNet.Contains(conn.SrcIP) &&
				// 	(f.DstPort == 0 || f.DstPort == uint32(conn.SrcPort)) {
				// 	// fmt.Printf("*** [Response] MATCHED DstIP (%s): %s\n", srcIPNet.String(), conn.dstIP.String())
				// 	// fmt.Printf("*** [Response] MATCHED SrcIP (%s): %s\n", dstIPNet.String(), conn.srcIP.String())
				// 	// fmt.Printf("*** [Response] MATCHED SrcPort (%d): %d\n", f.DstPort, conn.srcPort)
				// 	matchSrcIP = true
				// 	matchDstIP = true
				// 	matchPort = true
				// }
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

	xlog.Warnf("Discarding packet from srcIP %s to dstIP %s\n", conn.SrcIP.String(), conn.DstIP.String())

	return true // drop the pkt
}
