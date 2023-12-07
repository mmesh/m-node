//go:build linux
// +build linux

package start

import (
	"os"

	"github.com/spf13/viper"
	"mmesh.dev/m-lib/pkg/version"
	"mmesh.dev/m-lib/pkg/xlog"
)

func Main() {
	start()
	xlog.Infof("%s started on %s :-)", version.NODE_NAME, viper.GetString("host.id"))
	defer xlog.Logger().Close()
	<-done

	xlog.Infof("%s stopped on %s", version.NODE_NAME, viper.GetString("host.id"))

	os.Exit(0)
}
