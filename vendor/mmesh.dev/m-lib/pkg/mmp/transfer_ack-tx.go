package mmp

import (
	"context"
	"time"

	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-api-go/grpc/network/mmsp/transfer"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/logging"
)

func transferAckTx(ctx context.Context, payload *mmsp.Payload, errStr string) {
	var recvBytes int64

	if len(errStr) == 0 {
		if payload.Transfer.FileChunk == nil {
			return
		}

		if !payload.Transfer.FileChunk.IsLastChunk {
			return
		}
	}

	fileInfo, err := getFileInfo(payload.Transfer.Paths.DstFilePath)
	if err != nil {
		logging.Debugf("getFileInfo: %v", errors.Cause(err))
	}
	if fileInfo != nil {
		recvBytes = fileInfo.Size()
	}

	p := newFileAck(payload, recvBytes, errStr)
	TxControlQueue <- p
}

func newFileAck(p *mmsp.Payload, recvBytes int64, errStr string) *mmsp.Payload {
	mmID := viper.GetString("mm.id")

	return &mmsp.Payload{
		SrcID:       mmID,
		DstID:       p.RequesterID,
		Interactive: p.Interactive,
		PayloadType: mmsp.PayloadType_TRANSFER_ACK,
		Transfer: &transfer.Transfer{
			Paths: &transfer.Paths{
				// SrcFilePath: p.Transfer.Paths.SrcFilePath,
				DstFilePath: p.Transfer.Paths.DstFilePath,
			},
			FileChunk: nil,
			FileAck: &transfer.FileAck{
				RecvBytes: recvBytes,
				Timestamp: time.Now().Unix(),
				Err:       errStr,
			},
		},
	}
}
