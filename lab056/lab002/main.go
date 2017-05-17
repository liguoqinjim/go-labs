package main

import (
	"github.com/go-redis/redis"
	"lab056/lab002/conf"
	"fmt"
)

func main() {
	conf.ReadConf()

	client := redis.NewClient(&redis.Options{
		Addr:conf.ConnConf.Addr,
	})

	Get := func(client *redis.Client,key string) *redis.StringCmd{
		cmd := redis.NewStringCmd("get",key)
		client.Process(cmd)
		return cmd
	}

	v,err := Get(client,"key2").Result()
	fmt.Println(v,err)
}
