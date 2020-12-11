package main

import (
	"flag"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	db *gorm.DB
)

var (
	host     string
	port     string
	dbName   string
	user     string
	password string
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	pflag.StringVarP(&host, "host", "h", "127.0.0.1", "db host")
	pflag.StringVarP(&port, "port", "p", "3306", "db port")
	pflag.StringVarP(&dbName, "dbName", "d", "temp", "db name")
	pflag.StringVarP(&user, "user", "u", "root", "db user")
	pflag.StringVarP(&password, "password", "P", "", "db port")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	open()
}

//连接数据库
func open() {
	//mysql链接样例：username:password@tcp(host:port)/dbname?charset=utf8&parseTime=True&loc=Local&tls=skip-verify&autocommit=true
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user, password,
		host, port,
		dbName)
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("connnect mysql error:%v", err)
	} else {
		log.Println("connect mysql success")
	}
}

func main() {
	updates()
}

func updates() {
	var user User
	if err := db.First(&user).Error; err != nil {
		log.Fatalf("db.First error:%v", err)
	}
	log.Printf("user:%+v", user)

	//update
	ups := map[string]interface{}{
		"username": "new_name",
	}
	if err := db.Model(&user).Updates(ups).Error; err != nil {
		log.Fatalf("updates error:%v", err)
	}
	log.Printf("user:%+v", user)
}
