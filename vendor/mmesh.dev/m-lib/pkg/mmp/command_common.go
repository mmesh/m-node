package mmp

import (
	"mmesh.dev/m-api-go/grpc/network/mmsp/command"
)

const (
	commandTypeAction = iota
	commandTypeShell
)

const (
//commandStatusExecuted = "EXECUTED"
//commandStatusCanceled = "CANCELED"
//commandStatusRunning = "RUNNING"
//commandStatusFailed  = "FAILED"
)

func newCommand(name, c string, args ...string) *command.CommandExec {
	return &command.CommandExec{
		Cmd:  c,
		Args: args,
		UID:  0,
		GID:  0,
	}
}
