package utils

import (
	"fmt"
	"lab032/lab006/consts"
)

func init() {
	fmt.Println("hello utils package", consts.HELLO_MESSAGE)
}

func Add(a, b int) int {
	return a + b
}
