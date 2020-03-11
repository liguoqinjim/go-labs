package main

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once

func main() {

	for i, v := range make([]string, 3) {
		once.Do(once01)
		fmt.Println("count:", v, "---", i)
	}

	for i := 0; i < 3; i++ {
		go func() {
			once.Do(once02)
			fmt.Println("213")
		}()
	}
	time.Sleep(time.Hour)
}

func once01() {
	fmt.Println("once01")
}

func once02() {
	fmt.Println("once02")
}
