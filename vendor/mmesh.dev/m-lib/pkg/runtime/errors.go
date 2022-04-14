package runtime

import (
	"mmesh.dev/m-lib/pkg/xlog"
)

var ErrorEventsJobQueue = make(chan error, 128)

func ErrorEventsHandler(w *Wrkr) {
	xlog.Infof("Started worker %s", w.Name)
	w.Running = true

	for {
		select {
		case err := <-ErrorEventsJobQueue:
			xlog.Error(err)
		case <-w.QuitChan:
			w.WG.Done()
			w.Running = false
			xlog.Infof("Stopped worker %s", w.Name)
			return
		}
	}
}
