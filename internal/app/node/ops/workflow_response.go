package ops

import (
	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-api-go/grpc/resources/ops/workflow"
)

func newWorkflowResponse(p *mmsp.Payload) *mmsp.Payload {
	return &mmsp.Payload{
		SrcID:       p.DstID,
		DstID:       p.SrcID,
		RequesterID: p.RequesterID,
		Interactive: p.Interactive,
		PayloadType: mmsp.PayloadType_WORKFLOW_RESPONSE,
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
