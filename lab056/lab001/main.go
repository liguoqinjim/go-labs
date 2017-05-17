package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"lab056/lab001/conf"
)

func main() {
	conf.ReadConf()

	client := redis.NewClient(&redis.Options{
		Addr:     conf.ConnConf.Addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	//ping
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	//set get
	err = client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}
	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}
