package mmp

import (
	"context"
	"io"
	"os/exec"

	"github.com/creack/pty"
	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"x6a.dev/pkg/xlog"
)

var shs = newShellSession()

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

	shs.setShellSession(sID)
	iop := shs.getShellSessionIO(sID)
	if iop == nil {
		xlog.Errorf("shell io pipes not found for session from %s", sID)
		return
	}
	if iop.out.rp == nil {
		xlog.Errorf("shell output writer pipe not found for session from %s", sID)
		return
	}
	if iop.in.rp == nil {
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
		shellWriteOutput(mmID, payload, iop.out.rp)
		waitc <- struct{}{}
	}()

	xlog.Trace("Shell starting")

	go func() {
		err := ptyRun(c, iop.in.rp, iop.out.wp)
		if nil != err {
			xlog.Errorf("Unable to run shell %s: %v", cReq.Command.Cmd, err)
		}
		waitc <- struct{}{}
	}()

	<-waitc

	// Shut down shell session
	xlog.Trace("Shell terminated")
	shs.deleteShellSession(sID)

	// Finish cli session
	p := NewShellExit(mmID, payload)
	TxControlQueue <- p
}

func ptyRun(c *exec.Cmd, inrp *io.PipeReader, outwp *io.PipeWriter) error {
	// Start the command with a pty.
	ptmx, err := pty.Start(c)
	if err != nil {
		return err
	}
	// Make sure to close the pty at the end.
	defer func() { _ = ptmx.Close() }() // Best effort.

	// Handle pty size.
	// ch := make(chan os.Signal, 1)
	// signal.Notify(ch, syscall.SIGWINCH)
	// go func() {
	// 	for range ch {
	// 		if err := pty.InheritSize(os.Stdin, ptmx); err != nil {
	// 			log.Printf("error resizing pty: %s", err)
	// 		}
	// 	}
	// }()
	// ch <- syscall.SIGWINCH // Initial resize.

	// Copy stdin to the pty and the pty to stdout.
	go func() { _, _ = io.Copy(ptmx, inrp) }()
	_, _ = io.Copy(outwp, ptmx)

	return nil
}
