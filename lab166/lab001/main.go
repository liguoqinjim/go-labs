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

	//创建bloomFilter
	conn := pool.Get()
	m, k := bloom.EstimateParameters(1000, .01)
	//test_key是在redis的key的name的前缀
	bitSet := bloom.NewRedisBitSet("test_key", m, conn)
	b := bloom.New(m, k, bitSet)

	//判断是否exist
	b.Add([]byte("key1"))
	exists, _ := b.Exists([]byte("key1"))
	log.Println("key exists:", exists)
	doesNotExist, _ := b.Exists([]byte("key2"))
	log.Println("key2 exists:", doesNotExist)
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
