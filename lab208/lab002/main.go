package main

import (
	"github.com/spf13/viper"
	"log"
	"time"
)

func main() {
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath(".")
	v.SetConfigType("toml")

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("v.ReadInConfig error:%v", err)
	}

	//读取到map中
	mapConfig := make(map[string]interface{})
	if err := v.Unmarshal(&mapConfig); err != nil {
		log.Fatalf("v.Unmarshal error:%v", err)
	}
	for k, v := range mapConfig {
		log.Println(k, v)
	}

	//读取到对应的struct
	config := new(Config)
	if err := v.Unmarshal(config); err != nil {
		log.Fatalf("v.Unmarshal error:%v", err)
	}
	log.Printf("%+v", config)
}

type Config struct {
	Port  int
	Title string
	Owner struct {
		Name string
		Dob  time.Time
	}
	Database struct {
		Server        string
		Ports         []int
		ConnectionMax int
		Enabled       bool
	}
	Servers map[string]struct {
		IP string //注意:这里Ip IP都是可以解析的
		DC string
	}
	Clients struct {
		Data [][]interface{}
	}
	Server struct {
		Port int
	}
	Hosts []interface{}
}

//type CompanyInfomation struct {
//	Name                 string
//	MarketCapitalization int64
//	EmployeeNum          int64
//	Department           []interface{}
//	IsOpen               bool
//}
//
//type YamlSetting struct {
//	TimeStamp         string
//	Address           string
//	Postcode          int64
//	CompanyInfomation CompanyInfomation
//}
