package ipnet

import (
	"fmt"
	"net"

	"github.com/google/gopacket/layers"
	"mmesh.dev/m-api-go/grpc/resources/topology"
)

func CheckNetworkPolicy(p *topology.Policy) error {
	for _, nf := range p.NetworkFilters {
		if err := checkNetworkFilter(nf); err != nil {
			return err
		}
	}

	return nil
}

func checkNetworkFilter(nf *topology.Filter) error {
	// srcIPNet check
	if _, _, err := net.ParseCIDR(nf.SrcIPNet); err != nil {
		return fmt.Errorf("INVALID networkPolicy srcIPNet %s: %v", nf.SrcIPNet, err)
	}

	// dstIPNet check
	if _, _, err := net.ParseCIDR(nf.DstIPNet); err != nil {
		return fmt.Errorf("INVALID networkPolicy dstIPNet %s: %v", nf.DstIPNet, err)
	}

	// proto check
	if !validIPProtocol(nf.Proto) {
		return fmt.Errorf("INVALID networkPolicy proto %s", nf.Proto)
	}

	// dstPort check
	if IPProtocol(nf.Proto.String()) == layers.IPProtocolTCP ||
		IPProtocol(nf.Proto.String()) == layers.IPProtocolUDP {

		if nf.DstPort < 1 || nf.DstPort > 65535 {
			return fmt.Errorf("INVALID NetworkPolicy dstPort %d", nf.DstPort)
		}
	}

	return nil
}
