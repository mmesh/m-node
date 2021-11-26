//go:build darwin
// +build darwin

package start

import (
	"log"
	"os"
	"os/exec"

	"github.com/kardianos/service"
	"github.com/spf13/viper"
	"mmesh.dev/m-lib/pkg/version"
	"x6a.dev/pkg/xlog"
)

type serviceAction int

const (
	actionConsoleRun serviceAction = iota
	actionServiceStart
	actionServiceInstall
	actionServiceUninstall
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

func runAsService(action serviceAction) {
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
	case actionConsoleRun:
		err = s.Run()
	case actionServiceStart:
		err = s.Run()
	case actionServiceInstall:
		err = s.Install()
	case actionServiceUninstall:
		err = s.Uninstall()
	}
	if err != nil {
		logger.Error(err)
	}
}

func Main() {
	xlog.Infof("%s starting on %s :-)", version.NODE_NAME, viper.GetString("host.id"))
	runAsService(actionConsoleRun)
	<-done

	xlog.Infof("%s stopped on %s", version.NODE_NAME, viper.GetString("host.id"))
	os.Exit(0)
}

func ServiceStart() {
	xlog.Infof("Starting %s Service", version.NODE_NAME)
	runAsService(actionServiceStart)
	os.Exit(0)
}

func ServiceInstall() {
	xlog.Infof("Installing %s as Service", version.NODE_NAME)
	runAsService(actionServiceInstall)

	cmd := exec.Command("launchctl", "load", "/Library/LaunchDaemons/mmesh-node.plist")
	if err := cmd.Run(); err != nil {
		xlog.Warnf("Unable to load launchctl mmesh-node service, please check: %v", err)
	}

	os.Exit(0)
}

func ServiceUninstall() {
	xlog.Infof("Uninstalling %s Service", version.NODE_NAME)
	runAsService(actionServiceUninstall)

	// cmd := exec.Command("launchctl", "unload", "/Library/LaunchDaemons/mmesh-node.plist")
	// if err := cmd.Run(); err != nil {
	// 	xlog.Warnf("Unable to unload launchctl mmesh-node service, please check: %v", err)
	// }

	os.Exit(0)
}