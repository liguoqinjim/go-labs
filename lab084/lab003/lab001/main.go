package main

import (
	"github.com/robertkrimen/otto"
	"log"
)

func main() {
	//running something in the VM
	vm := otto.New()

	//set a string
	vm.Set("xyzzy", "Nothing happens")
	vm.Run(`
		console.log("xyzzy.length=" + xyzzy.length);
	`)

	//Get the value of an expression
	value, _ := vm.Run("xyzzy.length")
	{
		value, _ := value.ToInteger()
		log.Printf("xyzzy.length=%d", value)
	}
}
