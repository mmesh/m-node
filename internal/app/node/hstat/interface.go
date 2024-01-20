package hstat

import (
	"net/netip"
	"sync"
	"time"

	"mmesh.dev/m-api-go/grpc/resources/nstore"
	metricsdb_pb "mmesh.dev/m-api-go/grpc/resources/nstore/metricsdb"
	"mmesh.dev/m-api-go/grpc/resources/nstore/netdb"
	"mmesh.dev/m-api-go/grpc/resources/topology"
	"mmesh.dev/m-node/internal/app/node/kvstore/db/metricsdb"
)

type netTrafficData struct {
	addr       netip.Addr
	txBytes    uint64
	rxBytes    uint64
	droppedPkt bool
}

var netTrafficStatQueue = make(chan *netTrafficData, 256)

type Interface interface {
	GetHostMetrics() *topology.HostMetrics
	Close()

	dump()
	wrkr(nodeReq *topology.NodeReq)
}

type netTraffic struct {
	rx *netdb.TrafficCounter
	tx *netdb.TrafficCounter
}

type hstats struct {
	host       *topology.HostMetrics
	netTraffic *netTraffic
	// byAddr     map[netip.Addr]*netTraffic
	closeCh chan struct{}
	sync.RWMutex
}

func Init(nodeReq *topology.NodeReq) Interface {
	hs := &hstats{
		host: &topology.HostMetrics{},
		netTraffic: &netTraffic{
			rx: &netdb.TrafficCounter{},
			tx: &netdb.TrafficCounter{},
		},
		// byAddr:  make(map[netip.Addr]*netTraffic, 0),
		closeCh: make(chan struct{}, 1),
	}

	go hs.wrkr(nodeReq)

	return hs
}

func NewTrafficData(addr netip.Addr, txBytes, rxBytes uint64, droppedPkt bool) {
	netTrafficStatQueue <- &netTrafficData{
		addr:       addr,
		txBytes:    txBytes,
		rxBytes:    rxBytes,
		droppedPkt: droppedPkt,
	}
}

func (hs *hstats) GetHostMetrics() *topology.HostMetrics {
	return hs.host
}

func (hs *hstats) Close() {
	hs.closeCh <- struct{}{}
}

func (hs *hstats) dump() {
	hs.Lock()
	defer hs.Unlock()

	hmdps := []*metricsdb_pb.HostMetricDataPoint{
		{
			Timestamp: time.Now().UnixMilli(),
			TimeRange: nstore.TimeRange_TTL_1H,
			Metric:    metricsdb_pb.HostMetricType_NET_TX_BYTES,
			Value:     float64(hs.netTraffic.tx.Bytes),
		},
		{
			Timestamp: time.Now().UnixMilli(),
			TimeRange: nstore.TimeRange_TTL_1H,
			Metric:    metricsdb_pb.HostMetricType_NET_RX_BYTES,
			Value:     float64(hs.netTraffic.rx.Bytes),
		},
		// {
		// 	Timestamp: time.Now().UnixMilli(),
		// 	TimeRange: nstore.TimeRange_TTL_1H,
		// 	Metric:    metricsdb_pb.HostMetricType_NET_DROPPED_PKTS,
		// 	Value:     float64(hs.netTraffic.DroppedPkts),
		// },
		{
			Timestamp: time.Now().UnixMilli(),
			TimeRange: nstore.TimeRange_TTL_1H,
			Metric:    metricsdb_pb.HostMetricType_HOST_LOAD_AVG,
			Value:     hs.host.LoadAvg,
		},
		{
			Timestamp: time.Now().UnixMilli(),
			TimeRange: nstore.TimeRange_TTL_1H,
			Metric:    metricsdb_pb.HostMetricType_HOST_CPU_USAGE,
			Value:     float64(hs.host.CpuUsage),
		},
		{
			Timestamp: time.Now().UnixMilli(),
			TimeRange: nstore.TimeRange_TTL_1H,
			Metric:    metricsdb_pb.HostMetricType_HOST_MEM_USAGE,
			Value:     float64(hs.host.MemoryUsage),
		},
		{
			Timestamp: time.Now().UnixMilli(),
			TimeRange: nstore.TimeRange_TTL_1H,
			Metric:    metricsdb_pb.HostMetricType_HOST_DISK_USAGE,
			Value:     float64(hs.host.DiskUsage),
		},
	}

	hs.netTraffic = &netTraffic{
		rx: &netdb.TrafficCounter{},
		tx: &netdb.TrafficCounter{},
	}
	// hs.byAddr = make(map[netip.Addr]*netTraffic, 0)

	metricsdb.InputQueue <- hmdps
}

func (hs *hstats) wrkr(nodeReq *topology.NodeReq) {
	ticker := time.NewTicker(120 * time.Second) // 2 minutes
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			hs.updateSys(nodeReq)
			hs.dump()
		case ntd := <-netTrafficStatQueue:
			hs.updateNet(ntd.addr, ntd.txBytes, ntd.rxBytes)
		case <-hs.closeCh:
			return
		}
	}
}
