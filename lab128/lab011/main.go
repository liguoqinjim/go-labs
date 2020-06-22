package main

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"io/ioutil"
	"log"
	"strconv"
)

var dbConfig *DBConfig
var db *gorm.DB

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	readConf()
	connectDB()
}

func main() {

	//插入数据
	//insert()

	//查询
	query()

	db.Close()
}

func insert() {
	//创建10个用户
	//for i := 0; i < 10; i++ {
	//	u := &User{Username: fmt.Sprintf("admin_%d", i+1), Password: "pwd"}
	//	db.Create(u)
	//}

	//创建block
	for i := 0; i < 10; i++ {
		userId := i + 1

		for j := 0; j < 20; j++ {
			b := &Block{UserId: userId, Wxid: fmt.Sprintf("wxid_%d_%d", i, j)}
			db.Create(b)
		}
	}
}

func query() {
	userId := 2
	//查询userId为2的block
	var blocks []*Block
	db.Table((&Block{UserId: userId}).TableName()).Where("user_id=?", userId).Find(&blocks)

	for _, v := range blocks {
		log.Println(v)
	}
}

type Block struct {
	Id         int    `gorm:"column:id;type:int(10) unsigned;primary_key;AUTO_INCREMENT;not null" json:"id"`
	UserId     int    `gorm:"column:user_id;type:int(11);not null" json:"userId"`
	Wxid       string `gorm:"column:wxid;type:varchar(32);size:32;not null" json:"wxid"`
	BlockType  int    `gorm:"column:block_type;type:int(11);not null" json:"blockType"`
	BlockData1 string `gorm:"column:block_data1;type:varchar(32);size:32;not null" json:"blockData1"`
	BlockData2 string `gorm:"column:block_data2;type:varchar(32);size:32" json:"blockData2"`
}

func (b *Block) TableName() string {
	return "t_block_" + strconv.Itoa(b.UserId%10)
}

type User struct {
	Id       int    `gorm:"column:id;type:int(10) unsigned;primary_key;AUTO_INCREMENT;not null" json:"id"`
	Username string `gorm:"column:username;type:varchar(32);size:32" json:"username"`
	Password string `gorm:"column:password;type:varchar(64);size:64" json:"password"`
}

func (User) TableName() string {
	return "t_user"
}

func connectDB() {
	connectInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", dbConfig.Mysql.User, dbConfig.Mysql.Password, dbConfig.Mysql.Host, dbConfig.Mysql.Port, dbConfig.Mysql.DBName)
	var err error
	db, err = gorm.Open("mysql", connectInfo)
	if err != nil {
		log.Fatalf("open db error:%v", err)
	}
	log.Println("connect to db success")

	db.LogMode(true)
}

func readConf() {
	data, err := ioutil.ReadFile("../db_config.json")
	if err != nil {
		log.Fatalf("readFile error:%v", err)
	}

	err = json.Unmarshal(data, &dbConfig)
	if err != nil {
		log.Fatalf("GetDBConfig error:%v", err)
	} else {
		log.Printf("%+v", dbConfig)
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
