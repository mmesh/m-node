package transfer

import (
	"context"
	"time"

	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-api-go/grpc/network/mmsp/stream"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/logging"
	"mmesh.dev/m-lib/pkg/mmp/queuing"
)

func transferAckTx(ctx context.Context, payload *mmsp.Payload, errStr string) {
	var recvBytes int64

	if len(errStr) == 0 {
		if payload.TransferPDU.Transfer.FileChunk == nil {
			return
		}

		if !payload.TransferPDU.Transfer.FileChunk.IsLastChunk {
			return
		}
	}

	fileInfo, err := getFileInfo(payload.TransferPDU.Transfer.Paths.DstFilePath)
	if err != nil {
		logging.Debugf("getFileInfo: %v", errors.Cause(err))
	}
	if fileInfo != nil {
		recvBytes = fileInfo.Size()
	}

	p := newFileAck(payload, recvBytes, errStr)
	queuing.TxControlQueue <- p
}

func newFileAck(p *mmsp.Payload, recvBytes int64, errStr string) *mmsp.Payload {
	mmID := viper.GetString("mm.id")

	return &mmsp.Payload{
		SrcID:       mmID,
		DstID:       p.TransferPDU.Metadata.RequesterID,
		Type: mmsp.PDUType_TRANSFER,
		TransferPDU: &mmsp.TransferPDU{
			Metadata: &mmsp.StreamMetadata{
				Interactive: p.TransferPDU.Metadata.Interactive,
			},
			Type: mmsp.TransferMsgType_TRANSFER_ACK,
			Transfer: &stream.Transfer{
				Paths: &stream.Paths{
					// SrcFilePath: p.TransferPDU.Transfer.Paths.SrcFilePath,
					DstFilePath: p.TransferPDU.Transfer.Paths.DstFilePath,
				},
				FileChunk: nil,
				FileAck: &stream.FileAck{
					RecvBytes: recvBytes,
					Timestamp: time.Now().Unix(),
					Err:       errStr,
				},
			},
		},
	}

	// return &mmsp.Payload{
	// 	SrcID:       mmID,
	// 	DstID:       p.RequesterID,
	// 	Interactive: p.Interactive,
	// 	PayloadType: mmsp.PayloadType_TRANSFER_ACK,
	// 	Transfer: &transfer.Transfer{
	// 		Paths: &transfer.Paths{
	// 			// SrcFilePath: p.Transfer.Paths.SrcFilePath,
	// 			DstFilePath: p.Transfer.Paths.DstFilePath,
	// 		},
	// 		FileChunk: nil,
	// 		FileAck: &transfer.FileAck{
	// 			RecvBytes: recvBytes,
	// 			Timestamp: time.Now().Unix(),
	// 			Err:       errStr,
	// 		},
	// 	},
	// }
}
