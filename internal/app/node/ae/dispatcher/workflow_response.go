package dispatcher

import (
	"mmesh.dev/mmesh/internal/api/grpc/network/mmsp"
	"mmesh.dev/mmesh/internal/api/grpc/network/resources/ae/workflow"
	"mmesh.dev/mmesh/internal/pkg/mmp"
)

func newWorkflowResponse(p *mmsp.Payload) *mmsp.Payload {
	return &mmsp.Payload{
		SrcID:       p.DstID,
		DstID:       p.SrcID,
		RequesterID: p.RequesterID,
		Interactive: p.Interactive,
		PayloadType: mmp.PayloadTypeWorkflowResponse,
		Workflow: &workflow.Workflow{
			AccountID:    p.Workflow.AccountID,
			AccountToken: p.Workflow.AccountToken,
			ProjectID:    p.Workflow.ProjectID,
			WorkflowID:   p.Workflow.WorkflowID,
			Owner:        p.Workflow.Owner,
			Notify:       p.Workflow.Notify,
			Operations:   p.Workflow.Operations,
		},
	}
}
