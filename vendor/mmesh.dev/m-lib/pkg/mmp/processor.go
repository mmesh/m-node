package mmp

import (
	"context"

	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/logging"
	"mmesh.dev/m-lib/pkg/mmp/stream/protos/command"
	"mmesh.dev/m-lib/pkg/mmp/stream/protos/portfwd"
	"mmesh.dev/m-lib/pkg/mmp/stream/protos/transfer"
)

func Processor(ctx context.Context, p *mmsp.Payload) error {
	var err error

	switch p.Type {
	case mmsp.PDUType_COMMAND:
		err = command.Processor(ctx, p)
	case mmsp.PDUType_TRANSFER:
		err = transfer.Processor(ctx, p)
	case mmsp.PDUType_PORTFWD:
		err = portfwd.Processor(ctx, p)
	}

	if err != nil {
		logging.Debugf("Unable to process mmp payload: %v", errors.Cause(err))
		return err
	}

	return nil
}
