package main

import (
	"github.com/onrik/logrus/filename"
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
	filenameHook := filename.NewHook()
	filenameHook.Field = "custom_source_field" // Customize source field name
	log.AddHook(filenameHook)

	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	Test1()
	Test2()
}

func Test1() {
	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")
}

func Test2() {
	log.WithFields(log.Fields{
		"omg":    true,
		"number": 100,
	}).Error("The ice breaks!")
}
