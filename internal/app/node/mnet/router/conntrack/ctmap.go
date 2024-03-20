package conntrack

import (
	"sync"
	"time"

	"mmesh.dev/m-api-go/grpc/resources/nstore/netdb"
	"mmesh.dev/m-lib/pkg/xlog"
)

const ctTimeout = 3600 // 1 hour (idle connection timeout)

type ctMap struct {
	table   map[Connection]*netdb.ConnectionState
	closeCh chan struct{}
	sync.RWMutex
}

var conntrack *ctMap

func newMap() {
	conntrack = &ctMap{
		table:   make(map[Connection]*netdb.ConnectionState, 0),
		closeCh: make(chan struct{}, 1),
	}

	go conntrack.wrkr()
}

func (ctm *ctMap) getTable() map[Connection]*netdb.ConnectionState {
	ctm.RLock()
	defer ctm.RUnlock()

	return ctm.table
}

func (ctm *ctMap) outboundConnection(c *Connection, bytes uint64) {
	if c == nil {
		return
	}

	ctm.Lock()
	defer ctm.Unlock()

	conn := c.outbound()

	s, ok := ctm.table[conn]
	if !ok {
		s = &netdb.ConnectionState{
			Status:        netdb.ConnectionStatus_NEW,
			FirstSeen:     time.Now().UnixMilli(),
			OriginCounter: &netdb.TrafficCounter{},
			ReplyCounter:  &netdb.TrafficCounter{},
		}
		ctm.table[conn] = s
	}

	s.Timeout = time.Now().Add(ctTimeout * time.Second).UnixMilli()
	s.OriginCounter.Packets++
	s.OriginCounter.Bytes += bytes
}

func (ctm *ctMap) isActiveConnection(c *Connection, bytes uint64) bool {
	if c == nil {
		return false
	}

	ctm.Lock()
	defer ctm.Unlock()

	conn := c.reverse()

	s, ok := ctm.table[conn]
	if !ok {
		return false
	}

	s.Status = netdb.ConnectionStatus_ACTIVE
	s.Timeout = time.Now().Add(ctTimeout * time.Second).UnixMilli()
	s.ReplyCounter.Packets++
	s.ReplyCounter.Bytes += bytes

	return true
}

func (ctm *ctMap) cleanup() {
	ctm.Lock()
	defer ctm.Unlock()

	for c, s := range ctm.table {
		tm := time.UnixMilli(s.Timeout)
		if time.Since(tm) > ctTimeout {
			delete(ctm.table, c)
			xlog.Infof("[conntrack] Removed expired %s connection from %s:%d to %s:%d",
				c.Proto.String(),
				c.SrcIP.String(),
				c.SrcPort,
				c.DstAddr.String(),
				c.DstPort,
			)
		}
	}
}

func (ctm *ctMap) wrkr() {
	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			ctm.cleanup()
		case <-ctm.closeCh:
			return
		}
	}
}
