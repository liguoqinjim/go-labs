package main

import (
	"flag"
	"github.com/go-redis/redis/v7"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
)

var (
	address  string
	password string
	db       int
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
}

func main() {
	example()
}

func example() {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password, // no password set
		DB:       db,       // use default DB
	})

	//custom cmd
	Get := func(client *redis.Client, key string) *redis.StringCmd {
		cmd := redis.NewStringCmd("get", key)
		client.Process(cmd)
		return cmd
	}

	//call custom cmd 01
	if v, err := Get(client, "key2").Result(); err != nil {
		if err == redis.Nil {
			log.Println("key2 not exist")
		} else {
			log.Fatalf("Get error:%v", err)
		}
	} else {
		log.Println("key2=", v)
	}

	//call custom cmd 02
	if v, err := client.Do("get", "key2").Result(); err != nil {
		if err == redis.Nil {
			log.Println("key2 not exist")
		} else {
			log.Fatalf("Get error:%v", err)
		}
	} else {
		log.Println("key2=", v.(string))
	}
}
