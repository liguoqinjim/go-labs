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
	log.Println(conf)

	client := redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Password: conf.Password, // no password set
		DB:       conf.DB,       // use default DB
	})

	//ping
	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatalf("client.Ping error:%v", err)
	}
	log.Println("pong=", pong)

	//set get
	err = client.Set("key", "value123", 0).Err()
	if err != nil {
		log.Fatalf("client.Set error:%v", err)
	}
	val, err := client.Get("key").Result()
	if err != nil {
		log.Fatalf("client.Get error:%v", err)
	}
	log.Println("get key=", val)

	//get key not exist
	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		log.Println("key2 does not exist")
	} else if err != nil {
		log.Fatalf("client.Get error:%v", err)
	} else {
		log.Println("key2", val2)
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
