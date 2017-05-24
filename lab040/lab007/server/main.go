package main

import (
	"errors"
	"fmt"
	"lab040/lab007/li/rpc"
	"log"

	"git.apache.org/thrift.git/lib/go/thrift"
)

const (
	Addr = "127.0.0.1:19090"
)

type QuerySrvImpl struct {
}

func (this *QuerySrvImpl) QryUser(name string, age int32) (r *rpc.UserDemo, err error) {
	for _, v := range StudentMap {
		if v.Sname == name && v.Sage == int(age) {
			return &rpc.UserDemo{ID: int32(v.Sid), Name: v.Sname, Age: int32(v.Sage), Phone: v.Sphone}, nil
		}
	}

	return nil, errors.New("no user found")
}

func (this *QuerySrvImpl) QueryPhone(id int32) (r string, err error) {
	if s, ok := StudentMap[int(id)]; ok {
		return s.Sphone, nil
	}
	return "", errors.New("no userid found")
}

func main() {
	//初始化StudentMap
	initStudentMap()

	//服务器
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	serverTransport, err := thrift.NewTServerSocket(Addr)
	if err != nil {
		log.Fatal("serverTransport error ", err)
	}

	handler := new(QuerySrvImpl)
	processor := rpc.NewQuerySrvProcessor(handler) //rpc.NewQuerySvrProcessor这个方法是thrift编译的时候自动生成的

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory) //还有server2,server6方法，只是方法的参数是两个或者六个
	fmt.Println("thrift server in", Addr)
	server.Serve()
}

type Student struct {
	Sid    int
	Sname  string
	Sage   int
	Sphone string
}

var StudentMap map[int]*Student

func initStudentMap() {
	StudentMap = make(map[int]*Student)
	for i := 10001; i <= 10010; i++ {
		s := &Student{Sid: i, Sname: fmt.Sprintf("小明%d", (i - 10000)), Sage: 20 + (i - 10000), Sphone: fmt.Sprintf("138000%d", i)}
		StudentMap[i] = s
	}
}
