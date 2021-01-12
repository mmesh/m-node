package dispatcher

import (
	"context"

	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-lib/pkg/logging"
	"mmesh.dev/m-lib/pkg/mmp"
	"mmesh.dev/m-lib/pkg/runtime"
	"x6a.dev/pkg/xlog"
)

func WorkflowEventHandler(w *runtime.Wrkr) {
	xlog.Infof("Started worker %s", w.Name)
	w.Running = true

	for {
		select {
		case payload := <-mmp.WorkflowNodeQueue:
			xlog.Debugf("Received command on queue from %s", payload.SrcID)
			nodeWorkflowProcessor(context.TODO(), payload)
		case <-w.QuitChan:
			w.WG.Done()
			w.Running = false
			xlog.Infof("Stopped worker %s", w.Name)
			return
		}
	}
}

func nodeWorkflowProcessor(ctx context.Context, p *mmsp.Payload) error {
	var err error

	switch p.PayloadType {
	case mmp.PayloadTypeWorkflowExpedite:
		err = workflowExpedite(ctx, p)
	case mmp.PayloadTypeWorkflowSchedule:
		err = workflowSchedule(ctx, p)
	}

	if err != nil {
		logging.Debugf("Unable to process mmp workflow payload: %v", err)
		return err
	}

	return nil
}
