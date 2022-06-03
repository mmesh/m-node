package svcs

import (
	"context"
	"time"

	metrics_pb "mmesh.dev/m-api-go/grpc/resources/metrics"
	"mmesh.dev/m-api-go/grpc/resources/network"
	"mmesh.dev/m-lib/pkg/runtime"
	"mmesh.dev/m-lib/pkg/xlog"
	"mmesh.dev/m-node/internal/app/node/metrics"
	"mmesh.dev/m-node/internal/app/node/mnet"
)

var updateNodeMetrics = make(chan struct{}, 1)

func MetricsAgent(w *runtime.Wrkr) {
	xlog.Infof("Started worker %s", w.Name)
	w.Running = true

	quitMetrics := make(chan struct{}, 1)

	go func() {
		for {
			select {
			case <-updateNodeMetrics:
				n := getNodeMetrics()

				if _, err := w.NxNC.Metrics(context.TODO(), n); err != nil {
					xlog.Errorf("Unable to send metrics to controller: %v", err)
					mnet.LocalNode().Connection().Watcher() <- struct{}{}
					return
				}

				if _, err := w.NxNC.DataPointMetrics(context.TODO(), getDataPointMetrics(n)); err != nil {
					xlog.Errorf("Unable to send data point metrics to controller: %v", err)
					mnet.LocalNode().Connection().Watcher() <- struct{}{}
					return
				}

				xlog.Debug("Metrics updated")

				metrics.ClearNetworkMetrics()
			case <-quitMetrics:
				// xlog.Debug("Closing metrics processor")
				return
			}
		}
	}()

	go metricsAgentCtl()

	<-w.QuitChan
	quitMetrics <- struct{}{}

	w.WG.Done()
	w.Running = false
	xlog.Infof("Stopped worker %s", w.Name)
}

var metricsAgentCtlRun bool

func metricsAgentCtl() {
	if !metricsAgentCtlRun {
		metricsAgentCtlRun = true
		go func() {
			for {
				metrics.UpdateHostMetrics()
				time.Sleep(60 * time.Second)
			}
		}()
		go func() {
			for {
				time.Sleep(900 * time.Second)
				updateNodeMetrics <- struct{}{}
			}
		}()
	}
}

func getNodeMetrics() *network.Node {
	// n := netp2p.GetNode()
	n := mnet.LocalNode().NetworkNode()

	if n.Agent.Metrics == nil {
		n.Agent.Metrics = &network.AgentMetrics{}
	}

	// if n.Agent.IsRelay {
	// 	if err := metrics.ReadNetDevStats(); err != nil {
	// 		xlog.Errorf("Unable to read net-dev-stats: %v", err)
	// 	} else {
	// 		n.Agent.Metrics.NetDevStats = metrics.GetNetDevStats()
	// 	}
	// }

	// n.Agent.Metrics.NetworkMetrics = metrics.GetNetworkMetrics()
	// n.Agent.Metrics.NetworkTraffic = metrics.GetNetworkTraffic()
	n.Agent.Metrics.HostMetrics = metrics.GetHostMetrics()
	// n.Agent.Metrics.RelayMetrics = metrics.GetRelayMetrics()

	return n
}

func getDataPointMetrics(n *network.Node) *metrics_pb.DataPoints {
	mdp := &metrics_pb.DataPoints{
		AccountID:  n.AccountID,
		TenantID:   n.TenantID,
		NetID:      n.NetID,
		VRFID:      n.VRFID,
		NodeID:     n.NodeID,
		DataPoints: make([]*metrics_pb.DataPoint, 0),
	}

	dp := metrics.GetHostDataPoint(n)
	if dp != nil {
		mdp.DataPoints = append(mdp.DataPoints, dp)
	}

	dp = metrics.GetNetDataPoint(n)
	if dp != nil {
		mdp.DataPoints = append(mdp.DataPoints, dp)
	}

	return mdp
}
