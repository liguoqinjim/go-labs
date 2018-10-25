package main

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"io/ioutil"
	"log"
)

var dbConfig *DBConfig
var db *gorm.DB

var users = make(map[int]*User)

func init() {
	readConf()
}

func main() {
	connectDB()

	migration()

	insert()
	update()
	delete()

	db.Close()
}

func connectDB() {
	connectInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", dbConfig.Mysql.User, dbConfig.Mysql.Password, dbConfig.Mysql.Host, dbConfig.Mysql.Port, dbConfig.Mysql.DBName)
	var err error
	db, err = gorm.Open("mysql", connectInfo)
	if err != nil {
		log.Fatalf("open db error:%v", err)
	}
	log.Println("connect to db success")
}

func migration() {
	db.AutoMigrate(&User{})
}

func insert() {
	user := &User{Uid: 90001, Name: "tom", Age: 12}
	if db.NewRecord(user) {
		log.Printf("user is newRecord")
	} else {
		log.Printf("user is not newRecord")
	}

	if err := db.Create(user).Error; err != nil {
		log.Printf("insert error:%v", err)
	} else {
		log.Printf("insert success")
	}
	log.Printf("insert:user=%+v", user)

	if db.NewRecord(user) {
		log.Printf("user is newRecord")
	} else {
		log.Printf("user is not newRecord")
	}

	user2 := &User{Uid: 90002, Name: "kimi", Age: 14}
	user3 := &User{Uid: 90003, Name: "Kitty", Age: 15}
	db.Create(user2)
	db.Create(user3)

	users[user.Uid] = user
	users[user2.Uid] = user2
	users[user3.Uid] = user3
}

func update() {
	user := users[90001]
	user.Age += 100
	db.Save(&user)

	user2 := users[90002]
	log.Printf("db.First user2:%+v", user2)

	db.Model(&user2).Update("age", 212)
	log.Printf("update user2:%+v", user2)

	db.Model(&user2).Updates(map[string]interface{}{"name": "Kimi_changed", "age": 312})
	log.Printf("updates user2:%+v", user2)
}

func delete() {
	user3 := users[90003]
	db.Delete(&user3)
	log.Printf("db.Delete user3:%+v", user3)
}

type User struct {
	ID   int
	Uid  int
	Name string
	Age  int
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
