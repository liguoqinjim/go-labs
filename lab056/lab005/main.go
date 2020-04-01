package main

import (
	"flag"
	"fmt"
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
	//demo()
	demo2()
}

//存struct
func demo() {
	key := "temp01"
	l := &Lab{
		Name: "admin",
		Age:  13,
	}

	//set
	if r, err := client.Set(key, l, 0).Result(); err != nil {
		log.Fatalf("set error:%v", err)
	} else {
		log.Printf("set result:%s", r)
	}

	//get
	if r, err := client.Get(key).Result(); err != nil {
		log.Fatalf("get error:%v", err)
	} else {
		log.Printf("get result:%s", r)
	}
}

func demo2() {
	key := "temp02"
	l := &Lab{
		Name: "tom",
		Age:  24,
	}

	//set
	if r, err := client.HSet(key, l).Result(); err != nil {
		log.Fatalf("hset error:%v", err)
	} else {
		log.Printf("hset result:%v", r)
	}
}

type Lab struct {
	Name string
	Age  int
}

func (l Lab) MarshalBinary() (data []byte, err error) {
	r := fmt.Sprintf("%s;%d", l.Name, l.Age)
	return []byte(r), nil
}
