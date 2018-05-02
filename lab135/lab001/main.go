package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/bouk/monkey"
)

func main() {
	//patch fmt.Println方法，注意后面的func的参数和返回是要一样的
	monkey.Patch(fmt.Println, func(a ...interface{}) (n int, err error) {
		s := make([]interface{}, len(a))
		for i, v := range a {
			s[i] = strings.Replace(fmt.Sprint(v), "hell", "*bleep*", -1)
		}
		return fmt.Fprintln(os.Stdout, s...)
	})

	fmt.Println("what the hell?")
}
