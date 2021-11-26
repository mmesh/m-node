package utils

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

// http client and related generic functions

func ParsePostBody(body io.ReadCloser, obj interface{}) error {
	data, err := ioutil.ReadAll(io.LimitReader(body, 1048576))
	if err != nil {
		return err
	}

	if err := body.Close(); err != nil {
		return err
	}

	if err := json.Unmarshal(data, obj); err != nil {
		return err
	}

	return nil
}
