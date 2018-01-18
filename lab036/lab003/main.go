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
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 100,
	}).Error("The ice breaks!")
}
