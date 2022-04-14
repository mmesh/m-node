package svcs

import (
	"github.com/spf13/viper"
	"mmesh.dev/m-lib/pkg/runtime"
	"mmesh.dev/m-lib/pkg/xlog"
	"mmesh.dev/m-node/internal/app/node/k8s"
)

func KubernetesConnector(w *runtime.Wrkr) {
	xlog.Infof("Started worker %s", w.Name)
	w.Running = true

	quitKCtrlCh := make(chan struct{})

	go func() {
		if !viper.GetBool("agent.kubernetes.controller.enabled") {
			xlog.Info("Kubernetes controller not enabled")
			return
		}

		if err := k8s.KubernetesController(quitKCtrlCh); err != nil {
			xlog.Tracef("Unable to start kubernetes controller: %v", err)
			xlog.Info("Kubernetes environment not detected")
		} else {
			xlog.Info("Kubernetes controller running")
		}
	}()

	<-w.QuitChan
	close(quitKCtrlCh)

	w.WG.Done()
	w.Running = false
	xlog.Infof("Stopped worker %s", w.Name)
}
