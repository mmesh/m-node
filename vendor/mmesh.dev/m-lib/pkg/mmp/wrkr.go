package mmp

import (
	"context"

	"mmesh.dev/m-lib/pkg/mmp/queuing"
	"mmesh.dev/m-lib/pkg/runtime"
	"mmesh.dev/m-lib/pkg/xlog"
)

func Dispatcher(w *runtime.Wrkr) {
	xlog.Infof("Started worker %s", w.Name)
	w.Running = true

	close := make(chan struct{}, 1)

	go func() {
		for {
			select {
			case payload := <-queuing.RxControlQueue:
				Processor(context.TODO(), payload)
			case <-close:
				xlog.Debug("Closing mmp dispacher")
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
