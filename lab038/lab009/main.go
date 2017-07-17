package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"lab038/lab009/pb"
	"log"
)

func main() {
	mailItem1 := &pb.MailItem{ItemId: 1, ItemNum: 100}
	mailItem2 := &pb.MailItem{ItemId: 2, ItemNum: 200}
	mailItems := &pb.MailItems{MailItems: []*pb.MailItem{mailItem1, mailItem2}}
	back1 := &pb.Back{PlayerId: 1, MailItems: mailItems}
	fmt.Println("back1:", back1)

	back2 := &pb.Back{PlayerId: 2}
	fmt.Println("back2:", back2)

	//写文件
	out1, err := proto.Marshal(back1)
	if err != nil {
		log.Fatalln("failed to marshal1:", err)
	}
	if err := ioutil.WriteFile("tmp1.out", out1, 0644); err != nil {
		log.Fatalln("failed to writeFile1:", err)
	}
	out2, err := proto.Marshal(back2)
	if err != nil {
		log.Fatalln("failed to marshal2:", err)
	}
	if err := ioutil.WriteFile("tmp2.out", out2, 0644); err != nil {
		log.Fatalln("failed to writeFile2:", err)
	}

	//读文件
	in1, err := ioutil.ReadFile("tmp1.out")
	if err != nil {
		log.Fatalln("failed to readFile1:", err)
	}
	in2, err := ioutil.ReadFile("tmp2.out")
	if err != nil {
		log.Fatalln("failed to readFile2:", err)
	}

	//unmarshal
	back11 := new(pb.Back)
	if err := proto.Unmarshal(in1, back11); err != nil {
		log.Fatalln("failed to unmarshal1:", err)
	}
	fmt.Println("back11:", back11)
	back22 := new(pb.Back)
	if err := proto.Unmarshal(in2, back22); err != nil {
		log.Fatalln("failed to unmarshal2:", err)
	}
	fmt.Println("back22:", back22)
	if back22.MailItems == nil {
		fmt.Println("back22.MailItems is nil")
	} else {
		fmt.Println("back22.MailItems is not nil")
	}
}
