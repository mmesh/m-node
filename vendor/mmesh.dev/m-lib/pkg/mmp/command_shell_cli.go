package mmp

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh/terminal"
	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-api-go/grpc/network/mmsp/command"
	"mmesh.dev/m-lib/pkg/cli/output"
	"mmesh.dev/m-lib/pkg/mmp/term"
	"x6a.dev/pkg/msg"
)

var cliStdinWaitc = make(chan struct{}, 1)

var cliStdinActive bool

func newShellInput(srcID string, p *mmsp.Payload) *mmsp.Payload {
	cReq := p.Command.CommandResponse.RequestedCommand

	return &mmsp.Payload{
		SrcID:       srcID,
		DstID:       p.SrcID,
		PayloadType: mmsp.PayloadType_COMMAND_SHELL_INPUT,
		Command: &command.Command{
			CommandType: p.Command.CommandType,
			CommandRequest: &command.CommandRequest{
				Command:  cReq.Command,
				Schedule: cReq.Schedule,
				Stdin:    nil,
			},
			CommandResponse: nil,
		},
	}
}

func NewShellExit(srcID string, p *mmsp.Payload) *mmsp.Payload {
	return &mmsp.Payload{
		SrcID:       srcID,
		DstID:       p.SrcID,
		PayloadType: mmsp.PayloadType_COMMAND_SHELL_EXIT,
	}
}

func shellReadOutput(ctx context.Context, payload *mmsp.Payload) error {
	if len(payload.Command.CommandResponse.Stdout) > 0 {
		fmt.Print(string(payload.Command.CommandResponse.Stdout))
	}

	go func() {
		if !cliStdinActive {
			shellWriteInput(ctx, payload)
		}
	}()

	return nil
}

func shellWriteInput(ctx context.Context, payload *mmsp.Payload) {
	cliStdinActive = true

	mmID := viper.GetString("mm.id")

	// msg.Debug("Reading standard input..")

	fd := int(os.Stdin.Fd()) // fd 0 is stdin

	state, err := terminal.MakeRaw(fd)
	if err != nil {
		msg.Errorf("Unable to set stdin to raw mode: %v", err)
		os.Exit(1)
	}
	// msg.Debug("Terminal in raw mode.")

	defer func() {
		if err := terminal.Restore(fd, state); err != nil {
			msg.Errorf("Unable to restore terminal: %v", err)
			os.Exit(1)
		}
		// msg.Debug("Terminal restored.")
	}()

	go func() {
		for {
			var input []byte

			p := newShellInput(mmID, payload)
			for len(input) == 0 {
				input = term.ReadKey()
			}
			p.Command.CommandRequest.Stdin = input
			TxControlQueue <- p
		}
	}()

	<-cliStdinWaitc

	// msg.Debug("Closing standard input..")
}

func shellExit(ctx context.Context, p *mmsp.Payload) error {
	cliStdinWaitc <- struct{}{}

	time.Sleep(100 * time.Millisecond)

	//endOfTransmission()
	output.Disconnected()

	cliStdinActive = false

	os.Exit(0)

	return nil
}
