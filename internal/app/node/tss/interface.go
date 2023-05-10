package tss

import (
	"fmt"
	"os"
	"time"

	"github.com/dgraph-io/badger/v4"
	"mmesh.dev/m-api-go/grpc/resources/tss"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/xlog"
)

const (
	// Default BadgerDB discardRatio. It represents the discard ratio for the
	// BadgerDB GC.
	//
	// Ref: https://godoc.org/github.com/dgraph-io/badger#DB.RunValueLogGC
	badgerDiscardRatio = 0.5

	// Default BadgerDB GC interval
	badgerGCInterval = 10 * time.Minute
)

var RequestQueue = make(chan *tss.MetricsRequest, 128)

type DB interface {
	Query(req *tss.MetricsRequest) (*tss.NodeMetrics, error)
	// Set(mp *tss.DataPoint) error
	WriteBatch(dps []*tss.DataPoint) error
	Scan(keyPrefix []byte) ([]*tss.DataPoint, error)
	Last(keyPrefix []byte) (*tss.DataPoint, error)
	Close() error
}

type tsDB struct {
	db                   *badger.DB
	gcCloseCh            chan struct{}
	aggControllerCloseCh chan struct{}
}

func Open() (DB, error) {
	if err := os.MkdirAll(tssDir(), 0754); err != nil {
		return nil, errors.Wrapf(err, "[%v] function os.MkdirAll()", errors.Trace())
	}

	opts := badger.DefaultOptions(tssDir())
	opts.SyncWrites = true
	opts.Dir, opts.ValueDir = tssDir(), tssDir()
	opts.Logger = nil
	opts.InMemory = false
	opts.MetricsEnabled = false      // default: true
	opts.NumMemtables = 1            // default: 5
	opts.NumLevelZeroTables = 1      // default: 5
	opts.NumLevelZeroTablesStall = 3 // default: 15
	opts.NumCompactors = 2           // default: 4 (Run at least 2 compactors)
	opts.ValueLogFileSize = 1 << 20  // default: 1<<30 - 1 | min: 1<<20 (1MB)
	opts.MemTableSize = 2 << 20      // default: 64<<20 (64MB) | 2<<20 (2MB)
	opts.ValueThreshold = 1 << 18    // default: maxValueThreshold | 1<<18 (256KB)

	db, err := badger.Open(opts)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function badger.Open()", errors.Trace())
	}

	tsdb := &tsDB{
		db:                   db,
		gcCloseCh:            make(chan struct{}, 1),
		aggControllerCloseCh: make(chan struct{}, 1),
	}

	go tsdb.runGC()
	go tsdb.aggController()

	return tsdb, nil
}

func (tsdb *tsDB) Close() error {
	tsdb.aggControllerCloseCh <- struct{}{}
	tsdb.gcCloseCh <- struct{}{}

	return tsdb.db.Close()
}

func (tsdb *tsDB) WriteBatch(dps []*tss.DataPoint) error {
	wb := tsdb.db.NewWriteBatch()
	defer wb.Cancel()

	for _, p := range dps {
		dp := &dataPoint{p}

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

/*
func (tsdb *tsDB) Set(mp *tss.DataPoint) error {
	if err := tsdb.db.Update(func(txn *badger.Txn) error {
		dp := &dataPoint{mp}

		e, err := dp.newEntry()
		if err != nil {
			return errors.Wrapf(err, "[%v] function dp.newEntry()", errors.Trace())
		}

		if err := txn.SetEntry(e); err != nil {
			return errors.Wrapf(err, "[%v] function txn.SetEntry()", errors.Trace())
		}

		return nil
	}); err != nil {
		xlog.Debugf("[tss] Unable to set key %s for time range %s: %v",
			mp.Type.String(), mp.TTL.String(), err)
		return errors.Wrapf(err, "[%v] function tsdb.db.Update()", errors.Trace())
	}

	return nil
}
*/

func (tsdb *tsDB) Scan(keyPrefix []byte) ([]*tss.DataPoint, error) {
	// fmt.Printf("----- scan | keyPrefix=%s\n", keyPrefix)

	r := make([]*tss.DataPoint, 0)

	if err := tsdb.db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()

		for it.Seek(keyPrefix); it.ValidForPrefix(keyPrefix); it.Next() {
			item := it.Item()
			k := item.Key()

			v, err := item.ValueCopy(nil)
			if err != nil {
				xlog.Errorf("[tss] Unable to get value for key=%s: %v", k, err)
				continue
			}

			dp, err := decodeKV(k, v)
			if err != nil {
				xlog.Errorf("[tss] Unable to decode key=%s: %v", k, err)
				continue
			}

			// fmt.Printf("====== scan | key=%s | value=%v\n", k, dp.Value)

			r = append(r, dp)

			// if err := item.Value(func(v []byte) error {
			// 	fmt.Printf("[tss] key=%s, value=%s\n", k, v)

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

func (tsdb *tsDB) Last(keyPrefix []byte) (*tss.DataPoint, error) {
	var dp *tss.DataPoint

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
				xlog.Errorf("[tss] Unable to get value for key=%s: %v", k, err)
				continue
			}

			fmt.Printf("[tss] key=%s, value=%s\n", k, v)

			dp, err = decodeKV(k, v)
			if err != nil {
				xlog.Errorf("[tss] Unable to decode key=%s: %v", k, err)
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

// runGC triggers the garbage collection for the BadgerDB backend database. It
// should be run in a goroutine.
func (tsdb *tsDB) runGC() {
	ticker := time.NewTicker(badgerGCInterval)
	for {
		select {
		case <-ticker.C:
			if err := tsdb.db.RunValueLogGC(badgerDiscardRatio); err != nil {
				// don't report error when GC didn't result in any cleanup
				if err == badger.ErrNoRewrite {
					xlog.Debugf("[tss] No BadgerDB GC occurred: %v", err)
				} else {
					xlog.Errorf("[tss] Failed to GC BadgerDB: %v", err)
				}
			}

		case <-tsdb.gcCloseCh:
			return
		}
	}
}
