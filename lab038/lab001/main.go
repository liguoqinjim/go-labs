package main

import (
	"fmt"
	"io/ioutil"
	"lab038/lab001/pb"
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {
	personNumber := &pb.Person_PhoneNumber{Number: "13999999999", Type: pb.Person_WORK}
	person := pb.Person{Name: "xiaoming", Id: 23, Email: "136542728@qq.com", Phones: []*pb.Person_PhoneNumber{personNumber}}
	fmt.Println(person)

	//write to file
	out, err := proto.Marshal(&person)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	fmt.Println("out:", out)
	if err := ioutil.WriteFile("tmp.out", out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}

	//read from file
	in, err := ioutil.ReadFile("tmp.out")
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	person1 := &pb.Person{}
	if err := proto.Unmarshal(in, person1); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}
	fmt.Println(person1, person1.Id)
}
