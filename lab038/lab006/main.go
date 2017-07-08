package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"lab038/lab006/pb"
	"log"
)

func main() {
	items := make([]*pb.Item, 2)
	items[0] = &pb.Item{ItemId: 101, ItemNum: 2}
	items[1] = &pb.Item{ItemId: 102, ItemNum: 3}
	p1 := &pb.Player{PlayerId: 10001, PlayerHero: &pb.Hero{HeroId: 55, HeroLife: 100}, PlayerItem: items, PlayerGold: 1000}
	fmt.Println(p1)
	p2 := &pb.Player{PlayerId: 10002, PlayerGold: 3000}
	fmt.Println(p2)

	//write to file
	out1, err := proto.Marshal(p1)
	if err != nil {
		log.Fatalln("failed to encode p1:", err)
	}
	out2, err := proto.Marshal(p2)
	if err != nil {
		log.Fatalln("failed to encode p2:", err)
	}
	if err := ioutil.WriteFile("tmp1.out", out1, 0644); err != nil {
		log.Fatalln("failed to write file1:", err)
	}
	if err := ioutil.WriteFile("tmp2.out", out2, 0644); err != nil {
		log.Fatalln("failed to write file2:", err)
	}

	//read file
	in1, err := ioutil.ReadFile("tmp1.out")
	if err != nil {
		log.Fatalln("failed to read file1", err)
	}
	in2, err := ioutil.ReadFile("tmp2.out")
	if err != nil {
		log.Fatalln("failed to read file2:", err)
	}
	p11 := new(pb.Player)
	p22 := new(pb.Player)
	if err := proto.Unmarshal(in1, p11); err != nil {
		log.Fatalln("failed to parse1:", err)
	}
	if err := proto.Unmarshal(in2, p22); err != nil {
		log.Fatalln("failed to parse2:", err)
	}
	fmt.Println(p11)
	fmt.Println(p22)

	//判断为空的数据对应的结构体是否为nil
	if p22.PlayerHero == nil {
		fmt.Println("p2.PlayerHero = nil")
	}
	fmt.Printf("p2.PlayerItem,len[%d],cap[%d]\n", len(p22.PlayerItem), cap(p22.PlayerItem))
}
