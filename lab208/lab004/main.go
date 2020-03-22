package main

import (
	"flag"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
)

func main() {
	flag.Int("flagName", 1234, "help message for flagname")
	pflag.BoolP("testMode", "t", true, "是否进入测试模式")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	i := viper.GetInt("flagName") // retrieve value from viper
	log.Println("i=", i)

	t := viper.Get("testMode")
	log.Println("t=", t)
}
