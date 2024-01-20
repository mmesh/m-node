package alert

import (
	"fmt"
	"time"

	"mmesh.dev/m-api-go/grpc/resources/events"
	"mmesh.dev/m-api-go/grpc/resources/topology"
)

var cpuAlert, memoryAlert, diskAlert bool

func hostAlert(nr *topology.NodeReq, nodeName string) *events.Event {
	return &events.Event{
		AccountID:    nr.AccountID,
		AccountAlert: true,
		Timestamp:    time.Now().UnixMilli(),
		Source: &events.Source{
			Type: events.SourceType_NODE,
			Node: &events.SourceNode{
				AccountID: nr.AccountID,
				TenantID:  nr.TenantID,
				NodeID:    nr.NodeID,
				NodeName:  nodeName,
			},
		},
		Type:  events.EventType_ALERT,
		Class: events.Class_HOST,
		Group: events.Group_HOST_METRICS,
		// Component: string,
		// Severity: events.Severity,
		// ActionType: events.ActionType,
		// Summary: string,
		CustomDetails: make(map[string]string, 0),
	}
}

func HostUptimeAlert(nr *topology.NodeReq, nodeName string, uptime string) {
	e := hostAlert(nr, nodeName)
	e.Component = "Uptime"
	e.CustomDetails["Uptime"] = uptime
	e.Severity = events.Severity_WARNING
	e.Summary = fmt.Sprintf("[%s] REBOOT detected on node %s", e.Severity.String(), nodeName)
	e.ActionType = events.ActionType_TRIGGER

	newAlertEvent(e)
}

func HostCPUHighAlert(nr *topology.NodeReq, nodeName string, load string) {
	if cpuAlert {
		return
	}

	cpuAlert = true

	e := hostAlert(nr, nodeName)
	e.Component = "Load"
	e.CustomDetails["Load Average"] = load
	e.Severity = events.Severity_WARNING
	e.Summary = fmt.Sprintf("[%s] High LOAD average on node %s", e.Severity.String(), nodeName)
	e.ActionType = events.ActionType_TRIGGER

	newAlertEvent(e)
}

func HostCPULowAlert(nr *topology.NodeReq, nodeName string, load string) {
	if !cpuAlert {
		return
	}

	cpuAlert = false

	e := hostAlert(nr, nodeName)
	e.Component = "Load"
	e.CustomDetails["Load Average"] = load
	e.Severity = events.Severity_INFO
	e.Summary = fmt.Sprintf("[%s] Normal LOAD average on node %s", e.Severity.String(), nodeName)
	e.ActionType = events.ActionType_RESOLVE

	newAlertEvent(e)
}

func HostMemHighAlert(nr *topology.NodeReq, nodeName string, usage string) {
	if memoryAlert {
		return
	}

	memoryAlert = true

	e := hostAlert(nr, nodeName)
	e.Component = "Memory"
	e.CustomDetails["Memory"] = usage
	e.Severity = events.Severity_WARNING
	e.Summary = fmt.Sprintf("[%s] MEMORY usage above 90%% on node %s", e.Severity.String(), nodeName)
	e.ActionType = events.ActionType_TRIGGER

	newAlertEvent(e)
}

func HostMemLowAlert(nr *topology.NodeReq, nodeName string, usage string) {
	if !memoryAlert {
		return
	}

	memoryAlert = false

	e := hostAlert(nr, nodeName)
	e.Component = "Memory"
	e.CustomDetails["Memory"] = usage
	e.Severity = events.Severity_INFO
	e.Summary = fmt.Sprintf("[%s] MEMORY usage under 90%% on node %s", e.Severity.String(), nodeName)
	e.ActionType = events.ActionType_RESOLVE

	newAlertEvent(e)
}

func HostDiskHighAlert(nr *topology.NodeReq, nodeName string, usage string) {
	if diskAlert {
		return
	}

	diskAlert = true

	e := hostAlert(nr, nodeName)
	e.Component = "Disk"
	e.CustomDetails["Disk"] = usage
	e.Severity = events.Severity_WARNING
	e.Summary = fmt.Sprintf("[%s] DISK usage above 90%% on node %s", e.Severity.String(), nodeName)
	e.ActionType = events.ActionType_TRIGGER

	newAlertEvent(e)
}

func HostDiskLowAlert(nr *topology.NodeReq, nodeName string, usage string) {
	if !diskAlert {
		return
	}

	diskAlert = false

	e := hostAlert(nr, nodeName)
	e.Component = "Disk"
	e.CustomDetails["Disk"] = usage
	e.Severity = events.Severity_INFO
	e.Summary = fmt.Sprintf("[%s] DISK usage under 90%% on node %s", e.Severity.String(), nodeName)
	e.ActionType = events.ActionType_RESOLVE

	newAlertEvent(e)
}
