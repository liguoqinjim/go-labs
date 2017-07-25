package main

import "log"

func main() {
	//第一段
	log.Println("start1")

	i := 0
HELLO:
	i++
	if i < 5 {
		goto HELLO
	}
	log.Println("i=", i)

	log.Println("end1")

	//第二段
	j := 0
	for j < 10 {
		j++
		if j == 3 {
			goto END
		}
	}
END:
	log.Println("j=", j)
}
