package main

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"time"
)

var GlobalConfig *Config

func main() {
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath(".")
	v.SetConfigType("toml")

	//设置默认值
	v.SetDefault("port", "13131")
	v.SetDefault("hosts", "a b")
	v.SetDefault("server.port", 18082)

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
	GlobalConfig = new(Config)
	config := GlobalConfig
	if err := v.Unmarshal(config); err != nil {
		log.Fatalf("v.Unmarshal error:%v", err)
	}
	log.Printf("%+v", config)

	//监听变化
	v.OnConfigChange(func(in fsnotify.Event) {
		//解析到map
		if err := v.Unmarshal(&mapConfig); err != nil {
			log.Fatalf("v.Unmarshal error:%v", err)
		}
		for k, v := range mapConfig {
			log.Println(k, v)
		}

		//解析到struct
		config := new(Config)
		if err := v.Unmarshal(config); err != nil {
			log.Fatalf("v.Unmarshal error:%v", err)
		}
		log.Printf("config:%+v", config)
		GlobalConfig = config
		log.Printf("globalConfig:%+v", GlobalConfig)
	})
	v.WatchConfig()

	time.Sleep(time.Hour)
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
		ConnectionMax int `mapstructure:"connection_max"`
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
