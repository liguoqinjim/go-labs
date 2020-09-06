package main

import (
	"errors"
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
	//inserts()

	//recordNotFound()

	tableName()
}

//批量insert
func inserts() {
	users := make([]*User, 10)
	for i := 0; i < 10; i++ {
		users[i] = &User{
			Username: fmt.Sprintf("xiaoming_%d", i+1),
		}
	}

	if err := db.Create(users).Error; err != nil {
		log.Fatalf("db.Create error:%v", err)
	}
}

//RecordNotFound错误
func recordNotFound() {
	var user User
	if err := db.Where("id=?", 999).Take(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("record not found")
		} else {
			log.Fatalf("db error:%v", err)
		}
	} else {
		log.Printf("user:%v", user)
	}
}

//动态表名
func tableName() {
	msg1 := &Message{
		UserId:  1,
		Message: "message1",
	}
	msg2 := &Message{
		UserId:  2,
		Message: "message2",
	}

	//使用scope实现动态表名
	//DB.Scopes(UserTable(user)).Create(&user)
	db.Scopes(MessageTable(msg1)).Create(msg1)
	db.Scopes(MessageTable(msg2)).Create(msg2)
}
