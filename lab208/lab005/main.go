package main

import (
	"github.com/spf13/viper"
	"log"
)

func main() {
	v := viper.New()

	v.SetConfigName("app")
	v.AddConfigPath(".")
	v.SetConfigType("ini")

	//读取到map
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

	//读取到struct中
	config := new(Config)
	if err := v.Unmarshal(config); err != nil {
		log.Fatalf("v.Unmarshal error:%v", err)
	}
	log.Printf("%+v", config)
}

type Config struct {
	Default struct {
		Version     string
		MonitorPath string
		NewPath     string
	}

	Mysql struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
	}

	Kafka struct {
		BrokerList []string
	}

	ZkInfo struct {
		ZkConns []string
	} `mapstructure:"zk_info"`
}
