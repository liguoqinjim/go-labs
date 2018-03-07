package main

import (
	"fmt"
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

func MyLog(format string, detail ...interface{}) {
	log.Info(fmt.Sprintf(format, detail...))
}

func Test1() {
	MyLog("Test:%d", 1)
}

func Test2() {
	MyLog("Test:%d", 2)
}
