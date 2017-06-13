package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://g.cn/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	//读取body数据
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}
