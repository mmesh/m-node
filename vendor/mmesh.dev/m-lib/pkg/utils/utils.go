package utils

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	"x6a.dev/pkg/errors"
	"x6a.dev/pkg/msg"
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

// RandString generates a random ASCII string with at least one digit
// and one special character.
func RandString(length int) string {
	rand.Seed(time.Now().UnixNano())

	digits := "0123456789"
	specials := "~=+%^*/()[]{}/!@#$?|"
	all := "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		digits + specials
	buf := make([]byte, length)
	buf[0] = digits[rand.Intn(len(digits))]
	buf[1] = specials[rand.Intn(len(specials))]
	for i := 2; i < length; i++ {
		buf[i] = all[rand.Intn(len(all))]
	}
	rand.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})

	return string(buf) // E.g. "3i[g0|)z"
}
