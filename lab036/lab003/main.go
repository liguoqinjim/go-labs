package main

import (
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05", //输出时间的格式
		FullTimestamp:   true,                  //输出到tty的时候有用，为true的时候，会输出完整的时间，而不是从开始时间到日志输出的间隔
		ForceColors:     true,                  //为true的时候不检查是否是输出到tty
	})
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
}
