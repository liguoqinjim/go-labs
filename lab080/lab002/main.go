package main

import "log"

func main() {
	log.Println("start")

HELLO:
	for i := 0; i < 10; i++ {
		log.Println("i=", i)
		break HELLO //这里的break label，不会像goto一样，再从goto开始执行逻辑。而是会直接跳过这个循环
	}

	log.Println("end")
}
