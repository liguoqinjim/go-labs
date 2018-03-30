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

	//判断是否有表
	hasTable := db.HasTable("user_lab006")
	if hasTable {
		db.DropTable("user_lab006")
	}

	db.AutoMigrate(&User{})

	//insert data
	u1 := &User{Name: "Tom", Age: 12}
	u2 := &User{Name: "Kimi", Age: 13}
	u3 := &User{Name: "Alice", Age: 15}
	u4 := &User{Name: "Ben", Age: 17}
	u5 := &User{Name: "Mark", Age: 18}
	u6 := &User{Name: "Tom", Age: 1}
	db.Create(u1)
	db.Create(u2)
	db.Create(u3)
	db.Create(u4)
	db.Create(u5)
	db.Create(u6)

	var user User
	log.Printf("db.First user:%+v", user)

	//update
	user.Age += 100
	db.Debug().Save(&user)

}

type User struct {
	Id   int `gorm:"auto_increment"`
	Name string
	Age  int
}

func (User) TableName() string {
	return "user_lab006"
}
