package main

import (
	"errors"
	"fmt"
)

func funcA() error {
	defer func() {
		if r := recover(); r != nil {
			//r值就是之前调用panic的时候里面的参数
			fmt.Printf("panic recover! r: %v\n", r)
		}
	}()
	return funcB() //这里返回的就是空
}

func funcB() error {
	panic("foo")   //调用的panic的时候就停止执行了
	fmt.Println(1) //1不会打印，因为不会执行到这行
	return errors.New("success")
}

func main() {
	err := funcA()
	if err == nil {
		fmt.Println("err is nil")
	} else {
		fmt.Printf("err is %v\n", err)
	}
}
