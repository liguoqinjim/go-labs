package main

import (
	"context"
	"fmt"

	micro "github.com/micro/go-micro/v2"
	"li.com/greetersvc"
)

func main() {
	// Create a new service
	service := micro.NewService(micro.Name("greeter.client"))
	// Initialise the client and parse command line flags
	service.Init()

	// Create new greeter client
	greeterCli := greetersvc.NewGreeterService("com.li.svc.greeter", service.Client())

	// Call the greeter
	rsp, err := greeterCli.Hello(context.TODO(), &greetersvc.Request{Name: "River"})
	if err != nil {
		fmt.Println(err)
	}

	// Print response
	fmt.Println(rsp.Greeting)
}