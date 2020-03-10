package main

import (
	"github.com/casbin/casbin/v2"
	"log"
)

func main() {
	e, err := casbin.NewEnforcer("./basic_model.conf", "./basic_policy.csv")
	if err != nil {
		log.Fatalf("casbin.NewEnforcer error:%v", err)
	}

	sub := "alice" // the user that wants to access a resource.
	obj := "data1" // the resource that is going to be accessed.
	act := "read"  // the operation that the user performs on the resource.

	res, err := e.Enforce(sub, obj, act)
	if err != nil {
		log.Fatalf("e.Enforce error:%v", err)
	}

	if res {
		log.Println("permit alice to read data1")
	} else {
		log.Println("deny the request, show an error")
	}
}
