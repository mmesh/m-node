package conntrack

import (
	"sync"
	"time"

	"mmesh.dev/m-api-go/grpc/resources/nstore/netdb"
	"mmesh.dev/m-node/internal/app/node/kvstore/db/netflowdb"
)

type netflowState struct {
	direction netdb.ConnectionDirection
	traffic   *netdb.TrafficCounter
}

type netflowMap struct {
	table   map[Connection]*netflowState
	closeCh chan struct{}
	sync.RWMutex
}

var nfMap *netflowMap

func newNetflowMap() {
	nfMap = &netflowMap{
		table:   make(map[Connection]*netflowState, 0),
		closeCh: make(chan struct{}, 1),
	}

	go nfMap.wrkr()
}

func (nfm *netflowMap) outboundConnection(c *Connection, bytes uint64) {
	if c == nil {
		return
	}

	nfm.Lock()
	defer nfm.Unlock()

	conn := c.flow()

	s, ok := nfm.table[conn]
	if !ok {
		s = &netflowState{
			direction: netdb.ConnectionDirection_OUTGOING,
			traffic:   &netdb.TrafficCounter{},
		}
		nfm.table[conn] = s
	}

	s.traffic.Packets++
	s.traffic.Bytes += bytes
}

func (nfm *netflowMap) inboundConnection(c *Connection, bytes uint64) {
	if c == nil {
		return
	}

	nfm.Lock()
	defer nfm.Unlock()

	conn := c.flow()

	s, ok := nfm.table[conn]
	if !ok {
		s = &netflowState{
			direction: netdb.ConnectionDirection_INCOMING,
			traffic:   &netdb.TrafficCounter{},
		}
		nfm.table[conn] = s
	}

	s.traffic.Packets++
	s.traffic.Bytes += bytes
}

func (nfm *netflowMap) dump() {
	netflows := make([]*netdb.NetFlowEntry, 0)

	nfm.Lock()
	defer nfm.Unlock()

	for c, s := range nfm.table {
		netflows = append(netflows, &netdb.NetFlowEntry{
			Timestamp: time.Now().UnixMilli(),
			Flow: &netdb.Flow{
				Connection: c.GetNetConnection(),
				Direction:  s.direction,
			},
			Traffic: s.traffic,
		})
	}

	nfm.table = make(map[Connection]*netflowState, 0)

	netflowdb.InputQueue <- netflows
}

func (nfm *netflowMap) wrkr() {
	ticker := time.NewTicker(120 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			nfm.dump()
		case <-nfm.closeCh:
			return
		}
	}
}
