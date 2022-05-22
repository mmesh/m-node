package connection

import (
	"time"

	"mmesh.dev/m-lib/pkg/runtime"
	"mmesh.dev/m-lib/pkg/xlog"
)

func (s *session) connWatcher() {
	networkErrorHandlerRunning := false

	for {
		select {
		case <-s.watcherCh:
			if !networkErrorHandlerRunning {
				networkErrorHandlerRunning = true
				go func() {
					time.Sleep(3 * time.Second)
					xlog.Warn("Connection lost, reconnecting...")

					if err := s.connection.grpcClientConn.Close(); err != nil {
						xlog.Errorf("Unable to close gRPC network connection: %v", err)
					}

					s.connection = newConnection()
					runtime.NetworkWrkrReconnect(s.NetworkClient())

					networkErrorHandlerRunning = false
				}()
			}
		case <-s.endCh:
			return
		}
	}
}
