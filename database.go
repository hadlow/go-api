package main

import (
	bolt "go.etcd.io/bbolt"
)

type Database struct {
	connection *bolt.DB
	bucket []byte
}

func NewDatabase(dbPath string) (db *Database, closeCallback func() error, err error) {
	// Open the database connection
	connection, err := bolt.Open(dbPath, 0600, nil)

	if err != nil {
		return nil, nil, err
	}

	// Return the databsae object
	db = &Database{connection: connection, bucket: []byte("default")}
	closeCallback = connection.Close

	return
}

func (d *Database) SetBucket(bucket string) error {
	err := d.connection.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucket))

		return err
	})

	if err != nil {
		return err
	}

	d.bucket = []byte(bucket)

	return nil
}

func (d *Database) Get(key string) (value []byte, err error) {
	err = d.connection.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(d.bucket)

		value = b.Get([]byte(key))

		return nil
	})

	return
}

func (d *Database) Set(key string, value []byte) error {
	return d.connection.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(d.bucket)

		return b.Put([]byte(key), value)
	})
}

func (d *Database) Delete(key string) error {
	return d.connection.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(d.bucket)

		return b.Delete([]byte(key))
	})
}

func (d *Database) BulkDelete(keys []string) error {
	return d.connection.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(d.bucket)

		for _, key := range keys {
			err := b.Delete([]byte(key))

			if err != nil {
				return err
			}
		}

		return nil
	})
}