package utils

import (
	"crypto/sha256"
	"io"
	"os"

	"mmesh.dev/m-lib/pkg/errors"
	"github.com/zeebo/blake3"
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

func ChecksumBlake3(data []byte) ([]byte, error) {
	h := blake3.New()

	if _, err := h.Write(data); err != nil {
		return nil, errors.Wrapf(err, "[%v] function h.Write()", errors.Trace())
	}

	return h.Sum(nil), nil
}

/*
func ChecksumSHA256(data []byte) ([]byte, error) {
	h := sha256.New()

	if _, err := h.Write(data); err != nil {
		return nil, errors.Wrapf(err, "[%v] function h.Write()", errors.Trace())
	}

	return h.Sum(nil), nil
}
*/
