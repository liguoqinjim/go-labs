package main

import (
	"github.com/spf13/viper"
	"log"
)

func main() {
	viper.AutomaticEnv()
	log.Println(viper.Get("wx_test"))
}
