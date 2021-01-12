package metrics

import (
	"sync"

	"mmesh.dev/m-api-go/grpc/network/resources/metrics"
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

	for rmaddr, conns := range rm.relay {
		rmetrics[rmaddr] = &metrics.RelayMetrics{
			Connections: conns,
		}
	}

	return rmetrics
}

func (rm *relayMetricsMap) incr(rmaddr string) {
	rm.Lock()
	defer rm.Unlock()

	rm.relay[rmaddr]++
}

func (rm *relayMetricsMap) decr(rmaddr string) {
	rm.Lock()
	defer rm.Unlock()

	rm.relay[rmaddr]--
}

func GetRelayMetrics() map[string]*metrics.RelayMetrics {
	if relayMetrics == nil {
		relayMetrics = newRelayMetricsMap()
		return nil
	}

	return relayMetrics.getRelayMetrics()
}

func IncrRelayConns(rmaddr string) {
	if relayMetrics == nil {
		relayMetrics = newRelayMetricsMap()
	}

	relayMetrics.incr(rmaddr)
}

func DecrRelayConns(rmaddr string) {
	if relayMetrics == nil {
		relayMetrics = newRelayMetricsMap()
		return
	}

	relayMetrics.decr(rmaddr)
}
