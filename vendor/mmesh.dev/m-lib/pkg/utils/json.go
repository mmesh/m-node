package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

// ReadJsonFile parse json data from a file
func ReadJsonFile(jsonFile string) ([]byte, error) {
	var jsonBlob []byte

	if _, err := os.Stat(jsonFile); err == nil {
		jsonBlob, err = ioutil.ReadFile(jsonFile)
		if err != nil {
			return nil, errors.Wrapf(err, "[%v] function ioutil.ReadFile(jsonFile)", errors.Trace())
		}
	} else if os.IsNotExist(err) {
		//fmt.Printf("jsonFile not found, please, execute 'nx auth login' to authenticate")
		msg.Error("JSON file not found.")
		return nil, errors.Wrapf(err, "[%v] file %v not found", errors.Trace(), jsonFile)
	} else {
		return nil, errors.Wrapf(err, "[%v] file stat error", errors.Trace())
	}

	if !json.Valid(jsonBlob) {
		return nil, errors.Errorf("Invalid JSON (file %v)", jsonFile)
	}

	return jsonBlob, nil
}
