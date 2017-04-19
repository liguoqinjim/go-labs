package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	file, err := os.Create("tmp.out")
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(file)

	log.Info("aaa")
}
