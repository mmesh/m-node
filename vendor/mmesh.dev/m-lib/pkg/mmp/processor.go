package mmp

import (
	"context"

	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-lib/pkg/logging"
	"mmesh.dev/m-lib/pkg/runtime"
	"x6a.dev/pkg/xlog"
)

const (
	PayloadTypeNodeInit = iota
	PayloadTypeCancelSession

	PayloadTypeCommandShellExec
	PayloadTypeCommandShellExit
	PayloadTypeCommandShellInput
	PayloadTypeCommandShellOutput
	PayloadTypeCommandShellDisabled

	PayloadTypeTransferInit
	PayloadTypeTransferData
	PayloadTypeTransferAck
	PayloadTypeTransferDisabled

	PayloadTypePortFwdListen
	PayloadTypePortFwdDial
	PayloadTypePortFwdDialAck
	PayloadTypePortFwdEnd
	PayloadTypePortFwdData
	PayloadTypePortFwdDisabled

	PayloadTypeWorkflowExpedite
	PayloadTypeWorkflowSchedule
	PayloadTypeWorkflowResponse
)

var RecvCommandQueue = make(chan *mmsp.Payload, 128)
var SendCommandQueue = make(chan *mmsp.Payload, 128)

func CmdEventsHandler(w *runtime.Wrkr) {
	xlog.Infof("Started worker %s", w.Name)
	w.Running = true

	for {
		select {
		case payload := <-RecvCommandQueue:
			xlog.Debugf("Received command on queue from %s", payload.SrcID)
			Processor(context.Background(), payload)
		case <-w.QuitChan:
			w.WG.Done()
			w.Running = false
			xlog.Infof("Stopped worker %s", w.Name)
			return
		}
	}
}

func Processor(ctx context.Context, p *mmsp.Payload) error {
	var err error

	switch p.PayloadType {
	case PayloadTypeNodeInit:
		// err = mmpNodeInit(ctx, p)
		WorkflowControllerQueue <- p

	case PayloadTypeCommandShellExec:
		go shellExec(ctx, p)
	case PayloadTypeCommandShellExit:
		err = shellExit(ctx, p)
	case PayloadTypeCommandShellInput:
		err = shellReadInput(ctx, p)
	case PayloadTypeCommandShellOutput:
		err = shellReadOutput(ctx, p)
	case PayloadTypeCommandShellDisabled:
		shellUnavailable()

	case PayloadTypeTransferInit:
		err = transferInit(ctx, p)
	case PayloadTypeTransferData:
		err = transferDataRx(ctx, p)
	case PayloadTypeTransferAck:
		err = transferAckRx(ctx, p)
	case PayloadTypeTransferDisabled:
		transferUnavailable()

	case PayloadTypePortFwdListen:
		go portFwdListen(ctx, p)
	case PayloadTypePortFwdDial:
		go portFwdDial(ctx, p)
	case PayloadTypePortFwdDialAck:
		err = portFwdDialAck(ctx, p)
	case PayloadTypePortFwdEnd:
		err = portFwdEnd(ctx, p)
	case PayloadTypePortFwdData:
		err = portFwdReadData(ctx, p)
	case PayloadTypePortFwdDisabled:
		portFwdUnavailable()

	case PayloadTypeWorkflowExpedite:
		// err = workflowExpedite(ctx, p)
		WorkflowNodeQueue <- p
	case PayloadTypeWorkflowSchedule:
		// err = workflowSchedule(ctx, p)
		WorkflowNodeQueue <- p
	case PayloadTypeWorkflowResponse:
		// err = workflowResponse(ctx, p)
		WorkflowControllerQueue <- p
	}

	if err != nil {
		logging.Debugf("Unable to process mmp payload: %v", err)
		return err
	}

	return nil
}
