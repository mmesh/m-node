package mnet

import (
	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/common/datetime"
	"mmesh.dev/m-api-go/grpc/resources/network"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/xlog"
	"mmesh.dev/m-node/internal/app/node/mnet/connection"
	"mmesh.dev/m-node/internal/app/node/mnet/maddr"
	"mmesh.dev/m-node/internal/app/node/mnet/router"
)

func Init() error {
	vrfID := viper.GetString("node.subnet")
	agentID := viper.GetString("host.id")
	endpointID := viper.GetString("host.id")
	dnsName := viper.GetString("endpoint.dnsName")
	reqIPv4 := viper.GetString("endpoint.ipv4")
	port := viper.GetInt("agent.port")
	prio := int32(viper.GetInt("agent.priority"))
	isRelay := viper.GetBool("agent.relay")
	k8sGw := viper.GetBool("agent.kubernetes.controller.enabled")
	isMRS := viper.GetBool("agent.mrs")
	isNetworkingDisabled := viper.GetBool("agent.networking.disabled")
	localForwarding := true
	rtExported := viper.GetStringSlice("routes.export")
	rtImported := viper.GetStringSlice("routes.import")

	if k8sGw {
		isNetworkingDisabled = false
	}

	if isMRS {
		k8sGw = false
		isNetworkingDisabled = true
	}

	if isNetworkingDisabled {
		localForwarding = false
	}

	conn := connection.New()

	rtr := router.New(conn.GetExternalIPv4(), vrfID, port, isRelay, localForwarding, rtImported, rtExported)
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
	xlog.Debugf("p2pHostID: %s", rtr.P2PHost().ID().Pretty())

	localnode = &localNode{
		accountID: viper.GetString("node.account"),
		tenantID:  viper.GetString("node.tenant"),
		netID:     viper.GetString("node.network"),
		vrfID:     vrfID,
		nodeID:    viper.GetString("node.id"),
		agent: &network.Agent{
			AgentID:      agentID,
			P2PHostID:    rtr.P2PHost().ID().Pretty(),
			MAddrs:       maddrs,
			ExternalIPv4: conn.GetExternalIPv4(),
			Port:         int32(port),
			Priority:     prio,
			IsRelay:      isRelay,
			IsMRS:        isMRS,
			Routes: &network.Routes{
				Export: rtExported,
				Import: rtImported,
			},
			Options: &network.AgentOptions{
				DNSPort:      int32(viper.GetInt("agent.dns.port")),
				KubernetesGw: k8sGw,
				// Transport:          proto,
				NetworkingDisabled: isNetworkingDisabled,
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
		},
		endpoints: &endpointsMap{
			endpt: make(map[string]*network.NetworkEndpoint),
		},
		replicaSet: viper.GetBool("node.replicaSet"),
		connection: conn,
		router:     rtr,
	}

	if isNetworkingDisabled {
		if err := localnode.registerNode(); err != nil {
			return errors.Wrapf(err, "[%v] function localnode.registerNode()", errors.Trace())
		}
	} else {
		if _, err := localnode.AddNetworkEndpoint(endpointID, dnsName, reqIPv4); err != nil {
			return errors.Wrapf(err, "[%v] function localnode.AddNetworkEndpoint()", errors.Trace())
		}
	}

	xlog.Infof("Node %s initialized", agentID)

	return nil
}
