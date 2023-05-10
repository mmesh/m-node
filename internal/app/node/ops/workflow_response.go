package ops

import (
	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-api-go/grpc/resources/ops"
)

func newWorkflowResponse(pdu *mmsp.WorkflowPDU) *mmsp.Payload {
	mmID := viper.GetString("mm.id")

	return &mmsp.Payload{
		SrcID: mmID,
		Type:  mmsp.PDUType_WORKFLOW,
		WorkflowPDU: &mmsp.WorkflowPDU{
			Type: mmsp.WorkflowMsgType_WORKFLOW_RESPONSE,
			Workflow: &ops.Workflow{
				AccountID:   pdu.Workflow.AccountID,
				TenantID:    pdu.Workflow.TenantID,
				ProjectID:   pdu.Workflow.ProjectID,
				WorkflowID:  pdu.Workflow.WorkflowID,
				Name:        pdu.Workflow.Name,
				Description: pdu.Workflow.Description,
				Notify:      pdu.Workflow.Notify,
				TaskLogs:    pdu.Workflow.TaskLogs,
			},
		},
	}

	// return &mmsp.Payload{
	// 	SrcID:       p.DstID,
	// 	DstID:       p.SrcID,
	// 	RequesterID: p.RequesterID,
	// 	Interactive: p.Interactive,
	// 	PayloadType: mmsp.PayloadType_WORKFLOW_RESPONSE,
	// 	Workflow: &ops.Workflow{
	// 		AccountID: p.Workflow.AccountID,
	// 		TenantID:  p.Workflow.TenantID,
	// 		ProjectID:  p.Workflow.ProjectID,
	// 		WorkflowID: p.Workflow.WorkflowID,
	//      Name:       p.Workflow.Name,
	//      Description: p.Workflow.Description,
	// 		OwnerUserID:      p.Workflow.OwnerUserID,
	// 		Notify:     p.Workflow.Notify,
	// 		TaskLogs: p.Workflow.TaskLogs,
	// 	},
	// }
}
