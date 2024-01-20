package netflowdb

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

type netflowEntry struct {
	*netdb.NetFlowEntry
}

func (e *netflowEntry) newEntry() (*badger.Entry, error) {
	k := e.encodeKey()

	v, err := e.getValue()
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function e.getValue()", errors.Trace())
	}

	// ttl: 30 days
	return badger.NewEntry(k, v).WithTTL(30 * 24 * time.Hour), nil
}

func (e *netflowEntry) encodeKey() []byte {
	return []byte(fmt.Sprintf("%s:%d:%d:%d:%s:%d:%s:%d:%d",
		netflowPrefix,
		e.Timestamp,
		int(e.Flow.Connection.AF),
		int(e.Flow.Connection.Proto),
		e.Flow.Connection.SrcIP,
		e.Flow.Connection.SrcPort,
		e.Flow.Connection.DstIP,
		e.Flow.Connection.DstPort,
		int(e.Flow.Direction),
	))
}

func (e *netflowEntry) getValue() ([]byte, error) {
	var bValue bytes.Buffer

	err := gob.NewEncoder(&bValue).Encode(e.Traffic)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function gob.NewEncoder()", errors.Trace())
	}

	return bValue.Bytes(), nil
}

func decodeKey(k []byte) (*netdb.NetFlowEntry, error) {
	s := strings.Split(string(k), ":")

	if len(s) != 9 {
		return nil, fmt.Errorf("[netflowdb] malformed key")
	}

	if len(s[0]) == 0 || len(s[1]) == 0 || len(s[2]) == 0 || len(s[3]) == 0 ||
		len(s[4]) == 0 || len(s[5]) == 0 || len(s[6]) == 0 || len(s[7]) == 0 ||
		len(s[8]) == 0 {
		return nil, fmt.Errorf("[netflowdb] invalid key")
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

	dir, err := parseFlowDirection(s[8])
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function parseFlowDirection()", errors.Trace())
	}

	return &netdb.NetFlowEntry{
		Timestamp: tm,
		Flow: &netdb.Flow{
			Connection: &netdb.Connection{
				AF:      af,
				SrcIP:   s[4],
				DstIP:   s[6],
				Proto:   proto,
				SrcPort: uint32(srcPort),
				DstPort: uint32(dstPort),
			},
			Direction: dir,
		},
	}, nil
}

func decodeKV(k, v []byte) (*netdb.NetFlowEntry, error) {
	e, err := decodeKey(k)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function decodeKey()", errors.Trace())
	}

	var traffic netdb.TrafficCounter

	if err := gob.NewDecoder(bytes.NewBuffer(v)).Decode(&traffic); err != nil {
		return nil, errors.Wrapf(err, "[%v] function gob.NewDecoder()", errors.Trace())
	}

	e.Traffic = &traffic

	return e, nil
}

func parseFlowDirection(str string) (netdb.ConnectionDirection, error) {
	if len(str) == 0 {
		return netdb.ConnectionDirection_UNKNOWN_DIRECTION, fmt.Errorf("[netflowdb] invalid direction field")
	}

	dir, err := strconv.Atoi(str)
	if err != nil {
		return netdb.ConnectionDirection_UNKNOWN_DIRECTION, errors.Wrapf(err, "[%v] function strconv.Atoi()", errors.Trace())
	}

	switch dir {
	case int(netdb.ConnectionDirection_INCOMING):
		return netdb.ConnectionDirection_INCOMING, nil
	case int(netdb.ConnectionDirection_OUTGOING):
		return netdb.ConnectionDirection_OUTGOING, nil
	}

	return netdb.ConnectionDirection_UNKNOWN_DIRECTION, fmt.Errorf("[netflowdb] unknown direction field")
}
