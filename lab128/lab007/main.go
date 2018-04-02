package main

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"io/ioutil"
	"log"
	"time"
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

const (
	TABLE_NAME = "user_lab007"
)

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

	//判断是否有表
	hasTable := db.HasTable(TABLE_NAME)
	if hasTable {
		db.DropTable(TABLE_NAME)
	}

	db.Debug().AutoMigrate(&User{})
}

type User struct {
	Id        int
	LastLogin time.Time
}

func (User) TableName() string {
	return TABLE_NAME
}
