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
		case <-mqttConnectionWatcherCh:
			s.watcherCh <- struct{}{}
		case <-s.watcherCh:
			if !networkErrorHandlerRunning {
				networkErrorHandlerRunning = true
				go func() {
					time.Sleep(3 * time.Second)
					xlog.Warn("Connection lost, reconnecting...")

					// close grpc connection
					if err := s.connection.grpcClientConn.Close(); err != nil {
						xlog.Errorf("Unable to close gRPC network connection: %v", err)
					}

					// disconnect mqtt connection
					if s.connection.mqttClient != nil {
						s.connection.mqttClient.Disconnect(250)
						s.connection.mqttClient = nil
					}

					s.connection.new()
					runtime.NetworkWrkrReconnect(s.NetworkClient())

					// reconnect mqtt subscriptions
					if len(s.locationID) > 0 && !s.connection.node.Cfg.DisableNetworking {
						if err := s.NewRoutingSession(s.locationID); err != nil {
							xlog.Errorf("Unable to open a new MQTT routing session: %v", err)
							s.watcherCh <- struct{}{}
						}
					}

					networkErrorHandlerRunning = false
				}()
			}
		case <-s.endCh:
			return
		}
	}
}
