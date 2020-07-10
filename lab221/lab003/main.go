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
	pids        []string
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

	pids = strings.Split(pid, "_")
}

func main() {
	opentaobao.AppKey = appKey
	opentaobao.AppSecret = appSecret
	opentaobao.Router = "http://gw.api.taobao.com/router/rest"

	clickUrl := passwordParse("4yBa1xt6i0B")
	//short := short(clickUrl)
	password(clickUrl)

	//password(short)
}

func passwordParse(password string) string {
	res, err := opentaobao.Execute("taobao.tbk.sc.tpwd.convert", opentaobao.Parameter{
		"session":          accessToken,
		"password_content": password,
		"adzone_id":        pids[3],
		"site_id":          pids[2],
	})
	if err != nil {
		log.Fatalf("execute error:%+v", err)
	}

	log.Println("res", res)
	j, err := res.MarshalJSON()
	if err != nil {
		log.Fatalf("marshal json error:%v", err)
	}
	log.Printf("%s", j)
	clickUrl, err := res.Get("tbk_sc_tpwd_convert_response").Get("data").Get("click_url").String()
	if err != nil {
		log.Fatalf("string error:%v", err)
	}
	log.Println("clickUrl=", clickUrl)
	return clickUrl
}

func short(clickUrl string) string {
	url := clickUrl

	log.Println("url=", url)
	res, err := opentaobao.Execute("taobao.tbk.spread.get", opentaobao.Parameter{
		"requests": struct {
			Url string `json:"url"`
		}{Url: url},
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

	return short
}

//淘口令
func password(url string) {
	res, err := opentaobao.Execute("taobao.tbk.tpwd.create", opentaobao.Parameter{
		"text": "1", //这个参数在API里面被废除了，填了也不会有效果，但是不能不填，不填会报错
		"url":  url,
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
}
