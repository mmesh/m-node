package workflow

import (
	"context"

	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-lib/pkg/xlog"
	"mmesh.dev/m-node/internal/app/node/ops"
)

func Processor(ctx context.Context, pdu *mmsp.WorkflowPDU) {
	if pdu == nil {
		return
	}

	var err error

	switch pdu.Type {
	case mmsp.WorkflowMsgType_WORKFLOW_EXPEDITE:
		err = ops.WorkflowExpedite(ctx, pdu)
	case mmsp.WorkflowMsgType_WORKFLOW_SCHEDULE:
		err = ops.WorkflowSchedule(ctx, pdu)
	}

	if err != nil {
		xlog.Errorf("[mmp] Unable to process mmp workflowPDU (%s): %v",
			pdu.Type.String(), err)
	}
}
