package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"io/ioutil"
	"log"
)

var dbConfig *DBConfig

func init() {
	readConf()
}

func main() {
	mysql()
}

func mysql() {
	//mysql链接样例：username:password@tcp(host:port)/dbname?charset=utf8&parseTime=True&loc=Local&tls=skip-verify&autocommit=true
	connectInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", dbConfig.Mysql.User, dbConfig.Mysql.Password, dbConfig.Mysql.Host, dbConfig.Mysql.Port, dbConfig.Mysql.DBName)

	eng, err := xorm.NewEngine("mysql", connectInfo)
	if err != nil {
		log.Fatalf("xorm.NewEngine error:%v", err)
	}
	if err := eng.Ping(); err != nil {
		log.Fatalf("eng.Ping error:%v", err)
	} else {
		log.Println("connect mysql success")
	}
}

func readConf() {
	data, err := ioutil.ReadFile("../db_config.json")
	if err != nil {
		log.Fatalf("readFile error:%v", err)
	}

	err = json.Unmarshal(data, &dbConfig)
	if err != nil {
		log.Fatalf("GetDBConfig error:%v", err)
	}
}

type DBConfig struct {
	Mysql struct {
		Host     string `json:"Host"`
		Port     int    `json:"Port"`
		User     string `json:"User"`
		Password string `json:"Password"`
		DBName   string `json:"DBName"`
	} `json:"mysql"`
	Postgresql struct {
		Host     string `json:"Host"`
		Port     int    `json:"Port"`
		User     string `json:"User"`
		Password string `json:"Password"`
		DBName   string `json:"DBName"`
	} `json:"postgresql"`
	Sqlite3 struct {
		Path string `json:"Path"`
	} `json:"sqlite3"`
}
