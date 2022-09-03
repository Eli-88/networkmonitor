package db

import (
	"github.com/dgraph-io/badger/v3"
)

var _ KvDb = &badgerDb{}

type badgerDb struct {
	db *badger.DB
}

type keyvaluePair struct {
	key   []byte
	value []byte
}

func (k keyvaluePair) Key() []byte {
	return k.key
}

func (k keyvaluePair) Value() []byte {
	return k.value
}

func MakeBadgerDb(path string) (KvDb, error) {
	options := badger.DefaultOptions(path)
	options.Logger = nil // TODO: consider making a logger adapter for badger
	db, err := badger.Open(options)
	if err != nil {
		return nil, err
	}
	return &badgerDb{db: db}, nil
}

func (b *badgerDb) GetKvValue(key []byte) ([]byte, error) {
	var result []byte
	err := b.db.Update(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}
		result, err = item.ValueCopy(result)
		return err
	})

	if err != nil {
		return nil, err
	}
	return result, nil
}

func (b *badgerDb) GetAllKeyValue(applyFn ApplyFn) error {
	err := b.db.View(func(txn *badger.Txn) error {
		iter := txn.NewIterator(badger.DefaultIteratorOptions)
		defer iter.Close()
		for iter.Rewind(); iter.Valid(); iter.Next() {
			var key []byte
			key = iter.Item().KeyCopy(key)

			var value []byte

			value, err := iter.Item().ValueCopy(value)

			if err != nil {
				return err
			}

			// user may end the iteration prematurely by returning false
			if !applyFn(key, value) {
				return nil
			}
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

func (b *badgerDb) SetKvKeyValue(key []byte, value []byte) error {
	err := b.db.Update(func(txn *badger.Txn) error {
		return txn.Set(key, value)
	})
	return err
}

func (b *badgerDb) Close() {
	b.db.Close()
}
