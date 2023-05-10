package ops

import (
	"context"

	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-api-go/grpc/resources/ops"
	"mmesh.dev/m-lib/pkg/mmp/queuing"
	"mmesh.dev/m-lib/pkg/xlog"
)

// WorkflowExpedite executes a workflow.
// This function usually will be executed on target nodes.
func WorkflowExpedite(ctx context.Context, pdu *mmsp.WorkflowPDU) error {
	wf := pdu.Workflow

	if disabledOps {
		xlog.Alertf("[ops] Ops disabled on this node. Unauthorized workflow: %s", wf.WorkflowID)
		return nil
	}

	var taskLogs []*ops.TaskLog

	if !wf.Enabled {
		xlog.Warnf("[ops] Workflow %s not enabled", wf.WorkflowID)
		return nil
	}

	if len(wf.Tasks) == 0 {
		xlog.Warnf("[ops] Task not found in workflow %s", wf.WorkflowID)
		return nil
	}

	for _, t := range wf.Tasks {
		taskLog := runWorkflowTask(wf, t)
		taskLogs = append(taskLogs, taskLog)
	}

	wf.TaskLogs = taskLogs

	p := newWorkflowResponse(pdu)
	queuing.TxControlQueue <- p

	return nil
}
