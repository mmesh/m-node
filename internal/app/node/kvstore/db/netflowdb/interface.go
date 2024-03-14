package netflowdb

import (
	"fmt"

	"github.com/dgraph-io/badger/v4"
	"mmesh.dev/m-api-go/grpc/resources/nstore/netdb"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/xlog"
)

const netflowPrefix string = "netflow"
const maxEntries = 30

var InputQueue = make(chan []*netdb.NetFlowEntry, 128)
var RequestQueue = make(chan *netdb.TrafficMetricsRequest, 128)

type Interface interface {
	Query(r *netdb.TrafficMetricsRequest) (*netdb.TrafficMetricsResponse, error)
	WriteBatch(nfel []*netdb.NetFlowEntry) error
	Scan() ([]*netdb.NetFlowEntry, error)
	Close()
}
type netflowDB struct {
	db *badger.DB
}

func Open(db *badger.DB) Interface {
	nfdb := &netflowDB{
		db: db,
	}

	return nfdb
}

func (nfdb *netflowDB) Close() {
}

func (nfdb *netflowDB) WriteBatch(nfel []*netdb.NetFlowEntry) error {
	wb := nfdb.db.NewWriteBatch()
	defer wb.Cancel()

	for _, nfe := range nfel {
		nf := &netflowEntry{nfe}

		e, err := nf.newEntry()
		if err != nil {
			return errors.Wrapf(err, "[%v] function nf.newEntry()", errors.Trace())
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

func (nfdb *netflowDB) Scan() ([]*netdb.NetFlowEntry, error) {
	keyPrefix := []byte(fmt.Sprintf("%s:", netflowPrefix))

	// fmt.Printf("----- scan | keyPrefix=%s\n", keyPrefix)

	r := make([]*netdb.NetFlowEntry, 0)

	if err := nfdb.db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()

		for it.Seek(keyPrefix); it.ValidForPrefix(keyPrefix); it.Next() {
			item := it.Item()
			k := item.Key()

			v, err := item.ValueCopy(nil)
			if err != nil {
				xlog.Errorf("[netflowdb] Unable to get value for key=%s: %v", k, err)
				continue
			}

			nfe, err := decodeKV(k, v)
			if err != nil {
				xlog.Errorf("[netflowdb] Unable to decode key=%s: %v", k, err)
				continue
			}

			// fmt.Printf("====== scan | key=%s | value=%v\n", k, dp.Value)

			r = append(r, nfe)

			// if err := item.Value(func(v []byte) error {
			// 	fmt.Printf("[netflowdb] key=%s, value=%s\n", k, v)

			// 	nfe, err := decodeKV(k, v)
			// 	if err != nil {
			// 		return errors.Wrapf(err, "[%v] function decodeKV()", errors.Trace())
			// 	}

			// 	r = append(r, nfe)

			// 	return nil
			// }); err != nil {
			// 	return errors.Wrapf(err, "[%v] function item.Value()", errors.Trace())
			// }
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[%v] function nfdb.db.View()", errors.Trace())
	}

	return r, nil
}
