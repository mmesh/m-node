package mmp

import (
	"os"

	"github.com/dchest/blake2b"
	"x6a.dev/pkg/errors"
)

func blake2bChecksum(data []byte) ([]byte, error) {
	h, err := blake2b.New(nil)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function blake2b.New(nil)", errors.Trace())
	}

	if _, err := h.Write(data); err != nil {
		return nil, errors.Wrapf(err, "[%v] function h.Write(data)", errors.Trace())
	}

	return h.Sum(nil), nil
}

func getFileInfo(filePath string) (os.FileInfo, error) {
	fd, err := os.Open(filePath)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function os.Open()", errors.Trace())
	}
	defer fd.Close()

	fileInfo, err := fd.Stat()
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function fd.Stat()", errors.Trace())
	}

	return fileInfo, nil
}
