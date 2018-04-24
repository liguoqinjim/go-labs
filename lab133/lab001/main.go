package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"log"
)

func main() {
	t, err := tail.TailFile("test.log", tail.Config{Follow: true})
	if err != nil {
		log.Printf("tail error:%v", err)
	}

	for line := range t.Lines {
		fmt.Println(line.Text)
	}
}
