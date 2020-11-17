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

	getPresent()
}

func getPresent() {
	//mm_18678356_46636874_108815600271
	pids := strings.Split(pid, "_")
	//adzoneId := "108815600271"
	adzoneId := pids[3]
	log.Printf("accessToken=[%s]", accessToken)
	log.Printf("adzoneId=[%s]", adzoneId)
	res, err := opentaobao.Execute("taobao.tbk.dg.vegas.tlj.create", opentaobao.Parameter{
		"session":                  accessToken,
		"adzone_id":                adzoneId,
		"item_id":                  1,
		"total_num":                1,
		"name":                     "我的第一个淘礼金",
		"user_total_win_num_limit": 1,
		"security_switch":          true,
		"per_face":                 1,
		"send_start_time":          "2018-09-01 00:00:00",
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

//转链
func privilege(itemId string) (string, string) {
	res, err := opentaobao.Execute("taobao.tbk.privilege.get", opentaobao.Parameter{
		"session":     accessToken,
		"item_id":     itemId,
		"site_id":     pids[2],
		"adzone_id":   pids[3],
		"relation_id": rid,
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

	log.Println("itemUrl=", itemUrl)
	log.Println("couponUrl=", couponUrl)
	return itemUrl, couponUrl
}

//短连接
func short(itemUrl, couponUrl string) string {
	url := itemUrl
	if couponUrl != "" {
		url = couponUrl
	} else {
		//手动补上优惠券
		//url += "&activityId=4e77fdf019a8404695818dddb0929d46"
	}

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
