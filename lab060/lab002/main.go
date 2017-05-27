package main

import (
	"errors"
	"fmt"
)

func funcA() error { //这样在defer里面返回error的返回，上层收到的还是nil
	defer func() error {
		if r := recover(); r != nil {
			//r值就是之前调用panic的时候里面的参数
			fmt.Printf("panic recover! r: %v\n", r)
			str, ok := r.(string)
			if ok {
				return errors.New(str)
			} else {
				return errors.New("panic")
			}
		}
		return nil
	}()
	return funcB()
}

func funcA2() (err error) {
	defer func() error {
		if r := recover(); r != nil {
			//r值就是之前调用panic的时候里面的参数
			fmt.Printf("panic recover! r: %v\n", r)
			str, ok := r.(string)
			if ok {
				err = errors.New(str)
			} else {
				err = errors.New("panic")
			}
		}
		return nil
	}()
	return funcB()
}

func funcB() error {
	panic("foo")   //调用的panic的时候就停止执行了
	fmt.Println(1) //1不会打印，因为不会执行到这行
	return errors.New("success")
}

func main() {
	err := funcA2()
	if err == nil {
		fmt.Println("err is nil")
	} else {
		fmt.Printf("err is %v\n", err)
	}
}
