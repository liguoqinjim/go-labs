package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

var (
	db *gorm.DB
)

func init() {
	connectInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&allowNativePasswords=true",
		"root", "123456", "127.0.0.1", "3306", "db_test")
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

func main() {
	insertData()
}

func insertData() {
	types := []int{1, 2, 3, 4, 5}

	for _, ty := range types {
		for i := -24 * 7; i <= 0; i++ {
			data := &Test01{
				DataType: ty,
				Time:     time.Now().Add(time.Hour * time.Duration(i)),
				Num:      ty,
			}

			db.Create(data)
		}
	}
}

type Test01 struct {
	DataType int       `gorm:"column:data_type;type:int(11);not null" json:"data_type"`
	Time     time.Time `gorm:"column:time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"time"`
	Num      int       `gorm:"column:num;type:int(11);not null" json:"num"`
}

func (Test01) TableName() string {
	return "t_test01"
}
