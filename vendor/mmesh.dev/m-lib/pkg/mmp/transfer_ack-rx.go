package mmp

import (
	"context"
	"fmt"
	"os"

	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-lib/pkg/cli/output"
	"mmesh.dev/m-lib/pkg/logging"
	"mmesh.dev/m-lib/pkg/utils/colors"
	"x6a.dev/pkg/errors"
)

func transferAckRx(ctx context.Context, payload *mmsp.Payload) error {
	errStr := payload.Transfer.FileAck.Err
	recvBytes := payload.Transfer.FileAck.RecvBytes

	if len(errStr) > 0 {
		if payload.Interactive {
			logging.Errorf("%s", errStr)
			output.Disconnected()
			os.Exit(1)
		}
		return errors.Errorf("Unable to complete file transfer (%v bytes transferred): %s", recvBytes, errStr)
	}

	fmt.Println("\nTransfer completed: " + colors.DarkWhite(fmt.Sprintf("%v bytes", recvBytes)) + ".")

	if payload.Interactive {
		output.Disconnected()

		os.Exit(0)
	}

	return nil
}
