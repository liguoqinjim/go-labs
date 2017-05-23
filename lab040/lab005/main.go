package main

import (
	"lab040/lab005/lib"
	pb "lab040/lab005/message"
	"log"
	"net"
	"time"

	"golang.org/x/net/context"
)

func main() {
	//测试grpc
	s := lib.NewGRPCServer()

	listener, err := net.Listen("tcp", lib.GRPCPort)
	if err != nil {
		log.Fatalf("failed to connect:%v\n", err)
	}

	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("serve error %v\n", err)
		}
	}()

	time.Sleep(time.Second * 10)

	lib.NewGRPCClient()

	multResult, err := lib.GRPCCalClient.Mult(context.Background(), &pb.CalRequest{A: 2, B: 3})
	if err != nil {
		log.Fatal("multError", err)
	}
	log.Println("multResult:", multResult.Result)
}
