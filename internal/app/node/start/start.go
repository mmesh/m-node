package start

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/spf13/viper"
	"mmesh.dev/m-lib/pkg/runtime"
	"mmesh.dev/m-lib/pkg/version"
	"mmesh.dev/m-node/internal/app/node/connection"
	"mmesh.dev/m-node/internal/app/node/netp2p"
	"x6a.dev/pkg/xlog"
)

var done = make(chan struct{})

func start() {
	nxnc := connection.AgentConnect()
	initWrkrs(nxnc)
	runtime.StartWrkrs()

	cleanup()
}

func cleanup() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c

		xlog.Debug("Cleaning and finishing...")
		var wg sync.WaitGroup

		xlog.Debug("Closing workers...")
		wg.Add(1)
		runtime.StopWrkrs(&wg)

		wg.Wait()

		xlog.Debug("Closing agents connection handlers...")
		netp2p.Disconnect()

		time.Sleep(time.Second)

		close(done)
	}()
}

func Main() {
	start()
	xlog.Infof("%s started on %s :-)", version.NODE_NAME, viper.GetString("host.id"))
	<-done

	xlog.Infof("%s stopped on %s", version.NODE_NAME, viper.GetString("host.id"))
	os.Exit(0)
}
