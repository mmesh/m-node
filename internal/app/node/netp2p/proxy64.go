package netp2p

import (
	"fmt"
	"sync"
	"time"

	"github.com/google/gopacket/layers"
	"mmesh.dev/mmesh/internal/app/node/proxy"
	"mmesh.dev/mmesh/internal/pkg/ipnet"
	"x6a.dev/pkg/errors"
	"x6a.dev/pkg/xlog"
)

type vsID string // string(vip)
type connID string

type proxy64Map struct {
	vs map[vsID]*proxy64VS
	sync.RWMutex
}

type proxy64VS struct {
	vip6 string
	rs   map[connID]*proxy64RS
}

type proxy64RS struct {
	ipv4         string
	proto        layers.IPProtocol
	port         uint16
	lastActivity time.Time
}

var px64Map *proxy64Map

func newProxy64Map() *proxy64Map {
	return &proxy64Map{
		vs: make(map[vsID]*proxy64VS),
	}
}

func newProxy64VS(vip6 string) *proxy64VS {
	return &proxy64VS{
		vip6: vip6,
		rs:   make(map[connID]*proxy64RS),
	}
}

func newProxy64RS(ipv4 string, proto layers.IPProtocol, port uint16) *proxy64RS {
	return &proxy64RS{
		ipv4:  ipv4,
		proto: proto,
		port:  port,
	}
}

func getConnectionID(ipv4 string, proto layers.IPProtocol, port uint16) connID {
	return connID(fmt.Sprintf("mmesh64:%s:%s:%d", ipv4, proto.String(), port))
}

func (px64m *proxy64Map) setVS(vip6 string) {
	px64m.Lock()
	defer px64m.Unlock()

	if _, ok := px64m.vs[vsID(vip6)]; !ok {
		px64m.vs[vsID(vip6)] = newProxy64VS(vip6)

		if err := ip6AddrAdd(vip6); err != nil {
			xlog.Alertf("Proxy64: Unable to add address %s to interface: %s", vip6, errors.Cause(err))
		}
	}
}

func (px64m *proxy64Map) setConnection(vip6, ipv4 string, proto layers.IPProtocol, port uint16) connID {
	px64m.setVS(vip6)

	px64m.Lock()
	defer px64m.Unlock()

	vs := px64m.vs[vsID(vip6)]

	connID := getConnectionID(ipv4, proto, port)

	if _, ok := vs.rs[connID]; !ok {
		vs.rs[connID] = newProxy64RS(ipv4, proto, port)
	}

	vs.rs[connID].lastActivity = time.Now()

	return connID
}

func setProxy64(vip6, ipv4 string, proto layers.IPProtocol, port uint16) {
	if px64Map == nil {
		px64Map = newProxy64Map()
	}

	connID := px64Map.setConnection(vip6, ipv4, proto, port)

	ns := proxy.NamespaceNone
	svcName := string(connID)
	portName := ipnet.GetPortName(proto, port)

	if proxy.RunningPort(ns, svcName, portName) {
		return
	}

	proxy.SetPort(proxy.ServiceTypeProxy64, ns, svcName, ipv4, vip6, portName, proto, int32(port), ipnet.AddressFamilyIPv6)
	proxy.FwdSvc(ns, svcName)

	time.Sleep(500 * time.Millisecond)
}

func Proxy64GC() {
	if px64Map == nil {
		px64Map = newProxy64Map()
		return
	}

	px64Map.Lock()
	defer px64Map.Unlock()

	for vsID, vs := range px64Map.vs {
		for connID, rs := range vs.rs {
			if time.Since(rs.lastActivity) > 300*time.Second {
				xlog.Infof("Removing inactive connection %s", connID)

				ns := proxy.NamespaceNone
				svcName := string(connID)
				portName := ipnet.GetPortName(rs.proto, rs.port)

				proxy.DeletePort(ns, svcName, portName)

				delete(vs.rs, connID)
			}
		}

		if len(vs.rs) == 0 {
			if err := ip6AddrDel(vs.vip6); err != nil {
				xlog.Alertf("Proxy64: Unable to delete address %s from interface: %s", vs.vip6, errors.Cause(err))
			}

			delete(px64Map.vs, vsID)
		}
	}
}
