package metrics

import (
	"fmt"

	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/network/mmsp/alert"
	"mmesh.dev/m-lib/pkg/mmp"
)

var cpuAlert, memoryAlert, diskAlert bool

func hostAlert(accountID, nodeID string) *alert.EventPayload {
	mmID := viper.GetString("mm.id")

	return &alert.EventPayload{
		AccountID:    accountID,
		AccountAlert: true,
		SourceID:     mmID,
		// Component:    component,
		Group: "host-metrics",
		// Message:      msg,
		CustomDetails: map[string]string{
			"Account": accountID,
			"Tenant":  viper.GetString("node.tenant"),
			"Network": viper.GetString("node.network"),
			"Subnet":  viper.GetString("node.subnet"),
			"Node":    nodeID,
		},
		EventClass: alert.EventClass_HOST,
		EventType:  alert.EventType_ALERT,
	}
}

func hostUptimeAlert(uptime string) {
	accountID := viper.GetString("node.account")
	nodeID := viper.GetString("node.id")

	e := hostAlert(accountID, nodeID)
	e.Component = "Uptime"
	e.CustomDetails["Uptime"] = uptime
	e.Message = fmt.Sprintf("[%s] REBOOT detected on node %s", accountID, nodeID)
	e.Severity = alert.EventSeverity_WARNING
	e.ActionType = alert.AlertActionType_TRIGGER

	mmp.NewEvent(e)
}

func hostCPUHighAlert(load string) {
	if cpuAlert {
		return
	}

	cpuAlert = true

	accountID := viper.GetString("node.account")
	nodeID := viper.GetString("node.id")

	e := hostAlert(accountID, nodeID)
	e.Component = "Load"
	e.CustomDetails["Load Average"] = load
	e.Message = fmt.Sprintf("[%s] High LOAD average on node %s", accountID, nodeID)
	e.Severity = alert.EventSeverity_WARNING
	e.ActionType = alert.AlertActionType_TRIGGER

	mmp.NewEvent(e)
}

func hostCPULowAlert(load string) {
	if !cpuAlert {
		return
	}

	cpuAlert = false

	accountID := viper.GetString("node.account")
	nodeID := viper.GetString("node.id")

	e := hostAlert(accountID, nodeID)
	e.Component = "Load"
	e.CustomDetails["Load Average"] = load
	e.Message = fmt.Sprintf("[%s] Normal LOAD average on node %s", accountID, nodeID)
	e.Severity = alert.EventSeverity_INFO
	e.ActionType = alert.AlertActionType_RESOLVE

	mmp.NewEvent(e)
}

func hostMemHighAlert(usage string) {
	if memoryAlert {
		return
	}

	memoryAlert = true

	accountID := viper.GetString("node.account")
	nodeID := viper.GetString("node.id")

	e := hostAlert(accountID, nodeID)
	e.Component = "Memory"
	e.CustomDetails["Memory"] = usage
	e.Message = fmt.Sprintf("[%s] MEMORY usage above 90%% on node %s", accountID, nodeID)
	e.Severity = alert.EventSeverity_WARNING
	e.ActionType = alert.AlertActionType_TRIGGER

	mmp.NewEvent(e)
}

func hostMemLowAlert(usage string) {
	if !memoryAlert {
		return
	}

	memoryAlert = false

	accountID := viper.GetString("node.account")
	nodeID := viper.GetString("node.id")

	e := hostAlert(accountID, nodeID)
	e.Component = "Memory"
	e.CustomDetails["Memory"] = usage
	e.Message = fmt.Sprintf("[%s] MEMORY usage under 90%% on node %s", accountID, nodeID)
	e.Severity = alert.EventSeverity_INFO
	e.ActionType = alert.AlertActionType_RESOLVE

	mmp.NewEvent(e)
}

func hostDiskHighAlert(usage string) {
	if diskAlert {
		return
	}

	diskAlert = true

	accountID := viper.GetString("node.account")
	nodeID := viper.GetString("node.id")

	e := hostAlert(accountID, nodeID)
	e.Component = "Disk"
	e.CustomDetails["Disk"] = usage
	e.Message = fmt.Sprintf("[%s] DISK usage above 90%% on node %s", accountID, nodeID)
	e.Severity = alert.EventSeverity_WARNING
	e.ActionType = alert.AlertActionType_TRIGGER

	mmp.NewEvent(e)
}

func hostDiskLowAlert(usage string) {
	if !diskAlert {
		return
	}

	diskAlert = false

	accountID := viper.GetString("node.account")
	nodeID := viper.GetString("node.id")

	e := hostAlert(accountID, nodeID)
	e.Component = "Disk"
	e.CustomDetails["Disk"] = usage
	e.Message = fmt.Sprintf("[%s] DISK usage under 90%% on node %s", accountID, nodeID)
	e.Severity = alert.EventSeverity_INFO
	e.ActionType = alert.AlertActionType_RESOLVE

	mmp.NewEvent(e)
}
