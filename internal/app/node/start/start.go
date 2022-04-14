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
	"mmesh.dev/m-node/internal/app/node/connection"
	"mmesh.dev/m-node/internal/app/node/netp2p"
)

var done = make(chan struct{})

func start() {
	nxnc := connection.AgentConnect()
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
	netp2p.Disconnect()
	if err := connection.GRPCClientConn.Close(); err != nil {
		xlog.Errorf("Unable to close gRPC network connection: %v", err)
	}

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
	netp2p.Disconnect()
	if err := connection.GRPCClientConn.Close(); err != nil {
		xlog.Errorf("Unable to close gRPC network connection: %v", err)
	}

	update.RestartReady <- struct{}{}
}
