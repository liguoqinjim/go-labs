package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type ConfStruct struct {
	Key       string `json:"key"`
	EventName string `json:"event_name"`
	UrlFormat string `json:"url_format"`
	Value1    string `json:"value1"`
	Value2    string `json:"value2"`
	Value3    string `json:"value3"`
}

var conf = new(ConfStruct)

func main() {
	//读取配置文件
	data, err := ioutil.ReadFile("conf.json")
	if err != nil {
		log.Fatalf("ioutil.ReadFile error:%v", err)
	}
	log.Printf("data=\n%s", data)

	err = json.Unmarshal(data, conf)
	if err != nil {
		log.Fatalf("json.Unmarshal error:%v", err)
	}

	//拼接url
	u := fmt.Sprintf(conf.UrlFormat, conf.EventName, conf.Key)
	log.Printf("url=%s", u)

	//client
	//client := &http.Client{}
	//values := map[string]string{
	//	"value1": conf.Value1,
	//	"value2": conf.Value2,
	//	"value3": conf.Value3,
	//}
	//jsonStr, _ := json.Marshal(values)
	//
	//req, err := http.NewRequest(http.MethodPost, u, bytes.NewBuffer(jsonStr))
	//if err != nil {
	//	log.Fatalf("http.NewRequest error:%v", err)
	//}
	//req.Header.Add("Content-Type", "application/json")
	//
	//resp, err := client.Do(req)
	//if err != nil {
	//	log.Fatalf("client.Do error:%v", err)
	//}
	//d, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	log.Fatalf("ioutil.ReadAll error:%v", err)
	//}
	//log.Printf("resp:\n%s", d)
}
