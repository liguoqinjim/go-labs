package main

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"lab040/lab006/demo/rpc"
	"os"
)

const (
	Addr = "127.0.0.1:19090"
)

type RpcServiceImpl struct {
}

func (this *RpcServiceImpl) FunCall(callTime int64, funCode string, paramMap map[string]string) (r []string, err error) {
	fmt.Println("-->FunCall:", callTime, funCode, paramMap)

	for k, v := range paramMap {
		r = append(r, k+v)
	}
	return
}

func main() {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocalFactory := thrift.NewTBinaryProtocolFactoryDefault()

	serverTransport, err := thrift.NewTServerSocket(Addr)
	if err != nil {
		fmt.Println("Error!", err)
		os.Exit(1)
	}

	handler := &RpcServiceImpl{}
	processor := rpc.NewRpcServiceProcessor(handler)

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocalFactory)
	fmt.Println("thrift server in", Addr )
	server.Serve()
}
