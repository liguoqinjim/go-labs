package main

import (
	log "github.com/sirupsen/logrus"
)

func init() {
	//log.SetFormatter(&log.JSONFormatter{})

	//设置output
	//log.SetOutput(os.Stdout)

	//设置最低的日志等级，低于这个等级会被忽略
	log.SetLevel(log.DebugLevel)
}

func main() {
	log.WithFields(log.Fields{
		"id": "1",
	}).Info("this is 1st log")

	log.WithFields(log.Fields{
		"id": "2",
	}).Warn("this is 2nd log")

	log.WithFields(log.Fields{
		"id": "3",
	}).Error("this is 3rd log")

	log.Println("this is a 4th without fields")

	//A common pattern
	contextLogger := log.WithFields(log.Fields{
		"common": "pattern",
	})

	contextLogger.Info("this is 5th log")
	contextLogger.Info("this is 6th log")
}
