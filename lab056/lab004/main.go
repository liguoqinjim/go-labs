package main

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"io/ioutil"
	"log"
)

const (
	KeyName = "go-labs-set"
)

func main() {
	demo()
}

func demo() {
	conf := readConf()
	if conf == nil {
		log.Fatalf("conf is nil")
	}
	log.Println(conf)

	client := redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Password: conf.Password, // no password set
		DB:       1,             // 可以指定redis使用的db
	})

	//ping
	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatalf("client.Ping error:%v", err)
	}
	log.Println("pong=", pong)

	//incr
	result, err := client.Incr(KeyName).Result()
	if err != nil {
		log.Fatalf("client.Incr error:%v", err)
	}
	log.Println("result=", result)

	result, err = client.IncrBy(KeyName, 5).Result()
	if err != nil {
		log.Fatalf("client.IncrBy error:%v", err)
	}
	log.Println("result=", result)
}

func readConf() *Conf {
	data, err := ioutil.ReadFile("../conf.json")
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
