package mmp

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path/filepath"

	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/logging"
	"mmesh.dev/m-lib/pkg/utils"
	"mmesh.dev/m-lib/pkg/utils/colors"
	"mmesh.dev/m-lib/pkg/xlog"
)

func transferDataRx(ctx context.Context, p *mmsp.Payload) error {
	//if viper.GetBool("agent.management.disableTransfer") {
	if !mgmtAuth(p) {
		transferDisabled(p)
		return nil
	}

	logging.Tracef("Received transfer data request from %s", p.SrcID)

	if err := transferRx(p); err != nil {
		errCause := errors.Cause(err)
		transferAckTx(ctx, p, errCause.Error())
		return nil
		//return errors.Wrapf(err, "[%v] function transferRx()", errors.Trace())
	}
	transferAckTx(ctx, p, "")

	return nil
}

func transferRx(p *mmsp.Payload) error {
	var fd *os.File
	var err error

	srcFilePath := p.Transfer.Paths.SrcFilePath
	srcFileName := p.Transfer.Paths.SrcFileName
	srcFileMode := p.Transfer.Paths.SrcFileMode
	srcIsDir := p.Transfer.Paths.SrcIsDir
	dstFilePath := p.Transfer.Paths.DstFilePath

	filePath, err := checkDstFile(srcFileMode, srcIsDir, srcFileName, dstFilePath)
	if err != nil {
		return errors.Wrapf(err, "[%v] function checkDstFile()", errors.Trace())
	}

	// check filePath
	if len(filePath) == 0 {
		return errors.New("invalid file path")
	}

	if srcIsDir {
		fType := fmt.Sprintf("%s%s%s", colors.Blue("["), colors.DarkBlue("dir"), colors.Blue("]"))
		logging.Println(colors.DarkCyan(srcFilePath), fmt.Sprintf("%s %s %s", colors.Cyan("->"), colors.DarkWhite(filePath), fType))

		return nil
	}

	// check file chunk
	if p.Transfer.FileChunk == nil {
		return errors.New("invalid file chunk")
	}

	logging.Trace("Receiving file chunk..")

	chunkNumber := p.Transfer.FileChunk.ChunkNumber

	// check file chunk data
	if len(p.Transfer.FileChunk.Data) == 0 && !p.Transfer.FileChunk.IsLastChunk {
		return errors.New("no bytes received")
	}

	// check file chunk size
	if p.Transfer.FileChunk.SizeInBytes != int64(len(p.Transfer.FileChunk.Data)) {
		return errors.Errorf("%v == p.Transfer.FileChunk.SizeInBytes != int64(len(p.Transfer.FileChunk.Data)) == %v",
			p.Transfer.FileChunk.SizeInBytes, int64(len(p.Transfer.FileChunk.Data)))
	}

	// check file chunk checksum
	chunkChecksum, err := blake2bChecksum(p.Transfer.FileChunk.Data)
	if err != nil {
		return errors.Wrapf(err, "[%v] function blake2bChecksum(p.Transfer.FileChunk.Data)", errors.Trace())
	}
	xlog.Tracef("Received fileChunk #%v. Checksum: %x", chunkNumber, chunkChecksum)
	if !bytes.Equal(chunkChecksum, p.Transfer.FileChunk.ChunkChecksum) {
		logging.Error("Chunk invalid checksum")
		return errors.Errorf("invalid chunk %v. Data checksums mismatch", p.Transfer.FileChunk.ChunkNumber)
	}

	f := filePath + ".tmp"

	// open file
	if chunkNumber == 0 {
		// create file
		fd, err = os.OpenFile(f, os.O_CREATE|os.O_WRONLY, 0640)
		if err != nil {
			return errors.Wrapf(err, "[%v] function os.Create()", errors.Trace())
		}
	} else {
		fd, err = os.OpenFile(f, os.O_APPEND|os.O_WRONLY, 0640)
		if err != nil {
			return errors.Wrapf(err, "[%v] function os.Create()", errors.Trace())
		}
	}

	// write to file
	if _, err := fd.Write(p.Transfer.FileChunk.Data); err != nil {
		fd.Close() // ignore error; Write error takes precedence
		return errors.Wrapf(err, "[%v] function fd.Write()", errors.Trace())
	}

	// last file chunk
	if p.Transfer.FileChunk.IsLastChunk {
		// close file
		if err := fd.Close(); err != nil {
			return errors.Wrapf(err, "[%v] function fd.Close()", errors.Trace())
		}

		// rename file
		if err := os.Rename(f, filePath); err != nil {
			return errors.Wrapf(err, "[%v] function os.Rename(f, filePath)", errors.Trace())
		}

		// check file checksum
		fileChecksum, err := utils.ChecksumSHA256(filePath)
		if err != nil {
			return errors.Wrapf(err, "[%v] function utils.ChecksumSHA256(filePath)", errors.Trace())
		}
		logging.Tracef("File checksum: %x", fileChecksum)
		if !bytes.Equal(fileChecksum, p.Transfer.FileChunk.FileChecksum) {
			logging.Error("File invalid checksum")
			return errors.New("file checksum mismatch")
		}

		fType := fmt.Sprintf("%s%s%s", colors.Yellow("["), colors.DarkYellow("file"), colors.Yellow("]"))
		logging.Println(colors.DarkCyan(srcFilePath), fmt.Sprintf("%s %s %s", colors.Cyan("->"), colors.DarkWhite(filePath), fType))
	}

	return nil
}

