package main

import (
	"database/sql"
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

func init() {
	readConf()
}

func main() {
	connectDB()

	migration()
	tableExists()
	createTable()
	dropTable()

	//modifyColumn()
	//dropColumn()

	//修改默认的表命规则
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "t_" + defaultTableName
	}

	db.AutoMigrate(&User2{}, &Email{}, &Address{}, &Language{}, &CreditCard{})

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

	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Student{})

	db.Debug().AutoMigrate(&Class{})
}

func tableExists() {
	result := db.HasTable(&User{})
	log.Println("has table &User{}", result)

	result = db.HasTable("users")
	log.Println("has table user")
}

func createTable() {
	if err := db.Debug().CreateTable(&Temp{}).Error; err != nil {
		log.Printf("create table error:%v", err)
	} else {
		log.Println("create table success")
	}
}

func dropTable() {
	if err := db.Debug().DropTable(&Temp{}).Error; err != nil {
		log.Printf("drop table error:%v", err)
	} else {
		log.Println("drop table success")
	}
}

func modifyColumn() {
	db.Model(&User{}).ModifyColumn("udes", "int")
}

func dropColumn() {
	db.Model(&User{}).DropColumn("u_address")
}

type User struct {
	ID       string
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

func (Class) TableName() string {
	return "t_class"
}

type Temp struct {
	Id  int
	Tid int
}

type User2 struct {
	gorm.Model
	Birthday time.Time
	Age      int
	Name     string `gorm:"size:255"` // Default size for string is 255, reset it with this tag
	Num      int    `gorm:"AUTO_INCREMENT"`

	CreditCard CreditCard // One-To-One relationship (has one - use CreditCard's UserID as foreign key)
	Emails     []Email    // One-To-Many relationship (has many - use Email's UserID as foreign key)

	BillingAddress   Address // One-To-One relationship (belongs to - use BillingAddressID as foreign key)
	BillingAddressID sql.NullInt64

	ShippingAddress   Address // One-To-One relationship (belongs to - use ShippingAddressID as foreign key)
	ShippingAddressID int

	IgnoreMe  int        `gorm:"-"`                         // Ignore this field
	Languages []Language `gorm:"many2many:user_languages;"` // Many-To-Many relationship, 'user_languages' is join table
}

type Email struct {
	ID         int
	UserID     int    `gorm:"index"`                          // Foreign key (belongs to), tag `index` will create index for this column
	Email      string `gorm:"type:varchar(100);unique_index"` // `type` set sql type, `unique_index` will create unique index for this column
	Subscribed bool
}

type Address struct {
	ID       int
	Address1 string         `gorm:"not null;unique"` // Set field as not nullable and unique
	Address2 string         `gorm:"type:varchar(100);unique"`
	Post     sql.NullString `gorm:"not null"`
}

type Language struct {
	ID   int
	Name string `gorm:"index:idx_name_code"` // Create index with name, and will create combined index if find other fields defined same name
	Code string `gorm:"index:idx_name_code"` // `unique_index` also works
}

type CreditCard struct {
	gorm.Model
	UserID uint
	Number string
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
