package metricsdb

import (
	"mmesh.dev/m-api-go/grpc/resources/nstore/metricsdb"
)

const maxEntries = 30

func aggregateMetrics(hm []*metricsdb.HostMetrics) []*metricsdb.HostMetrics {
	// if possible, it's always needed to run this function
	if len(hm) < 2 {
		return hm
	}

	tmStart := hm[0].Timestamp
	tmEnd := hm[len(hm)-1].Timestamp

	aggInterval := (tmEnd - tmStart) / maxEntries

	hmm := make(map[int64]*metricsdb.HostMetrics, 0)

	for i := 0; i < maxEntries; i++ {
		tm := tmStart + (int64(i) * aggInterval)

		hmm[tm] = &metricsdb.HostMetrics{
			Timestamp: tm + (aggInterval / 2),
			Data:      make(map[string]*metricsdb.MetricValue, 0),
		}
	}

	for tm, aggHM := range hmm {
		for _, hme := range hm {
			if hme.Timestamp < tm {
				continue
			}

			if hme.Timestamp > (tm + aggInterval) {
				continue
			}

			for mt, mv := range hme.Data {
				if mt == metricsdb.HostMetricType_HOST_CPU_USAGE.String() ||
					mt == metricsdb.HostMetricType_HOST_MEM_USAGE.String() ||
					mt == metricsdb.HostMetricType_HOST_DISK_USAGE.String() ||
					mt == metricsdb.HostMetricType_HOST_LOAD_AVG.String() {
					// these metrics are average values (% usage)

					if _, ok := aggHM.Data[mt]; !ok {
						aggHM.Data[mt] = &metricsdb.MetricValue{
							Value: mv.Value,
						}
					}

					v := (aggHM.Data[mt].Value + mv.Value) / 2 // get average
					aggHM.Data[mt].Value = v
				} else if mt == metricsdb.HostMetricType_NET_RX_BYTES.String() ||
					mt == metricsdb.HostMetricType_NET_TX_BYTES.String() {
					// these metrics are sums (bytes or packets)

					if _, ok := aggHM.Data[mt]; !ok {
						aggHM.Data[mt] = &metricsdb.MetricValue{}
					}

					aggHM.Data[mt].Value += mv.Value
				}
			}
		}
	}

	nhm := make([]*metricsdb.HostMetrics, 0)

	for _, aggHM := range hmm {
		isZero := true

		for _, mv := range aggHM.Data {
			if mv.Value > 0 {
				isZero = false
				break
			}
		}

		if isZero {
			continue
		}

		nhm = append(nhm, aggHM)
	}

	return nhm
}
