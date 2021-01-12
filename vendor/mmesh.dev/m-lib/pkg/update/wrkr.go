package update

import (
	"time"

	"github.com/spf13/viper"
	"mmesh.dev/m-lib/pkg/runtime"
	"x6a.dev/pkg/xlog"
)

var checkUpdate = make(chan struct{}, 1)

func UpdateAgent(w *runtime.Wrkr) {
	xlog.Infof("Started worker %s", w.Name)
	w.Running = true

	quitUpdate := make(chan struct{}, 1)

	go func() {
		appName := viper.GetString("mm.app")
		if len(appName) == 0 {
			return
		}

		for {
			select {
			case <-checkUpdate:
				xlog.Debug("Checking for updates...")
				if err := Update(appName); err != nil {
					xlog.Errorf("Unable to check for updates: %v", err)
				}
			case <-quitUpdate:
				xlog.Debug("Closing update agent")
				return
			}
		}
	}()

	go updateAgentCtl()

	<-w.QuitChan
	quitUpdate <- struct{}{}

	w.WG.Done()
	w.Running = false
	xlog.Infof("Stopped worker %s", w.Name)
}

var updateAgentCtlRun bool

func updateAgentCtl() {
	if !updateAgentCtlRun {
		updateAgentCtlRun = true

		if !viper.GetBool("maintenance.autoUpdate") {
			xlog.Info("Auto-update disabled")
			return
		}

		mHour := viper.GetInt("maintenance.schedule.hour")
		mMin := viper.GetInt("maintenance.schedule.minute")

		xlog.Infof("Auto-update enabled, scheduled at %02d:%02d", mHour, mMin)

		go func() {
			for {
				tm := time.Now()
				h := tm.Hour()
				m := tm.Minute()

				if mHour == h && mMin == m {
					checkUpdate <- struct{}{}
				}

				time.Sleep(time.Minute)
			}
		}()
	}
}
