package main

import (
	"github.com/robertkrimen/otto"
	"log"
)

func main() {
	vm := otto.New()

	script, err := vm.Compile("js/test1.js", nil)
	//script, err := vm.Compile("js/test2.js", nil)
	if err != nil {
		log.Fatalf("compile error:%v", err)
	}

	//run
	_, err = vm.Run(script)
	if err != nil {
		log.Fatalf("run script error:%v", err)
	}

	_, err = vm.Run(`
		myTest()
	`)
	if err != nil {
		log.Fatalf("run error:%v", err)
	}
}
