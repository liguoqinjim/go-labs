package main

import (
	"log"
	"runtime"
)

func main() {
	pc, file, line, ok := runtime.Caller(0)
	log.Println("pc=", pc)
	log.Println("file=", file)
	log.Println("line=", line)
	log.Println("ok=", ok)

	A()
}

func A() {
	B()
}

func B() {
	pc, file, line, ok := runtime.Caller(1)
	log.Println("pc=", pc)
	log.Println("file=", file)
	log.Println("line=", line)
	log.Println("ok=", ok)
}
