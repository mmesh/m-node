package mnet

import (
	"context"
	"time"

	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/network/nac"
	"mmesh.dev/m-api-go/grpc/resources/topology"
	"mmesh.dev/m-lib/pkg/errors"
)

func (ln *localNode) NodeReq() *topology.NodeReq {
	return &topology.NodeReq{
		AccountID: ln.node.AccountID,
		TenantID:  ln.node.TenantID,
		NetID:     ln.node.NetID,
		SubnetID:  ln.node.SubnetID,
		NodeID:    ln.node.NodeID,
	}
}

func (ln *localNode) Node() *topology.Node {
	ln.node.Endpoints = ln.endpoints.endpt

	return ln.node
}

func (ln *localNode) NewCfg(ncfg *topology.NodeCfg) error {
	ln.node.Cfg = ncfg

	setCfgVars(ncfg)

	return nil
}

func (ln *localNode) registerNode() error {
	n := ln.Node()

	nrReq := &nac.NodeRegRequest{
		Node: n,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	r, err := ln.Connection().NetworkClient().RegisterNode(ctx, nrReq)
	if err != nil {
		return errors.Wrapf(err, "[%v] function ln.Connection().NetworkClient().RegisterNode()", errors.Trace())
	}

	if ln.initialized {
		return nil
	}

	// initialize rib data structure
	mq := ln.Connection().RIBDataMsgRxQueue()
	ln.Router().RIB().Initialize(mq, n, r.RoutingDomain, r.NetworkPolicy)

	// subscribe to mqtt routing topics
	if err := ln.Connection().NewRoutingSession(r.RoutingDomain.LocationID); err != nil {
		return errors.Wrapf(err, "[%v] function ln.Connection().NewRoutingSession()", errors.Trace())
	}

	ln.initialized = true

	return nil
}

func setCfgVars(cfg *topology.NodeCfg) {
	viper.Set("nodeName", cfg.NodeName)

	viper.Set("management.disableExec", cfg.Management.DisableExec)
	viper.Set("management.disableTransfer", cfg.Management.DisableTransfer)
	viper.Set("management.disablePortForwarding", cfg.Management.DisablePortForwarding)
	viper.Set("management.disableOps", cfg.Management.DisableOps)

	viper.Set("maintenance.autoUpdate", cfg.Maintenance.AutoUpdate)
	viper.Set("maintenance.schedule.hour", int(cfg.Maintenance.Schedule.Hour))
	viper.Set("maintenance.schedule.minute", int(cfg.Maintenance.Schedule.Minute))
}
