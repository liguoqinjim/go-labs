package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	viper.AutomaticEnv()

	path := viper.Get("path")
	fmt.Println("path=", path)

	rootPath := viper.Get("ROOT_PATH")
	fmt.Println("root_path=", rootPath)
}
