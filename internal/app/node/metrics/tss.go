package metrics

import (
	"time"

	"mmesh.dev/m-api-go/grpc/resources/metrics"
	"mmesh.dev/m-api-go/grpc/resources/tss"
)

func GetTSDataPoints() []*tss.DataPoint {
	if hm == nil {
		hm = &metrics.HostMetrics{}
	}

	if networkMetrics == nil {
		networkMetrics = newNetworkMetricsMap()
	}

	return []*tss.DataPoint{
		{
			Timestamp: time.Now().UnixMilli(),
			TimeRange: tss.TimeRange_TTL_1H,
			Metric:    tss.MetricType_NET_RX_BYTES,
			Value:     float64(networkMetrics.networkMetrics.RxTotalBytes),
		},
		{
			Timestamp: time.Now().UnixMilli(),
			TimeRange: tss.TimeRange_TTL_1H,
			Metric:    tss.MetricType_NET_TX_BYTES,
			Value:     float64(networkMetrics.networkMetrics.TxTotalBytes),
		},
		// {
		// 	Timestamp: time.Now().UnixMilli(),
		// 	TimeRange: tss.TimeRange_TTL_1H,
		// 	Metric:    tss.MetricType_NET_DROPPED_PKTS,
		// 	Value:     float64(networkMetrics.networkMetrics.DroppedPkts),
		// },
		{
			Timestamp: time.Now().UnixMilli(),
			TimeRange: tss.TimeRange_TTL_1H,
			Metric:    tss.MetricType_HOST_LOAD_AVG,
			Value:     hm.LoadAvg,
		},
		{
			Timestamp: time.Now().UnixMilli(),
			TimeRange: tss.TimeRange_TTL_1H,
			Metric:    tss.MetricType_HOST_CPU_USAGE,
			Value:     float64(hm.CpuUsage),
		},
		{
			Timestamp: time.Now().UnixMilli(),
			TimeRange: tss.TimeRange_TTL_1H,
			Metric:    tss.MetricType_HOST_MEM_USAGE,
			Value:     float64(hm.MemoryUsage),
		},
		{
			Timestamp: time.Now().UnixMilli(),
			TimeRange: tss.TimeRange_TTL_1H,
			Metric:    tss.MetricType_HOST_DISK_USAGE,
			Value:     float64(hm.DiskUsage),
		},
	}
}
