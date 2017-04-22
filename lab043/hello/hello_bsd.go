// +build darwin dragonfly freebsd netbsd openbsd

package hello

import "fmt"

func SayHello() {
	fmt.Println("hello linux")
}
