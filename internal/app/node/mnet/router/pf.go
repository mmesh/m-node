package router

import (
	"net/netip"

	"mmesh.dev/m-api-go/grpc/resources/topology"
	"mmesh.dev/m-lib/pkg/ipnet"
	"mmesh.dev/m-lib/pkg/xlog"
	"mmesh.dev/m-node/internal/app/node/mnet/router/conntrack"
)

func (r *router) packetFilter(conn *conntrack.Connection, pktlen int) bool {
	// check ipDst rib entry
	if err := r.RIB().CheckIPDst(&conn.DstAddr); err != nil {
		xlog.Warnf("[pf] Unable to check IP dstAddr: %v", err)
		return true // drop the pkt
	}

	// check configured network filters
	if r.policyFilter(conn) {
		// packet dropped by policy

		// check conntrack table
		if !conn.IsActive(pktlen) {
			// packet dropped by conntrack
			xlog.Warnf("[pf] Dropping %s packet from %s:%d to %s:%d",
				conn.Proto.String(),
				conn.SrcIP.String(),
				conn.SrcPort,
				conn.DstAddr.String(),
				conn.DstPort,
			)
			return true // drop the pkt
		}
	}

	return false // accept the pkt
}

func (r *router) policyFilter(conn *conntrack.Connection) bool {
	// filter specific proto packet types
	if conn.ProtoFilter() {
		return true // invalid proto packet type, drop the pkt
	}

	// get network policy
	p := r.RIB().GetPolicy(r.subnetID)
	if p == nil {
		return true // no policy, drop the pkt
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

		if srcIPNet.Contains(conn.SrcIP) &&
			dstIPNet.Contains(conn.DstIP) &&
			(f.DstPort == 0 || f.DstPort == uint32(conn.DstPort)) {
			// fmt.Printf("*** MATCHED SrcIP (%s): %s\n", srcIPNet.String(), conn.SrcIP.String())
			// fmt.Printf("*** MATCHED DstIP (%s): %s\n", dstIPNet.String(), conn.DstIP.String())
			// fmt.Printf("*** MATCHED DstPort (%d): %d\n", f.DstPort, conn.DstPort)
			matchSrcIP = true
			matchDstIP = true
			matchPort = true
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

	return true // drop the pkt
}
