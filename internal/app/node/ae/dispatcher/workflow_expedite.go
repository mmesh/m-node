package dispatcher

import (
	"context"

	"mmesh.dev/mmesh/internal/api/grpc/network/mmsp"
	"mmesh.dev/mmesh/internal/api/grpc/network/resources/ae/operation"
	"mmesh.dev/mmesh/internal/pkg/mmp"
	"x6a.dev/pkg/xlog"
)

// workflowExpedite executes a workflow.
// This function usually will be executed on target nodes.
func workflowExpedite(ctx context.Context, payload *mmsp.Payload) error {
	var ops []*operation.Operation

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
	mmp.SendCommandQueue <- p

	return nil
}
