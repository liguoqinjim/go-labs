package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"io/ioutil"
	"lab038/lab008/pb"
	"log"
)

type ErrorInfo struct {
}

type Info struct {
	InfoId int
	Info   string
}

func (i *Info) Reset() {

}
func (i *Info) String() string {
	return fmt.Sprintf("%d,%s", i.InfoId, i.Info)
}
func (i *Info) ProtoMessage() {

}

func main() {
	es1 := &pb.ErrorStatus{Message: "error1"}
	info1 := &Info{InfoId: 23, Info: "hello"}
	any, err := ptypes.MarshalAny(info1)
	if err != nil {
		log.Fatalln("failed to marshalAny:", err)
	}
	es1.Details[0] = any
	fmt.Println(es1)

	//write to file
	out1, err := proto.Marshal(es1)
	if err := ioutil.WriteFile("tmp1.out", out1, 0644); err != nil {
		log.Fatalln("failed to writeFile:", err)
	}
}
