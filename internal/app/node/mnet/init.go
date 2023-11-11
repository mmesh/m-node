package mnet

import (
	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/resources/topology"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/mm"

	// "mmesh.dev/m-lib/pkg/mmid"
	"mmesh.dev/m-lib/pkg/xlog"
	"mmesh.dev/m-node/internal/app/node/mnet/connection"
	"mmesh.dev/m-node/internal/app/node/mnet/maddr"
	"mmesh.dev/m-node/internal/app/node/mnet/router"
)

func Init() error {
	hostID := viper.GetString("host.id")
	port := viper.GetInt("port")
	dnsPort := viper.GetInt("dnsPort")
	reqIPv4 := viper.GetString("ipv4")
	rtExported := viper.GetStringSlice("routes.export")
	rtImported := viper.GetStringSlice("routes.import")

	conn := connection.New()

	n := conn.Node()

	mmID, err := mm.GetID(&topology.NodeReq{
		AccountID: n.AccountID,
		TenantID:  n.TenantID,
		NetID:     n.NetID,
		SubnetID:  n.SubnetID,
		NodeID:    n.NodeID,
	})
	if err != nil {
		return errors.Wrapf(err, "[%v] function mm.GetID()", errors.Trace())
	}

	viper.Set("mm.id", mmID.String())

	if n.Agent.DevMode {
		viper.Set("version.branch", "dev")
	}

	setCfgVars(n.Cfg)

	nodeName := n.Cfg.NodeName
	viper.Set("nodeName", nodeName)

	localForwarding := true

	if n.Cfg.KubernetesGw {
		n.Cfg.DisableNetworking = false
	}

	if n.Cfg.DisableNetworking {
		localForwarding = false
	}

	n.Agent.ExternalIPv4 = conn.GetExternalIPv4()
	n.Agent.CanRelay = len(n.Agent.ExternalIPv4) > 0

	rtr := router.New(n.Agent.ExternalIPv4, n.SubnetID, port, localForwarding, rtImported, rtExported)
	if err := rtr.Init(); err != nil {
		return errors.Wrapf(err, "[%v] function r.Init()", errors.Trace())
	}

	maddrs := maddr.GetGlobalUnicastAddrStrings(rtr.P2PHost().Addrs()...)

	if len(maddrs) > 0 {
		xlog.Info("Node multi-addresses:")
		for _, ma := range maddrs {
			xlog.Infof(" => %s", ma)
		}
	}
	xlog.Debugf("p2pHostID: %s", rtr.P2PHost().ID().String())

	n.Agent.P2PHostID = rtr.P2PHost().ID().String()
	n.Agent.Hostname = hostID
	n.Agent.Port = int32(port)
	n.Agent.DNSPort = int32(dnsPort)
	n.Agent.MAddrs = maddrs
	n.Agent.Routes = &topology.Routes{
		Export: rtExported,
		Import: rtImported,
	}
	n.Agent.Metrics = &topology.AgentMetrics{}
	n.Endpoints = make(map[string]*topology.Endpoint)

	localnode = &localNode{
		node: n,
		endpoints: &endpointsMap{
			endpt: make(map[string]*topology.Endpoint),
		},
		connection: conn,
		router:     rtr,
	}

	if n.Cfg.DisableNetworking {
		if err := localnode.registerNode(); err != nil {
			return errors.Wrapf(err, "[%v] function localnode.registerNode()", errors.Trace())
		}
	} else {
		endpointID := nodeName
		dnsName := nodeName
		if _, err := localnode.AddNetworkEndpoint(endpointID, dnsName, reqIPv4); err != nil {
			return errors.Wrapf(err, "[%v] function localnode.AddNetworkEndpoint()", errors.Trace())
		}
	}

	xlog.Infof("Node %s initialized", nodeName)

	return nil
}
