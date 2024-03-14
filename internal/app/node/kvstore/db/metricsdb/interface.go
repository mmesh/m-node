package metricsdb

import (
	"fmt"

	"github.com/dgraph-io/badger/v4"
	"mmesh.dev/m-api-go/grpc/resources/nstore/metricsdb"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/xlog"
)

const hostMetricsPrefix string = "host"

var InputQueue = make(chan []*metricsdb.HostMetricDataPoint, 128)
var RequestQueue = make(chan *metricsdb.HostMetricsRequest, 128)

type Interface interface {
	Query(r *metricsdb.HostMetricsRequest) (*metricsdb.HostMetricsResponse, error)
	// QueryAll(r *metricsdb.HostMetricsRequest) (*metricsdb.HostMetricsResponse, error)
	WriteBatch(dps []*metricsdb.HostMetricDataPoint) error
	Scan(keyPrefix []byte) ([]*metricsdb.HostMetricDataPoint, error)
	Last(keyPrefix []byte) (*metricsdb.HostMetricDataPoint, error)
	Close()
}
type tsDB struct {
	db                   *badger.DB
	aggControllerCloseCh chan struct{}
}

func Open(db *badger.DB) Interface {
	tsdb := &tsDB{
		db:                   db,
		aggControllerCloseCh: make(chan struct{}, 1),
	}

	go tsdb.aggController()

	return tsdb
}

func (tsdb *tsDB) Close() {
	tsdb.aggControllerCloseCh <- struct{}{}
}

func (tsdb *tsDB) WriteBatch(dps []*metricsdb.HostMetricDataPoint) error {
	wb := tsdb.db.NewWriteBatch()
	defer wb.Cancel()

	for _, p := range dps {
		dp := &hostMetricDataPoint{p}

		e, err := dp.newEntry()
		if err != nil {
			return errors.Wrapf(err, "[%v] function dp.newEntry()", errors.Trace())
		}

		if err := wb.SetEntry(e); err != nil {
			return errors.Wrapf(err, "[%v] function wb.SetEntry()", errors.Trace())
		}
	}
	if err := wb.Flush(); err != nil {
		return errors.Wrapf(err, "[%v] function wb.Flush()", errors.Trace())
	}

	return nil
}

func (tsdb *tsDB) Scan(keyPrefix []byte) ([]*metricsdb.HostMetricDataPoint, error) {
	// fmt.Printf("----- scan | keyPrefix=%s\n", keyPrefix)

	r := make([]*metricsdb.HostMetricDataPoint, 0)

	if err := tsdb.db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()

		for it.Seek(keyPrefix); it.ValidForPrefix(keyPrefix); it.Next() {
			item := it.Item()
			k := item.Key()

			v, err := item.ValueCopy(nil)
			if err != nil {
				xlog.Errorf("[metricsdb] Unable to get value for key=%s: %v", k, err)
				continue
			}

			dp, err := decodeKV(k, v)
			if err != nil {
				xlog.Errorf("[metricsdb] Unable to decode key=%s: %v", k, err)
				continue
			}

			// fmt.Printf("====== scan | key=%s | value=%v\n", k, dp.Value)

			r = append(r, dp)

			// if err := item.Value(func(v []byte) error {
			// 	fmt.Printf("[metricsdb] key=%s, value=%s\n", k, v)

			// 	dp, err := decodeKV(k, v)
			// 	if err != nil {
			// 		return errors.Wrapf(err, "[%v] function decodeKV()", errors.Trace())
			// 	}

			// 	r = append(r, dp)

			// 	return nil
			// }); err != nil {
			// 	return errors.Wrapf(err, "[%v] function item.Value()", errors.Trace())
			// }
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[%v] function tsdb.db.View()", errors.Trace())
	}

	return r, nil
}

func (tsdb *tsDB) Last(keyPrefix []byte) (*metricsdb.HostMetricDataPoint, error) {
	var dp *metricsdb.HostMetricDataPoint

	if err := tsdb.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 1
		opts.Reverse = true

		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Seek(keyPrefix); it.ValidForPrefix(keyPrefix); it.Next() {
			item := it.Item()
			k := item.Key()

			v, err := item.ValueCopy(nil)
			if err != nil {
				xlog.Errorf("[metricsdb] Unable to get value for key=%s: %v", k, err)
				continue
			}

			fmt.Printf("[metricsdb] key=%s, value=%s\n", k, v)

			dp, err = decodeKV(k, v)
			if err != nil {
				xlog.Errorf("[metricsdb] Unable to decode key=%s: %v", k, err)
				continue
			}

			break
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[%v] function tsdb.db.View()", errors.Trace())
	}

	return dp, nil
}
