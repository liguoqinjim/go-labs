package main

import "fmt"

func main() {
	//map
	imap := make(map[int]int)
	imap[1] = 1
	imap[2] = 2
	fmt.Println("imap=", imap)

	var a interface{}
	a = imap
	imapp := a.(map[int]int) //接口查询
	imapp[3] = 3
	fmt.Println("imap=", imap)

	//slice
	islice := make([]int, 3)
	islice[0] = 1
	islice[1] = 2
	islice[2] = 3
	fmt.Println("islice=", islice)
	var b interface{}
	b = islice
	islicep := b.([]int)
	islicep[0] = 99
	fmt.Println("islice=", islice)
}
