package transfer

import (
	"context"
	"fmt"
	"os"

	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/logging"
	"mmesh.dev/m-lib/pkg/mmp/stream/utils/cli"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

func transferAckRx(ctx context.Context, p *mmsp.Payload) error {
	errStr := p.TransferPDU.Transfer.FileAck.Err
	recvBytes := p.TransferPDU.Transfer.FileAck.RecvBytes

	if len(errStr) > 0 {
		if p.TransferPDU.Metadata.Interactive {
			logging.Errorf("%s", errStr)
			cli.Disconnected()
			os.Exit(1)
		}
		return errors.Errorf("Unable to complete file transfer (%v bytes transferred): %s", recvBytes, errStr)
	}

	fmt.Println("\nTransfer completed: " + colors.DarkWhite(fmt.Sprintf("%v bytes", recvBytes)) + ".")

	if p.TransferPDU.Metadata.Interactive {
		cli.Disconnected()

		os.Exit(0)
	}

	return nil
}
