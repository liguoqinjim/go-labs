package main

import (
	"fmt"
	"os"
)

func main() {
	//遍历所有的环境变量
	for n, e := range os.Environ() {
		fmt.Println(e)
		if n > 5 {
			break //打印前5个
		}
	}

	//得到某个环境变量
	fmt.Println()
	goroot := os.Getenv("GOROOT")
	fmt.Println("GOROOT=", goroot)

	//设置环境变量
	//注：os.SetEnv _only_ changes the environment definition inside your process
	os.Setenv("test", "1")
}
