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
	//set()
	//hash()
	list()
}

//redis list
func list() {
	key := "key_list"

	//list加入元素
	if r, err := client.LPush(key, "redis").Result(); err != nil {
		log.Fatalf("lpush error:%v", err)
	} else {
		log.Println("lpush result:", r)
	}

	//lpush
	if r, err := client.LPush(key, "mongodb", "mysql").Result(); err != nil {
		log.Fatalf("lpush error:%v", err)
	} else {
		log.Println("lpush result:", r)
	}

	//rpush
	if r, err := client.RPush(key, "pg").Result(); err != nil {
		log.Fatalf("rpush error:%v", err)
	} else {
		log.Println("rpush result:", r)
	}

	//在列表的元素前或者后插入元素
	if r, err := client.LInsert(key, "before", "mysql", "db2").Result(); err != nil {
		log.Fatalf("linsert error:%v", err)
	} else {
		log.Println("linsert result:", r)
	}
	if r, err := client.LInsertBefore(key, "mysql", "oracle").Result(); err != nil {
		log.Fatalf("linsert error:%v", err)
	} else {
		log.Println("linsert result:", r)
	}

	//有多少个元素
	var length int64
	if r, err := client.LLen(key).Result(); err != nil {
		log.Fatalf("llen error:%v", err)
	} else {
		log.Println("llen result:", r)
		length = r
	}

	//得到范围内的元素
	if r, err := client.LRange(key, 0, length).Result(); err != nil {
		log.Fatalf("lrange error:%v", err)
	} else {
		log.Println("lrange result:", r)
	}

	//修改index的元素
	if r, err := client.LSet(key, 2, "mariadb").Result(); err != nil {
		log.Fatalf("lset error:%v", err)
	} else {
		log.Println("lset result:", r)
	}

	//得到index的元素
	if r, err := client.LIndex(key, 2).Result(); err != nil {
		log.Fatalf("lindex error:%v", err)
	} else {
		log.Println("lindex result:", r)
	}
}

//redis hash
func hash() {
	key := "key_hash"

	//hash加入元素(键值对)
	if r, err := client.HMSet(key, "name", "test", "type", "hash", "lab", "056").Result(); err != nil {
		log.Fatalf("hmset error:%v", err)
	} else {
		log.Println("hmset result:", r)
	}

	//加入错误长度的参数 (ERR wrong number of arguments for HMSET)
	if r, err := client.HMSet(key, "sub_lab").Result(); err != nil {
		log.Println("hmset error:", err)
	} else {
		log.Println("hmset result:", r)
	}

	//获取hash所有字段和值
	if r, err := client.HGetAll(key).Result(); err != nil {
		log.Fatalf("hgetall error:%v", err)
	} else {
		log.Println("hgetall result:", r)
	}

	//获取hash的所有字段
	if r, err := client.HKeys(key).Result(); err != nil {
		log.Fatalf("hkeys error:%v", err)
	} else {
		log.Println("hkeys result:", r)
	}

	//获取hash的所有值
	if r, err := client.HVals(key).Result(); err != nil {
		log.Fatalf("hvals error:%v", err)
	} else {
		log.Println("hvals result:", r)
	}

	//获取hash的某个字段的值
	if r, err := client.HGet(key, "name").Result(); err != nil {
		log.Fatalf("hget error:%v", err)
	} else {
		log.Println("hget result:", r)
	}

	//查看hash有多少个字段
	if r, err := client.HLen(key).Result(); err != nil {
		log.Fatalf("hlen error:%v", err)
	} else {
		log.Println("hlen result:", r)
	}
}

//redis集合
func set() {
	key := "key_set"

	//添加元素到set
	if r, err := client.SAdd(key, 1, 2, 3, 4, 5).Result(); err != nil {
		log.Fatalf("sadd error:%v", err)
	} else {
		log.Println("sadd result:", r)
	}

	//添加重复元素 (不会报错，但是result为0)
	if r, err := client.SAdd(key, 1).Result(); err != nil {
		log.Fatalf("sadd error:%v", err)
	} else {
		log.Println("sdd result:", r)
	}

	//查看元素个数
	if r, err := client.SCard(key).Result(); err != nil {
		log.Fatalf("scard error:%v", err)
	} else {
		log.Println("scard result:", r)
	}

	//删除元素
	if r, err := client.SRem(key, 3).Result(); err != nil {
		log.Fatalf("srem error:%v", err)
	} else {
		log.Println("srem result:", r)
	}

	//判断元素是否存在
	if r, err := client.SIsMember(key, 3).Result(); err != nil {
		log.Fatalf("sismember error:%v", err)
	} else {
		log.Println("sismember result:", r)
	}

	//随机移除元素
	if r, err := client.SPop(key).Result(); err != nil {
		log.Fatalf("spop error:%v", err)
	} else {
		log.Println("spop result:", r)
	}

	//查看所有元素
	if r, err := client.SMembers(key).Result(); err != nil {
		log.Fatalf("smember error:%v", err)
	} else {
		log.Println("smember result:", r)
	}
}
