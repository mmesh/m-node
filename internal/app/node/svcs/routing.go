package svcs

import (
	"time"

	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-lib/pkg/mmp/queuing"
	"mmesh.dev/m-lib/pkg/runtime"
	"mmesh.dev/m-lib/pkg/xlog"
	"mmesh.dev/m-node/internal/app/node/mmsp/protos/routing"
	"mmesh.dev/m-node/internal/app/node/mnet"
)

// RoutingAgent runs routing engine
func RoutingAgent(w *runtime.Wrkr) {
	xlog.Infof("Started worker %s", w.Name)
	w.Running = true

	endCh := make(chan struct{}, 2)

	go func() {
		mmID := viper.GetString("mm.id")

		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				if mnet.LocalNode().Node().Cfg.DisableNetworking ||
					mnet.LocalNode().Router() == nil {
					continue
				}

				if !routing.ServiceEnabled {
					time.Sleep(20 * time.Minute)
				}

				xlog.Debug("Sending node-LSA")

				lsa := mnet.LocalNode().GetNodeLSA()
				if lsa == nil {
					continue
				}

				queuing.TxControlQueue <- &mmsp.Payload{
					SrcID: mmID,
					Type:  mmsp.PDUType_ROUTING,
					RoutingPDU: &mmsp.RoutingPDU{
						Type: mmsp.RoutingMsgType_ROUTING_LSA,
						LSA:  lsa,
					},
				}
			case <-endCh:
				// xlog.Warn("Closing rtRequest send stream")
				return
			}
		}
	}()

	<-w.QuitChan

	endCh <- struct{}{}

	w.WG.Done()
	w.Running = false
	xlog.Infof("Stopped worker %s", w.Name)
}
