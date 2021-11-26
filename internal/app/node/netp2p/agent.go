package netp2p

import (
	"bufio"
	"os"
	"sync"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/common/datetime"
	"mmesh.dev/m-api-go/grpc/network/mmnp/routing"
	"mmesh.dev/m-api-go/grpc/resources/network"
	"mmesh.dev/m-api-go/grpc/rpc"
	"x6a.dev/pkg/errors"
	"x6a.dev/pkg/xlog"
)

const (
	mmProtocolID = "/mmesh/1.0.0"
)

type endpointsMap struct {
	endpt map[string]*network.NetworkEndpoint
	sync.RWMutex
}

type RoutingTable struct {
	RT *routing.RoutingTable
	sync.RWMutex
}

type streamsMap struct {
	tunnel map[string]*bufio.ReadWriter
	sync.RWMutex
}

type mmAgent struct {
	agentID    string
	accountID  string
	tenantID   string
	netID      string
	vrfID      string
	nodeID     string
	agent      *network.Agent
	endpoints  *endpointsMap
	replicaSet bool
	mainIPv4   string
	p2pHost    host.Host
	nxnc       rpc.NetworkAPIClient
	rt         *RoutingTable
	linkStatus *routing.LinkStatus
	// rtRequestQueue chan struct{}
	streams        *streamsMap
	closeInterface chan struct{}
}

var mma *mmAgent

var nodeIPv6 string

func newEndpointsMap() *endpointsMap {
	return &endpointsMap{
		endpt: make(map[string]*network.NetworkEndpoint),
	}
}

func newRoutingTable() *RoutingTable {
	return &RoutingTable{
		RT: &routing.RoutingTable{},
	}
}

func newStreamsMap() *streamsMap {
	return &streamsMap{
		tunnel: make(map[string]*bufio.ReadWriter),
	}
}

func InitMMAgent(externalIPv4 string, nxnc rpc.NetworkAPIClient) error {
	var p2pHost host.Host
	var p2pHostID string
	var err error

	agentID := viper.GetString("host.id")
	endpointID := viper.GetString("host.id")
	dnsName := viper.GetString("endpoint.dnsName")
	reqIPv4 := viper.GetString("endpoint.ipv4")
	port := int32(viper.GetInt("agent.port"))
	prio := int32(viper.GetInt("agent.priority"))
	isRelay := viper.GetBool("agent.relay")

	endpoints := newEndpointsMap()
	rt := newRoutingTable()
	linkStatus := &routing.LinkStatus{}
	streams := newStreamsMap()
	closeInterface := make(chan struct{})

	if mma == nil {
		if err := ifDown(); err != nil {
			xlog.Alertf("Unable to reset interface: %v", err)
			os.Exit(1)
		}

		if isRelay && len(externalIPv4) > 0 {
			xlog.Info("Initializing relay host...")
			p2pHost, err = newP2PHost(port, p2pHostTypeRelayHost)
		} else if len(externalIPv4) > 0 {
			xlog.Info("Initializing basic host...")
			p2pHost, err = newP2PHost(port, p2pHostTypeBasicHost)
		} else {
			xlog.Info("Initializing hidden host...")
			p2pHost, err = newP2PHost(port, p2pHostTypeHiddenHost)
		}
		if err != nil {
			return errors.Wrapf(err, "[%v] function newP2PHost()", errors.Trace())
		}

		// Set a function as stream handler.
		// This function is called when a peer connects, and starts a stream with this protocol.
		// Only applies on the receiving side.
		p2pHost.SetStreamHandler(mmProtocolID, handleStream)

		p2pHostID = p2pHost.ID().Pretty()

		if len(p2pHost.Addrs()) > 0 {
			xlog.Info("This endpoint's multiaddresses:")
			for _, la := range p2pHost.Addrs() {
				xlog.Infof(" => %v", la)
			}
		}
		xlog.Infof("p2pHostID: %s", p2pHostID)
	} else {
		externalIPv4 = mma.agent.ExternalIPv4
		reqIPv4 = mma.mainIPv4

		p2pHost = mma.p2pHost
		p2pHostID = p2pHost.ID().Pretty()

		endpoints = mma.endpoints
		rt = mma.rt
		linkStatus = mma.linkStatus
		streams = mma.streams
		closeInterface = mma.closeInterface
	}

	al := addrList(p2pHost.Addrs())

	agent := &network.Agent{
		AgentID:      agentID,
		NxHostID:     p2pHostID,
		MAddrs:       al.strings(),
		ExternalIPv4: externalIPv4,
		Port:         port,
		Priority:     prio,
		IsRelay:      isRelay,
		IsMRS:        viper.GetBool("agent.mrs"),
		Routes: &network.Routes{
			Export: viper.GetStringSlice("routes.export"),
			Import: viper.GetStringSlice("routes.import"),
		},
		Options: &network.AgentOptions{
			DNSPort:      int32(viper.GetInt("agent.dns.port")),
			KubernetesGw: viper.GetBool("agent.kubernetes.controller.enabled"),
		},
		Metrics: &network.AgentMetrics{},
		Management: &network.AgentManagement{
			DisableExec:           viper.GetBool("agent.management.disableExec"),
			DisableTransfer:       viper.GetBool("agent.management.disableTransfer"),
			DisablePortForwarding: viper.GetBool("agent.management.disablePortForwarding"),
			DisableOps:            viper.GetBool("agent.management.disableOps"),
		},
		Maintenance: &network.AgentMaintenance{
			AutoUpdate: viper.GetBool("maintenance.autoUpdate"),
			Schedule: &datetime.DateTime{
				Hour:   int32(viper.GetInt("maintenance.schedule.hour")),
				Minute: int32(viper.GetInt("maintenance.schedule.minute")),
			},
		},
		Healthy: true,
	}

	mma = &mmAgent{
		agentID:    agentID,
		accountID:  viper.GetString("node.account"),
		tenantID:   viper.GetString("node.tenant"),
		netID:      viper.GetString("node.network"),
		vrfID:      viper.GetString("node.subnet"),
		nodeID:     viper.GetString("node.id"),
		agent:      agent,
		endpoints:  endpoints,
		replicaSet: viper.GetBool("node.replicaSet"),
		p2pHost:    p2pHost,
		nxnc:       nxnc,
		rt:         rt,
		linkStatus: linkStatus,
		// rtRequestQueue: make(chan struct{}),
		streams:        streams,
		closeInterface: closeInterface,
	}

	ipv4, err := AddNetworkEndpoint(endpointID, dnsName, reqIPv4)
	if err != nil {
		return errors.Wrapf(err, "[%v] function AddNetworkEndpoint()", errors.Trace())
	}

	mma.mainIPv4 = ipv4

	xlog.Debugf("Agent %s initialized", agentID)

	return nil
}

func AgentConfigured() bool {
	if mma != nil {
		return true
	}

	return false
}

func UpdateNetworkClient(nxnc rpc.NetworkAPIClient) {
	mma.nxnc = nxnc
}

func GetLinkStatusConnections() int32 {
	return mma.linkStatus.Connections
}

func Disconnect() {
	xlog.Info("Disconnecting...")
	mma.closeInterface <- struct{}{}
	if iface != nil {
		if err := ifaceClose(); err != nil {
			xlog.Warnf("Unable to close interface %s: %v", iface.name, err)
		}
	}
	if err := mma.p2pHost.Close(); err != nil {
		xlog.Alertf("Unable to close p2pHost handler: %v", err)
		os.Exit(1)
	}
}
