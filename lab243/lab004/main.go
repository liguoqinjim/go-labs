package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
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
	pflag.StringVarP(&port, "port", "p", "5432", "postgres db port")
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
	//postgresql链接样例：host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		host, user, password, dbName, port)
	var err error
	log.Printf("dsn=[%s]", dsn)

	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("connect postgresql error:%v", err)
	} else {
		log.Println("connect postgresql success")
	}
}

func main() {
	//create()
	//find()
}

func create() {
	s := &TStudent{
		Sid:   1,
		Sname: "小明",
		Saddress: sql.NullString{
			String: "上海",
			Valid:  true,
		},
	}

	if err := db.Create(s).Error; err != nil {
		log.Fatalf("db.Create error:%v", err)
	}
	log.Printf("create success:%v", s)
}

func find() {
	var ss []*TStudent
	if err := db.Find(&ss).Error; err != nil {
		log.Fatalf("db.Find error:%v", err)
	}

	for _, s := range ss {
		log.Printf("%v", s)
	}
}
