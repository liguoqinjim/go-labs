package main

import (
	"encoding/base64"
	"log"
)

func main() {
	a := `{"secId":7,"tags":{"aaaaabbbbbb58dccccc":{"1":"辣鸡","2":"1","3":"特殊用户","4":"增加"}}}`
	ae := base64.URLEncoding.EncodeToString([]byte(a))
	log.Println(ae)

	s := "eyJzZWNJZCI6NywidGFncyI6eyJhYWFhYWJiYmJiYjU4ZGNjY2NjIjp7IjEiOiLovqPpuKEiLCIyIjoiMSIsIjMiOiLnibnmrornlKjmiLciLCI0Ijoi5aKe5YqgIn19fQ"
	log.Println(s)
	log.Println(len(s))
	log.Println(len(s) / 4)

	//补上=号 ,正确的base64的长度一定是4的倍数
	if len(s)%4 != 0 {
		len := 4*(len(s)/4+1) - len(s)
		log.Println("l=", len)
		for i := 0; i < len; i++ {
			s += "="
		}
	}
	log.Println(s)

	sd, err := base64.URLEncoding.DecodeString(s)
	if err != nil {
		log.Fatalf("decode error:%v", err)
	}
	log.Printf("sd=%s", sd)
}
