package main

import (
	"flag"
	"github.com/bculberson/bloom"
	"github.com/garyburd/redigo/redis"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"time"
)

var (
	n    = 100000
	fp   = 0.001
	info string
)

func init() {
	flag.String("redis", "127.0.0.1:6379", "redis connection info")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	info = viper.GetString("redis") // retrieve value from viper
}

func main() {
	//demo1()

	demo2()
}

func demo1() {
	//redis连接池
	pool := &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", info) },
	}

	//创建过滤器
	conn := pool.Get()
	m, k := bloom.EstimateParameters(uint(n), fp)
	//bf_prefix_是在redis的key的name的前缀
	bitSet := bloom.NewRedisBitSet("bf_prefix_", m, conn)
	filter := bloom.New(m, k, bitSet)

	filter.Add([]byte("hello"))
	filter.Add([]byte("world"))

	//判断是否存在
	if ok, err := filter.Exists([]byte("hello")); err != nil {
		log.Fatalf("filter.Exists error:%v", err)
	} else {
		if ok {
			log.Println("hello is in the filter")
		} else {
			log.Println("hello not in the filter")
		}
	}
}

func demo2() {
	//redis连接池
	pool := &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", info) },
	}

	//创建过滤器
	conn := pool.Get()
	m, k := bloom.EstimateParameters(uint(n), fp)
	//bf_prefix_是在redis的key的name的前缀
	bitSet := bloom.NewRedisBitSet("bf_prefix_", m, conn)
	filter := bloom.New(m, k, bitSet)

	//判断是否存在
	if ok, err := filter.Exists([]byte("hello")); err != nil {
		log.Fatalf("filter.Exists error:%v", err)
	} else {
		if ok {
			log.Println("hello is in the filter")
		} else {
			log.Println("hello not in the filter")
		}
	}
}
