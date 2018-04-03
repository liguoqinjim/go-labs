package main

import (
	"log"
	"sync"
)

type UserInfo struct {
	Name string
	Age  int
}

func (userInfo *UserInfo) GetKey() string {
	return userInfo.Name
}

func main() {
	//sync.Map
	var m sync.Map

	u1 := &UserInfo{Name: "tom", Age: 11}

	//store
	m.Store(u1.GetKey(), u1)

	//load
	u, ok := m.Load("tom")
	if !ok { //false的时候是没有找到
		log.Println("load key tom not found")
	} else {
		log.Printf("load key tom found:%+v", u.(*UserInfo))
	}

	//LoadOrStore
	u2 := &UserInfo{Name: "ben", Age: 12}
	vv, ok := m.LoadOrStore(u2.GetKey(), u2)
	if !ok {
		log.Printf("loadOrStore key ben stored:%+v", vv)
	} else {
		log.Printf("loadOrStore key ben loaded:%+v", vv)
	}

	//delete
	u3 := &UserInfo{Name: "alice", Age: 13}
	m.Store("alice", u3)
	m.Delete("ben")

	//range
	m.Range(func(key, value interface{}) bool {
		log.Printf("range key=%s,value=%+v", key, value.(*UserInfo))
		return true
	})
}
