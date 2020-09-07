package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"log"
)

var (
	db *gorm.DB
)

var (
	host     string
	host2    string
	port     string
	port2    string
	dbName   string
	user     string
	password string
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	pflag.StringVarP(&host, "host", "h", "127.0.0.1", "db host")
	pflag.StringVarP(&host2, "host2", "n", "127.0.0.1", "db host2")
	pflag.StringVarP(&port, "port", "p", "3306", "db port")
	pflag.StringVarP(&port2, "port2", "l", "3307", "db port2")
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
	dsn1 := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user, password,
		host, port,
		dbName)
	var err error
	log.Println("dsn1=", dsn1)
	db, err = gorm.Open(mysql.Open(dsn1), &gorm.Config{})
	if err != nil {
		log.Fatalf("connnect mysql error:%v", err)
	} else {
		log.Println("connect mysql success")
	}

	dsn2 := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user, password,
		host2, port2,
		dbName)
	log.Println("dsn2=", dsn2)

	if err := db.Use(dbresolver.Register(dbresolver.Config{
		Replicas: []gorm.Dialector{mysql.Open(dsn2)},
		Policy:   dbresolver.RandomPolicy{}})); err != nil {
		log.Fatalf("dbresolver error:%v", err)
	}

	//DB.Use(dbresolver.Register(dbresolver.Config{
	//	// use `db2` as sources, `db3`, `db4` as replicas
	//	Sources:  []gorm.Dialector{mysql.Open("db2_dsn")},
	//	Replicas: []gorm.Dialector{mysql.Open("db3_dsn"), mysql.Open("db4_dsn")},
	//	// sources/replicas load balancing policy
	//	Policy: dbresolver.RandomPolicy{},
	//}).Register(dbresolver.Config{
	//	// use `db1` as sources (DB's default connection), `db5` as replicas for `User`, `Address`
	//	Replicas: []gorm.Dialector{mysql.Open("db5_dsn")},
	//}, &User{}, &Address{}).Register(dbresolver.Config{
	//	// use `db6`, `db7` as sources, `db8` as replicas for `orders`, `Product`
	//	Sources:  []gorm.Dialector{mysql.Open("db6_dsn"), mysql.Open("db7_dsn")},
	//	Replicas: []gorm.Dialector{mysql.Open("db8_dsn")},
	//}, "orders", &Product{}, "secondary"))
}

func main() {
	team := query()
	if team != nil {
		update(team)
	}
}

func query() *Team {
	var team Team
	if err := db.Where("id=?", 10).Take(&team).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("team not found")
		} else {
			log.Fatalln("db error")
		}
	} else {
		log.Println("team=", team)
		return &team
	}

	return nil
}

func update(team *Team) {
	db.Model(team).Select("team_name").Updates(Team{TeamName: "第一个队伍"})
}
