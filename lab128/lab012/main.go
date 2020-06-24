package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"io/ioutil"
	"log"
)

var dbConfig *DBConfig
var db *gorm.DB

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	readConf()
	connectDB()
}

func main() {
	//query1()
	query2()

	db.Close()
}

//查询一列(多行)
func query1() {
	//查询出所有的userId
	var ids []int
	db.Table(User{}.TableName()).Pluck("id", &ids)
	log.Println(ids)
}

//查询一列(单行),不能用pluck，pluck要传入struct或者slice
func query2() {
	var id int
	row := db.Table(User{}.TableName()).Select("id").Where("username='tom1'").Row()

	if row == nil {
		log.Println("row is nil")
	} else {
		log.Println("row is not nil", row)
	}

	if err := row.Scan(id); err != nil {
		if err == sql.ErrNoRows {
			log.Println("not found record")
		}
		log.Fatalf("scan error:%v", err)
	}
	log.Println(id)
}

type AuthCode struct {
	Id       int `gorm:"column:id;type:int(10) unsigned;primary_key;AUTO_INCREMENT;not null" json:"id"`
	AutoType int `gorm:"column:auto_type;type:int(11);not null" json:"autoType"`
}

func (AuthCode) TableName() string {
	return "t_auth_code"
}

type Crm struct {
	Id         int    `gorm:"column:id;type:int(10) unsigned;primary_key;AUTO_INCREMENT;not null" json:"id"`
	AuthCodeId int    `gorm:"column:auth_code_id;type:int(11);not null" json:"authCodeId"`
	CrmName    string `gorm:"column:crm_name;type:varchar(32);size:32;not null" json:"crmName"`
}

func (Crm) TableName() string {
	return "t_crm"
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
