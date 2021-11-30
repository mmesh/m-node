//go:build windows
// +build windows

package update

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"mmesh.dev/m-lib/pkg/logging"
	"mmesh.dev/m-lib/pkg/version"
	"mmesh.dev/m-lib/pkg/xlog"
)

func restartProcess(app, exe string) {
	if app == version.CLI_NAME {
		fmt.Printf("Binary updated to latest version :-)\n\n")
		os.Exit(0)
	}

	RestartRequest <- struct{}{}

	xlog.Info("Binary updated to latest version, spawning a new process...")

	<-RestartReady

	if !strings.Contains(strings.Join(os.Args, " "), "service-start") {
		c := []string{"cmd.exe", "/C", "start", "mmesh-node service", exe}
		c = append(c, os.Args[1:]...)

		xlog.Infof(" -> %s", strings.Join(c, " "))

		if err := exec.Command(c[0], c[1:]...).Start(); err != nil {
			logging.Alertf("Unable to restart main process: %v", err)
			return
		}
	}

	os.Exit(2)
}
