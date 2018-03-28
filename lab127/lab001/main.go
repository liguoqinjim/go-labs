package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

func main() {
	//open the database
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatalf("open db error:%v", err)
	}
	defer db.Close()

	//using buckets
	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte("MyBucket"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		_ = b

		return nil
	})

	//set key/value
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		err := b.Put([]byte("answer"), []byte("42"))
		return err
	})

	//get key/value
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		v := b.Get([]byte("answer"))
		log.Printf("answer=%s", v)
		return nil
	})

	//delete key
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		err := b.Delete([]byte("answer"))
		return err
	})

	//get deleted key
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		v := b.Get([]byte("answer"))
		if v == nil {
			log.Println("key not exist")
		} else {
			log.Println("key exist")
		}

		return nil
	})
}
