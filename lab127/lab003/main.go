package main

import (
	"github.com/boltdb/bolt"
	"log"
)

func main() {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatalf("open error:%v", err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte("MyBucket"))
		if err != nil {
			return err
		}

		b.Put([]byte("one"), []byte("1"))
		b.Put([]byte("two"), []byte("2"))
		b.Put([]byte("three"), []byte("3"))
		b.Put([]byte("four"), []byte("4"))
		b.Put([]byte("five"), []byte("5"))

		return nil
	})

	//遍历
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))

		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			log.Printf("key=%s,value=%s", k, v)
		}

		return nil
	})
}
