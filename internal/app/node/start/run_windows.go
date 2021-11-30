//go:build windows
// +build windows

package start

import (
	"log"
	"os"

	"github.com/kardianos/service"
	"github.com/spf13/viper"
	"mmesh.dev/m-lib/pkg/version"
	"mmesh.dev/m-lib/pkg/xlog"
)

type windowsAction int

const (
	windowsActionConsoleRun windowsAction = iota
	windowsActionServiceStart
	windowsActionServiceInstall
	windowsActionServiceUninstall
)

type program struct{}

func (p *program) Start(s service.Service) error {
	go p.run()

	return nil
}

func (p *program) run() {
	start()
}

func (p *program) Stop(s service.Service) error {
	go finish()

	return nil
}

func runAsWindowsService(action windowsAction) {
	svcConfig := &service.Config{
		Name:        version.NODE_NAME,
		DisplayName: version.NODE_NAME,
		Description: "mmesh-node",
		Arguments:   []string{"service-start"},
		Option: service.KeyValue{
			"OnFailure": "restart",
		},
	}

	prg := &program{}

	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}

	logger, err := s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}

	switch action {
	case windowsActionConsoleRun:
		err = s.Run()
	case windowsActionServiceStart:
		err = s.Run()
	case windowsActionServiceInstall:
		err = s.Install()
	case windowsActionServiceUninstall:
		err = s.Uninstall()
	}
	if err != nil {
		logger.Error(err)
	}
}

func Main() {
	xlog.Infof("%s starting on %s :-)", version.NODE_NAME, viper.GetString("host.id"))
	runAsWindowsService(windowsActionConsoleRun)
	<-done

	xlog.Infof("%s stopped on %s", version.NODE_NAME, viper.GetString("host.id"))
	os.Exit(0)
}

func ServiceStart() {
	xlog.Infof("Starting %s Windows Service", version.NODE_NAME)
	runAsWindowsService(windowsActionServiceStart)
	os.Exit(0)
}

func ServiceInstall() {
	xlog.Infof("Installing %s as Windows Service", version.NODE_NAME)
	runAsWindowsService(windowsActionServiceInstall)
	os.Exit(0)
}

func ServiceUninstall() {
	xlog.Infof("Uninstalling %s Windows Service", version.NODE_NAME)
	runAsWindowsService(windowsActionServiceUninstall)
	os.Exit(0)
}
