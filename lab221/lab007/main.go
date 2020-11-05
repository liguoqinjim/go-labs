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

	//getItemInfo(616202233158)
	//getItemIdFromPassword("PszQc9opkx5")
	itemPrivilege("616202233158")
}

func getItemIdFromPassword(password string) {
	res, err := opentaobao.Execute("taobao.tbk.sc.tpwd.convert", opentaobao.Parameter{
		"session":          accessToken,
		"password_content": password,
		"adzone_id":        "110775100090",
		"site_id":          "198700363",
	})

	if err != nil {
		log.Fatalf("execute error:%+v", err)
	}

	j, err := res.MarshalJSON()
	if err != nil {
		log.Fatalf("marshal json error:%v", err)
	}
	log.Printf("%s", j)

	itemId, err := res.Get("tbk_sc_tpwd_convert_response").Get("data").Get("num_iid").String()
	if err != nil {
		log.Fatalf("res get error:%v", err)
	}
	log.Println("itemId=", itemId)
}

func getItemInfo(itemId int) {
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

	results, err := res.Get("tbk_item_info_get_response").Get("results").Get("n_tbk_item").Array()
	if err != nil {
		log.Fatalf("res get error:%v", err)
	}

	for _, result := range results {
		log.Printf("result:%v", result)
		r := result.(map[string]interface{})
		log.Println("image:", r["pict_url"])
		log.Println("title", r["title"])
	}
}

func itemPrivilege(itemId string) {
	pids := strings.Split(pid, "_")
	res, err := opentaobao.Execute("taobao.tbk.privilege.get", opentaobao.Parameter{
		"session":   accessToken,
		"site_id":   pids[2],
		"adzone_id": pids[3],
		"item_id":   itemId,
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
