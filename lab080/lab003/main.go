package main

import "log"

func main() {
HELLO:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			log.Println("j=", j)
			if j < 1 {
				continue HELLO
			}
		}
	}
}
