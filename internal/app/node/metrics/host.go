package metrics

import (
	"fmt"
	"time"

	linuxproc "github.com/c9s/goprocinfo/linux"
	"mmesh.dev/m-api-go/grpc/resources/metrics"
)

var hm *metrics.HostMetrics

func UpdateHostMetrics() {
	if hm == nil {
		hm = &metrics.HostMetrics{}
	}

	uptime, err := linuxproc.ReadUptime("/proc/uptime")
	if err != nil {
		return
	}
	cpuInfo, err := linuxproc.ReadCPUInfo("/proc/cpuinfo")
	if err != nil {
		return
	}
	loadAvg, err := linuxproc.ReadLoadAvg("/proc/loadavg")
	if err != nil {
		return
	}
	memInfo, err := linuxproc.ReadMemInfo("/proc/meminfo")
	if err != nil {
		return
	}
	diskInfo, err := linuxproc.ReadDisk("/")
	if err != nil {
		return
	}

	hm.Uptime = uptime.GetTotalDuration().String()
	if uptime.GetTotalDuration() < time.Duration(5*time.Minute) {
		go hostUptimeAlert(uptime.GetTotalDuration().String())
	}

	hm.LoadAvg = float32(loadAvg.Last1Min)
	if (loadAvg.Last1Min / float64(cpuInfo.NumCore())) > 1.70 {
		hm.CpuPressure = true
		go hostCPUHighAlert(fmt.Sprintf("%f", loadAvg.Last1Min/float64(cpuInfo.NumCore())))
	} else {
		hm.CpuPressure = false
		go hostCPULowAlert(fmt.Sprintf("%f", loadAvg.Last1Min/float64(cpuInfo.NumCore())))
	}

	hm.MemoryUsage = 100 - ((memInfo.MemAvailable * 100) / memInfo.MemTotal)
	if hm.MemoryUsage > 90 {
		hm.MemoryPressure = true
		go hostMemHighAlert(fmt.Sprintf("%d%%", hm.MemoryUsage))
	} else {
		hm.MemoryPressure = false
		go hostMemLowAlert(fmt.Sprintf("%d%%", hm.MemoryUsage))
	}

	hm.DiskUsage = (diskInfo.Used * 100) / diskInfo.All
	if hm.DiskUsage > 90 {
		hm.DiskPressure = true
		go hostDiskHighAlert(fmt.Sprintf("%d%%", hm.DiskUsage))
	} else {
		hm.DiskPressure = false
		go hostDiskLowAlert(fmt.Sprintf("%d%%", hm.DiskUsage))
	}
}

func GetHostMetrics() *metrics.HostMetrics {
	return hm
}
