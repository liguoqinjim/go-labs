package main

import (
	"context"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	//example01()
	//example02()
	//example03()
	example04()
}

func example01() {
	//读取yaml文件
	v := viper.New()
	//设置读取的配置文件
	v.SetConfigName("config")
	//添加读取的配置文件路径
	v.AddConfigPath(".")
	//windows环境下为%GOPATH，linux环境下为$GOPATH
	v.AddConfigPath("$GOPATH/src/")
	//设置配置文件类型
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("err:%s\n", err)
	}

	fmt.Printf(
		`
		TimeStamp:%s
		CompanyInfomation.Name:%s
		CompanyInfomation.Department:%s `,
		v.Get("TimeStamp"),
		v.Get("CompanyInfomation.Name"),
		v.Get("CompanyInfomation.Department"),
	)

	/*
		result:
		TimeStamp:2018-10-18 10:09:23
		CompanyInfomation.Name:Sunny
		CompanyInfomation.Department:[Finance Design Program Sales]
	*/
}

func example02() {
	//读取yaml文件
	v := viper.New()
	//设置读取的配置文件
	v.SetConfigName("config")
	//添加读取的配置文件路径
	v.AddConfigPath(".")
	//windows环境下为%GOPATH，linux环境下为$GOPATH
	v.AddConfigPath("$GOPATH/src/")
	//设置配置文件类型
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("err:%s\n", err)
	}

	fmt.Printf(
		`
		TimeStamp:%s
		CompanyInfomation.Name:%s
		CompanyInfomation.Department:%s `,
		v.Get("TimeStamp"),
		v.Get("CompanyInfomation.Name"),
		v.Get("CompanyInfomation.Department"),
	)

	parseYaml(v)
}

type CompanyInfomation struct {
	Name                 string
	MarketCapitalization int64
	EmployeeNum          int64
	Department           []interface{}
	IsOpen               bool
}

type YamlSetting struct {
	TimeStamp         string
	Address           string
	Postcode          int64
	CompanyInfomation CompanyInfomation
}

func parseYaml(v *viper.Viper) {
	var yamlObj YamlSetting
	if err := v.Unmarshal(&yamlObj); err != nil {
		fmt.Printf("err:%s", err)
	}
	fmt.Println(yamlObj)
	/*
		result:
		{2018-10-18 10:09:23 Shenzhen 518000 {Sunny 50000000 200 [Finance Design Program Sales] false}}
	*/
}

//命令行参数
func example03() {
	pflag.String("hostAddress", "127.0.0.1", "Server running address")
	pflag.Int64("port", 8080, "Server running port")
	pflag.Parse()

	viper.BindPFlags(pflag.CommandLine)
	fmt.Printf("hostAddress :%s , port:%s", viper.GetString("hostAddress"), viper.GetString("port"))

	/*
			example:
			go run main2.go --hostAddress=192.192.1.10 --port=9000
			help:
			Usage of /tmp/go-build183981952/b001/exe/main:
		     --hostAddress string   Server running address (default "127.0.0.1")
		     --port int             Server running port (default 8080)
	*/
}

//监听配置文件变化
func example04() {
	//读取yaml文件
	v := viper.New()
	//设置读取的配置文件
	v.SetConfigName("config")
	//添加读取的配置文件路径
	v.AddConfigPath(".")
	//设置配置文件类型
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("err:%s\n", err)
	}
	fmt.Printf("config.Address=%s", v.Get("Address"))

	//创建一个信道等待关闭（模拟服务器环境）
	ctx, _ := context.WithCancel(context.Background())
	//cancel可以关闭信道
	//ctx, cancel := context.WithCancel(context.Background())
	//设置监听回调函数
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("config is change :%s \n", e.String())
		fmt.Printf("config.Address=%s", v.Get("Address"))
		//cancel()
	})
	//开始监听
	v.WatchConfig()
	//信道不会主动关闭，可以主动调用cancel关闭
	<-ctx.Done()
}
