package main

import (
	"github.com/natefinch/lumberjack"
	"io"
	"log"
	"os"
)

func main() {
	log.SetOutput(io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   "foo.log",
		MaxSize:    1,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   false,
	}))

	for i := 0; i < 2000; i++ {
		log.Println("hello world hello world hello world hello world hello world hello world hello world hello world hello world hello world hello world hello world ")
		log.Println("hello world hello world hello world hello world hello world hello world hello world hello world hello world hello world hello world hello world ")
		log.Println("hello world hello world hello world hello world hello world hello world hello world hello world hello world hello world hello world hello world ")
		log.Println("hello world hello world hello world hello world hello world hello world hello world hello world hello world hello world hello world hello world ")
		log.Println("hello world hello world hello world hello world hello world hello world hello world hello world hello world hello world hello world hello world ")
	}
}
