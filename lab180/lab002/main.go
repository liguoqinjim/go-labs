package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"io/ioutil"
	"log"
	"time"
)

var dbConfig *DBConfig
var engine *xorm.Engine

func init() {
	readConf()
}

func main() {
	connectDB()

	example()
}

func connectDB() {
	//mysql链接样例：username:password@tcp(host:port)/dbname?charset=utf8&parseTime=True&loc=Local&tls=skip-verify&autocommit=true
	connectInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", dbConfig.Mysql.User, dbConfig.Mysql.Password, dbConfig.Mysql.Host, dbConfig.Mysql.Port, dbConfig.Mysql.DBName)

	var err error
	engine, err = xorm.NewEngine("mysql", connectInfo)
	if err != nil {
		log.Fatalf("xorm.NewEngine error:%v", err)
	}
	if err := engine.Ping(); err != nil {
		log.Fatalf("eng.Ping error:%v", err)
	} else {
		log.Println("connect mysql success")
	}
}

func example() {
	//创建表
	if err := engine.Sync2(new(User)); err != nil {
		log.Fatalf("engine.Sync2 error:%v", err)
	}

	//insert数据
	u := &User{Name: "tom", Age: 18}
	engine.Insert(u)

	//insert多个数据
	us := make([]*User, 10)
	for i := 0; i < 10; i++ {
		us[i] = &User{Name: fmt.Sprintf("tom%d", i), Age: 18 + i}
	}
	engine.Insert(us)
}

type User struct {
	Id      int64
	Name    string
	Salt    string
	Age     int
	Passwd  string    `xorm:"varchar(200)"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
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
