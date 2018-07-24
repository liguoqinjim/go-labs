package main

import "log"

//参数pings只能send值
func ping(pings chan<- string, msg string) {
	pings <- msg

	//因为pings只能send，所有下面这行代码编译的时候就不报错
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
	log.Println(<-pongs)
}
