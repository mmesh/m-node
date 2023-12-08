package conntrack

type Interface interface {
	OutboundConnection(c *Connection, pktlen int)
	IsValidConnection(c *Connection, pktlen int) bool
	Close()
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
}

func (api *api) IsValidConnection(c *Connection, pktlen int) bool {
	if conntrack == nil {
		return false
	}

	return conntrack.isValidConnection(c, uint64(pktlen))
}

func (api *api) Close() {
	if conntrack == nil {
		return
	}

	conntrack.closeCh <- struct{}{}
}
