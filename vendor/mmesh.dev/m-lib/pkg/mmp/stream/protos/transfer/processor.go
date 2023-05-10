package transfer

import (
	"context"
	"fmt"

	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/logging"
)

func Processor(ctx context.Context, p *mmsp.Payload) error {
	if p.TransferPDU == nil {
		return fmt.Errorf("null pdu")
	}

	var err error

	switch p.TransferPDU.Type {
	case mmsp.TransferMsgType_TRANSFER_INIT:
		err = transferInit(ctx, p)
	case mmsp.TransferMsgType_TRANSFER_DATA:
		err = transferDataRx(ctx, p)
	case mmsp.TransferMsgType_TRANSFER_ACK:
		err = transferAckRx(ctx, p)
	case mmsp.TransferMsgType_TRANSFER_DISABLED:
		transferUnavailable()
	}

	if err != nil {
		logging.Errorf("Unable to process mmp transferPDU payload: %v", errors.Cause(err))
		return err
	}

	return nil
}
