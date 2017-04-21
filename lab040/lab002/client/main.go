package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "lab040/lab002/helloworld"
	"log"
)

const (
	address     = "localhost:50051"
	defaultName = "mayday"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	name := defaultName
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
