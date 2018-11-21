package main

import (
	"encoding/json"
	"github.com/bculberson/bloom"
	"github.com/garyburd/redigo/redis"
	"io/ioutil"
	"log"
	"time"
)

func main() {
	example()
}

func example() {
	pool := &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", readAddr()) },
	}

	conn := pool.Get()
	m, k := bloom.EstimateParameters(1000, .01)
	bitSet := bloom.NewRedisBitSet("test_key", m, conn)
	b := bloom.New(m, k, bitSet)
	b.Add([]byte("some key"))
	exists, _ := b.Exists([]byte("some key"))
	log.Println(exists)
	doesNotExist, _ := b.Exists([]byte("some other key"))
	log.Println(doesNotExist)

	b.Add([]byte("helloworld"))
	log.Println(b.Exists([]byte("helloworld")))

	b.Add([]byte("helloworld1"))
	log.Println(b.Exists([]byte("helloworld1")))

	b.Add([]byte("iris"))
	log.Println(b.Exists([]byte("iris")))
}

func readAddr() string {
	data, err := ioutil.ReadFile("../conf.json")
	if err != nil {
		log.Fatalf("ioutil.ReadFile error:%v", err)
	}

	var conf = &Conf{}
	if err := json.Unmarshal(data, conf); err != nil {
		log.Fatalf("json.Unmarshal error:%v", err)
	}

	return conf.Addr
}

type Conf struct {
	Addr string `json:"addr"`
}
