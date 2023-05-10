package svcs

import (
	"context"
	"os"
	"time"

	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-api-go/grpc/resources/topology"
	"mmesh.dev/m-lib/pkg/mmp/queuing"
	"mmesh.dev/m-lib/pkg/runtime"
	"mmesh.dev/m-lib/pkg/xlog"
	"mmesh.dev/m-node/internal/app/node/metrics"
	"mmesh.dev/m-node/internal/app/node/mmsp/protos/routing"
	"mmesh.dev/m-node/internal/app/node/mnet"
	"mmesh.dev/m-node/internal/app/node/tss"
)

var updateNodeMetrics = make(chan struct{}, 1)

func MetricsAgent(w *runtime.Wrkr) {
	xlog.Infof("Started worker %s", w.Name)
	w.Running = true

	quitMetrics := make(chan struct{}, 1)

	tsdb, err := tss.Open()
	if err != nil {
		xlog.Alertf("[tss] Unable to open time-series store: %v", err)
		os.Exit(1)
	}

	go func() {
		mmID := viper.GetString("mm.id")

		tickerTSS := time.NewTicker(2 * time.Minute)
		defer tickerTSS.Stop()

		for {
			select {
			case req := <-tss.RequestQueue:
				nm, err := tsdb.Query(req)
				if err != nil {
					xlog.Errorf("[tss] Unable to get metrics: %v", err)
					continue
				}

				queuing.TxControlQueue <- &mmsp.Payload{
					SrcID:           mmID,
					DstControllerID: req.ControllerID,
					Type:            mmsp.PDUType_NODEMGMT,
					NodeMgmtPDU: &mmsp.NodeMgmtPDU{
						Type:                mmsp.NodeMgmtMsgType_NODE_METRICS_RESPONSE,
						NodeMetricsResponse: nm,
					},
				}

			case <-tickerTSS.C:
				xlog.Debug("[tss] Writing metrics..")

				nr := mnet.LocalNode().NodeReq()
				metrics.UpdateHostMetrics(nr)
				dpList := metrics.GetTSDataPoints()
				// fmt.Println(dpList)
				metrics.ClearNetworkMetrics()

				if err := tsdb.WriteBatch(dpList); err != nil {
					xlog.Errorf("[tss] Unable to write metrics to tss: %v", err)
				}

			case <-updateNodeMetrics:
				if routing.ServiceEnabled {
					n := getNodeMetrics()

					if _, err := w.NxNC.Metrics(context.TODO(), n); err != nil {
						xlog.Errorf("Unable to send metrics to controller: %v", err)
						mnet.LocalNode().Connection().Watcher() <- struct{}{}
						return
					}
				}

				// metrics.ClearNetworkMetrics()
			case <-quitMetrics:
				xlog.Debug("[tss] Closing metrics processor")
				return
			}
		}
	}()

	go metricsAgentCtl()

	<-w.QuitChan
	quitMetrics <- struct{}{}

	if err := tsdb.Close(); err != nil {
		xlog.Errorf("[tss] Unable to close tsDB: %v", err)
	} else {
		xlog.Info("[tss] Database closed")
	}

	w.WG.Done()
	w.Running = false
	xlog.Infof("Stopped worker %s", w.Name)
}

var metricsAgentCtlRun bool

func metricsAgentCtl() {
	if !metricsAgentCtlRun {
		metricsAgentCtlRun = true
		// go func() {
		// 	for {
		// 		nr := mnet.LocalNode().NodeReq()
		// 		metrics.UpdateHostMetrics(nr)
		// 		time.Sleep(60 * time.Second)
		// 	}
		// }()
		go func() {
			for {
				time.Sleep(900 * time.Second)
				updateNodeMetrics <- struct{}{}
			}
		}()
	}
}

func getNodeMetrics() *topology.Node {
	n := mnet.LocalNode().Node()

	if n.Agent.Metrics == nil {
		n.Agent.Metrics = &topology.AgentMetrics{}
	}

	// n.Agent.Metrics.NetworkMetrics = metrics.GetNetworkMetrics()
	// n.Agent.Metrics.NetworkTraffic = metrics.GetNetworkTraffic()
	n.Agent.Metrics.HostMetrics = metrics.GetHostMetrics()

	return n
}
