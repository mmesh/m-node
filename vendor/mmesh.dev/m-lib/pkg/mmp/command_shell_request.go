package mmp

import (
	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-api-go/grpc/network/mmsp/command"
	"mmesh.dev/m-api-go/grpc/network/resources/iam/auth"
)

func NewShellRequest(authKey *auth.AuthKey, dstID string, c string, args ...string) {
	mmID := viper.GetString("mm.id")

	p := &mmsp.Payload{
		SrcID:         mmID,
		DstID:         dstID,
		AuthKey:       authKey,
		PSK:           viper.GetString("agent.management.auth.psk"),
		SecurityToken: viper.GetString("agent.management.auth.securityToken"),
		PayloadType:   PayloadTypeCommandShellExec,
		Command: &command.Command{
			CommandType: commandTypeShell,
			CommandRequest: &command.CommandRequest{
				Command:  newCommand(c, c, args...),
				Schedule: nil,
				Stdin:    nil,
			},
		},
	}

	SendCommandQueue <- p
}
