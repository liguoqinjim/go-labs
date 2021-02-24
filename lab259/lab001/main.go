package main

import (
	"fmt"
	"log"
)

func foo() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("捕获到的错误：%s\n", r)
		}
	}()

	var a, b int

	a, b = 1, 1
	c := 3 / (a - b)
	fmt.Println(a, b, c)
}

func main() {
	//当前的包
	{
		//foo()
	}

	//别的包
	{
		//defer func() {
		//	if r := recover(); r != nil {
		//		fmt.Printf("捕获到的错误：%s\n", r)
		//	}
		//}()
		//
		////sub.SubFoo()
		//sub.SubFoo2()
	}

	//fatal
	{
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("捕获到的错误：%s\n", r)
			}
		}()

		log.Fatalf("fatal退出")
	}
}
