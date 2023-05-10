package tss

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/dgraph-io/badger/v4"
	"mmesh.dev/m-api-go/grpc/resources/tss"
	"mmesh.dev/m-lib/pkg/errors"
)

type dataPoint struct {
	*tss.DataPoint
}

func (dp *dataPoint) newEntry() (*badger.Entry, error) {
	k := dp.encodeKey()

	// fmt.Printf("----- newEntry | k=%s\n", k)

	v, err := dp.getValue()
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function dp.getValue()", errors.Trace())
	}

	return badger.NewEntry(k, v).WithTTL(getTTL(dp.TimeRange)), nil
}

func (dp *dataPoint) encodeKey() []byte {
	return []byte(fmt.Sprintf("%s:%d", encodeKeyPrefix(dp.TimeRange, dp.Metric), dp.Timestamp))
}

func (dp *dataPoint) getValue() ([]byte, error) {
	var bValue bytes.Buffer

	err := gob.NewEncoder(&bValue).Encode(&dp.Value)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function gob.NewEncoder()", errors.Trace())
	}

	return bValue.Bytes(), nil
}

func encodeKeyPrefix(timeRange tss.TimeRange, metric tss.MetricType) string {
	return fmt.Sprintf("%d:%d", int(timeRange), int(metric))
}

func decodeKey(k []byte) (*tss.DataPoint, error) {
	s := strings.Split(string(k), ":")

	if len(s) != 3 {
		return nil, fmt.Errorf("[tss] malformed key")
	}

	if len(s[0]) == 0 || len(s[1]) == 0 || len(s[2]) == 0 {
		return nil, fmt.Errorf("[tss] invalid key")
	}

	tm, err := strconv.ParseInt(s[2], 10, 64)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function strconv.ParseInt()", errors.Trace())
	}

	ttl, err := parseMetricTTL(s[0])
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function parseMetricTTL()", errors.Trace())
	}

	mt, err := parseMetricType(s[1])
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function parseMetricType()", errors.Trace())
	}

	return &tss.DataPoint{
		Timestamp: tm,
		TimeRange: ttl,
		Metric:    mt,
	}, nil
}

func decodeKV(k, v []byte) (*tss.DataPoint, error) {
	dp, err := decodeKey(k)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function decodeKey()", errors.Trace())
	}

	var value float64

	if err := gob.NewDecoder(bytes.NewBuffer(v)).Decode(&value); err != nil {
		return nil, errors.Wrapf(err, "[%v] function gob.NewDecoder()", errors.Trace())
	}

	dp.Value = value

	return dp, nil
}

func parseMetricTTL(str string) (tss.TimeRange, error) {
	if len(str) == 0 {
		return tss.TimeRange_TTL_UNDEFINED, fmt.Errorf("[tss] invalid timeRange key")
	}

	ttl, err := strconv.Atoi(str)
	if err != nil {
		return tss.TimeRange_TTL_UNDEFINED, errors.Wrapf(err, "[%v] function strconv.Atoi()", errors.Trace())
	}

	switch ttl {
	case int(tss.TimeRange_TTL_1H):
		return tss.TimeRange_TTL_1H, nil
	case int(tss.TimeRange_TTL_6H):
		return tss.TimeRange_TTL_6H, nil
	case int(tss.TimeRange_TTL_12H):
		return tss.TimeRange_TTL_12H, nil
	case int(tss.TimeRange_TTL_24H):
		return tss.TimeRange_TTL_24H, nil
	case int(tss.TimeRange_TTL_7D):
		return tss.TimeRange_TTL_7D, nil
	case int(tss.TimeRange_TTL_14D):
		return tss.TimeRange_TTL_14D, nil
	case int(tss.TimeRange_TTL_30D):
		return tss.TimeRange_TTL_30D, nil
	case int(tss.TimeRange_TTL_365D):
		return tss.TimeRange_TTL_365D, nil
	}

	return tss.TimeRange_TTL_UNDEFINED, fmt.Errorf("[tss] unknown timeRange key")
}

func parseMetricType(str string) (tss.MetricType, error) {
	if len(str) == 0 {
		return tss.MetricType_UNDEFINED, fmt.Errorf("[tss] invalid metricType key")
	}

	mt, err := strconv.Atoi(str)
	if err != nil {
		return tss.MetricType_UNDEFINED, errors.Wrapf(err, "[%v] function strconv.Atoi()", errors.Trace())
	}

	switch mt {
	case int(tss.MetricType_NET_RX_BYTES):
		return tss.MetricType_NET_RX_BYTES, nil
	case int(tss.MetricType_NET_TX_BYTES):
		return tss.MetricType_NET_TX_BYTES, nil
	case int(tss.MetricType_HOST_LOAD_AVG):
		return tss.MetricType_HOST_LOAD_AVG, nil
	case int(tss.MetricType_HOST_CPU_USAGE):
		return tss.MetricType_HOST_CPU_USAGE, nil
	case int(tss.MetricType_HOST_MEM_USAGE):
		return tss.MetricType_HOST_MEM_USAGE, nil
	case int(tss.MetricType_HOST_DISK_USAGE):
		return tss.MetricType_HOST_DISK_USAGE, nil
	}

	return tss.MetricType_UNDEFINED, fmt.Errorf("[tss] unknown metricType key")
}

func getTTL(ttl tss.TimeRange) time.Duration {
	switch ttl {
	case tss.TimeRange_TTL_1H:
		return 1 * time.Hour
	case tss.TimeRange_TTL_6H:
		return 6 * time.Hour
	case tss.TimeRange_TTL_12H:
		return 12 * time.Hour
	case tss.TimeRange_TTL_24H:
		return 24 * time.Hour
	case tss.TimeRange_TTL_7D:
		return 7 * 24 * time.Hour
	case tss.TimeRange_TTL_14D:
		return 14 * 24 * time.Hour
	case tss.TimeRange_TTL_30D:
		return 30 * 24 * time.Hour
	case tss.TimeRange_TTL_365D:
		return 365 * 24 * time.Hour
	}

	return 1 * time.Hour
}
