package command

import (
	"context"
	"os/exec"

	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-lib/pkg/mmp/stream/protos/command/term"
	"mmesh.dev/m-lib/pkg/mmp/stream/utils/auth"
	"mmesh.dev/m-lib/pkg/mmp/stream/utils/streaming"
	"mmesh.dev/m-lib/pkg/mmp/queuing"
	"mmesh.dev/m-lib/pkg/xlog"
)

var shs = streaming.NewIOSessions()

func shellExec(ctx context.Context, p *mmsp.Payload) {
	if !auth.AgentMgmtAuth(p) {
		shellExecDisabled(p)
		return
	}

	// switch p.CommandPDU.Command.Type {
	// case command.CommandType_SHELL:
	// 	xlog.Debugf("Received shell exec request from %s", p.SrcID)
	// 	runShell(p)
	// }

	runShell(p)
}

func shellExecDisabled(payload *mmsp.Payload) {
	mmID := viper.GetString("mm.id")

	p := &mmsp.Payload{
		SrcID: mmID,
		DstID: payload.SrcID,
		Type: mmsp.PDUType_COMMAND,
		CommandPDU: &mmsp.CommandPDU{
			Type: mmsp.CommandMsgType_COMMAND_DISABLED,
		},
	}

	// p := &mmsp.Payload{
	// 	SrcID:       mmID,
	// 	DstID:       payload.SrcID,
	// 	PayloadType: mmsp.PayloadType_COMMAND_SHELL_DISABLED,
	// }

	queuing.TxControlQueue <- p
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

	cReq := payload.CommandPDU.Command.Request

	xlog.Debugf("Received request to execute command: %v", cReq.Exec.Cmd)

	// Logic to execute the command.
	c := exec.Command(cReq.Exec.Cmd, cReq.Exec.Args...)

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
			xlog.Errorf("Unable to run shell %s: %v", cReq.Exec.Cmd, err)
		}
		waitc <- struct{}{}
	}()

	<-waitc

	// Shut down shell session
	xlog.Trace("Shell terminated")
	shs.DeleteIOSession(sID)

	// Finish cli session
	p := NewShellExit(mmID, payload)
	queuing.TxControlQueue <- p
}

func NewShellExit(srcID string, p *mmsp.Payload) *mmsp.Payload {
	return &mmsp.Payload{
		SrcID: srcID,
		DstID: p.SrcID,
		Type: mmsp.PDUType_COMMAND,
		CommandPDU: &mmsp.CommandPDU{
			Type: mmsp.CommandMsgType_COMMAND_EXIT,
		},
	}

	// return &mmsp.Payload{
	// 	SrcID:       srcID,
	// 	DstID:       p.SrcID,
	// 	PayloadType: mmsp.PayloadType_COMMAND_SHELL_EXIT,
	// }
}
