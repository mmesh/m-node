package metrics

import (
	"sync"

	"mmesh.dev/m-api-go/grpc/resources/metrics"
)

type relayMetricsMap struct {
	relay map[string]int32
	sync.RWMutex
}

var relayMetrics *relayMetricsMap

func newRelayMetricsMap() *relayMetricsMap {
	return &relayMetricsMap{
		relay: make(map[string]int32),
	}
}

func (rm *relayMetricsMap) getRelayMetrics() map[string]*metrics.RelayMetrics {
	var rmetrics = make(map[string]*metrics.RelayMetrics)

	rm.RLock()
	defer rm.RUnlock()

	for rP2PHostID, conns := range rm.relay {
		rmetrics[rP2PHostID] = &metrics.RelayMetrics{
			Connections: conns,
		}
	}

	return rmetrics
}

func (rm *relayMetricsMap) incr(rP2PHostID string) {
	rm.Lock()
	defer rm.Unlock()

	rm.relay[rP2PHostID]++
}

func (rm *relayMetricsMap) decr(rP2PHostID string) {
	rm.Lock()
	defer rm.Unlock()

	rm.relay[rP2PHostID]--
}

func GetRelayMetrics() map[string]*metrics.RelayMetrics {
	if relayMetrics == nil {
		relayMetrics = newRelayMetricsMap()
		return nil
	}

	return relayMetrics.getRelayMetrics()
}

func IncrRelayConns(rP2PHostID string) {
	if relayMetrics == nil {
		relayMetrics = newRelayMetricsMap()
	}

	relayMetrics.incr(rP2PHostID)
}

func DecrRelayConns(rP2PHostID string) {
	if relayMetrics == nil {
		relayMetrics = newRelayMetricsMap()
		return
	}

	relayMetrics.decr(rP2PHostID)
}
