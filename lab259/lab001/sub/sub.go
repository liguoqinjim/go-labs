package sub

import "fmt"

//除以0
func SubFoo() {
	var a, b int

	a, b = 1, 1
	c := 3 / (a - b)
	fmt.Println(a, b, c)
}

func SubFoo2() {
	panic("主动panic")
}
