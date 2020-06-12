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

	//转链
	res, err := opentaobao.Execute("taobao.tbk.privilege.get", opentaobao.Parameter{
		"session":   accessToken,
		"item_id":   605457131248,
		"site_id":   pids[2],
		"adzone_id": pids[3],
	})
	if err != nil {
		log.Fatalf("execute error:%+v", err)
	}

	dm, err := res.Get("tbk_privilege_get_response").Get("result").Get("data").Map()
	if err != nil {
		log.Fatalf("repsonse get error:%v", err)
	}
	itemUrl := ""
	couponUrl := ""
	for k, v := range dm {
		log.Println(k, v)

		if k == "item_url" {
			itemUrl = v.(string)
		} else if k == "coupon_click_url" {
			couponUrl = v.(string)
		}
	}

	//短连接
	itemUrl += "&activityId=664a1b9713744a08b112831d579f66a3"
	log.Println("itemUrl=", itemUrl)
	couponUrl += "&activityId=b4700d16496e45e78bcb93c55fe095ff"
	log.Println("couponUrl=", couponUrl)

	res, err = opentaobao.Execute("taobao.tbk.spread.get", opentaobao.Parameter{
		"requests": struct {
			Url string `json:"url"`
		}{Url: couponUrl},
	})
	if err != nil {
		log.Fatalf("execute error:%+v", err)
	}
	log.Println("短连接:", res)
	shorts, err := res.Get("tbk_spread_get_response").Get("results").Get("tbk_spread").Array()
	if err != nil {
		log.Fatalf("array error:%v", err)
	}
	short := ""
	for _, s := range shorts {
		v := s.(map[string]interface{})
		short = v["content"].(string)
	}
	if err != nil {
		log.Fatalf("短连接error:%v", err)
	}
	log.Println("short:", short)

	//淘口令
	res, err = opentaobao.Execute("taobao.tbk.tpwd.create", opentaobao.Parameter{
		"text": "1", //这个参数在API里面被废除了，填了也不会有效果，但是不能不填，不填会报错
		"url":  couponUrl,
	})
	if err != nil {
		log.Fatalf("execute error:%+v", err)
	}
	log.Println("淘口令:", res)
	tpwd, err := res.Get("tbk_tpwd_create_response").Get("data").Get("model").String()
	if err != nil {
		log.Fatalf("淘口令 res error:%v", err)
	}
	log.Println("tpwd:", tpwd)

	//用短连接生成淘口令
	res, err = opentaobao.Execute("taobao.tbk.tpwd.create", opentaobao.Parameter{
		"text": "1", //这个参数在API里面被废除了，填了也不会有效果，但是不能不填，不填会报错
		"url":  short,
	})
	if err != nil {
		log.Fatalf("execute error:%+v", err)
	}
	log.Println("淘口令2:", res)
	tpwd, err = res.Get("tbk_tpwd_create_response").Get("data").Get("model").String()
	if err != nil {
		log.Fatalf("淘口令 res error:%v", err)
	}
	log.Println("tpwd:", tpwd)
}
