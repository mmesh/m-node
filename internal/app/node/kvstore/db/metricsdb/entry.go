package metricsdb

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/dgraph-io/badger/v4"
	"mmesh.dev/m-api-go/grpc/resources/nstore"
	"mmesh.dev/m-api-go/grpc/resources/nstore/metricsdb"
	"mmesh.dev/m-lib/pkg/errors"
)

type hostMetricDataPoint struct {
	*metricsdb.HostMetricDataPoint
}

func (dp *hostMetricDataPoint) newEntry() (*badger.Entry, error) {
	k := dp.encodeKey()

	// fmt.Printf("----- newEntry | k=%s\n", k)

	v, err := dp.getValue()
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function dp.getValue()", errors.Trace())
	}

	return badger.NewEntry(k, v).WithTTL(getTTL(dp.TimeRange)), nil
}

func (dp *hostMetricDataPoint) encodeKey() []byte {
	return []byte(fmt.Sprintf("%s:%d", encodeKeyPrefix(dp.TimeRange, dp.Metric), dp.Timestamp))
}

/*
func (dp *hostMetricDataPoint) encodeKey() []byte {
	return []byte(fmt.Sprintf("%s:%d:%d:%d",
		hostMetricsPrefix,
		int(dp.TimeRange),
		int(dp.Metric),
		dp.Timestamp,
	))
}
*/

func (dp *hostMetricDataPoint) getValue() ([]byte, error) {
	var bValue bytes.Buffer

	err := gob.NewEncoder(&bValue).Encode(&dp.Value)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function gob.NewEncoder()", errors.Trace())
	}

	return bValue.Bytes(), nil
}

func encodeKeyPrefix(timeRange nstore.TimeRange, metric metricsdb.HostMetricType) string {
	return fmt.Sprintf("%s:%d:%d", hostMetricsPrefix, int(timeRange), int(metric))
}

func decodeKey(k []byte) (*metricsdb.HostMetricDataPoint, error) {
	s := strings.Split(string(k), ":")

	if len(s) != 4 {
		return nil, fmt.Errorf("[metricsdb] malformed key")
	}

	if len(s[0]) == 0 || len(s[1]) == 0 || len(s[2]) == 0 || len(s[3]) == 0 {
		return nil, fmt.Errorf("[metricsdb] invalid key")
	}

	tm, err := strconv.ParseInt(s[3], 10, 64)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function strconv.ParseInt()", errors.Trace())
	}

	ttl, err := parseMetricTTL(s[1])
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function parseMetricTTL()", errors.Trace())
	}

	mt, err := parseMetricType(s[2])
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function parseMetricType()", errors.Trace())
	}

	return &metricsdb.HostMetricDataPoint{
		Timestamp: tm,
		TimeRange: ttl,
		Metric:    mt,
	}, nil
}

func decodeKV(k, v []byte) (*metricsdb.HostMetricDataPoint, error) {
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

func getTTL(ttl nstore.TimeRange) time.Duration {
	switch ttl {
	case nstore.TimeRange_TTL_1H:
		return 1 * time.Hour
	case nstore.TimeRange_TTL_6H:
		return 6 * time.Hour
	case nstore.TimeRange_TTL_12H:
		return 12 * time.Hour
	case nstore.TimeRange_TTL_24H:
		return 24 * time.Hour
	case nstore.TimeRange_TTL_7D:
		return 7 * 24 * time.Hour
	case nstore.TimeRange_TTL_14D:
		return 14 * 24 * time.Hour
	case nstore.TimeRange_TTL_30D:
		return 30 * 24 * time.Hour
	case nstore.TimeRange_TTL_365D:
		return 365 * 24 * time.Hour
	}

	return 1 * time.Hour
}

func parseMetricTTL(str string) (nstore.TimeRange, error) {
	if len(str) == 0 {
		return nstore.TimeRange_TTL_UNDEFINED, fmt.Errorf("[kvstore] invalid timeRange key")
	}

	ttl, err := strconv.Atoi(str)
	if err != nil {
		return nstore.TimeRange_TTL_UNDEFINED, errors.Wrapf(err, "[%v] function strconv.Atoi()", errors.Trace())
	}

	switch ttl {
	case int(nstore.TimeRange_TTL_1H):
		return nstore.TimeRange_TTL_1H, nil
	case int(nstore.TimeRange_TTL_6H):
		return nstore.TimeRange_TTL_6H, nil
	case int(nstore.TimeRange_TTL_12H):
		return nstore.TimeRange_TTL_12H, nil
	case int(nstore.TimeRange_TTL_24H):
		return nstore.TimeRange_TTL_24H, nil
	case int(nstore.TimeRange_TTL_7D):
		return nstore.TimeRange_TTL_7D, nil
	case int(nstore.TimeRange_TTL_14D):
		return nstore.TimeRange_TTL_14D, nil
	case int(nstore.TimeRange_TTL_30D):
		return nstore.TimeRange_TTL_30D, nil
	case int(nstore.TimeRange_TTL_365D):
		return nstore.TimeRange_TTL_365D, nil
	}

	return nstore.TimeRange_TTL_UNDEFINED, fmt.Errorf("[kvstore] unknown timeRange key")
}

func parseMetricType(str string) (metricsdb.HostMetricType, error) {
	if len(str) == 0 {
		return metricsdb.HostMetricType_UNDEFINED, fmt.Errorf("[metricsdb] invalid metricType key")
	}

	mt, err := strconv.Atoi(str)
	if err != nil {
		return metricsdb.HostMetricType_UNDEFINED, errors.Wrapf(err, "[%v] function strconv.Atoi()", errors.Trace())
	}

	switch mt {
	case int(metricsdb.HostMetricType_NET_RX_BYTES):
		return metricsdb.HostMetricType_NET_RX_BYTES, nil
	case int(metricsdb.HostMetricType_NET_TX_BYTES):
		return metricsdb.HostMetricType_NET_TX_BYTES, nil
	case int(metricsdb.HostMetricType_HOST_LOAD_AVG):
		return metricsdb.HostMetricType_HOST_LOAD_AVG, nil
	case int(metricsdb.HostMetricType_HOST_CPU_USAGE):
		return metricsdb.HostMetricType_HOST_CPU_USAGE, nil
	case int(metricsdb.HostMetricType_HOST_MEM_USAGE):
		return metricsdb.HostMetricType_HOST_MEM_USAGE, nil
	case int(metricsdb.HostMetricType_HOST_DISK_USAGE):
		return metricsdb.HostMetricType_HOST_DISK_USAGE, nil
	}

	return metricsdb.HostMetricType_UNDEFINED, fmt.Errorf("[metricsdb] unknown metricType key")
}
