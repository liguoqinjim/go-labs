package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

func main() {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatalf("open error:%v", err)
	}
	defer db.Close()

	//创建bucket
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("users"))
		return err
	})

	//sequence
	u1 := &User{Username: "xiaoming"}
	u2 := &User{Username: "xiaohong"}
	err = CreateUser(db, u1)
	if err != nil {
		log.Printf("CreateUser error:%v", err)
	}

	err = CreateUser(db, u2)
	if err != nil {
		log.Printf("CreateUser error:%v", err)
	}

	//读取
	err = GetUser(db, 2)
	if err != nil {
		log.Printf("GetUser userid:%d error:%v", 2, err)
	}
	err = GetUser(db, 3)
	if err != nil {
		log.Printf("GetUser userid:%d error:%v", 3, err)
	}
}

func CreateUser(db *bolt.DB, u *User) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("users"))

		id, _ := b.NextSequence()
		u.ID = int(id)
		log.Printf("CreateUser UserId:%d", u.ID)

		buf, err := json.Marshal(u)
		if err != nil {
			return err
		}

		return b.Put(itob(u.ID), buf)
	})
}

func GetUser(db *bolt.DB, id int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("users"))

		u := b.Get(itob(id))
		if u == nil {
			log.Printf("GetUser userid:%d not found", id)
			return fmt.Errorf("not found userId:%d", id)
		} else {
			user := &User{}
			err := json.Unmarshal(u, user)
			if err != nil {
				return err
			}

			log.Printf("GetUser userid:%d username:%s", id, user.Username)
			return nil
		}
	})
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
