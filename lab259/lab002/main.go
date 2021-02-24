package main

import (
	"log"
	"time"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("recover:%s", r)
		}
	}()

	log.Println(div(2, 2))

	//别的goruninte中的panic
	//{
	//	go func() {
	//		log.Println(div(2, 0))
	//	}()
	//	time.Sleep(time.Second * 2)
	//}

	//当前goroutine里面的是可以捕获的
	{
		go func() {
			defer func() {
				if r := recover(); r != nil {
					log.Printf("recover:%s", r)
				}
			}()

			log.Println(div(2, 0))
		}()
		time.Sleep(time.Second * 2)
	}

	//log.Println(div(2, 0))

}

func div(a, b int) int {
	return a / b
}
