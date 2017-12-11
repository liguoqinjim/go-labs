package main

import (
	"log"
)

func main() {
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	log.Println(<-messages)
	log.Println(<-messages)

	//错误代码
	messages <- "test"
	messages <- "test2"
	messages <- "test3" //这里会一直等前面的chan有空间，也就锁死了
	log.Println(<-messages)
	//上面这样超过channel的size会出错
}
