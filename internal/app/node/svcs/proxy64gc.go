package svcs

import (
	"time"

	"mmesh.dev/m-lib/pkg/runtime"
	"mmesh.dev/m-node/internal/app/node/netp2p"
	"x6a.dev/pkg/xlog"
)

var proxy64GCch = make(chan struct{}, 1)

func Proxy64GC(w *runtime.Wrkr) {
	xlog.Infof("Started worker %s", w.Name)
	w.Running = true

	endCh := make(chan struct{}, 1)

	go func() {
		for {
			select {
			case <-proxy64GCch:
				xlog.Info("Running mmesh64 garbage collector...")
				netp2p.Proxy64GC()
			case <-endCh:
				xlog.Debug("Closing mmesh64 garbage collector")
				return
			}
		}
	}()

	go proxy64GCCtl()

	<-w.QuitChan

	endCh <- struct{}{}

	w.WG.Done()
	w.Running = false
	xlog.Infof("Stopped worker %s", w.Name)
}

var proxy64GCCtlRun bool

func proxy64GCCtl() {
	if !proxy64GCCtlRun {
		proxy64GCCtlRun = true
		go func() {
			for {
				proxy64GCch <- struct{}{}
				time.Sleep(120 * time.Second)
			}
		}()
	}
}
