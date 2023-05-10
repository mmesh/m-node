package transfer

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-api-go/grpc/network/mmsp/stream"
	"mmesh.dev/m-api-go/grpc/resources/iam/auth"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/logging"
	mmp_auth "mmesh.dev/m-lib/pkg/mmp/stream/utils/auth"
	"mmesh.dev/m-lib/pkg/mmp/queuing"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

const fileChunk_size = 1048576

func NewTransfer(authKey *auth.AuthKey, srcFilePath, dstFilePath, srcID, dstID string, interactive bool) {
	logging.Debugf("New file transfer requested: srcID %s, dstID %s", srcID, dstID)

	p := newTransferInit(authKey, srcFilePath, dstFilePath, srcID, dstID, interactive)
	queuing.TxControlQueue <- p
}

func newTransferInit(authKey *auth.AuthKey, srcFilePath, dstFilePath, srcID, dstID string, interactive bool) *mmsp.Payload {
	// data transfer must to be initiated by srcID,
	// so srcID becomes also dstID for grpc routing to work properly;
	// srcNodeID and dstNodeID are used to preserve the right transfer flow
	mmID := viper.GetString("mm.id")

	return &mmsp.Payload{
		SrcID: mmID,
		DstID: srcID,
		Type:  mmsp.PDUType_TRANSFER,
		TransferPDU: &mmsp.TransferPDU{
			Metadata: &mmsp.StreamMetadata{
				RequesterID:   mmID,
				Interactive:   interactive,
				AuthKey:       authKey,
				PSK:           viper.GetString("management.auth.psk"),
				SecurityToken: viper.GetString("management.auth.securityToken"),
			},
			Type: mmsp.TransferMsgType_TRANSFER_INIT,
			Transfer: &stream.Transfer{
				Paths: &stream.Paths{
					SrcNodeID:   srcID,
					SrcFilePath: srcFilePath,
					DstNodeID:   dstID,
					DstFilePath: dstFilePath,
				},
				FileChunk: nil,
				FileAck:   nil,
			},
		},
	}

	// return &mmsp.Payload{
	// 	SrcID:         mmID,
	// 	DstID:         srcID,
	// 	RequesterID:   mmID,
	// 	Interactive:   interactive,
	// 	AuthKey:       authKey,
	// 	PSK:           viper.GetString("management.auth.psk"),
	// 	SecurityToken: viper.GetString("management.auth.securityToken"),
	// 	PayloadType:   mmsp.PayloadType_TRANSFER_INIT,
	// 	Transfer: &transfer.Transfer{
	// 		Paths: &transfer.Paths{
	// 			SrcNodeID:   srcID,
	// 			SrcFilePath: srcFilePath,
	// 			DstNodeID:   dstID,
	// 			DstFilePath: dstFilePath,
	// 		},
	// 	},
	// }
}

func transferInit(ctx context.Context, payload *mmsp.Payload) error {
	if !mmp_auth.AgentMgmtAuth(payload) {
		transferDisabled(payload)
		return nil
	}

	srcFilePath := payload.TransferPDU.Transfer.Paths.SrcFilePath
	dstFilePath := payload.TransferPDU.Transfer.Paths.DstFilePath
	dstID := payload.TransferPDU.Transfer.Paths.DstNodeID
	requesterID := payload.TransferPDU.Metadata.RequesterID
	interactive := payload.TransferPDU.Metadata.Interactive

	logging.Debugf("Starting file transfer to dstID %s", dstID)

	srcFiles, err := filepath.Glob(srcFilePath)
	if err != nil {
		transferAckTx(ctx, payload, err.Error())
		return errors.Wrapf(err, "[%v] function filepath.Glob()", errors.Trace())
	}

	if len(srcFiles) == 0 {
		transferAckTx(ctx, payload, "File not found")
		return fmt.Errorf("transferInit(): file not found")
	}

	for _, srcFile := range srcFiles {
		if err := transferSrcFile(srcFile, dstFilePath, dstID, requesterID, interactive); err != nil {
			logging.Trace(err)
			logging.Warn(errors.Cause(err))
			continue
		}
	}

	return nil
}

func transferSrcFile(srcFile, dstFilePath, dstID, requesterID string, interactive bool) error {
	var srcIsDir bool

	fd, err := os.Open(srcFile)
	if err != nil {
		return errors.Wrapf(err, "[%v] function os.Open()", errors.Trace())
	}
	defer fd.Close()

	fileInfo, err := fd.Stat()
	if err != nil {
		return errors.Wrapf(err, "[%v] function fd.Stat()", errors.Trace())
	}

	if fileInfo != nil {
		srcFileName := fileInfo.Name()
		srcFileMode := uint32(fileInfo.Mode())

		if fileInfo.IsDir() {
			dstFilePath := filepath.Join(dstFilePath, srcFileName)

			srcIsDir = true

			fType := fmt.Sprintf("%s%s%s", colors.Blue("["), colors.DarkBlue("dir"), colors.Blue("]"))
			logging.Println(colors.DarkCyan(srcFile), fmt.Sprintf("%s %s %s", colors.Cyan("->"), colors.DarkWhite(dstFilePath), fType))

			if err := transferDataTx(srcFileMode, srcIsDir, srcFile, srcFileName, dstFilePath, dstID, requesterID, interactive); err != nil {
				return errors.Wrapf(err, "[%v] function transferDataTx()", errors.Trace())
			}

			dirContents, err := fd.Readdir(0)
			if err != nil {
				return errors.Wrapf(err, "[%v] function fd.Readdir()", errors.Trace())
			}
			for _, fInfo := range dirContents {
				srcFile := filepath.Join(srcFile, fInfo.Name())

				if err := transferSrcFile(srcFile, dstFilePath, dstID, requesterID, interactive); err != nil {
					logging.Error(err)
					continue
				}
			}
		} else {
			//dstFilePath := filepath.Join(dstFilePath, srcFileName)

			srcIsDir = false

			dstFile := filepath.Join(dstFilePath, srcFileName)
			fType := fmt.Sprintf("%s%s%s", colors.Yellow("["), colors.DarkYellow("file"), colors.Yellow("]"))
			logging.Println(colors.DarkCyan(srcFile), fmt.Sprintf("%s %s %s", colors.Cyan("->"), colors.DarkWhite(dstFile), fType))

			if err := transferDataTx(srcFileMode, srcIsDir, srcFile, srcFileName, dstFilePath, dstID, requesterID, interactive); err != nil {
				return errors.Wrapf(err, "[%v] function transferDataTx()", errors.Trace())
			}
		}
	}

	return nil
}
