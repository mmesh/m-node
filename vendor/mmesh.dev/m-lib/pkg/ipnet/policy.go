package ipnet

import (
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/google/gopacket/layers"
	"mmesh.dev/m-api-go/grpc/network/resources/network"
)

const (
	Policy_ACCEPT string = "ACCEPT"
	Policy_DROP   string = "DROP"
)

func CheckNetworkPolicy(p *network.Policy) error {
	if err := checkPolicy(p.DefaultPolicy); err != nil {
		return err
	}

	for _, nf := range p.NetworkFilters {
		if err := checkNetworkFilter(nf); err != nil {
			return err
		}
	}

	return nil
}

func checkPolicy(policy string) error {
	// policy check
	p := strings.ToUpper(policy)
	if p != Policy_ACCEPT && p != Policy_DROP {
		return fmt.Errorf("INVALID networkPolicy policy %s", policy)
	}

	return nil
}

func checkNetworkFilter(nf *network.Filter) error {
	nf.LastModified = time.Now().Unix()

	// policy check
	if err := checkPolicy(nf.Policy); err != nil {
		return err
	}

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
	if IPProtocol(nf.Proto) == layers.IPProtocolTCP ||
		IPProtocol(nf.Proto) == layers.IPProtocolUDP {

		if nf.DstPort < 1 || nf.DstPort > 65535 {
			return fmt.Errorf("INVALID NetworkPolicy dstPort %d", nf.DstPort)
		}
	}

	return nil
}
