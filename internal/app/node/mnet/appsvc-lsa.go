package mnet

import (
	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-api-go/grpc/network/routing"
	"mmesh.dev/m-lib/pkg/mmp/queuing"
)

func (ln *localNode) SendAppSvcLSAs(mmID string) {
	if ln == nil {
		return
	}

	if ln.Node().Cfg.DisableNetworking || ln.Router() == nil {
		return
	}

	if !ln.initialized {
		return
	}

	n := ln.Node()
	if n == nil {
		return
	}

	for _, as := range ln.Router().RIB().GetNodeAppSvcs() {
		lsa := &routing.LSA{
			Type: routing.LSAType_APPSVC_LSA,
			AppSvcLSA: &routing.AppSvcLSA{
				AppSvc:      as,
				P2PHostID:   n.Agent.P2PHostID,
				Priority:    n.Cfg.Priority,
				IPv6:        ln.Router().GlobalIPv6(),
				Connections: ln.Router().GetConnections(),
			},
		}

		queuing.TxControlQueue <- &mmsp.Payload{
			SrcID: mmID,
			Type:  mmsp.PDUType_ROUTING,
			RoutingPDU: &mmsp.RoutingPDU{
				Type: mmsp.RoutingMsgType_ROUTING_LSA,
				LSA:  lsa,
			},
		}
	}
}
