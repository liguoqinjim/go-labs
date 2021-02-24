package main

import (
	"fmt"
	"lab259/lab001/sub"
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
	//foo()

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("捕获到的错误：%s\n", r)
		}
	}()

	//sub.SubFoo()
	sub.SubFoo2()
}
