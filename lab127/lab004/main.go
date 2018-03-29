package main

import (
	"bytes"
	"github.com/boltdb/bolt"
	"log"
)

func main() {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatalf("open error:%v", err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte("MyBucket"))
		if err != nil {
			return err
		}

		b.Put([]byte("1231"), []byte("1231"))
		b.Put([]byte("1232"), []byte("1232"))
		b.Put([]byte("1233"), []byte("1233"))
		b.Put([]byte("1211"), []byte("1211"))

		return nil
	})
	if err != nil {
		log.Fatalf("db.Update error:%v", err)
	}

	//prefix scan
	db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket([]byte("MyBucket")).Cursor()

		prefix := []byte("123")
		for k, v := c.Seek(prefix); k != nil && bytes.HasPrefix(k, prefix); k, v = c.Next() {
			log.Printf("key=%s value=%s", k, v)
		}

		return nil
	})

	//range scan
	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte("Event"))
		if err != nil {
			return err
		}

		b.Put([]byte("1990-01-01T00:00:00Z"), []byte("1"))
		b.Put([]byte("1991-01-01T00:00:00Z"), []byte("2"))
		b.Put([]byte("1992-01-01T00:00:00Z"), []byte("3"))
		b.Put([]byte("2000-01-01T00:00:00Z"), []byte("4"))
		return nil
	})

	db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket([]byte("Event")).Cursor()

		min := []byte("1990-01-01T00:00:00Z")
		max := []byte("1999-01-01T00:00:00Z")

		// Iterate over the 90's.
		for k, v := c.Seek(min); k != nil && bytes.Compare(k, max) <= 0; k, v = c.Next() {
			log.Printf("%s: %s\n", k, v)
		}

		return nil
	})

	//forEach
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))

		b.ForEach(func(k, v []byte) error {
			log.Printf("key=%s value=%s", k, v)
			return nil
		})

		return nil
	})
}
