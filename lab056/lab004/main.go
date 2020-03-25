package main

import (
	"flag"
	"github.com/go-redis/redis/v7"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	address  string
	password string
	db       int
	client   *redis.Client
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
	//demo()

	//concurrent()

	keyNotExist()
}

func demo() {
	key := "key_incr"

	//incr
	if r, err := client.Incr(key).Result(); err != nil {
		log.Fatalf("incr error:%v", err)
	} else {
		log.Println("incr result:", r)
	}

	//incrby
	if r, err := client.IncrBy(key, 5).Result(); err != nil {
		log.Fatalf("incrby error:%v", err)
	} else {
		log.Println("incrby result:", r)
	}

	//get
	if r, err := client.Get(key).Result(); err != nil {
		log.Fatalf("get error:%v", err)
	} else {
		log.Println("get result:", r)
	}
}

func concurrent() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	key := "key_incr2"

	for i := 0; i < 100; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				if _, err := client.Incr(key).Result(); err != nil {
					log.Fatalf("incr error:%v", err)
				}
			}
		}()
	}

	log.Println("key=", client.Get(key).String())

	<-sigs
}

func keyNotExist() {
	if r, err := client.FlushDB().Result(); err != nil {
		log.Fatalf("flushdb error:%v", err)
	} else {
		log.Println("flushdb result:", r)
	}

	key := "not_exist_key"

	if r, err := client.Incr(key).Result(); err != nil {
		log.Fatalf("incr error:%v", err)
	} else {
		log.Println("incr result:", r)
	}
}
