package dispatcher

import (
	"context"

	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"x6a.dev/pkg/xlog"
)

// workflowSchedule configure the local cron with workflow-related operations.
// This function usually will be executed on agents.
func workflowSchedule(ctx context.Context, payload *mmsp.Payload) error {
	wf := payload.Workflow

	if wf.Enabled {
		xlog.Infof("Scheduling workflow %s", wf.WorkflowID)
	} else {
		xlog.Infof("Removing disabled workflow %s", wf.WorkflowID)
	}

	if wf.Triggers.Schedule.DateTime != nil {
		atdCommandQueue <- payload
	}

	if len(wf.Triggers.Schedule.Crontab) > 0 {
		cronCommandQueue <- payload
	}

	return nil
}
