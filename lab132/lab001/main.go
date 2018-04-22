package main

import (
	"lab132/lab001/ping"
	"log"
)

func main() {
	ok := ping.Ping("baidu.com", 5)
	log.Println("ping result", ok)

	ok = ping.Ping("115.239.210.27", 5)
	log.Println("ping result2", ok)
}
