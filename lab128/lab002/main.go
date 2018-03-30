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

	//Migration，没有表则创建
	db.AutoMigrate(&User{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Student{})

	//先用Debug在调用，可以打印出具体的sql
	db.Debug().AutoMigrate(&Class{})

	//查看是否有表
	result := db.HasTable(&User{})
	log.Println("has table &User{}", result)
	result = db.HasTable("users")
	log.Println("has table user")

	//create table
	db.CreateTable(&Temp{})

	//drop table
	db.DropTable(&Temp{})

	//ModifyColumn
	db.Model(&User{}).ModifyColumn("udes", "int")

	//DropColumn
	db.Model(&User{}).DropColumn("u_address")

	//Add Foreign Key
	db.Model(&User{}).AddForeignKey("stu_id", "students(id)", "RESTRICT", "RESTRICT")

	//Index
	db.Model(&User{}).AddIndex("idx_user_name", "uname")
	db.Model(&User{}).AddUniqueIndex("idx_user_name_age", "uage")
}

type User struct {
	Id       int
	Uid      int
	Uname    string
	Uage     int
	StuId    int
	Udes     string
	UAddress string
}

type Student struct {
	Id    int
	Sno   int
	Sname string
	Sage  int
}

type Class struct {
	Id  int `gorm:"AUTO_INCREMENT"`
	Cno string
}

//指定表命
func (Class) TableName() string {
	return "t_class"
}

type Temp struct {
	Id  int
	Tid int
}
