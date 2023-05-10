package transfer

import (
	"io"
	"os"
	"time"

	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-api-go/grpc/network/mmsp/stream"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/logging"
	"mmesh.dev/m-lib/pkg/mmp/queuing"
	"mmesh.dev/m-lib/pkg/utils"
)

func transferDataTx(srcFileMode uint32, srcIsDir bool, srcFile, srcFileName, dstFilePath, dstID, requesterID string, interactive bool) error {
	mmID := viper.GetString("mm.id")

	if srcIsDir {
		p := newDir(srcFileMode, srcFile, srcFileName, dstFilePath, mmID, dstID, requesterID, interactive)

		queuing.TxControlQueue <- p
		return nil
	}

	var chunkNumber int64 = 0

	fileChecksum, err := utils.ChecksumSHA256(srcFile)
	if err != nil {
		return errors.Wrapf(err, "[%v] function utils.ChecksumSHA256()", errors.Trace())
	}
	logging.Tracef("File checksum: %x", fileChecksum)

	fd, err := os.Open(srcFile)
	if err != nil {
		return errors.Wrapf(err, "[%v] function os.Open()", errors.Trace())
	}
	defer fd.Close()

	transferring := true
	isLastChunk := false

	for transferring {
		buf := make([]byte, fileChunk_size)

		n, err := fd.Read(buf)
		if err == io.EOF {
			// read done
			// logging.Trace("End-of-File")
			transferring = false
			isLastChunk = true
			err = nil
		}
		if err != nil {
			return errors.Wrapf(err, "[%v] function fd.Read()", errors.Trace())
		}

		chunkChecksum, err := blake2bChecksum(buf[0:n])
		if err != nil {
			return errors.Wrapf(err, "[%v] function blake2bChecksum()", errors.Trace())
		}

		logging.Tracef("FileChunk #%v checksum: %x", chunkNumber, chunkChecksum)
		p := newFileChunk(srcFile, srcFileName, dstFilePath, mmID, dstID, requesterID, chunkNumber, int64(n),
			chunkChecksum, fileChecksum, buf[0:n], isLastChunk, interactive)

		queuing.TxControlQueue <- p

		chunkNumber++
	}

	return nil
}

func newFileChunk(srcFilePath, srcFileName, dstFilePath, srcID, dstID, requesterID string, chunkNumber,
	sizeInBytes int64, chunkChecksum, fileChecksum, data []byte,
	isLastChunk, interactive bool) *mmsp.Payload {

	return &mmsp.Payload{
		SrcID:       srcID,
		DstID:       dstID,
		Type: mmsp.PDUType_TRANSFER,
		TransferPDU: &mmsp.TransferPDU{
			Metadata: &mmsp.StreamMetadata{
				RequesterID: requesterID,
				Interactive: interactive,
				PSK:           viper.GetString("management.auth.psk"),
				SecurityToken: viper.GetString("management.auth.securityToken"),
			},
			Type: mmsp.TransferMsgType_TRANSFER_DATA,
			Transfer: &stream.Transfer{
				Paths: &stream.Paths{
					SrcFilePath: srcFilePath,
					SrcFileName: srcFileName,
					SrcIsDir:    false,
					DstFilePath: dstFilePath,
				},
				FileChunk: &stream.FileChunk{
					SizeInBytes:   sizeInBytes,
					SendTime:      time.Now().Unix(),
					ChunkChecksum: chunkChecksum,
					Data:          data,
					ChunkNumber:   chunkNumber,
					IsLastChunk:   isLastChunk,
					FileChecksum:  fileChecksum,
				},
				FileAck: nil,
			},
		},
	}

	// return &mmsp.Payload{
	// 	SrcID:         srcID,
	// 	DstID:         dstID,
	// 	RequesterID:   requesterID,
	// 	Interactive:   interactive,
	// 	PSK:           viper.GetString("management.auth.psk"),
	// 	SecurityToken: viper.GetString("management.auth.securityToken"),
	// 	PayloadType:   mmsp.PayloadType_TRANSFER_DATA,
	// 	Transfer: &transfer.Transfer{
	// 		Paths: &transfer.Paths{
	// 			SrcFilePath: srcFilePath,
	// 			SrcFileName: srcFileName,
	// 			SrcIsDir:    false,
	// 			DstFilePath: dstFilePath,
	// 		},
	// 		FileChunk: &transfer.FileChunk{
	// 			SizeInBytes:   sizeInBytes,
	// 			SendTime:      time.Now().Unix(),
	// 			ChunkChecksum: chunkChecksum,
	// 			Data:          data,
	// 			ChunkNumber:   chunkNumber,
	// 			IsLastChunk:   isLastChunk,
	// 			FileChecksum:  fileChecksum,
	// 		},
	// 		FileAck: nil,
	// 	},
	// }
}

func newDir(srcFileMode uint32, srcFilePath, srcFileName, dstFilePath, srcID, dstID, requesterID string, interactive bool) *mmsp.Payload {
	return &mmsp.Payload{
		SrcID:       srcID,
		DstID:       dstID,
		Type: mmsp.PDUType_TRANSFER,
		TransferPDU: &mmsp.TransferPDU{
			Metadata: &mmsp.StreamMetadata{
				RequesterID: requesterID,
				Interactive: interactive,
				PSK:           viper.GetString("management.auth.psk"),
				SecurityToken: viper.GetString("management.auth.securityToken"),
			},
			Type: mmsp.TransferMsgType_TRANSFER_DATA,
			Transfer: &stream.Transfer{
				Paths: &stream.Paths{
					SrcFilePath: srcFilePath,
					SrcFileName: srcFileName,
					SrcFileMode: srcFileMode,
					SrcIsDir:    true,
					DstFilePath: dstFilePath,
				},
				FileChunk: nil,
				FileAck: nil,
			},
		},
	}

	// return &mmsp.Payload{
	// 	SrcID:         srcID,
	// 	DstID:         dstID,
	// 	RequesterID:   requesterID,
	// 	Interactive:   interactive,
	// 	PSK:           viper.GetString("management.auth.psk"),
	// 	SecurityToken: viper.GetString("management.auth.securityToken"),
	// 	PayloadType:   mmsp.PayloadType_TRANSFER_DATA,
	// 	Transfer: &transfer.Transfer{
	// 		Paths: &transfer.Paths{
	// 			SrcFilePath: srcFilePath,
	// 			SrcFileName: srcFileName,
	// 			SrcFileMode: srcFileMode,
	// 			SrcIsDir:    true,
	// 			DstFilePath: dstFilePath,
	// 		},
	// 		FileAck: nil,
	// 	},
	// }
}
