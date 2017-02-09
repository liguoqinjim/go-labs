package main

import "fmt"

func main() {
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)

	//messages <- "test"
	//messages <- "test2"
	//messages <- "test3"
	//fmt.Println(<-messages)
	//上面这样超过channel的size会出错
}
