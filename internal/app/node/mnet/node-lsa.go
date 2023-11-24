package mnet

import (
	"github.com/google/uuid"
	"mmesh.dev/m-api-go/grpc/network/routing"
)

func (ln *localNode) GetNodeLSA() *routing.LSA {
	if ln == nil {
		return nil
	}

	if ln.Node().Cfg.DisableNetworking || ln.Router() == nil {
		return nil
	}

	if !ln.initialized {
		return nil
	}

	n := ln.Node()
	if n == nil {
		return nil
	}

	rd := ln.Router().RIB().RoutingDomain()
	if rd == nil {
		return nil
	}

	endpoints := make(map[string]*routing.IPAddr, 0)

	for _, e := range n.Endpoints {
		dnsName := uuid.New().String()

		if len(e.DNSName) > 0 {
			dnsName = e.DNSName
		}

		endpoints[dnsName] = &routing.IPAddr{
			IPv4: e.IPv4,
			IPv6: e.IPv6,
		}
	}

	isRelay := false
	if n.Agent.CanRelay && !n.Cfg.DisableRelay {
		isRelay = true
	}

	lsa := &routing.LSA{
		Type: routing.LSAType_NODE_LSA,
		NodeLSA: &routing.NodeLSA{
			AccountID:      n.AccountID,
			TenantID:       n.TenantID,
			NetID:          n.Cfg.NetID,
			SubnetID:       n.Cfg.SubnetID,
			NodeID:         n.NodeID,
			NetworkCIDR:    rd.NetworkCIDR,
			SubnetCIDR:     rd.SubnetCIDR,
			P2PHostID:      n.Agent.P2PHostID,
			MAddrs:         n.Agent.MAddrs,
			ExternalIPv4:   n.Agent.ExternalIPv4,
			Port:           n.Agent.Port,
			Priority:       n.Cfg.Priority,
			IsRelay:        isRelay,
			Endpoints:      endpoints,
			ExportedRoutes: n.Agent.Routes.Export,
			Connections:    ln.Router().GetConnections(),
		},
	}

	return lsa
}
