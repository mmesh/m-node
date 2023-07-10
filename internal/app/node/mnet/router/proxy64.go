package router

import (
	"fmt"
	"sync"
	"time"

	"github.com/google/gopacket/layers"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/ipnet"
	"mmesh.dev/m-lib/pkg/xlog"
	"mmesh.dev/m-node/internal/app/node/metrics"
	"mmesh.dev/m-node/internal/app/node/mnet/proxy"
)

type proxy64VSID string // string(vip)
type proxy64ConnID string

type proxy64Map struct {
	vs      map[proxy64VSID]*proxy64VS
	closeCh chan struct{}
	sync.RWMutex
}

type proxy64VS struct {
	vip6 string
	rs   map[proxy64ConnID]*proxy64RS
}

type proxy64RS struct {
	ipv4         string
	proto        layers.IPProtocol
	port         uint16
	lastActivity time.Time
}

// var px64Map *proxy64Map
/*
func newProxy64Map() *proxy64Map {
	return &proxy64Map{
		vs: make(map[proxy64VSID]*proxy64VS),
	}
}
*/

func newProxy64VS(vip6 string) *proxy64VS {
	return &proxy64VS{
		vip6: vip6,
		rs:   make(map[proxy64ConnID]*proxy64RS),
	}
}

func newProxy64RS(ipv4 string, proto layers.IPProtocol, port uint16) *proxy64RS {
	return &proxy64RS{
		ipv4:  ipv4,
		proto: proto,
		port:  port,
	}
}

func getProxy64ConnectionID(ipv4 string, proto layers.IPProtocol, port uint16) proxy64ConnID {
	return proxy64ConnID(fmt.Sprintf("mmesh64:%s:%s:%d", ipv4, proto.String(), port))
}

func (r *router) setProxy64VS(vip6 string) {
	r.proxy64.Lock()
	defer r.proxy64.Unlock()

	if _, ok := r.proxy64.vs[proxy64VSID(vip6)]; !ok {
		r.proxy64.vs[proxy64VSID(vip6)] = newProxy64VS(vip6)

		if err := r.networkInterface.ip6AddrAdd(vip6); err != nil {
			xlog.Alertf("Proxy64: Unable to add address %s to interface: %s", vip6, errors.Cause(err))
		}
	}
}

func (r *router) setProxy64Connection(vip6, ipv4 string, proto layers.IPProtocol, port uint16) proxy64ConnID {
	r.setProxy64VS(vip6)

	r.proxy64.Lock()
	defer r.proxy64.Unlock()

	vs := r.proxy64.vs[proxy64VSID(vip6)]

	connID := getProxy64ConnectionID(ipv4, proto, port)

	if _, ok := vs.rs[connID]; !ok {
		vs.rs[connID] = newProxy64RS(ipv4, proto, port)
	}

	vs.rs[connID].lastActivity = time.Now()

	return connID
}

func (r *router) setProxy64(vip6, ipv4 string, proto layers.IPProtocol, port uint16) {
	connID := r.setProxy64Connection(vip6, ipv4, proto, port)

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

func (r *router) proxy64Forward(ipHdr *ipHeader, pkt []byte) bool {
	// if ipHdr.af == ipnet.AddressFamilyIPv4 {
	// 	return false
	// }

	// ipv6 traffic
	dstIPv4, err := ipnet.GetIPv4Encap(ipHdr.dstIP.String())
	if err != nil {
		// ipv6 addr does not encapsulate an ipv4 addr
		return false
	}

	// if err := ipHdr.parseProtocol(pkt); err != nil {
	// 	// unable to decode packet, so proxy64 not possible
	// 	xlog.Warnf("Unable to parse IP protocol: %v", errors.Cause(err))
	// 	return false
	// }

	if ipHdr.proto != layers.IPProtocolTCP {
		// for the moment, only tcp is supported
		return false
	}

	if ipHdr.dstAddr != r.ipv6 {
		// packet is not for this node
		return false
	}

	// forward to local proxy64
	r.setProxy64(ipHdr.dstIP.String(), dstIPv4, ipHdr.proto, ipHdr.dstPort)

	// write to TUN interface
	if r.networkInterface != nil {
		if _, err = r.networkInterface.write(pkt); err != nil {
			xlog.Warnf("Unable to write packet to network interface: %v", err)
			return true
		}

		// update metrics
		go metrics.UpdateNetworkMetric(ipHdr.srcIP.String(), 0, uint64(len(pkt)), false)
	}

	return true
}

func (r *router) proxy64GC() {
	ticker := time.NewTicker(2 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			xlog.Debug("Running mmesh64 garbage collector...")

			r.proxy64.Lock()
			defer r.proxy64.Unlock()

			for vsID, vs := range r.proxy64.vs {
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
					if err := r.networkInterface.ip6AddrDel(vs.vip6); err != nil {
						xlog.Alertf("Proxy64: Unable to delete address %s from interface: %s", vs.vip6, errors.Cause(err))
					}

					delete(r.proxy64.vs, vsID)
				}
			}
		case <-r.proxy64.closeCh:
			xlog.Debug("Closing mmesh64 garbage collector")
			return
		}
	}
}
