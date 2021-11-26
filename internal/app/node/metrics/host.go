package metrics

import (
	"fmt"
	"time"

	"mmesh.dev/m-api-go/grpc/resources/metrics"
	"mmesh.dev/m-api-go/grpc/resources/network"
)

var hm *metrics.HostMetrics

func GetHostMetrics() *metrics.HostMetrics {
	return hm
}

func GetHostDataPoint(n *network.Node) *metrics.DataPoint {
	if hm == nil {
		hm = &metrics.HostMetrics{}
		return nil
	}

	return &metrics.DataPoint{
		Timestamp:     time.Now().Unix(),
		Measurement:   metrics.Measurement_HOST,
		AccountID:     n.AccountID,
		TenantID:      n.TenantID,
		NetID:         n.NetID,
		VRFID:         n.VRFID,
		NodeID:        n.NodeID,
		HostLoadAvg:   hm.LoadAvg,
		HostCpuUsage:  hm.CpuUsage,
		HostMemUsage:  hm.MemoryUsage,
		HostDiskUsage: hm.DiskUsage,
	}
}

func uptimeStr(uptime uint64) string {
	var s string

	days := uptime / (60 * 60 * 24)

	if days == 1 {
		s = fmt.Sprintf("%d day", days)
	} else {
		s = fmt.Sprintf("%d days", days)
	}

	minutes := uptime / 60
	hours := minutes / 60
	hours %= 24
	minutes %= 60

	return fmt.Sprintf("%s, %d hours, %02d minutes", s, hours, minutes)
}
