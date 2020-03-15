package main

import (
	"fmt"
	"net/rpc"
	"os"
)

type Args struct {
	I, J int
}

//func (this *MyMethod) Mult(args *Args, reply *int) error {
//	if args == nil || reply == nil {
//		return errors.New("nil parameters !")
//	}
//	*reply = args.I*args.J
//	return nil
//}

type DivResult struct {
	Quo, Rem int
}

//func (this *MyMethod) Div(args *Args, reply *DivResult) {
//	reply.Quo = args.I / args.J
//	reply.Rem = args.J % args.J
//}

func main() {
	pClient, err := rpc.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Fprintf(os.Stderr, "err : %s\n", err.Error())
		return
	}

	// 同步RPC
	var multResult int
	err = pClient.Call("MyMethod.Mult", &Args{2, 7}, &multResult)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err : %s\n", err.Error())
		return
	}
	fmt.Println("Mult返回", multResult)

	// 异步RPC
	var divResult DivResult
	pCall := pClient.Go("MyMethod.Div", &Args{5, 2}, &divResult, nil)
	if pCall != nil {
		if replyCall, ok := <-pCall.Done; ok {
			fmt.Println("Div异步收到返回", replyCall)
			fmt.Println("Div异步收到返回", divResult.Quo, divResult.Rem)
		}
	}
}
