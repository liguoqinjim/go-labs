package main

import "fmt"

//参数pings只能send值
func ping(pings chan<- string, msg string) {
	pings <- msg

	//msg2 := <-pings //would be a compile-time error
}

//参数pings只能receive值
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
