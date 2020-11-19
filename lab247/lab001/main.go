package main

import "fmt"

func main() {
	lab001()
	lab002()

	lab003()
	lab004()
	lab005()
}

func lab001() {
	const max = 3

	number := [max]int{5, 6, 7}
	var ptrs [max]*int //指针数组
	//将number数组的值的地址赋给ptrs
	for i := 0; i < max; i++ {
		ptrs[i] = &number[i]
	}
	for i, x := range ptrs {
		fmt.Printf("指针数组：索引:%d 值:%+v 值的内存地址:%d\n", i, *x, x)
	}
}

func lab002() {
	const max = 3

	number := [max]int{5, 6, 7}
	var ptrs [max]*int //指针数组
	//将number数组的值的地址赋给ptrs
	for i, x := range &number {
		ptrs[i] = &x
	}
	for i, x := range ptrs {
		fmt.Printf("指针数组：索引:%d 值:%d 值的内存地址:%d\n", i, *x, x)
	}
}

type Meta struct {
	Id   int
	Name string
}

func lab003() {
	const max = 3

	metas := [max]Meta{{
		Id:   1,
		Name: "001",
	}, {Id: 2, Name: "002"}, {Id: 3, Name: "003"}}
	var ptrs [max]*Meta //指针数组

	//将meta数组的值的地址赋给ptrs
	for i := 0; i < max; i++ {
		ptrs[i] = &metas[i]
	}
	for i, x := range ptrs {
		fmt.Printf("指针数组：索引:%d 值:%v 值的内存地址:%p\n", i, *x, x)
		//fmt.Printf("指针数组：索引:%d 值:%v 值的内存地址:%d\n", i, *x, x)
	}
}

func lab004() {
	const max = 3

	metas := [max]*Meta{{
		Id:   1,
		Name: "001",
	}, {Id: 2, Name: "002"}, {Id: 3, Name: "003"}}
	var ptrs [max]*Meta //指针数组

	//将meta数组的值的地址赋给ptrs
	for i := 0; i < max; i++ {
		ptrs[i] = metas[i]
	}
	for i, x := range ptrs {
		fmt.Printf("指针数组：索引:%d 值:%v 值的内存地址:%p\n", i, *x, x)
		//fmt.Printf("指针数组：索引:%d 值:%v 值的内存地址:%d\n", i, *x, x)
	}
}

func lab005() {
	const max = 3

	metas := [max]Meta{{
		Id:   1,
		Name: "001",
	}, {Id: 2, Name: "002"}, {Id: 3, Name: "003"}}
	var ptrs [max]*Meta //指针数组

	//将meta数组的值的地址赋给ptrs
	for i, x := range &metas {
		ptrs[i] = &x
	}
	for i, x := range ptrs {
		fmt.Printf("指针数组：索引:%d 值:%v 值的内存地址:%p\n", i, *x, x)
		//fmt.Printf("指针数组：索引:%d 值:%v 值的内存地址:%d\n", i, *x, x)
	}
}
