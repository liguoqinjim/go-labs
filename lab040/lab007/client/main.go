package main

import (
	"log"
	"net"

	"lab040/lab007/li/rpc"

	"fmt"

	"git.apache.org/thrift.git/lib/go/thrift"
)

const (
	ServerIP   = "127.0.0.1"
	ServerPort = "19090"
)

func main() {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTSocket(net.JoinHostPort(ServerIP, ServerPort))
	if err != nil {
		log.Fatal("transport error ", err)
	}

	useTransport := transportFactory.GetTransport(transport)
	client := rpc.NewQuerySrvClientFactory(useTransport, protocolFactory)
	if err := transport.Open(); err != nil {
		log.Fatal("open transport error ", err)
	}
	defer transport.Close()

	userDemo, err := client.QryUser("小明1", 21)
	if err != nil {
		log.Fatal("QryUser error ", err)
	} else {
		fmt.Println("userDemo", userDemo)
	}

	phone, err := client.QueryPhone(10002)
	if err != nil {
		log.Fatal("QueryPhone error ", err)
	} else {
		fmt.Println("phone", phone)
	}
}
