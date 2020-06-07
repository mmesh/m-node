package svcs

import (
	"time"

	"mmesh.dev/mmesh/internal/app/node/connection"
	"mmesh.dev/mmesh/internal/app/node/netp2p"
	"mmesh.dev/mmesh/internal/pkg/grpc/client"
	"mmesh.dev/mmesh/internal/pkg/runtime"
	"x6a.dev/pkg/xlog"
)

var networkErrorEventsQueue = make(chan struct{}, 128)
var networkDisconnectQueue = make(chan struct{}, 128)

func NetworkErrorEventsHandler(w *runtime.Wrkr) {
	xlog.Infof("Started worker %s", w.Name)
	w.Running = true

	networkErrorHandlerRunning := false
	for {
		select {
		case <-networkErrorEventsQueue:

			if !networkErrorHandlerRunning {
				networkErrorHandlerRunning = true
				go func() {
					//networkErrorHandlerRunning = true
					if err := client.NxClientConn.Close(); err != nil {
						xlog.Alertf("Unable to close gRPC network connection: %v", err)
					}
					nxnc := connection.AgentConnect()
					runtime.NetworkWrkrReconnect(nxnc)
					time.Sleep(3 * time.Second)
					networkErrorHandlerRunning = false
				}()
			}
		case <-networkDisconnectQueue:
			xlog.Debug("Closing network connection handlers...")
			netp2p.Disconnect()
		case <-w.QuitChan:
			w.WG.Done()
			w.Running = false
			xlog.Infof("Stopped worker %s", w.Name)
			return
		}
	}
}
