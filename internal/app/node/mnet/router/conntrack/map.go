package conntrack

import (
	"sync"
	"time"

	"mmesh.dev/m-lib/pkg/xlog"
)

type ctStatus uint8

const (
	ctStatusNew ctStatus = iota
	ctStatusActive
)

const ctTimeout = 120

type ctState struct {
	status        ctStatus
	timeout       time.Time
	originCounter *ctCounter
	replyCounter  *ctCounter
}

type ctCounter struct {
	packets uint64
	bytes   uint64
}

type ctMap struct {
	table   map[Connection]*ctState
	closeCh chan struct{}
	sync.RWMutex
}

var conntrack *ctMap

func newMap() {
	conntrack = &ctMap{
		table:   make(map[Connection]*ctState),
		closeCh: make(chan struct{}, 1),
	}

	go conntrack.wrkr()
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
		s = &ctState{
			status:        ctStatusNew,
			originCounter: &ctCounter{},
			replyCounter:  &ctCounter{},
		}
		ctm.table[conn] = s
	}

	s.timeout = time.Now().Add(ctTimeout * time.Second)
	s.originCounter.packets++
	s.originCounter.bytes += bytes
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

	s.status = ctStatusActive
	s.timeout = time.Now().Add(ctTimeout * time.Second)
	s.replyCounter.packets++
	s.replyCounter.bytes += bytes

	return true
}

func (ctm *ctMap) cleanup() {
	ctm.Lock()
	defer ctm.Unlock()

	for c, s := range ctm.table {
		if time.Since(s.timeout) > ctTimeout {
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
