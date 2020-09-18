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
		Password: password,
		DB:       db,
	})
}

func main() {
	log.Println(client.Ping())

	//prefixKey()
	prefixKeyCount()
}

func prefixKey() {
	result := client.Keys("u:*")
	log.Println(result)
}

func prefixKeyCount() {
	//eval "return #redis.call('keys', 'prefix-*')" 0
	if result, err := client.Eval("return #redis.call('keys', 'u:*')", nil).Int(); err != nil {
		log.Fatalf("eval error:%v", err)
	} else {
		log.Println(result)
	}
}
