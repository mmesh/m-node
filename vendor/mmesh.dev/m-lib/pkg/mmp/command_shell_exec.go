package mmp

import (
	"context"
	"os/exec"

	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-lib/pkg/mmp/streaming"
	"mmesh.dev/m-lib/pkg/mmp/term"
	"mmesh.dev/m-lib/pkg/xlog"
)

var shs = streaming.NewIOSessions()

func shellExec(ctx context.Context, p *mmsp.Payload) {
	if !mgmtAuth(p) {
		shellExecDisabled(p)
		return
	}

	switch p.Command.CommandType {
	case commandTypeShell:
		xlog.Debugf("Received shell exec request from %s", p.SrcID)
		runShell(p)
	}
}

func runShell(payload *mmsp.Payload) {
	waitc := make(chan struct{})

	mmID := viper.GetString("mm.id")
	sID := payload.SrcID

	shs.SetIOSession(sID)
	iop := shs.GetIOSessionIO(sID)
	if iop == nil {
		xlog.Errorf("shell io pipes not found for session from %s", sID)
		return
	}
	if iop.Out.RP == nil {
		xlog.Errorf("shell output writer pipe not found for session from %s", sID)
		return
	}
	if iop.In.RP == nil {
		xlog.Errorf("shell input writer pipe not found for session from %s", sID)
		return
	}

	cReq := payload.Command.CommandRequest

	xlog.Debugf("Received request to execute command: %v", cReq.Command.Cmd)

	//Logic to execute the command.
	c := exec.Command(cReq.Command.Cmd, cReq.Command.Args...)

	// TODO: remote user/perms
	// c.SysProcAttr = &syscall.SysProcAttr{
	// 	Credential: &syscall.Credential{
	// 		Uid: uint32(65534),
	// 		Gid: uint32(65534),
	// 	},
	// }

	go func() {
		shellWriteOutput(mmID, payload, iop.Out.RP)
		waitc <- struct{}{}
	}()

	xlog.Trace("Shell starting")

	go func() {
		err := term.PTYRun(c, iop.In.RP, iop.Out.WP)
		if nil != err {
			xlog.Errorf("Unable to run shell %s: %v", cReq.Command.Cmd, err)
		}
		waitc <- struct{}{}
	}()

	<-waitc

	// Shut down shell session
	xlog.Trace("Shell terminated")
	shs.DeleteIOSession(sID)

	// Finish cli session
	p := NewShellExit(mmID, payload)
	TxControlQueue <- p
}
