package conntrack

import (
	"time"

	"mmesh.dev/m-api-go/grpc/resources/nstore/netdb"
)

var RequestQueue = make(chan *netdb.ConntrackTableRequest, 128)

type Interface interface {
	OutboundConnection(c *Connection, pktlen int)
	Close()

	GetTable(r *netdb.ConntrackTableRequest) *netdb.ConntrackTableResponse
}
type api struct{}

func Ctrl() Interface {
	return &api{}
}

func (api *api) OutboundConnection(c *Connection, pktlen int) {
	if conntrack == nil {
		newMap()
	}
	conntrack.outboundConnection(c, uint64(pktlen))

	// store netflow
	if nfMap == nil {
		newNetflowMap()
	}
	nfMap.outboundConnection(c, uint64(pktlen))
}

func (api *api) Close() {
	if conntrack != nil {
		conntrack.closeCh <- struct{}{}
	}

	if nfMap != nil {
		nfMap.closeCh <- struct{}{}
	}
}

func (api *api) GetTable(r *netdb.ConntrackTableRequest) *netdb.ConntrackTableResponse {
	cttr := &netdb.ConntrackTableResponse{
		AccountID: r.Request.AccountID,
		TenantID:  r.Request.TenantID,
		NodeID:    r.Request.NodeID,
		QueryID:   r.Request.QueryID,
		CtTable:   make([]*netdb.ConntrackEntry, 0),
		Timestamp: time.Now().UnixMilli(),
	}

	if conntrack == nil {
		return cttr
	}

	for c, state := range conntrack.getTable() {
		cttr.CtTable = append(cttr.CtTable, &netdb.ConntrackEntry{
			Timestamp: time.Now().UnixMilli(),
			Connection: &netdb.Connection{
				AF:      c.GetAddressFamily(),
				SrcIP:   c.SrcIP.String(),
				DstIP:   c.DstAddr.String(),
				Proto:   c.GetProtocol(),
				SrcPort: uint32(c.SrcPort),
				DstPort: uint32(c.DstPort),
			},
			State: state,
		})
	}

	return cttr
}
