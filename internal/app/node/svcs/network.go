package svcs

import (
	"time"

	"mmesh.dev/m-lib/pkg/runtime"
	"mmesh.dev/m-node/internal/app/node/connection"
	"mmesh.dev/m-node/internal/app/node/netp2p"
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
					time.Sleep(3 * time.Second)
					if err := connection.GRPCClientConn.Close(); err != nil {
						xlog.Alertf("Unable to close gRPC network connection: %v", err)
					}
					nxnc := connection.AgentConnect()
					runtime.NetworkWrkrReconnect(nxnc)
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
