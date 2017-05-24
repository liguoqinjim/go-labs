package main

import (
	"io/ioutil"
	"log"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//读取配置文件
	data, err := ioutil.ReadFile("db_config.json")
	if err != nil {
		log.Fatal("readfile error ", err)
	}
	fmt.Println("data", data)
}
