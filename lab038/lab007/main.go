package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"lab038/lab007/pb"
	"log"
)

func main() {
	mi1 := &pb.MapItem{MapItemId: 10001, MapItemType: pb.MapItemType_MONSTER}
	mi2 := &pb.MapItem{MapItemId: 10002}
	fmt.Println(mi1, mi1.MapItemType)
	fmt.Println(mi2, mi2.MapItemType)

	//write to file
	out1, err := proto.Marshal(mi1)
	if err != nil {
		log.Fatalln("failed to encode1:", err)
	}
	out2, err := proto.Marshal(mi2)
	if err != nil {
		log.Fatalln("failed to encode2:", err)
	}
	if err := ioutil.WriteFile("tmp1.out", out1, 0644); err != nil {
		log.Fatalln("failed to write file1:", err)
	}
	if err := ioutil.WriteFile("tmp2.out", out2, 06444); err != nil {
		log.Fatalln("failed to write file2:", err)
	}

	//read file
	mi11 := new(pb.MapItem)
	mi22 := new(pb.MapItem)
	in1, err := ioutil.ReadFile("tmp1.out")
	if err != nil {
		log.Fatalln("failed to read file1:", err)
	}
	in2, err := ioutil.ReadFile("tmp2.out")
	if err != nil {
		log.Fatalln("failed to read file2:", err)
	}
	if err := proto.Unmarshal(in1, mi11); err != nil {
		log.Fatalln("failed to parse1:", err)
	}
	if err := proto.Unmarshal(in2, mi22); err != nil {
		log.Fatalln("failed to parse2:", err)
	}
	fmt.Println(mi11, mi11.MapItemType)
	fmt.Println(mi22, mi22.MapItemType)
}
