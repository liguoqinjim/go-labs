package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type UserInfo struct {
	Name string
	Age  int
}

func (userInfo *UserInfo) GetKey() string {
	return userInfo.Name
}

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	//sync.Map
	var m sync.Map

	u1 := &UserInfo{Name: "tom", Age: 11}
	u2 := &UserInfo{Name: "ben", Age: 12}
	u3 := &UserInfo{Name: "alice", Age: 13}

	//store
	m.Store(u1.GetKey(), u1)
	m.Store(u2.GetKey(), u2)
	m.Store(u3.GetKey(), u3)

	m.Range(func(key, value interface{}) bool {
		log.Printf("range key=%s,value=%+v", key, value)
		if key.(string) == "tom" {
			return false
		} else {
			return true
		}
	})

	go func() {
		m.Range(func(key, value interface{}) bool {
			log.Printf("range2 key=%s,value=%+v", key, value)

			time.Sleep(time.Second * 1)
			return true
		})
		log.Printf("range2 end")
	}()

	go func(m sync.Map) {
		for i := 1; i <= 5; i++ {
			u := &UserInfo{Name: fmt.Sprintf("kimi0%d", i), Age: i}
			m.Store(u.GetKey(), u)
			log.Printf("store key=%s", u.Name)
			time.Sleep(time.Second * 1)
		}
		log.Printf("store end")
	}(m)

	go func(m sync.Map) {
		names := []string{"tom", "ben", "alice"}
		for _, n := range names {
			log.Printf("delete key=%s", n)
			m.Delete(n)
			time.Sleep(time.Second)
		}
		log.Printf("delete end")
	}(m)

	<-sigs
	log.Println("end")
}
