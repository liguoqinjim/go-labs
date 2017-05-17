package main

import (
	"fmt"
	"lab056/lab002/conf"

	"github.com/go-redis/redis"
)

func main() {
	conf.ReadConf()

	client := redis.NewClient(&redis.Options{
		Addr: conf.ConnConf.Addr,
	})

	Get := func(client *redis.Client, key string) *redis.StringCmd {
		cmd := redis.NewStringCmd("get", key)
		client.Process(cmd)
		return cmd
	}

	v, err := Get(client, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key dose not exits")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("v", v)
	}
}
