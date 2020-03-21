package main

import (
	"github.com/willf/bloom"
	"log"
	"os"
)

const (
	n        = 1000000
	fp       = 0.001
	filename = "filter.tmp"
)

func main() {
	write()

	read()
}

func write() {
	filter := bloom.NewWithEstimates(n, fp)
	log.Println("fp=", filter.EstimateFalsePositiveRate(n))

	filter.AddString("hello")
	filter.AddString("world")

	if filter.TestString("hello") {
		log.Println("hello is in the filter")
	} else {
		log.Println("hello not in the filter")
	}

	//写入文件
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalf("os.Open error:%v", err)
	}
	defer f.Close()

	num, err := filter.WriteTo(f)
	if err != nil {
		log.Fatalf("filter save error:%v", err)
	}
	log.Println("written bytes:", num)
}

func read() {
	filter := bloom.NewWithEstimates(n, fp)

	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("os.Open error:%v", err)
	}
	defer f.Close()

	num, err := filter.ReadFrom(f)
	if err != nil {
		log.Fatalf("filter.ReadFrom error:%v", err)
	}
	log.Println("read bytes:", num)

	if filter.TestString("hello") {
		log.Println("hello is in the filter")
	} else {
		log.Println("hello not in the filter")
	}
}
