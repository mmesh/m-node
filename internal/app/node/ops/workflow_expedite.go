package ops

import (
	"context"

	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-api-go/grpc/resources/ops/workflow"
	"mmesh.dev/m-lib/pkg/mmp"
	"x6a.dev/pkg/xlog"
)

// WorkflowExpedite executes a workflow.
// This function usually will be executed on target nodes.
func WorkflowExpedite(ctx context.Context, payload *mmsp.Payload) error {
	var ops []*workflow.Operation

	wf := payload.Workflow

	if !wf.Enabled {
		xlog.Warnf("Workflow %s not enabled", wf.WorkflowID)
		return nil
	}

	for _, j := range wf.Jobs {
		for _, t := range j.Tasks {
			for _, a := range t.Actions {
				op := runWorkflowAction(wf, a)
				ops = append(ops, op)
			}
		}
	}

	wf.Operations = ops

	p := newWorkflowResponse(payload)
	mmp.TxControlQueue <- p

	return nil
}
