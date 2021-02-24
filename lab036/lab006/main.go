package main

import (
	log "github.com/sirupsen/logrus"
	_ "lab036/lab006/logger"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("panic get error:%v", r)
		}
	}()

	log.Println("log01")
	log.Infof("this is info")
	log.Debugf("this is debug")
	log.Errorf("this is error")

	log.Panicf("this is a panic")
}
