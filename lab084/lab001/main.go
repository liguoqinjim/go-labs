package main

import (
	"github.com/robertkrimen/otto"
	"log"
)

func main() {
	//running something in the VM
	vm := otto.New()
	vm.Run(`
    abc = 2 + 2;
    console.log("The value of abc is " + abc); // 4
	`)

	//set a value
	vm.Set("abc", 5)
	vm.Run(`console.log("The value of abc is " + abc);`)

	//Get a value out of the VM
	if value, err := vm.Get("abc"); err == nil {
		if value_int, err := value.ToInteger(); err == nil {
			log.Printf("value=%d", value_int)
		} else {
			log.Fatalf("value.ToInteget error:%v", err)
		}
	} else {
		log.Fatalf("vm.Get error:%v", err)
	}

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

	//Set a Go function
	vm.Set("sayHello", func(call otto.FunctionCall) otto.Value {
		log.Printf("Hello, %s.\n", call.Argument(0).String())
		return otto.Value{}
	})

	//Set a Go function returns something useful
	vm.Set("twoPlus", func(call otto.FunctionCall) otto.Value {
		right, _ := call.Argument(0).ToInteger()
		result, _ := vm.ToValue(2 + right)
		return result
	})

	//Use the functions in Javascript
	result, _ := vm.Run(`
		sayHello("Tom");
		sayHello();

		result = twoPlus(2.0);
	`)
	log.Println("result=", result)
}
