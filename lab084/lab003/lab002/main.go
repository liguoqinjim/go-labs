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

	function getData()
	{
	  console.log("getData调用")
	  var info={"name":"oec2003","age":25};
	  return info;
	}
	`)

	//Get the value of an expression
	value, _ := vm.Run("getData()")
	{
		if value.IsObject() {
			log.Println("返回的是对象类型")
		} else {
			log.Println("返回的不是对象类型")
		}

		obj := value.Object()
		nameValue, _ := obj.Get("name")
		name, _ := nameValue.ToString()
		log.Println("name=", name)
		ageValue, _ := obj.Get("age")
		age, _ := ageValue.ToInteger()
		log.Println("age=", age)
	}
}
