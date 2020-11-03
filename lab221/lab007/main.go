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
	rid         string
	pids        []string
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	pflag.StringVarP(&appKey, "appKey", "k", "", "set appKey")
	pflag.StringVarP(&appSecret, "appSecret", "s", "", "set appSecret")
	pflag.StringVarP(&accessToken, "accessToken", "t", "", "set accessToken")
	pflag.StringVarP(&pid, "pid", "p", "", "pid")
	pflag.StringVarP(&rid, "rid", "r", "", "rid")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	if appKey == "" || appSecret == "" {
		log.Fatalf("need appKey and appSecret")
	}

	pids = strings.Split(pid, "_")
}

func main() {
	opentaobao.AppKey = appKey
	opentaobao.AppSecret = appSecret
	opentaobao.Router = "http://gw.api.taobao.com/router/rest"

	getItemInfo()
}

var (
	itemUrl = "https://detail.tmall.com/item.htm?id=603941978639"
	itemId  = "603941978639"
)

func getRebateInfo() {

}

func getItemInfo() {
	res, err := opentaobao.Execute("taobao.tbk.item.info.get", opentaobao.Parameter{
		"num_iids": itemId,
	})

	if err != nil {
		log.Fatalf("execute error:%+v", err)
	}

	j, err := res.MarshalJSON()
	if err != nil {
		log.Fatalf("marshal json error:%v", err)
	}
	log.Printf("%s", j)
}
