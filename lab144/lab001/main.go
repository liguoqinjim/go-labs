package main

import (
	"github.com/go-ini/ini"
	"log"
)

const rawdata = `
[author]
E-MAIL = u@gogs.io
TEST = nihao
`

func main() {
	cfg, err := ini.Load(
		[]byte(rawdata),
		"full.ini",
	)

	if err != nil {
		log.Fatalf("ini.Load error:%v", err)
	}

	//读取
	sec, err := cfg.GetSection("author")
	if err != nil {
		log.Fatalf("cfg.GetSection error:%v", err)
	}

	//判断有没有TEST这个key
	yes := sec.Haskey("TEST")
	log.Println("是否有TEST这个key:", yes)

	//得到key值
	v := sec.Key("TEST").String()
	log.Println("TEST的值:", v)

	//得到E-mail的值(可以看到这个值会被第二个加载的配置文件顶掉)
	v = sec.Key("E-MAIL").String()
	log.Println("E-MAIL的值:", v)

	//得到GITHUB的值(可以看到这个值里面原本的%(NAME)会被NAME这个KEY的值替换掉)
	v = sec.Key("GITHUB").String()
	log.Println("GITHUB的值:", v)

	//读取two_lines的值
	sec2 := cfg.Section("advance")
	v = sec2.Key("two_lines").String()
	log.Println("two_lines的值:", v)
}
