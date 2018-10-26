package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

var log = logrus.New()

func main() {
	//输出到文件
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	if err != nil {
		log.Errorf("Failed to log to file err:%v", err)
	} else {
		log.Out = file
	}

	log.WithFields(logrus.Fields{
		"id": "1",
	}).Info("this is 1st log")

	log.WithFields(logrus.Fields{
		"id": "2",
	}).Warn("this is 2nd log")

	log.WithFields(logrus.Fields{
		"id": "3",
	}).Error("this is 3rd log")
}
