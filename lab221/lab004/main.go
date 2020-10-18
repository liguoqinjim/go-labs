package main

import (
	"encoding/json"
	"flag"
	"fmt"
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

	//url := inviteCode()
	//url += "&rtag=weituandui"
	//
	//password(url)

	getRIDs()
}

//生成邀请码
func inviteCode() string {
	res, err := opentaobao.Execute("taobao.tbk.sc.invitecode.get", opentaobao.Parameter{
		"session":      accessToken,
		"relation_id":  2563991004,
		"relation_app": "common",
		"code_type":    1,
	})
	if err != nil {
		log.Fatalf("execute error:%+v", err)
	}
	log.Println("inviteCode:", res)

	code, err := res.Get("tbk_sc_invitecode_get_response").Get("data").Get("inviter_code").String()
	if err != nil {
		log.Fatalf("inviteCode error:%v", err)
	}
	log.Println("code=", code)

	url := fmt.Sprintf("https://mos.m.taobao.com/inviter/register?inviterCode=%s&src=pub&app=common", code)
	log.Println("url=", url)

	return url
}

func getRIDs() {
	res, err := opentaobao.Execute("taobao.tbk.sc.publisher.info.get", opentaobao.Parameter{
		"session":      accessToken,
		"info_type":    1,
		"page_no":      1,
		"page_size":    10,
		"relation_app": "common",
		"external_id":  "EhqtuVSomfQU",
	})
	if err != nil {
		log.Fatalf("execute error:%+v", err)
	}
	log.Println("rids:", res)
	j, err := res.MarshalJSON()
	if err != nil {
		log.Fatalf("marshalJSON error:%v", err)
	}
	log.Printf("rids.json:%s", j)

	result, err := res.Get("tbk_sc_publisher_info_get_response").Get("data").Get("inviter_list").Get("map_data").Array()
	log.Println(result)

	for _, r := range result {
		a := r.(map[string]interface{})

		if _, ok := a["rtag"]; ok {
			log.Println(a)
			rtag := a["rtag"].(string)
			relationId := a["relation_id"].(json.Number)
			createDate := a["create_date"].(string)

			log.Println(rtag, relationId, createDate)
		}
	}

	log.Println(err)
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
