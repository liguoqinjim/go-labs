package main

import (
	"fmt"
)

func main() {
	map1 := make(map[int]int)
	map1[1] = 1
	fmt.Println("map1:", map1)
	handleMap(map1)
	fmt.Println("map1:", map1)
}

func handleMap(m map[int]int) {
	m[1] = 2
	fmt.Println("handleMap:", m)
}
