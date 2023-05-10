//go:build !windows
// +build !windows

package metrics

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/resources/metrics"
	"mmesh.dev/m-api-go/grpc/resources/topology"
)

func UpdateHostMetrics(nr *topology.NodeReq) {
	nodeName := viper.GetString("nodeName")

	if hm == nil {
		hm = &metrics.HostMetrics{}
	}

	hostInfo, err := host.Info()
	if err != nil {
		return
	}
	// cpuCount, err := cpu.Counts(true)
	// if err != nil {
	// 	return
	// }
	cpuPercent, err := cpu.Percent(0, false)
	if err != nil {
		return
	}
	loadAvg, err := load.Avg()
	if err != nil {
		return
	}
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		return
	}
	diskInfo, err := disk.Usage("/")
	if err != nil {
		return
	}

	hm.OS = hostInfo.OS

	if hostInfo.Uptime > 0 {
		hm.Uptime = uptimeStr(hostInfo.Uptime)
		if hostInfo.Uptime < 300 { // uptime < 300 seconds
			go hostUptimeAlert(nr, nodeName, hm.Uptime)
		}
	}

	hm.LoadAvg = loadAvg.Load5
	// if (loadAvg.Load5 / float64(cpuCount)) > 2.00 {
	// 	hm.CpuPressure = true
	// 	go hostCPUHighAlert(nr, nodeName, fmt.Sprintf("%f", loadAvg.Load5/float64(cpuCount)))
	// } else {
	// 	hm.CpuPressure = false
	// 	go hostCPULowAlert(nr, nodeName, fmt.Sprintf("%f", loadAvg.Load5/float64(cpuCount)))
	// }

	hm.CpuUsage = uint64(cpuPercent[0])
	if hm.CpuUsage > 90 {
		hm.CpuPressure = true
		go hostCPUHighAlert(nr, nodeName, fmt.Sprintf("%d%%", hm.CpuUsage))
	} else {
		hm.CpuPressure = false
		go hostCPULowAlert(nr, nodeName, fmt.Sprintf("%d%%", hm.CpuUsage))
	}

	hm.MemoryUsage = uint64(memInfo.UsedPercent)
	if hm.MemoryUsage > 90 {
		hm.MemoryPressure = true
		go hostMemHighAlert(nr, nodeName, fmt.Sprintf("%d%%", hm.MemoryUsage))
	} else {
		hm.MemoryPressure = false
		go hostMemLowAlert(nr, nodeName, fmt.Sprintf("%d%%", hm.MemoryUsage))
	}

	hm.DiskUsage = uint64(diskInfo.UsedPercent)
	if hm.DiskUsage > 90 {
		hm.DiskPressure = true
		go hostDiskHighAlert(nr, nodeName, fmt.Sprintf("%d%%", hm.DiskUsage))
	} else {
		hm.DiskPressure = false
		go hostDiskLowAlert(nr, nodeName, fmt.Sprintf("%d%%", hm.DiskUsage))
	}
}
