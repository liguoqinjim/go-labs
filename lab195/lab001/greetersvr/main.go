package main

import (
	"context"
	"fmt"
	"li.com/greetersvc"

	micro "github.com/micro/go-micro/v2"
)

// Greeter .
type Greeter struct{}

// Hello .
func (g *Greeter) Hello(ctx context.Context, req *greetersvc.Request, rsp *greetersvc.Response) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("com.li.svc.greeter"),
	)
	// Init will parse the command line flags.
	service.Init()
	// Register handler
	greetersvc.RegisterGreeterHandler(service.Server(), new(Greeter))
	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
