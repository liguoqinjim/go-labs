/**
 * Created by Administrator on 13-12-31.
 */
package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"os"
)

type Args struct {
	I, J int
}

type MyMethod int

func (this *MyMethod) Mult(args *Args, reply *int) error {
	if args == nil || reply == nil {
		return errors.New("nil paramters !")
	}
	fmt.Println("Mult方法收到参数", args.I, args.J)
	*reply = args.I * args.J
	return nil
}

type DivResult struct {
	Quo, Rem int
}

func (this *MyMethod) Div(args *Args, reply *DivResult) error {
	if args == nil || reply == nil {
		return errors.New("nil paramters !")
	}
	if args.J == 0 {
		return errors.New("/0 !")
	}
	fmt.Println("Div收到参数", args.I, args.J)
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
		fmt.Fprintf(os.Stderr, "error %s\n", err.Error())
		return
	}
	server.Accept(listener)
}
