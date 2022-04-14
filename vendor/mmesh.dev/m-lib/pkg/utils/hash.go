package utils

import (
	"crypto/sha256"
	"io"
	"os"

	"mmesh.dev/m-lib/pkg/errors"
)

func ChecksumSHA256(filePath string) ([]byte, error) {
	fd, err := os.Open(filePath)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function os.Open(filePath)", errors.Trace())
	}
	defer fd.Close()

	h := sha256.New()
	if _, err := io.Copy(h, fd); err != nil {
		return nil, errors.Wrapf(err, "[%v] function io.Copy(h, fd)", errors.Trace())
	}

	return h.Sum(nil), nil
}
