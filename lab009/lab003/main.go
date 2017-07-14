package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type DBConf struct {
	DBHost string
	DBUser string
	DBPwd  string
	DBName string
}

//
//"DBHost": "DBHost",
//"DBUser": "DBUser",
//"DBPwd": "DBPwd",
//"DBName": "DBName"

var Conf *DBConf

func init() {
	readConf()
}

func readConf() *DBConf {
	file, err := os.Open("db_config.json")
	if err != nil {
		log.Fatalln("failed to open file:", err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln("failed to read data:", err)
	}

	dbConf := new(DBConf)
	if err := json.Unmarshal(data, dbConf); err != nil {
		log.Fatalln("failed to parse json:", err)
	}

	Conf = dbConf

	return dbConf
}

func main() {
	fmt.Println(Conf)
}
