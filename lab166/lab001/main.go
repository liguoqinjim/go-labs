package main

import (
	"github.com/bculberson/bloom"
	"github.com/garyburd/redigo/redis"
	"time"
)

func main() {
	pool := &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", addr) },
	}


	conn := pool.Get()
	m, k := bloom.EstimateParameters(1000, .01)
	bitSet := bloom.NewRedisBitSet("test_key", m, conn)
	b := bloom.New(m, k, bitSet)
	b.Add([]byte("some key"))
	exists, _ := b.Exists([]byte("some key"))
	doesNotExist, _ := b.Exists([]byte("some other key"))
}
