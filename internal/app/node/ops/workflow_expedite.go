package ops

import (
	"context"

	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-api-go/grpc/resources/ops/workflow"
	"mmesh.dev/m-lib/pkg/mmp"
	"mmesh.dev/m-lib/pkg/xlog"
)

// WorkflowExpedite executes a workflow.
// This function usually will be executed on target nodes.
func WorkflowExpedite(ctx context.Context, payload *mmsp.Payload) error {
	wf := payload.Workflow

	if disabledOps {
		xlog.Alertf("Ops disabled on this node. Unauthorized workflow: %s", wf.WorkflowID)
		return nil
	}

	var ops []*workflow.Operation

	if !wf.Enabled {
		xlog.Warnf("Workflow %s not enabled", wf.WorkflowID)
		return nil
	}

	for _, j := range wf.Jobs {
		for _, t := range j.Tasks {
			for _, a := range t.Actions {
				op := runWorkflowAction(wf, j.Name, t.Name, a)
				ops = append(ops, op)
			}
		}
	}

	wf.Operations = ops

	p := newWorkflowResponse(payload)
	mmp.TxControlQueue <- p

	return nil
}
