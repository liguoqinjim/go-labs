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

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	readConf()
	connectDB()
}

func main() {
	query()

	db.Close()
}

func query() {
	//查询所有user
	var users []*User
	if err := db.Find(&users).Error; err != nil {
		log.Fatalf("db.Find error:%v", err)
	} else {
		for _, v := range users {
			log.Println(v)
		}
	}

	//关联查询company, 使用association key
	log.Println("----------------------------------------------------------------------------------------")
	user := users[0]
	if err := db.Model(&user).Association("Companies").Find(&user.Companies).Error; err != nil {
		log.Fatalf("db.Association error:%v", err)
	} else {
		log.Println("user.companies", user.Companies)
	}

	//association key使用where

	//使用preload
	log.Println("----------------------------------------------------------------------------------------")
	if err := db.Preload("Companies").Find(user).Error; err != nil {
		log.Fatalf("db.Preload error:%v", err)
	} else {
		log.Println("user.companies:", user.Companies)
	}

	//使用preload where
	log.Println("----------------------------------------------------------------------------------------")
	u3 := &User{}
	if err := db.Preload("Companies").Where("id=3").Find(u3).Error; err != nil {
		log.Fatalf("preload where error:%v", err)
	} else {
		log.Println("u3:", u3)
	}

	//preload where查询多个
	var us []*User
	if err := db.Preload("Companies").Where("id<3").Find(&us).Error; err != nil {
		log.Fatalf("preload where error:%v", err)
	} else {
		for _, v := range us {
			log.Println(v)
		}
	}
}

type Company struct {
	Id          int    `gorm:"column:id;type:int(10) unsigned;primary_key;AUTO_INCREMENT;not null" json:"id"`
	UserId      int    `gorm:"column:user_id;type:int(11);not null" json:"userId"`
	CompanyName string `gorm:"column:company_name;type:varchar(32);size:32;not null" json:"companyName"`
}

func (Company) TableName() string {
	return "t_company"
}

type User struct {
	Id       int    `gorm:"column:id;type:int(10) unsigned;primary_key;AUTO_INCREMENT;not null" json:"id"`
	Username string `gorm:"column:username;type:varchar(32);size:32" json:"username"`

	//联合查询
	Companies []Company `gorm:"ASSOCIATION_FOREIGNKEY:id"`
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
