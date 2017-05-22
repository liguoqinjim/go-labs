package main

import (
	"fmt"
	"net/rpc"
)

type Args struct {
	I int
	J int
}

type DivResult struct {
	Quo, Rem int
}

func main() {
	pClient, err := rpc.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Printf("error %v\n", err)
		return
	}

	//同步rpc
	var multResult int
	err = pClient.Call("MyMethod.Mult", &Args{2, 7}, &multResult)
	if err != nil {
		fmt.Printf("error %v\n", err)
		return
	}
	fmt.Println("Mult返回", multResult)

	//异步rpc
	var divResult DivResult
	pCall := pClient.Go("MyMethod.Div", &Args{5, 2}, &divResult, nil)
	if pCall != nil {
		if replyCall, ok := <-pCall.Done; ok {
			fmt.Println("Div异步收到返回", replyCall)
			fmt.Println("Div异步收到返回", divResult.Quo, divResult.Rem)
		}
	}
}
