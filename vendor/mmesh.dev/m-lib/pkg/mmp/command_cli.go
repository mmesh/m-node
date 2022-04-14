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
	"mmesh.dev/m-api-go/grpc/resources/iam/auth"
	"mmesh.dev/m-lib/pkg/mmp/cli"
	"mmesh.dev/m-lib/pkg/mmp/term"
	"mmesh.dev/m-lib/pkg/utils/colors"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

var cliStdinWaitc = make(chan struct{}, 1)

var cliStdinActive bool

func NewShellRequest(authKey *auth.AuthKey, dstID string, c string, args ...string) {
	mmID := viper.GetString("mm.id")

	p := &mmsp.Payload{
		SrcID:         mmID,
		DstID:         dstID,
		AuthKey:       authKey,
		PSK:           viper.GetString("agent.management.auth.psk"),
		SecurityToken: viper.GetString("agent.management.auth.securityToken"),
		PayloadType:   mmsp.PayloadType_COMMAND_SHELL_EXEC,
		Command: &command.Command{
			CommandType: command.CommandType_SHELL,
			CommandRequest: &command.CommandRequest{
				Command: &command.CommandExec{
					Cmd:  c,
					Args: args,
					UID:  0,
					GID:  0,
				},
				Schedule: nil,
				Stdin:    nil,
			},
		},
	}

	TxControlQueue <- p
}

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
	cli.Disconnected()

	cliStdinActive = false

	os.Exit(0)

	return nil
}

func shellUnavailable() {
	text := "Shell exec unauthorized or disabled here"

	alert := fmt.Sprintf("%s%s%s", colors.Black("["), colors.Red(text), colors.Black("]"))
	fmt.Println(alert)

	cli.Disconnected()

	os.Exit(0)
}
