package main

import (
	"fmt"
	"net"
	"net/rpc"

	"errors"
)

type MyMethod int

type Args struct {
	I int
	J int
}

type DivResult struct {
	Quo, Rem int
}

func (m *MyMethod) Mult(args *Args, reply *int) error {
	if args == nil || reply == nil {
		return errors.New("nil param")
	}
	*reply = args.I * args.J
	return nil
}

func (m *MyMethod) Div(args *Args, reply *DivResult) error {
	if args == nil || reply == nil {
		return errors.New("nil param")
	}
	if args.J == 0 {
		return errors.New("J cannot be 0")
	}
	reply.Quo = args.I / args.J
	reply.Rem = args.I % args.J
	return nil
}

func main() {
	mm := new(MyMethod)
	server := rpc.NewServer()
	server.Register(mm)

	listener, err := net.Listen("tcp", ":7777")
	defer listener.Close()

	if err != nil {
		fmt.Printf("error %v\n", err)
	}
	server.Accept(listener)
}
