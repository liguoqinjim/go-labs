package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"lab038/lab004/pb"
	"log"
)

func main() {
	s1 := &pb.Student{Id: 1, Scores: []int32{60, 70, 80}, Age: 20}
	fmt.Println(s1)

	s2 := &pb.Student{Id: 2, Age: 30}
	fmt.Println(s2)

	//write to file
	out1, err := proto.Marshal(s1)
	if err != nil {
		log.Fatalln("Failed to encode s1:", err)
	}
	fmt.Println("out1:", out1)
	out2, err := proto.Marshal(s2)
	if err != nil {
		log.Fatalln("Failed to encode s2:", err)
	}
	fmt.Println("out2:", out2)
	if err := ioutil.WriteFile("tmp1.out", out1, 0644); err != nil {
		log.Fatalln("Failed to write file1:", err)
	}
	if err := ioutil.WriteFile("tmp2.out", out2, 0644); err != nil {
		log.Fatalln("Failed to write file2:", err)
	}

	//read from file
	in1, err := ioutil.ReadFile("tmp1.out")
	if err != nil {
		log.Fatalln("Failed to read file1:", err)
	}
	in2, err := ioutil.ReadFile("tmp2.out")
	if err != nil {
		log.Fatalln("Failed to read file2:", err)
	}
	s11 := new(pb.Student)
	s22 := new(pb.Student)
	if err := proto.Unmarshal(in1, s11); err != nil {
		log.Fatalln("Failed to s11:", err)
	}
	if err := proto.Unmarshal(in2, s22); err != nil {
		log.Fatalln("Failed to s22:", err)
	}
	fmt.Println(s11)
	fmt.Println(s22)
}
