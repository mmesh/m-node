package mnet

import (
	"time"

	"github.com/google/uuid"
	"mmesh.dev/m-api-go/grpc/resources/nstore"
	"mmesh.dev/m-api-go/grpc/resources/nstore/hsecdb"
	"mmesh.dev/m-api-go/grpc/resources/nstore/metricsdb"
	"mmesh.dev/m-api-go/grpc/resources/topology"
	"mmesh.dev/m-lib/pkg/xlog"
	"mmesh.dev/m-node/internal/app/node/hsec"
	"mmesh.dev/m-node/internal/app/node/kvstore"
)

func (ln *localNode) Metrics(kvs kvstore.Interface) *topology.AgentMetrics {
	if ln == nil {
		return nil
	}

	if !ln.initialized {
		return nil
	}

	n := ln.Node()
	if n == nil {
		return nil
	}

	am := &topology.AgentMetrics{
		LastUpdated: time.Now().UnixMilli(),
		HostMetrics: ln.Stats().GetHostMetrics(),
		SysMetrics: &topology.SysMetrics{
			OsPkgs:      0,
			Vulns:       &hsecdb.VulnTotals{},
			HostMetrics: make([]*metricsdb.HostMetrics, 0),
		},
	}

	req := &nstore.DataRequest{
		AccountID: n.AccountID,
		TenantID:  n.TenantID,
		NodeID:    n.NodeID,
		QueryID:   uuid.New().String(),
	}

	hmr, err := kvs.HostMetrics().Query(&metricsdb.HostMetricsRequest{
		Request:   req,
		Type:      metricsdb.HostMetricsQueryType_QUERY_LOAD_AVG,
		TimeRange: nstore.TimeRange_TTL_1H,
	})
	if err != nil {
		xlog.Errorf("Unable to get host metrics: %v", err)
	}

	if hmr != nil {
		am.SysMetrics.HostMetrics = hmr.Metrics
	}

	hsecSummary, err := hsec.GetSummary(req)
	if err != nil {
		xlog.Errorf("Unable to get hsec summary: %v", err)
	}

	if hsecSummary != nil {
		am.SysMetrics.OsPkgs = hsecSummary.TotalOSPkgs
		am.SysMetrics.Vulns = hsecSummary.Vulns
	}

	return am
}
