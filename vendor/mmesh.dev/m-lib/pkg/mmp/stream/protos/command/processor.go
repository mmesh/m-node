package command

import (
	"context"
	"fmt"

	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/logging"
)

func Processor(ctx context.Context, p *mmsp.Payload) error {
	if p.CommandPDU == nil {
		return fmt.Errorf("null pdu")
	}

	var err error

	switch p.CommandPDU.Type {
	case mmsp.CommandMsgType_COMMAND_EXEC:
		go shellExec(ctx, p)
	case mmsp.CommandMsgType_COMMAND_EXIT:
		err = shellExit(ctx, p)
	case mmsp.CommandMsgType_COMMAND_INPUT:
		err = shellReadInput(ctx, p)
	case mmsp.CommandMsgType_COMMAND_OUTPUT:
		err = shellReadOutput(ctx, p)
	case mmsp.CommandMsgType_COMMAND_DISABLED:
		shellUnavailable()
	}

	if err != nil {
		logging.Errorf("Unable to process mmp commandPDU payload: %v", errors.Cause(err))
		return err
	}

	return nil
}
