package main

import (
	"flag"
	"github.com/go-redis/redis/v7"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"time"
)

var (
	address  string
	password string
	db       int

	client *redis.Client
)

func init() {
	pflag.StringVarP(&address, "address", "a", "localhost:6379", "redis address")
	pflag.StringVarP(&password, "password", "p", "", "redis auth")
	pflag.IntVarP(&db, "db", "d", 0, "db")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		log.Fatalf("viper.BindPFlags error:%v", err)
	}

	client = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password, // no password set
		DB:       db,       // use default DB
	})
}

func main() {
	example()

	dateKey()
}

func example() {
	//ping
	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatalf("client.Ping error:%v", err)
	}
	log.Println("pong=", pong)

	//set
	if err = client.Set("key", "1", 0).Err(); err != nil {
		log.Fatalf("client.Set error:%v", err)
	}

	//get
	if val, err := client.Get("key").Result(); err != nil {
		log.Fatalf("client.Get error:%v", err)
	} else {
		log.Println("get key=", val)
	}

	//get key not exist
	if val2, err := client.Get("key2").Result(); err != nil {
		if err == redis.Nil {
			log.Println("key2 not exist")
		} else {
			log.Fatalf("redis error:%v", err)
		}
	} else {
		log.Println("get key2=", val2)
	}
}

//直接存time.Now()
func dateKey() {
	key := "date1"
	if r, err := client.Set(key, time.Now(), 0).Result(); err != nil {
		log.Fatalf("client.Set error:%v", err)
	} else {
		log.Println("set result:", r)
	}

	//string
	if r, err := client.Get(key).Result(); err != nil {
		log.Fatalf("client.Get.Result error:%v", err)
	} else {
		log.Println("get result:", r)
	}

	//time
	if r, err := client.Get(key).Time(); err != nil {
		log.Fatalf("client.Get.Time error:%v", err)
	} else {
		log.Println("get result:", r)
	}
}
