package rib

import (
	"strings"

	"mmesh.dev/m-api-go/grpc/network/nac"
	"mmesh.dev/m-api-go/grpc/network/routing"
)

func (r *ribData) DNSQuery(dnsName string) (ipv4, ipv6 string) {
	r.RLock()
	defer r.RUnlock()

	for addr, re := range r.rib.RoutingTable {
		if re.SubnetID != r.rib.RoutingDomain.SubnetID &&
			r.rib.RoutingDomain.Scope != nac.RoutingScope_NETWORK {
			continue
		}

		if re.DNSName == dnsName {
			ip := strings.Split(addr, "/")[0]

			switch re.AddressFamily {
			case routing.AddressFamily_IP4:
				if len(ipv4) == 0 {
					ipv4 = ip
				}
			case routing.AddressFamily_IP6:
				if len(ipv6) == 0 {
					ipv6 = ip
				}
			}
		}

		if len(ipv4) > 0 && len(ipv6) > 0 {
			break
		}
	}

	return ipv4, ipv6
}
