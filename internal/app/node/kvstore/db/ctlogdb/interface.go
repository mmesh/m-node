package ctlogdb

import (
	"fmt"

	"github.com/dgraph-io/badger/v4"
	"mmesh.dev/m-api-go/grpc/resources/nstore/netdb"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/xlog"
)

const ctLogPrefix string = "ctlog"

var InputQueue = make(chan *netdb.ConntrackLogEntry, 128)
var RequestQueue = make(chan *netdb.ConntrackLogRequest, 128)

type Interface interface {
	Set(ctLogEntry *netdb.ConntrackLogEntry) error
	Scan() ([]*netdb.ConntrackLogEntry, error)
	Last(n int) ([]*netdb.ConntrackLogEntry, error)
	Query(r *netdb.ConntrackLogRequest) (*netdb.ConntrackLogResponse, error)
	Close()
}
type ctlogDB struct {
	db *badger.DB
}

func Open(db *badger.DB) Interface {
	ndb := &ctlogDB{
		db: db,
	}

	return ndb
}

func (ndb *ctlogDB) Close() {
}

func (ndb *ctlogDB) Set(ctLogEntry *netdb.ConntrackLogEntry) error {
	if err := ndb.db.Update(func(txn *badger.Txn) error {
		e := &conntrackLogEntry{ctLogEntry}

		dbEntry, err := e.newEntry()
		if err != nil {
			return errors.Wrapf(err, "[%v] function e.newEntry()", errors.Trace())
		}

		if err := txn.SetEntry(dbEntry); err != nil {
			return errors.Wrapf(err, "[%v] function txn.SetEntry()", errors.Trace())
		}

		return nil
	}); err != nil {
		return errors.Wrapf(err, "[%v] function ndb.db.Update()", errors.Trace())
	}

	return nil
}

func (ndb *ctlogDB) Scan() ([]*netdb.ConntrackLogEntry, error) {
	keyPrefix := []byte(fmt.Sprintf("%s:", ctLogPrefix))

	r := make([]*netdb.ConntrackLogEntry, 0)

	if err := ndb.db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()

		for it.Seek(keyPrefix); it.ValidForPrefix(keyPrefix); it.Next() {
			item := it.Item()
			k := item.Key()

			v, err := item.ValueCopy(nil)
			if err != nil {
				xlog.Errorf("[ctlogdb] Unable to get value for key=%s: %v", k, err)
				continue
			}

			e, err := decodeKV(k, v)
			if err != nil {
				xlog.Errorf("[ctlogdb] Unable to decode key=%s: %v", k, err)
				continue
			}

			// fmt.Printf("====== scan | key=%s | value=%v\n", k, dp.Value)

			r = append(r, e)
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[%v] function ndb.db.View()", errors.Trace())
	}

	return r, nil
}

func (ndb *ctlogDB) Last(n int) ([]*netdb.ConntrackLogEntry, error) {
	keyPrefix := []byte(ctLogPrefix)

	r := make([]*netdb.ConntrackLogEntry, n)

	if err := ndb.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = n
		opts.Reverse = true

		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Seek(keyPrefix); it.ValidForPrefix(keyPrefix); it.Next() {
			item := it.Item()
			k := item.Key()

			v, err := item.ValueCopy(nil)
			if err != nil {
				xlog.Errorf("[ctlogdb] Unable to get value for key=%s: %v", k, err)
				continue
			}

			e, err := decodeKV(k, v)
			if err != nil {
				xlog.Errorf("[ctlogdb] Unable to decode key=%s: %v", k, err)
				continue
			}

			// fmt.Printf("[ctlogdb] key=%s, value=%s\n", k, v)

			r = append(r, e)
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[%v] function ndb.db.View()", errors.Trace())
	}

	return r, nil
}
