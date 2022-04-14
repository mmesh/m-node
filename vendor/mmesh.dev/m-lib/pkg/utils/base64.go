package utils

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"

	"mmesh.dev/m-lib/pkg/errors"
)

// FileToB64 read and convert a file to base64
func FileToB64(file string) (string, error) {
	var blob []byte

	if _, err := os.Stat(file); err == nil {
		blob, err = ioutil.ReadFile(file)
		if err != nil {
			return "", errors.Wrapf(err, "[%v] function ioutil.ReadFile(file)", errors.Trace())
		}
	} else if os.IsNotExist(err) {
		fmt.Printf("file %v not found", file)
		return "", errors.Wrapf(err, "[%v] file %v not found", errors.Trace(), file)
	} else {
		return "", errors.Wrapf(err, "[%v] file stat error", errors.Trace())
	}

	return base64.URLEncoding.EncodeToString(blob), nil
}
