package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"lab038/lab005/pb"
	"log"
)

func main() {
	pi1 := &pb.PlayerItem{PlayerId: 10001, Items: map[int32]int32{1: 2, 2: 3}, ItemCnt: 2}
	pi2 := &pb.PlayerItem{PlayerId: 10002, ItemCnt: 1}
	fmt.Println(pi1)
	fmt.Println(pi2)
	out1, err := proto.Marshal(pi1)
	if err != nil {
		log.Fatalln("failed to encode pi1:", err)
	}
	out2, err := proto.Marshal(pi2)
	if err != nil {
		log.Fatalln("failed to encode pi2:", err)
	}

	//write to file
	if err := ioutil.WriteFile("tmp1.out", out1, 0644); err != nil {
		log.Fatalln("failed to write file1:", err)
	}
	if err := ioutil.WriteFile("tmp2.out", out2, 0664); err != nil {
		log.Fatalln("failed to write file2:", err)
	}

	//read file
	in1, err := ioutil.ReadFile("tmp1.out")
	if err != nil {
		log.Fatalln("failed to read file1:", err)
	}
	in2, err := ioutil.ReadFile("tmp2.out")
	if err != nil {
		log.Fatalln("failed to read file2:", err)
	}
	p11 := new(pb.PlayerItem)
	p22 := new(pb.PlayerItem)
	if err := proto.Unmarshal(in1, p11); err != nil {
		log.Fatalln("failed to parse in1:", err)
	}
	if err := proto.Unmarshal(in2, p22); err != nil {
		log.Fatalln("failed to parse in2:", err)
	}
	fmt.Println(p11)
	fmt.Println(p22)

	//判断map是否是nil
	if p22.Items == nil {
		fmt.Println("p22.Items = nil")
	} else {
		fmt.Println("p22.Items != nil")
	}
	fmt.Printf("p22.Items:len[%d]\n", len(p22.Items))
}
