package main

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"io/ioutil"
	"log"
)

func main() {
	example()
}

func example() {
	conf := readConf()
	if conf == nil {
		log.Fatalf("conf is nil")
	}

	client := redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Password: conf.Password,
		DB:       conf.DB,
	})

	Get := func(client *redis.Client, key string) *redis.StringCmd {
		cmd := redis.NewStringCmd("get", key)
		client.Process(cmd)
		return cmd
	}

	v, err := Get(client, "key2").Result()
	if err == redis.Nil {
		log.Println("key dose not exits")
	} else if err != nil {
		log.Fatalf("Get error:%v", err)
	} else {
		log.Println("v", v)
	}
}

func readConf() *Conf {
	data, err := ioutil.ReadFile("conf.json")
	if err != nil {
		log.Fatalf("ioutil.ReadFile error:%v", err)
	}

	var conf = &Conf{}
	if err := json.Unmarshal(data, conf); err != nil {
		log.Fatalf("json.Unmarshal error:%v", err)
	}

	return conf
}

type Conf struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DB       int    `json:"DB"`
}
