package main

import (
	"flag"
	"github.com/nilorg/go-opentaobao"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"strings"
)

var (
	appKey      string
	appSecret   string
	accessToken string
	pid         string
)

func init() {
	pflag.StringVarP(&appKey, "appKey", "k", "", "set appKey")
	pflag.StringVarP(&appSecret, "appSecret", "s", "", "set appSecret")
	pflag.StringVarP(&accessToken, "accessToken", "t", "", "set accessToken")
	pflag.StringVarP(&pid, "pid", "p", "", "pid")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	if appKey == "" || appSecret == "" {
		log.Fatalf("need appKey and appSecret")
	}
}

func main() {
	opentaobao.AppKey = appKey
	opentaobao.AppSecret = appSecret
	opentaobao.Router = "http://gw.api.taobao.com/router/rest"

	pids := strings.Split(pid, "_")

	res, err := opentaobao.Execute("taobao.tbk.privilege.get", opentaobao.Parameter{
		"session":   accessToken,
		"item_id":   600996379667,
		"site_id":   pids[2],
		"adzone_id": pids[3],
	})
	if err != nil {
		log.Fatalf("execute error:%+v", err)
	}

	log.Println()
	dm, err := res.Get("tbk_privilege_get_response").Get("result").Get("data").Map()
	if err != nil {
		log.Fatalf("repsonse get error:%v", err)
	}
	for k, v := range dm {
		log.Println(k, v)
	}

	//短连接
	itemUrl, err := res.Get("tbk_privilege_get_response").Get("result").Get("data").Get("item_url").String()
	if err != nil {
		log.Fatalf("get itemUrl error:%v", err)
	}
	itemUrl += "&activityId=664a1b9713744a08b112831d579f66a3"
	log.Println("itemUrl=", itemUrl)

	res, err = opentaobao.Execute("taobao.tbk.spread.get", opentaobao.Parameter{
		"requests": struct {
			Url string `json:"url"`
		}{Url: itemUrl},
	})
	if err != nil {
		log.Fatalf("execute error:%+v", err)
	}

	log.Println("短连接:", res)

	//淘口令
	res, err = opentaobao.Execute("taobao.tbk.tpwd.create", opentaobao.Parameter{
		"text":"12345",
		"url":itemUrl,
	})
	if err != nil {
		log.Fatalf("execute error:%+v", err)
	}
	log.Println("淘口令:", res)
}
