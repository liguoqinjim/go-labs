package main

import "os"

func main() {
	//注意：newpath也要跟着文件名
	err := os.Rename("../data/1/1.txt", "../data/3/1.txt")
	if err != nil {
		panic(err)
	}
}
