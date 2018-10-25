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
	Select()

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
	if db.HasTable(&User{}) {
		db.DropTable(&User{})
	}

	db.AutoMigrate(&User{})
}

func insert() {
	user := &User{Uid: 90001, Name: "tom", Age: 12}
	user2 := &User{Uid: 90002, Name: "kimi", Age: 14}
	user3 := &User{Uid: 90003, Name: "Kitty", Age: 15}
	db.Create(user)
	db.Create(user2)
	db.Create(user3)

	users[user.Uid] = user
	users[user2.Uid] = user2
	users[user3.Uid] = user3
}

func Select() {
	user := &User{}
	db.Debug().First(user)
	log.Printf("user:%v", user)
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
