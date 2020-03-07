package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var (
	endpoint        string
	accessKeyId     string
	accessKeySecret string
)

func init() {
	readConfig()
}

func main() {

}

func readConfig() {
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath(".")
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("err:%s\n", err)
		os.Exit(1)
	}

	endpoint = v.GetString("Endpoint")
	accessKeyId = v.GetString("AccessKeyId")
	accessKeySecret = v.GetString("AccessKeySecret")
}
