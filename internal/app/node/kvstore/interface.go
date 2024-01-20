package kvstore

import (
	"os"
	"time"

	"github.com/dgraph-io/badger/v4"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/xlog"
	"mmesh.dev/m-node/internal/app/node/kvstore/db/ctlogdb"
	"mmesh.dev/m-node/internal/app/node/kvstore/db/metricsdb"
	"mmesh.dev/m-node/internal/app/node/kvstore/db/netflowdb"
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

type Interface interface {
	HostMetrics() metricsdb.Interface
	NetCtLog() ctlogdb.Interface
	Netflow() netflowdb.Interface
	Close() error

	runGC()
}

type kvStore struct {
	db        *badger.DB
	gcCloseCh chan struct{}

	hostMetrics metricsdb.Interface
	netCtLog    ctlogdb.Interface
	netflow     netflowdb.Interface
}

func Open() (Interface, error) {
	if err := os.MkdirAll(dbDir(), 0754); err != nil {
		return nil, errors.Wrapf(err, "[%v] function os.MkdirAll()", errors.Trace())
	}

	opts := badger.DefaultOptions(dbDir())
	opts.SyncWrites = true
	opts.Dir, opts.ValueDir = dbDir(), dbDir()
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

	kvs := &kvStore{
		db:          db,
		gcCloseCh:   make(chan struct{}, 1),
		hostMetrics: metricsdb.Open(db),
		netCtLog:    ctlogdb.Open(db),
		netflow:     netflowdb.Open(db),
	}

	go kvs.runGC()

	return kvs, nil
}

func (kvs *kvStore) HostMetrics() metricsdb.Interface {
	return kvs.hostMetrics
}

func (kvs *kvStore) NetCtLog() ctlogdb.Interface {
	return kvs.netCtLog
}

func (kvs *kvStore) Netflow() netflowdb.Interface {
	return kvs.netflow
}

func (kvs *kvStore) Close() error {
	kvs.hostMetrics.Close()
	kvs.netCtLog.Close()
	kvs.netflow.Close()

	kvs.gcCloseCh <- struct{}{}

	return kvs.db.Close()
}

// runGC triggers the garbage collection for the BadgerDB backend database. It
// should be run in a goroutine.
func (kvs *kvStore) runGC() {
	ticker := time.NewTicker(badgerGCInterval)
	for {
		select {
		case <-ticker.C:
			if err := kvs.db.RunValueLogGC(badgerDiscardRatio); err != nil {
				// don't report error when GC didn't result in any cleanup
				if err == badger.ErrNoRewrite {
					xlog.Debugf("[kvstore] No BadgerDB GC occurred: %v", err)
				} else {
					xlog.Errorf("[kvstore] Failed to GC BadgerDB: %v", err)
				}
			}
		case <-kvs.gcCloseCh:
			return
		}
	}
}
