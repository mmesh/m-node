package svcs

import (
	"context"
	"time"

	"mmesh.dev/m-api-go/grpc/network/resources/network"
	"mmesh.dev/m-lib/pkg/runtime"
	"mmesh.dev/m-node/internal/app/node/metrics"
	"mmesh.dev/m-node/internal/app/node/netp2p"
	"x6a.dev/pkg/xlog"
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

				_, err := w.NxNC.Metrics(context.Background(), n)
				if err != nil {
					xlog.Errorf("Unable to send metrics to mmesh controller: %v", err)
					networkErrorEventsQueue <- struct{}{}
					return
				}
				xlog.Debug("Metrics updated")

				metrics.ClearNetworkMetrics()
			case <-quitMetrics:
				xlog.Debug("Closing metrics processor")
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
				time.Sleep(300 * time.Second)
				updateNodeMetrics <- struct{}{}
			}
		}()
	}
}

func getNodeMetrics() *network.Node {
	n := netp2p.GetNode()

	if n.Agent.Metrics == nil {
		n.Agent.Metrics = &network.AgentMetrics{}
	}

	if n.Agent.IsRelay {
		if err := metrics.ReadNetDevStats(); err != nil {
			xlog.Errorf("Unable to read net-dev-stats: %v", err)
		} else {
			n.Agent.Metrics.NetDevStats = metrics.GetNetDevStats()
		}
	}

	n.Agent.Metrics.NetworkMetrics = metrics.GetNetworkMetrics()
	n.Agent.Metrics.NetworkTraffic = metrics.GetNetworkTraffic()
	n.Agent.Metrics.HostMetrics = metrics.GetHostMetrics()
	n.Agent.Metrics.RelayMetrics = metrics.GetRelayMetrics()

	return n
}
