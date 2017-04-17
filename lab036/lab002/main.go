package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	//以json的形式输出
	log.SetFormatter(&log.JSONFormatter{})

	log.SetOutput(os.Stdout)

	//输出日志级别
	log.SetLevel(log.WarnLevel)
}

func main() {
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	//输出fatal日志的同时，程序会退出
	log.WithFields(log.Fields{
		"omg":    true,
		"number": 100,
	}).Fatal("The ice breaks")

	//生成一个全局logger
	contextLogger := log.WithFields(log.Fields{
		"common": "this is a common field",
		"other":  "I also should be logged always",
	})

	contextLogger.Info("I'll be logged with common and other field")
	contextLogger.Info("Me too")
}
