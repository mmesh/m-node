package router

import (
	"bufio"
	"os"
	"sync"

	libp2pHost "github.com/libp2p/go-libp2p-core/host"
	"mmesh.dev/m-api-go/grpc/network/mmnp/routing"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/xlog"
	"mmesh.dev/m-node/internal/app/node/mnet/p2p"
	"mmesh.dev/m-node/internal/app/node/mnet/p2p/host"
)

type Interface interface {
	Init() error
	P2PHost() libp2pHost.Host
	IPv4() string
	SetIPv4(ipv4 string)
	IPv6() string
	SetIPv6(ipv6 string)
	IP4AddrAdd(ipv4 string) error
	IP4AddrDel(ipv4 string) error
	IP6AddrAdd(ipv6 string) error
	IP6AddrDel(ipv6 string) error
	GetLinkStatusConnections() int32
	UpdateRoutingTable(newRT *routing.RoutingTable)
	RtrQueue() chan struct{}
	Disconnect()
}

type router struct {
	p2pHost libp2pHost.Host

	vrfID string

	port         int
	isRelay      bool
	externalIPv4 string

	ipv4 string
	ipv6 string

	rib     *routingTable
	routes  *routeMap
	streams *streamMap
	proxy64 *proxy64Map

	networkInterface *networkInterface
	localForwarding  bool

	rtrQueue chan struct{}

	linkStatus *routing.LinkStatus
}

type routingTable struct {
	global *routing.RoutingTable
	sync.RWMutex
}

type cidrIPDst string

type routeMap struct {
	local    map[cidrIPDst]bool
	imported []string
	exported []string

	sync.RWMutex
}

type streamMap struct {
	tunnel map[string]*bufio.ReadWriter
	sync.RWMutex
}

func (d cidrIPDst) string() string {
	return string(d)
}

func New(externalIPv4, vrfID string, port int, isRelay, localForwarding bool, rtImported, rtExported []string) Interface {
	return &router{
		vrfID:        vrfID,
		port:         port,
		isRelay:      isRelay,
		externalIPv4: externalIPv4,
		rib: &routingTable{
			global: &routing.RoutingTable{},
		},
		routes: &routeMap{
			local:    make(map[cidrIPDst]bool),
			imported: rtImported,
			exported: rtExported,
		},
		streams: &streamMap{
			tunnel: make(map[string]*bufio.ReadWriter),
		},
		proxy64: &proxy64Map{
			vs:      make(map[proxy64VSID]*proxy64VS),
			closeCh: make(chan struct{}, 1),
		},
		localForwarding: localForwarding,
		rtrQueue:        make(chan struct{}, 8),
		linkStatus:      &routing.LinkStatus{},
	}
}

func (r *router) Init() error {
	var p2pHost libp2pHost.Host
	var err error

	if r.localForwarding {
		if err := r.ifDown(); err != nil {
			xlog.Alertf("Unable to reset interface: %v", err)
			os.Exit(1)
		}

		if err := r.ifUp(); err != nil {
			return errors.Wrapf(err, "[%v] function r.ifUp()", errors.Trace())
		}

		go r.proxy64GC()
	}

	if r.isRelay && len(r.externalIPv4) > 0 {
		xlog.Info("Initializing relay host...")
		p2pHost, err = host.New(host.P2PHostTypeRelayHost, r.port)
	} else if len(r.externalIPv4) > 0 {
		xlog.Info("Initializing basic host...")
		p2pHost, err = host.New(host.P2PHostTypeBasicHost, r.port)
	} else {
		xlog.Info("Initializing hidden host...")
		p2pHost, err = host.New(host.P2PHostTypeHiddenHost, r.port)
	}
	if err != nil {
		return errors.Wrapf(err, "[%v] function libp2p.NewHost()", errors.Trace())
	}

	// if len(p2pHost.Addrs()) > 0 {
	// 	xlog.Info("This endpoint's multi-addresses:")
	// 	for _, ma := range p2pHost.Addrs() {
	// 		xlog.Infof(" => %s", ma.String())
	// 	}
	// }
	// xlog.Infof("p2pHostID: %s", p2pHost.ID().Pretty())

	r.p2pHost = p2pHost

	// Set a function as stream handler.
	// This function is called when a peer connects, and starts a stream with this protocol.
	// Only applies on the receiving side.
	r.p2pHost.SetStreamHandler(p2p.ProtocolID, r.handleStream)

	return nil
}

func (r *router) P2PHost() libp2pHost.Host {
	return r.p2pHost
}

func (r *router) IPv4() string {
	return r.ipv4
}

func (r *router) SetIPv4(ipv4 string) {
	r.ipv4 = ipv4
}

func (r *router) IPv6() string {
	return r.ipv6
}

func (r *router) SetIPv6(ipv6 string) {
	r.ipv6 = ipv6
}

func (r *router) IP4AddrAdd(ipv4 string) error {
	if r.networkInterface == nil || !r.localForwarding {
		return nil
	}

	return r.networkInterface.ip4AddrAdd(ipv4)
}

func (r *router) IP4AddrDel(ipv4 string) error {
	if r.networkInterface == nil || !r.localForwarding {
		return nil
	}

	return r.networkInterface.ip4AddrDel(ipv4)
}

func (r *router) IP6AddrAdd(ipv6 string) error {
	if r.networkInterface == nil || !r.localForwarding {
		return nil
	}

	return r.networkInterface.ip6AddrAdd(ipv6)
}

func (r *router) IP6AddrDel(ipv6 string) error {
	if r.networkInterface == nil || !r.localForwarding {
		return nil
	}

	return r.networkInterface.ip6AddrDel(ipv6)
}

func (r *router) RtrQueue() chan struct{} {
	return r.rtrQueue
}

func (r *router) GetLinkStatusConnections() int32 {
	return r.linkStatus.Connections
}

func (r *router) Disconnect() {
	xlog.Info("Disconnecting...")

	if r.localForwarding && r.networkInterface != nil {
		r.networkInterface.closeCh <- struct{}{}

		if err := r.networkInterface.close(); err != nil {
			xlog.Warnf("Unable to close interface %s: %v", r.networkInterface.devName(), err)
		}

		r.proxy64.closeCh <- struct{}{}
	}

	if err := r.p2pHost.Close(); err != nil {
		xlog.Warnf("Unable to close p2pHost handler: %v", err)
		os.Exit(1)
	}
}
