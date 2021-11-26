package metrics

import (
	"sync"
	"time"

	"mmesh.dev/m-api-go/grpc/resources/metrics"
	"mmesh.dev/m-api-go/grpc/resources/network"
)

type networkMetricsMap struct {
	networkMetrics *metrics.NetworkMetrics
	networkTraffic map[string]*metrics.NetworkMetrics
	// netDevStats    *metrics.NetworkMetrics
	sync.RWMutex
}

var networkMetrics *networkMetricsMap

func newNetworkMetricsMap() *networkMetricsMap {
	return &networkMetricsMap{
		networkMetrics: &metrics.NetworkMetrics{},
		networkTraffic: make(map[string]*metrics.NetworkMetrics),
		// netDevStats:    &metrics.NetworkMetrics{},
	}
}

func (m *networkMetricsMap) zero() {
	m.Lock()
	defer m.Unlock()

	m.networkMetrics.TxTotalBytes = 0
	m.networkMetrics.RxTotalBytes = 0
	m.networkMetrics.TxTotalPkts = 0
	m.networkMetrics.RxTotalPkts = 0
	m.networkMetrics.DroppedPkts = 0

	for addr := range m.networkTraffic {
		delete(m.networkTraffic, addr)
	}
}

func (m *networkMetricsMap) update(addr string, tx, rx uint64, droppedPkt bool) {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.networkTraffic[addr]; !ok {
		m.networkTraffic[addr] = &metrics.NetworkMetrics{}
	}

	m.networkMetrics.TxTotalBytes += tx
	m.networkMetrics.RxTotalBytes += rx

	m.networkTraffic[addr].TxTotalBytes += tx
	m.networkTraffic[addr].RxTotalBytes += rx

	if tx > 0 {
		m.networkMetrics.TxTotalPkts++
		m.networkTraffic[addr].TxTotalPkts++
	}
	if rx > 0 {
		m.networkMetrics.RxTotalPkts++
		m.networkTraffic[addr].RxTotalPkts++
	}

	if droppedPkt {
		m.networkMetrics.DroppedPkts++
		m.networkTraffic[addr].DroppedPkts++
	}
}

// func (m *networkMetricsMap) setNetDevStats(txBytes, rxBytes, txPkts, rxPkts uint64) {
// 	m.Lock()
// 	defer m.Unlock()

// 	m.netDevStats.TxTotalBytes = txBytes
// 	m.netDevStats.RxTotalBytes = rxBytes
// 	m.netDevStats.TxTotalPkts = txPkts
// 	m.netDevStats.RxTotalPkts = rxPkts
// }

// func setNetDevStats(txBytes, rxBytes, txPkts, rxPkts uint64) {
// 	if networkMetrics == nil {
// 		networkMetrics = newNetworkMetricsMap()
// 	}

// 	networkMetrics.setNetDevStats(txBytes, rxBytes, txPkts, rxPkts)
// }

func ClearNetworkMetrics() {
	if networkMetrics == nil {
		networkMetrics = newNetworkMetricsMap()
		return
	}

	networkMetrics.zero()
}

func UpdateNetworkMetric(addr string, tx, rx uint64, droppedPkt bool) {
	if networkMetrics == nil {
		networkMetrics = newNetworkMetricsMap()
	}

	networkMetrics.update(addr, tx, rx, droppedPkt)
}

func GetNetworkMetrics() *metrics.NetworkMetrics {
	if networkMetrics == nil {
		networkMetrics = newNetworkMetricsMap()
		return nil
	}

	return networkMetrics.networkMetrics
}

func GetNetworkTraffic() map[string]*metrics.NetworkMetrics {
	if networkMetrics == nil {
		networkMetrics = newNetworkMetricsMap()
		return nil
	}

	return networkMetrics.networkTraffic
}

// func GetNetDevStats() *metrics.NetworkMetrics {
// 	if networkMetrics == nil {
// 		networkMetrics = newNetworkMetricsMap()
// 		return nil
// 	}

// 	return networkMetrics.netDevStats
// }

func GetNetDataPoint(n *network.Node) *metrics.DataPoint {
	if networkMetrics == nil {
		networkMetrics = newNetworkMetricsMap()
		return nil
	}

	return &metrics.DataPoint{
		Timestamp:      time.Now().Unix(),
		Measurement:    metrics.Measurement_NET,
		AccountID:      n.AccountID,
		TenantID:       n.TenantID,
		NetID:          n.NetID,
		VRFID:          n.VRFID,
		NodeID:         n.NodeID,
		NetRxBytes:     networkMetrics.networkMetrics.RxTotalBytes,
		NetTxBytes:     networkMetrics.networkMetrics.TxTotalBytes,
		NetDroppedPkts: networkMetrics.networkMetrics.DroppedPkts,
	}
}
