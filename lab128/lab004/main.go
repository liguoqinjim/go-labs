package main

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"io/ioutil"
	"log"
)

type DBConfig struct {
	DBHost string
	DBUser string
	DBPwd  string
	DBName string
}

func GetDBConfig(data []byte) *DBConfig {
	var config DBConfig
	err := json.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("GetDBConfig error:%v", err)
	}

	return &config
}

func main() {
	data, err := ioutil.ReadFile("db_config.json")
	if err != nil {
		log.Fatalf("readFile error:%v", err)
	}
	dbConfig := GetDBConfig(data)

	connectInfo := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbConfig.DBUser, dbConfig.DBPwd, dbConfig.DBHost, dbConfig.DBName)
	db, err := gorm.Open("mysql", connectInfo)
	if err != nil {
		log.Fatalf("open db error:%v", err)
	}
	defer db.Close()

	db.Debug().AutoMigrate(&User{})

	//insert :Create
	u1 := &User{Name: "tom", Age: 12}
	db.Debug().Create(u1)
	log.Printf("insert:u1=%+v", u1)

	//insert :NewRecord
	u2 := &User{Name: "kimi", Age: 13}
	result := db.Debug().NewRecord(u2)
	log.Printf("newRecord:u2=%+v,result=%t", u2, result)
}

type User struct {
	Id   int `gorm:"auto_increment"`
	Name string
	Age  int
}
