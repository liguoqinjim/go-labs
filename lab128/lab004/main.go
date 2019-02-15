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
	//查看第一个值
	user := &User{}
	db.Debug().First(user)
	log.Printf("user:%v", user)

	//按主键查找(First有两个参数)
	user2 := &User{}
	db.Debug().First(user2, 2)
	log.Printf("user2:%v", user2)

	//查看最后一个值
	user3 := &User{}
	db.Debug().Last(user3)
	log.Printf("user3:%v", user3)

	//查询所有
	var users []User
	db.Debug().Find(&users)
	log.Printf("users:%v", users)

	//where plain sql
	//这里用first，就会返回第一个match where条件的值
	var user4 User
	db.Debug().Where("name = ?", "Tom").First(&user4)
	log.Printf("where name=? First user4:%+v", user4)

	//查找多个用Find
	var users2 []User
	db.Debug().Where("name = ?", "Tom").Find(&users2)
	log.Printf("where name=? Find users2:%+v", users2)

	//多个条件
	var user5 User
	db.Debug().Where("name = ? and age = ?", "Tom", 1).First(&user5)
	log.Printf("where name=? and age=? First user5:%+v", user5)

	//where in的条件
	var users3 []User
	db.Debug().Where("name in (?)", []string{"Kimi", "Alice"}).Find(&users3)
	log.Printf("where name in (?) Find users3:%+v", &users3)
}

type User struct {
	ID   int
	Uid  int
	Name string
	Age  int
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
