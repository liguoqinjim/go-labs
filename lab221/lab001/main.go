package main

import (
	"flag"
	"github.com/nilorg/go-opentaobao"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
)

var (
	appKey    string
	appSecret string
	pid       string
)

func init() {
	pflag.StringVarP(&appKey, "appKey", "k", "", "set appKey")
	pflag.StringVarP(&appSecret, "appSecret", "s", "", "set appSecret")
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

	res, err := opentaobao.Execute("taobao.tbk.coupon.convert", opentaobao.Parameter{
		"item_id":   618386269125,
		"adzone_id": pid,
		"url":       "https://uland.taobao.com/quan/detail?sellerId=2207780422672&activityId=679ab82a8fc743179e87c6d159cdef6f",
	})

	if err != nil {
		log.Fatalf("execute error:%+v", err)
	}

	log.Println("商品数量:", res.Get("tbk_item_get_response").Get("total_results").MustInt())
	var imtes []interface{}
	imtes, _ = res.Get("tbk_item_get_response").Get("results").Get("n_tbk_item").Array()
	for _, v := range imtes {
		log.Println("======")
		item := v.(map[string]interface{})
		log.Println("商品名称:", item["title"])
		log.Println("商品价格:", item["reserve_price"])
		log.Println("商品链接:", item["item_url"])
	}
}
