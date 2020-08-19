package main

import (
	"bytes"
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

	limitLua()
}

/*
try {
	String luaScript = buildLuaScript();
	RedisScript<Number> redisScript = new DefaultRedisScript<>(luaScript, Number.class);
	Number count = limitRedisTemplate.execute(redisScript, keys, limitCount, limitPeriod);
	logger.info("Access try count is {} for name={} and key = {}", count, name, key);
	if (count != null && count.intValue() <= limitCount) {
		return pjp.proceed();
	} else {
		throw new RuntimeException("You have been dragged into the blacklist");
	}
} catch (Throwable e) {
	if (e instanceof RuntimeException) {
		throw new RuntimeException(e.getLocalizedMessage());
	}
	throw new RuntimeException("server exception");
}
*/

func limitLua() {
	script := luaScript()
	//第二个参数是限流的种类，可以是IP，或者是别的
	keys := []string{"1.2.3.4", "IP"}

	//多少秒内的最大次数
	var limitCount int64 = 5
	var limitPeriod int64 = 60

	if result, err := client.Eval(script, keys, limitCount, limitPeriod).Result(); err != nil {
		log.Fatalf("client.Eval error:%v", err)
	} else {
		count := result.(int64)
		if count <= limitCount {
			log.Printf("success,now times[%d/%d]", count, limitCount)
		} else {
			log.Println("You have been dragged into the blacklist")
		}
	}
}

/*
public String buildLuaScript() {
	StringBuilder lua = new StringBuilder();
	lua.append("local c");
	lua.append("\nc = redis.call('get',KEYS[1])");
	// 调用不超过最大值，则直接返回
	lua.append("\nif c and tonumber(c) > tonumber(ARGV[1]) then");
	lua.append("\nreturn c;");
	lua.append("\nend");
	// 执行计算器自加
	lua.append("\nc = redis.call('incr',KEYS[1])");
	lua.append("\nif tonumber(c) == 1 then");
	// 从第一次调用开始限流，设置对应键值的过期
	lua.append("\nredis.call('expire',KEYS[1],ARGV[2])");
	lua.append("\nend");
	lua.append("\nreturn c;");
	return lua.toString();
}
*/
func luaScript() string {
	var buffer bytes.Buffer

	buffer.WriteString("local c")
	buffer.WriteString("\nc = redis.call('get',KEYS[1])")
	buffer.WriteString("\nif c and tonumber(c) > tonumber(ARGV[1]) then")
	buffer.WriteString("\nreturn c;")
	buffer.WriteString("\nend")
	buffer.WriteString("\nc = redis.call('incr',KEYS[1])")
	buffer.WriteString("\nif tonumber(c) == 1 then")
	buffer.WriteString("\nredis.call('expire',KEYS[1],ARGV[2])")
	buffer.WriteString("\nend")
	buffer.WriteString("\nreturn c;")

	return buffer.String()
}
