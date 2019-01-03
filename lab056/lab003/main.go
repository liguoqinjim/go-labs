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
		DB:       conf.DB,       // use default DB
	})

	//ping
	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatalf("client.Ping error:%v", err)
	}
	log.Println("pong=", pong)

	//添加到set
	r, err := client.SAdd("", 1, 2, 3).Result()
	if err != nil {
		log.Fatalf("client.SAdd error:%v", err)
	}
	log.Printf("client.SAdd result:%d", r)

	//添加一个重复的
	r, err = client.SAdd(KeyName, 1).Result()
	if err != nil {
		log.Fatalf("client.SAdd error:%v", err)
	}
	log.Printf("client.SAdd result:%d", r)

	//添加一个不重复的
	r, err = client.SAdd(KeyName, 5).Result()
	if err != nil {
		log.Fatalf("client.SAdd error:%v", err)
	}
	log.Printf("client.SAdd result:%d", r)
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
