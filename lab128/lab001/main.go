package main

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"io/ioutil"
	"log"
)

var dbConfig *DBConfig

func init() {
	readConf()
}

func main() {
	mysql()

	postgresql()

	sqlite3()
}

func mysql() {
	//username:password@tcp(host:port)/dbname?charset=utf8&parseTime=True&loc=Local&tls=skip-verify&autocommit=true
	connectInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", dbConfig.Mysql.User, dbConfig.Mysql.Password, dbConfig.Mysql.Host, dbConfig.Mysql.Port, dbConfig.Mysql.DBName)
	db, err := gorm.Open("mysql", connectInfo)
	if err != nil {
		log.Fatalf("connnect mysql error:%v", err)
	}
	defer db.Close()

	log.Println("connect mysql success")
}

func postgresql() {
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s", dbConfig.Postgresql.Host, dbConfig.Postgresql.Port, dbConfig.Postgresql.User, dbConfig.Postgresql.DBName, dbConfig.Postgresql.Password))
	if err != nil {
		log.Fatalf("connect postgresql error:%v", err)
	}
	defer db.Close()

	log.Println("connect postgresql success")
}

func sqlite3() {
	db, err := gorm.Open("sqlite3", dbConfig.Sqlite3.Path)
	if err != nil {
		log.Fatalf("connect sqlite3 error:%v", err)
	}
	defer db.Close()

	log.Println("connect sqlite3 success")
}

func readConf() {
	data, err := ioutil.ReadFile("db_config.json")
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
