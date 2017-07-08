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

	}
}
