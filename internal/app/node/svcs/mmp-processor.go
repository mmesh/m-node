package svcs

import (
	"context"

	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-lib/pkg/logging"
	"mmesh.dev/m-lib/pkg/mmp"
	"mmesh.dev/m-lib/pkg/runtime"
	"mmesh.dev/m-node/internal/app/node/ops"
	"x6a.dev/pkg/xlog"
)

func MMPProcessor(w *runtime.Wrkr) {
	xlog.Infof("Started worker %s", w.Name)
	w.Running = true

	close := make(chan struct{}, 1)

	go func() {
		for {
			select {
			case payload := <-mmp.NodeQueue:
				xlog.Debugf("Received command on queue from %s", payload.SrcID)
				mmpProcessor(context.TODO(), payload)
			case <-close:
				xlog.Debug("Closing mmp processor")
				return
			}
		}
	}()

	<-w.QuitChan

	close <- struct{}{}

	w.WG.Done()
	w.Running = false
	xlog.Infof("Stopped worker %s", w.Name)
}

func mmpProcessor(ctx context.Context, p *mmsp.Payload) error {
	var err error

	switch p.PayloadType {
	case mmsp.PayloadType_WORKFLOW_EXPEDITE:
		err = ops.WorkflowExpedite(ctx, p)
	case mmsp.PayloadType_WORKFLOW_SCHEDULE:
		err = ops.WorkflowSchedule(ctx, p)
	}

	if err != nil {
		logging.Debugf("Unable to process mmp workflow payload: %v", err)
		return err
	}

	return nil
}
