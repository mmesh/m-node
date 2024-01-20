package ctlogdb

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/dgraph-io/badger/v4"
	"mmesh.dev/m-api-go/grpc/resources/nstore/netdb"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-node/internal/app/node/kvstore/utils"
)

type conntrackLogEntry struct {
	*netdb.ConntrackLogEntry
}

func (e *conntrackLogEntry) newEntry() (*badger.Entry, error) {
	k := e.encodeKey()

	v, err := e.getValue()
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function e.getValue()", errors.Trace())
	}

	// ttl: 30 days
	return badger.NewEntry(k, v).WithTTL(30 * 24 * time.Hour), nil
}

func (e *conntrackLogEntry) encodeKey() []byte {
	return []byte(fmt.Sprintf("%s:%d:%d:%d:%s:%d:%s:%d",
		ctLogPrefix,
		e.Timestamp,
		int(e.Connection.AF),
		int(e.Connection.Proto),
		e.Connection.SrcIP,
		e.Connection.SrcPort,
		e.Connection.DstIP,
		e.Connection.DstPort,
	))
}

func (e *conntrackLogEntry) getValue() ([]byte, error) {
	var bValue bytes.Buffer

	err := gob.NewEncoder(&bValue).Encode(&e.Status)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function gob.NewEncoder()", errors.Trace())
	}

	return bValue.Bytes(), nil
}

func decodeKey(k []byte) (*netdb.ConntrackLogEntry, error) {
	s := strings.Split(string(k), ":")

	if len(s) != 8 {
		return nil, fmt.Errorf("[ctlogdb] malformed key")
	}

	if len(s[0]) == 0 || len(s[1]) == 0 || len(s[2]) == 0 || len(s[3]) == 0 ||
		len(s[4]) == 0 || len(s[5]) == 0 || len(s[6]) == 0 || len(s[7]) == 0 {
		return nil, fmt.Errorf("[ctlogdb] invalid key")
	}

	tm, err := strconv.ParseInt(s[1], 10, 64)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function strconv.ParseInt()", errors.Trace())
	}

	af, err := utils.ParseAddressFamily(s[2])
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function utils.ParseAddressFamily()", errors.Trace())
	}

	proto, err := utils.ParseProto(s[3])
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function utils.ParseProto()", errors.Trace())
	}

	srcPort, err := strconv.Atoi(s[5])
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function strconv.Atoi()", errors.Trace())
	}

	dstPort, err := strconv.Atoi(s[7])
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function strconv.Atoi()", errors.Trace())
	}

	return &netdb.ConntrackLogEntry{
		Timestamp: tm,
		Connection: &netdb.Connection{
			AF:      af,
			SrcIP:   s[4],
			DstIP:   s[6],
			Proto:   proto,
			SrcPort: uint32(srcPort),
			DstPort: uint32(dstPort),
		},
	}, nil
}

func decodeKV(k, v []byte) (*netdb.ConntrackLogEntry, error) {
	e, err := decodeKey(k)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function decodeKey()", errors.Trace())
	}

	var value netdb.ConnectionStatus

	if err := gob.NewDecoder(bytes.NewBuffer(v)).Decode(&value); err != nil {
		return nil, errors.Wrapf(err, "[%v] function gob.NewDecoder()", errors.Trace())
	}

	e.Status = value

	return e, nil
}
