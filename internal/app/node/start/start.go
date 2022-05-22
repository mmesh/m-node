package start

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"mmesh.dev/m-lib/pkg/runtime"
	"mmesh.dev/m-lib/pkg/update"
	"mmesh.dev/m-lib/pkg/xlog"
	"mmesh.dev/m-node/internal/app/node/mnet"
)

var done = make(chan struct{})

func start() {
	if err := mnet.Init(); err != nil {
		xlog.Alertf("Unable to initialize node: %s", err)
		os.Exit(1)
	}
	nxnc := mnet.LocalNode().Connection().NetworkClient()
	initWrkrs(nxnc)
	runtime.StartWrkrs()

	go restart()

	cleanup()
}

func cleanup() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		finish()
	}()
}

func finish() {
	xlog.Debug("Cleaning and finishing...")
	var wg sync.WaitGroup

	xlog.Debug("Closing workers...")
	wg.Add(1)
	runtime.StopWrkrs(&wg)
	wg.Wait()

	xlog.Debug("Closing agent connection handlers...")
	mnet.LocalNode().Close()

	time.Sleep(time.Second)

	close(done)
}

func restart() {
	<-update.RestartRequest

	var wg sync.WaitGroup

	xlog.Debug("Closing workers...")
	wg.Add(1)
	runtime.StopWrkrs(&wg)
	wg.Wait()

	xlog.Debug("Closing agent connection handlers...")
	mnet.LocalNode().Close()

	update.RestartReady <- struct{}{}
}
