package main

import (
	"flag"
	"github.com/nilorg/go-opentaobao"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
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

	//解析淘口令
	res, err := opentaobao.Execute("taobao.tbk.tpwd.parse", opentaobao.Parameter{
		"password_content": "￥4iwZ1GmsnLc￥",
	})
	if err != nil {
		log.Fatalf("execute error:%+v,%+v", err, res)
	}

	log.Println("res1:", res)

	return

	//三个接口
	//https://open.taobao.com/api.htm?spm=a219a.7386797.0.0.32e2669al3VfB0&source=search&docId=42646&docType=2
	//https://open.taobao.com/api.htm?spm=a219a.7386797.0.0.32e2669al3VfB0&source=search&docId=32932&docType=2
	//https://open.taobao.com/api.htm?spm=a219a.7386797.0.0.32e2669al3VfB0&source=search&docId=43873&docType=2

	//解析淘口令-授权
	//pids := strings.Split(pid, "_")
	//res, err := opentaobao.Execute("taobao.tbk.sc.tpwd.convert", opentaobao.Parameter{
	//	"session":          accessToken,
	//	"password_content": "￥4iwZ1GmsnLc￥",
	//	"adzone_id":        pids[3],
	//	"site_id":          pids[2],
	//})
	//if err != nil {
	//	log.Fatalf("execute error:%+v", err)
	//}
	//
	//log.Println("res", res)
}
