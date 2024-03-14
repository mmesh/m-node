package metricsdb

import (
	"fmt"
	"math"
	"sort"
	"time"

	"mmesh.dev/m-api-go/grpc/resources/nstore/metricsdb"
	"mmesh.dev/m-lib/pkg/errors"
)

func (tsdb *tsDB) Query(r *metricsdb.HostMetricsRequest) (*metricsdb.HostMetricsResponse, error) {
	hmr := &metricsdb.HostMetricsResponse{
		AccountID: r.Request.AccountID,
		TenantID:  r.Request.TenantID,
		NodeID:    r.Request.NodeID,
		QueryID:   r.Request.QueryID,
		Metrics:   make([]*metricsdb.HostMetrics, 0),
		Timestamp: time.Now().UnixMilli(),
	}

	for _, mt := range getHostMetricTypes(r.Type) {
		keyPrefix := []byte(fmt.Sprintf("%s:", encodeKeyPrefix(r.TimeRange, mt)))

		dpl, err := tsdb.Scan(keyPrefix)
		if err != nil {
			return nil, errors.Wrapf(err, "[%v] function tsdb.Scan()", errors.Trace())
		}

		// tvl := getTimeValues(dpl, mt)
		hml := getHostMetrics(dpl, mt)

		hmr.Metrics = append(hmr.Metrics, hml...)
	}

	hmr.Metrics = aggregateMetrics(hmr.Metrics)

	sort.Slice(hmr.Metrics, func(i, j int) bool {
		return hmr.Metrics[i].Timestamp < hmr.Metrics[j].Timestamp
	})

	return hmr, nil
}

func getHostMetrics(dpl []*metricsdb.HostMetricDataPoint, mt metricsdb.HostMetricType) []*metricsdb.HostMetrics {
	hml := make([]*metricsdb.HostMetrics, 0)

	for _, dp := range dpl {
		switch mt {
		case metricsdb.HostMetricType_NET_RX_BYTES:
			// convert bytes to kbytes
			dp.Value /= 1000
		case metricsdb.HostMetricType_NET_TX_BYTES:
			// convert bytes to kbytes
			dp.Value /= -1000
		}

		e := &metricsdb.HostMetrics{
			Timestamp: dp.Timestamp,
			Data:      make(map[string]*metricsdb.MetricValue, 0),
		}

		e.Data[mt.String()] = &metricsdb.MetricValue{
			Value: roundFloat(dp.Value, 2),
		}

		hml = append(hml, e)
	}

	return hml
}

func getHostMetricTypes(t metricsdb.HostMetricsQueryType) []metricsdb.HostMetricType {
	switch t {
	case metricsdb.HostMetricsQueryType_QUERY_NET_USAGE:
		return []metricsdb.HostMetricType{
			metricsdb.HostMetricType_NET_RX_BYTES,
			metricsdb.HostMetricType_NET_TX_BYTES,
		}
	case metricsdb.HostMetricsQueryType_QUERY_LOAD_AVG:
		return []metricsdb.HostMetricType{
			metricsdb.HostMetricType_HOST_LOAD_AVG,
		}
	case metricsdb.HostMetricsQueryType_QUERY_CPU_USAGE:
		return []metricsdb.HostMetricType{
			metricsdb.HostMetricType_HOST_CPU_USAGE,
		}
	case metricsdb.HostMetricsQueryType_QUERY_MEM_USAGE:
		return []metricsdb.HostMetricType{
			metricsdb.HostMetricType_HOST_MEM_USAGE,
		}
	case metricsdb.HostMetricsQueryType_QUERY_DISK_USAGE:
		return []metricsdb.HostMetricType{
			metricsdb.HostMetricType_HOST_DISK_USAGE,
		}
	}

	return nil
}

func roundFloat(number float64, decimalPlace int) float64 {
	// Calculate the 10 to the power of decimal place
	temp := math.Pow(10, float64(decimalPlace))
	// Multiply floating-point number with 10**decimalPlace and round it
	// Divide the rounded number with 10**decimalPlace to get decimal place rounding
	return math.Round(number*temp) / temp
}

/*
func getTimeValues(dpl []*metricsdb.HostMetricDataPoint, m metricsdb.HostMetricType) []*metricsdb.TimeValue {
	tvl := make([]*metricsdb.TimeValue, 0)

	for _, dp := range dpl {
		switch m {
		case metricsdb.HostMetricType_NET_RX_BYTES:
			// convert bytes to kbytes
			dp.Value /= 1000
		case metricsdb.HostMetricType_NET_TX_BYTES:
			// convert bytes to kbytes
			dp.Value /= -1000
		}
		tvl = append(tvl, &metricsdb.TimeValue{
			Timestamp: dp.Timestamp,
			Value:     roundFloat(dp.Value, 2),
		})
	}

	return tvl
}
*/

/*
func (tsdb *tsDB) QueryAll(r *metricsdb.HostMetricsRequest) (*metricsdb.HostMetricsResponse, error) {
	hmr := &metricsdb.HostMetricsResponse{
		AccountID: r.Request.AccountID,
		TenantID:  r.Request.TenantID,
		NodeID:    r.Request.NodeID,
		QueryID:   r.Request.QueryID,
		Metrics:   make(map[string]*metricsdb.HostMetrics, 0),
		Timestamp: time.Now().UnixMilli(),
	}

	for _, tr := range timeRanges() {
		metrics := newHostMetrics()

		for _, m := range metricList {
			keyPrefix := []byte(fmt.Sprintf("%s:", encodeKeyPrefix(tr, m)))

			dpl, err := tsdb.Scan(keyPrefix)
			if err != nil {
				return nil, errors.Wrapf(err, "[%v] function tsdb.Scan()", errors.Trace())
			}

			tvl := getTimeValues(dpl, m)

			switch m {
			case metricsdb.HostMetricType_NET_RX_BYTES:
				metrics.NetUsage.NetRxBytes = tvl
			case metricsdb.HostMetricType_NET_TX_BYTES:
				metrics.NetUsage.NetTxBytes = tvl
			case metricsdb.HostMetricType_HOST_LOAD_AVG:
				metrics.LoadAvg = tvl
			case metricsdb.HostMetricType_HOST_CPU_USAGE:
				metrics.CpuUsage = tvl
			case metricsdb.HostMetricType_HOST_MEM_USAGE:
				metrics.MemUsage = tvl
			case metricsdb.HostMetricType_HOST_DISK_USAGE:
				metrics.DiskUsage = tvl
			}
		}

		hmr.Metrics[tr.String()] = metrics
	}

	return hmr, nil
}
*/

/*
func timeRanges() []nstore.TimeRange {
	return []nstore.TimeRange{
		nstore.TimeRange_TTL_1H,
		nstore.TimeRange_TTL_6H,
		nstore.TimeRange_TTL_12H,
		nstore.TimeRange_TTL_24H,
		nstore.TimeRange_TTL_7D,
		nstore.TimeRange_TTL_14D,
		nstore.TimeRange_TTL_30D,
		nstore.TimeRange_TTL_365D,
	}
}
*/

/*
var metricList []metricsdb.HostMetricType = []metricsdb.HostMetricType{
	metricsdb.HostMetricType_NET_RX_BYTES,
	metricsdb.HostMetricType_NET_TX_BYTES,
	metricsdb.HostMetricType_HOST_LOAD_AVG,
	metricsdb.HostMetricType_HOST_CPU_USAGE,
	metricsdb.HostMetricType_HOST_MEM_USAGE,
	metricsdb.HostMetricType_HOST_DISK_USAGE,
}
*/
