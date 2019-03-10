package main

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"io/ioutil"
	"log"
	"strconv"
	"time"
)

var dbConfig *DBConfig
var db *gorm.DB

func init() {
	readConf()
}

func main() {
	connectDB()

	//模拟插入数据
	t := int(time.Now().Add(-time.Minute * 5).Round(time.Minute * 5).Unix())

	for i := t; i <= t+(60*5*100); i += 60 * 5 {
		go func() {
			log.Println("start", i)
			for j := 0; j < 2000; j++ {
				heatMap := &HeatMap{Time: i, ServiceCode: strconv.Itoa(j), WifiCount: 10, AuditCount: 20, NetbarCount: 30, HotSpotCount: 40, Total: 100}
				db.Create(heatMap)
			}
			log.Println("i=", i)
		}()
	}

	time.Sleep(time.Hour)

	db.Close()
}

func connectDB() {
	connectInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&allowNativePasswords=true", dbConfig.Mysql.User, dbConfig.Mysql.Password, dbConfig.Mysql.Host, dbConfig.Mysql.Port, dbConfig.Mysql.DBName)
	var err error
	db, err = gorm.Open("mysql", connectInfo)
	if err != nil {
		log.Fatalf("open db error:%v", err)
	}

	//db.DB().SetConnMaxLifetime(time.Minute * 5)
	//db.DB().SetMaxIdleConns(5)
	//db.DB().SetMaxOpenConns(5)

	log.Println("connect to db success")
}

type HeatMap struct {
	Id           int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT;not null"`
	Time         int    `gorm:"column:time;type:int(11);not null"`
	ServiceId    int    `gorm:"column:service_id;type:int(11);not null"`
	ServiceCode  string `gorm:"column:service_code;type:varchar(64);size:64;not null"`
	WifiCount    int    `gorm:"column:wifi_count;type:int(11)"`
	HotSpotCount int    `gorm:"column:hotSpot_count;type:int(11)"`
	AuditCount   int    `gorm:"column:audit_count;type:int(11)"`
	NetbarCount  int    `gorm:"column:netbar_count;type:int(11)"`
	Total        int    `gorm:"column:total;type:int(11)"`
}

func (HeatMap) TableName() string {
	return "t_heat_map"
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
