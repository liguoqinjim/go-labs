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

	//判断是否有表
	hasTable := db.HasTable("user_lab005")
	if hasTable {
		db.Debug().DropTable("user_lab005")
	}

	db.Debug().AutoMigrate(&User{})

	//insert data
	u1 := &User{Name: "Tom", Age: 12}
	u2 := &User{Name: "Kimi", Age: 13}
	u3 := &User{Name: "Alice", Age: 15}
	u4 := &User{Name: "Ben", Age: 17}
	u5 := &User{Name: "Mark", Age: 18}
	u6 := &User{Name: "Tom", Age: 1}
	db.Create(u1)
	db.Create(u2)
	db.Create(u3)
	db.Create(u4)
	db.Create(u5)
	db.Create(u6)

	//select
	//First
	var user User
	db.Debug().First(&user)
	log.Printf("db.First user:%+v", user)

	//注意：这里在user上再使用Last，select语句里面的id会是user的id，所以结果不是最后一个id的对应的值
	db.Debug().Last(&user)
	log.Printf("db.Last user:%+v", user)

	//Last
	var user2 User
	db.Debug().Last(&user2)
	log.Printf("db.Last user2:%+v", user2)

	//Find all
	var users []User
	db.Debug().Find(&users)
	log.Printf("db.Find users:%+v", users)

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
	Id   int `gorm:"auto_increment"`
	Name string
	Age  int
}

func (User) TableName() string {
	return "user_lab005"
}
