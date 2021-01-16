package metrics

import (
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

	hm.LoadAvg = float32(loadAvg.Last1Min)
	if (loadAvg.Last1Min / float64(cpuInfo.NumCore())) > 1 {
		hm.CpuPressure = true
	} else {
		hm.CpuPressure = false
	}

	hm.MemoryUsage = 100 - ((memInfo.MemAvailable * 100) / memInfo.MemTotal)
	if hm.MemoryUsage > 90 {
		hm.MemoryPressure = true
	} else {
		hm.MemoryPressure = false
	}

	hm.DiskUsage = (diskInfo.Used * 100) / diskInfo.All
	if hm.DiskUsage > 90 {
		hm.DiskPressure = true
	} else {
		hm.DiskPressure = false
	}
}

func GetHostMetrics() *metrics.HostMetrics {
	return hm
}
