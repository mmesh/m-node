package tss

import (
	"fmt"
	"math"
	"time"

	"mmesh.dev/m-api-go/grpc/resources/tss"
	"mmesh.dev/m-lib/pkg/errors"
)

var timeRangeList []tss.TimeRange = []tss.TimeRange{
	tss.TimeRange_TTL_1H,
	tss.TimeRange_TTL_6H,
	tss.TimeRange_TTL_12H,
	tss.TimeRange_TTL_24H,
	tss.TimeRange_TTL_7D,
	tss.TimeRange_TTL_14D,
	tss.TimeRange_TTL_30D,
	tss.TimeRange_TTL_365D,
}

var metricList []tss.MetricType = []tss.MetricType{
	tss.MetricType_NET_RX_BYTES,
	tss.MetricType_NET_TX_BYTES,
	tss.MetricType_HOST_LOAD_AVG,
	tss.MetricType_HOST_CPU_USAGE,
	tss.MetricType_HOST_MEM_USAGE,
	tss.MetricType_HOST_DISK_USAGE,
}

func newHostMetrics() *tss.HostMetrics {
	return &tss.HostMetrics{
		NetUsage: &tss.NetMetrics{},
	}
}

func (tsdb *tsDB) Query(req *tss.MetricsRequest) (*tss.NodeMetrics, error) {
	nm := &tss.NodeMetrics{
		AccountID: req.AccountID,
		TenantID:  req.TenantID,
		// NetID:     req.NetID,
		// SubnetID:  req.SubnetID,
		NodeID:    req.NodeID,
		QueryID:   req.QueryID,
		Metrics:   make(map[string]*tss.HostMetrics, 0),
		Timestamp: time.Now().UnixMilli(),
	}

	for _, tr := range timeRangeList {
		metrics := newHostMetrics()

		for _, m := range metricList {
			keyPrefix := []byte(fmt.Sprintf("%s:", encodeKeyPrefix(tr, m)))

			dpl, err := tsdb.Scan(keyPrefix)
			if err != nil {
				return nil, errors.Wrapf(err, "[%v] function tsdb.Scan()", errors.Trace())
			}

			tvl := getTimeValues(dpl, m)

			switch m {
			case tss.MetricType_NET_RX_BYTES:
				metrics.NetUsage.NetRxBytes = tvl
			case tss.MetricType_NET_TX_BYTES:
				metrics.NetUsage.NetTxBytes = tvl
			case tss.MetricType_HOST_LOAD_AVG:
				metrics.LoadAvg = tvl
			case tss.MetricType_HOST_CPU_USAGE:
				metrics.CpuUsage = tvl
			case tss.MetricType_HOST_MEM_USAGE:
				metrics.MemUsage = tvl
			case tss.MetricType_HOST_DISK_USAGE:
				metrics.DiskUsage = tvl
			}
		}

		nm.Metrics[tr.String()] = metrics
	}

	return nm, nil
}

func getTimeValues(dpl []*tss.DataPoint, m tss.MetricType) []*tss.TimeValue {
	tvl := make([]*tss.TimeValue, 0)

	for _, dp := range dpl {
		switch m {
		case tss.MetricType_NET_RX_BYTES:
			// convert bytes to kbytes
			dp.Value /= 1000
		case tss.MetricType_NET_TX_BYTES:
			// convert bytes to kbytes
			dp.Value /= -1000
		}
		tvl = append(tvl, &tss.TimeValue{
			Timestamp: dp.Timestamp,
			Value:     roundFloat(dp.Value, 2),
		})
	}

	return tvl
}

func roundFloat(number float64, decimalPlace int) float64 {
	// Calculate the 10 to the power of decimal place
	temp := math.Pow(10, float64(decimalPlace))
	// Multiply floating-point number with 10**decimalPlace and round it
	// Divide the rounded number with 10**decimalPlace to get decimal place rounding
	return math.Round(number*temp) / temp
}