func checkDstFile(srcFileMode uint32, srcIsDir bool, srcFileName, dstFilePath string) (string, error) {
	var filePath string
	var dstExists bool

	if srcIsDir {
		filePath = dstFilePath
	} else {
		filePath = filepath.Join(dstFilePath, srcFileName)
	}

	fileInfo, err := os.Stat(dstFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			// doesn't exist yet - new file/dir
			dstExists = false
		} else {
			// filePath may or may not exist. See err for details.
			//logging.Errorf("filePath may or may not exist. See err for details: %v", err)
			return "", errors.Wrapf(err, "[%v] function os.Stat()", errors.Trace())
		}
	} else {
		dstExists = true
	}

	if !dstExists {
		if srcIsDir {
			logging.Debugf("Detected recursive transfer, creating directory %s", dstFilePath)
			err := os.MkdirAll(dstFilePath, os.FileMode(srcFileMode))
			if err != nil {
				return "", errors.Wrapf(err, "[%v] function os.MkdirAll()", errors.Trace())
			}
		} else {
			return "", errors.Errorf("destination directory %s not found", dstFilePath)
		}
	} else {
		if srcIsDir && !fileInfo.IsDir() {
			return "", errors.Errorf("destination path %s is an existing non-directory", dstFilePath)
		}

		if !srcIsDir && !fileInfo.IsDir() {
			return "", errors.Errorf("destination path %s is not a directory", dstFilePath)
		}
	}

	return filePath, nil
}

/*
func writeChunk(filePath string, chunkNumber int64, data []byte) error {
	f := fmt.Sprintf("%s.%v", filePath+".tmp", chunkNumber)

	// create file
	fd, err := os.OpenFile(f, os.O_CREATE|os.O_WRONLY, 0640)
	if err != nil {
		return errors.Wrapf(err, "[%v] function os.Create()", errors.Trace())
	}

	// write to file
	if _, err := fd.Write(data); err != nil {
		fd.Close() // ignore error; Write error takes precedence
		return errors.Wrapf(err, "[%v] function fd.Write()", errors.Trace())
	}

	// close file
	if err := fd.Close(); err != nil {
		return errors.Wrapf(err, "[%v] function fd.Close()", errors.Trace())
	}

	return nil
}

func writeFile(filePath string, lastChunkNumber int64) error {
	var i int64

	// create file
	fd, err := os.Create(filePath)
	if err != nil {
		return errors.Wrapf(err, "[%v] function os.Create()", errors.Trace())
	}
	defer fd.Close()

	for i = 0; i <= lastChunkNumber; i++ {
		f := fmt.Sprintf("%s.%v", filePath+".tmp", i)

		for !fileExists(f) {
			time.Sleep(500 + time.Millisecond)
		}

		// open file chunk
		chunkFd, err := os.Open(f)
		if err != nil {
			return errors.Wrapf(err, "[%v] function os.Create()", errors.Trace())
		}
		defer func() {
			if err := chunkFd.Close(); err != nil {
				logging.Errorf("Unable to close file %s: %v", f, err)
			}
			if err := os.Remove(f); err != nil {
				logging.Errorf("Unable to remove file %s: %v", f, err)
				return
			}
			logging.Tracef("Deleted temporary file %s", f)
		}()

		buf := make([]byte, fileChunk_size*2)
		n, err := chunkFd.Read(buf)
		if err == io.EOF {
			// read done
			logging.Tracef("File %s: EOF detected", f)
			continue
		}
		if err != nil {
			return errors.Wrapf(err, "[%v] function chunkFd.Read(buf)", errors.Trace())
		}

		// write to file
		if _, err := fd.Write(buf[0:n]); err != nil {
			fd.Close() // ignore error; Write error takes precedence
			return errors.Wrapf(err, "[%v] function fd.Write()", errors.Trace())
		}
	}

	if err := fd.Close(); err != nil {
		return errors.Wrapf(err, "[%v] function fd.Close()", errors.Trace())
	}

	return nil
}
*/
