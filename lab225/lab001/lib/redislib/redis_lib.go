/**
* Created by GoLand.
* User: link1st
* Date: 2019-07-25
* Time: 14:18
 */

package redislib

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"log"
)

var (
	client *redis.Client
)

func ExampleNewClient() {
	client = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		//Password:     "",
		DB:           1,
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConns"),
	})

	pong, err := client.Ping().Result()
	log.Println("初始化redis:", pong, err)
	// Output: PONG <nil>
}

func GetClient() (c *redis.Client) {
	return client
}
