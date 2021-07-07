package k8s

import (
	"fmt"

	"github.com/spf13/viper"
	v1 "k8s.io/api/core/v1"
	"mmesh.dev/m-api-go/grpc/network/mmsp/alert"
	"mmesh.dev/m-lib/pkg/mmid"
	"mmesh.dev/m-lib/pkg/mmp"
)

// managePodEvent is the business logic of the controller.
func (c *controller) managePodEvent(p *v1.Pod, evt eventType) error {
	clusterName := p.ObjectMeta.ClusterName
	ns := p.ObjectMeta.Namespace
	podName := p.ObjectMeta.Name
	// generateName := p.ObjectMeta.GenerateName

	if len(clusterName) == 0 {
		clusterName = "n/a"
	}

	var componentName string
	if p.ObjectMeta.OwnerReferences != nil {
		if len(p.ObjectMeta.OwnerReferences) > 0 {
			componentName = p.ObjectMeta.OwnerReferences[0].Name
		}
	}
	if len(componentName) == 0 {
		componentName = podName
	}

	if len(ns) == 0 || len(podName) == 0 || len(componentName) == 0 {
		return nil
	}

	mmID := viper.GetString("mm.id")
	accountID := mmid.MMNodeID(mmID).AccountID()

	eventPayload := &alert.EventPayload{
		AccountID:    accountID,
		AccountAlert: true,
		SourceID:     mmID,
		Component:    fmt.Sprintf("kubernetes:%s:%s:%s", clusterName, ns, componentName),
		Group:        fmt.Sprintf("kubernetes:%s:%s", clusterName, ns),
		// Message:      msg,
		CustomDetails: map[string]string{
			"Account":         accountID,
			"Backend":         "kubernetes",
			"Cluster":         clusterName,
			"Namespace":       ns,
			"Parent Resource": componentName,
		},
		EventClass: alert.EventClass_KUBERNETES,
	}

	switch evt {
	case eventAdd:
		msg := fmt.Sprintf("[%s] Pod %s ADDED to kubernetes namespace %s", accountID, podName, ns)
		eventPayload.Message = msg
		eventPayload.EventType = alert.EventType_CHANGE
		eventPayload.Severity = alert.EventSeverity_INFO
		mmp.NewEvent(eventPayload)
	case eventUpdate:
		var msg string
		switch p.Status.Phase {
		case v1.PodPending:
			msg = fmt.Sprintf("[%s] Pod %s PENDING in kubernetes namespace %s", accountID, podName, ns)
			eventPayload.Severity = alert.EventSeverity_WARNING
			eventPayload.ActionType = alert.AlertActionType_TRIGGER
		case v1.PodRunning:
			msg = fmt.Sprintf("[%s] Pod %s RUNNING in kubernetes namespace %s", accountID, podName, ns)
			eventPayload.Severity = alert.EventSeverity_INFO
			eventPayload.ActionType = alert.AlertActionType_RESOLVE
		case v1.PodSucceeded:
			msg = fmt.Sprintf("[%s] Pod %s SUCCEEDED in kubernetes namespace %s", accountID, podName, ns)
			eventPayload.Severity = alert.EventSeverity_INFO
			eventPayload.ActionType = alert.AlertActionType_RESOLVE
		case v1.PodFailed:
			msg = fmt.Sprintf("[%s] Pod %s FAILED in kubernetes namespace %s", accountID, podName, ns)
			eventPayload.Severity = alert.EventSeverity_ERROR
			eventPayload.ActionType = alert.AlertActionType_TRIGGER
		case v1.PodUnknown:
			msg = fmt.Sprintf("[%s] Pod %s UNKNOWN in kubernetes namespace %s", accountID, podName, ns)
			eventPayload.Severity = alert.EventSeverity_WARNING
			eventPayload.ActionType = alert.AlertActionType_TRIGGER
		}
		eventPayload.Message = msg
		eventPayload.EventType = alert.EventType_ALERT
		mmp.NewEvent(eventPayload)
	case eventDelete:
		msg := fmt.Sprintf("[%s] Pod %s DELETED from kubernetes namespace %s", accountID, podName, ns)
		eventPayload.Message = msg
		eventPayload.EventType = alert.EventType_CHANGE
		eventPayload.Severity = alert.EventSeverity_WARNING
		mmp.NewEvent(eventPayload)
	}

	return nil
}
