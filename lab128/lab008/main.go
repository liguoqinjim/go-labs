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

var dbConfig *DBConfig
var db *gorm.DB

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

type User struct {
	Id       int    `gorm:"column:id`
	Username string `gorm:"column:username`
}

func (User) TableName() string {
	return "t_user"
}

type Bill struct {
	Id     int `gorm:"column:id"`
	UserId int `gorm:"column:user_id"`
}

func (Bill) TableName() string {
	return "t_bill"
}

func init() {
	readConf()
}

func main() {
	connectDB()

	//lab001()
	
	if err := lab002(); err != nil {
		log.Printf("manual transaction error:%v", err)
	}
}

func lab001() {
	//username := "admin"
	username := "admin001" //会报错,username是varchar(6)的
	user := &User{
		Username: username,
	}

	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			log.Printf("tx.Create user error:%v", err)
			return err
		}

		log.Printf("new user:%v", user)

		bill := &Bill{
			UserId: user.Id,
		}

		if err := tx.Create(bill).Error; err != nil {
			log.Printf("tx.Create bill error:%v", err)
			return err
		}

		return nil
	})

	if err != nil {
		log.Fatalf("db.Transaction error:%v", err)
	}
}

//manual
func lab002() error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		log.Printf("tx.Error:%v", err)
	}

	username := "admin001" //会报错,username是varchar(6)的
	user := &User{
		Username: username,
	}
	if err := tx.Create(user).Error; err != nil {
		log.Println("create user rollback")
		tx.Rollback()
		return err
	}

	bill := &Bill{
		UserId: user.Id,
	}
	if err := tx.Create(bill); err != nil {
		log.Println("create bill rollback")
		tx.Rollback()
	}

	return tx.Commit().Error
}

func connectDB() {
	connectInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&allowNativePasswords=true", dbConfig.Mysql.User, dbConfig.Mysql.Password, dbConfig.Mysql.Host, dbConfig.Mysql.Port, dbConfig.Mysql.DBName)
	var err error
	db, err = gorm.Open("mysql", connectInfo)
	if err != nil {
		log.Fatalf("open db error:%v", err)
	}

	db.BlockGlobalUpdate(true)
	db.DB().SetConnMaxLifetime(time.Minute * 5)
	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxOpenConns(5)

	log.Println("connect to db success")
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
