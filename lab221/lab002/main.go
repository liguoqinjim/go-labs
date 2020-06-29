package main

import (
	"flag"
	"github.com/nilorg/go-opentaobao"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"time"
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

	layout := "2006-01-02 15:04:05"
	//timeNow := time.Unix(1592409600, 0)
	timeNow := time.Unix(1593344338-60*18, 0)

	startTime := timeNow.Format(layout)
	endTime := timeNow.Add(time.Minute * 20).Format(layout)

	//timeNow := time.Now()
	res, err := opentaobao.Execute("taobao.tbk.sc.order.details.get", opentaobao.Parameter{
		"session":    accessToken,
		"start_time": startTime,
		"end_time":   endTime,
		"page_zie":   20,
	})
	if err != nil {
		log.Fatalf("execute error:%+v", err)
	}

	log.Println("res", res)
	dm, err := res.Get("tbk_sc_order_details_get_response").Get("data").Get("results").Get("publisher_order_dto").Array()
	if err != nil {
		//log.Fatalf("repsonse get error:%v", err)
		log.Printf("response get error:%v", err)
	}

	log.Println("dm.length=", len(dm))
	for _, v := range dm {
		item := v.(map[string]interface{})

		for k, v := range item {
			log.Println(k, v)
		}
	}
}
