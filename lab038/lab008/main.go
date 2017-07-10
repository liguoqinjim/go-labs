package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	_ "github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"io/ioutil"
	"lab038/lab008/pb"
	"log"
)

func main() {
	es1 := &pb.ErrorStatus{Message: "error1"}
	es2 := &pb.ErrorStatus{Message: "error2"}
	player := &pb.Player{PlayerId: 10001, PlayerName: "firstPlayer", PlayerGold: 2000}
	playerHero := &pb.PlayerHero{Id: 90001, HeroId: 3, HeroLv: 1, HeroLife: 100}

	//处理es1
	data1, err := proto.Marshal(player)
	if err != nil {
		log.Fatalln("failed to marshal player:", err)
	}
	es1.Params = &any.Any{
		TypeUrl: "/" + proto.MessageName(player),
		Value:   data1,
	}
	//处理es2
	data2, err := proto.Marshal(playerHero)
	if err != nil {
		log.Fatalln("failed to marshal playerHero:", err)
	}
	es2.Params = &any.Any{
		TypeUrl: "/" + proto.MessageName(playerHero),
		Value:   data2,
	}

	//写文件
	out1, err := proto.Marshal(es1)
	if err != nil {
		log.Fatalln("failed to marshal es1:", err)
	}
	if err := ioutil.WriteFile("tmp1.out", out1, 0644); err != nil {
		log.Fatalln("failed to write file1:", err)
	}
	out2, err := proto.Marshal(es2)
	if err != nil {
		log.Fatalln("failed to marshal es2:", err)
	}
	if err := ioutil.WriteFile("tmp2.out", out2, 0644); err != nil {
		log.Fatalln("failed to write file2:", err)
	}

	//读文件
	in1, err := ioutil.ReadFile("tmp1.out")
	if err != nil {
		log.Fatalln("failed to read file1:", err)
	}
	in2, err := ioutil.ReadFile("tmp2.out")
	if err != nil {
		log.Fatalln("failed to read file2:", err)
	}

	//parse
	es11 := new(pb.ErrorStatus)
	es22 := new(pb.ErrorStatus)
	if err := proto.Unmarshal(in1, es11); err != nil {
		log.Fatalln("failed to parse1:", err)
	}
	if err := proto.Unmarshal(in2, es22); err != nil {
		log.Fatalln("failed to parse2:", err)
	}
	fmt.Println(es11)
	fmt.Println(es22)
	//parse any
	player2 := new(pb.Player)
	playerHero2 := new(pb.PlayerHero)
	if err := ptypes.UnmarshalAny(es11.Params, player2); err != nil {
		log.Fatalln("failed to parse any1:", err)
	}
	if err := ptypes.UnmarshalAny(es22.Params, playerHero2); err != nil {
		log.Fatalln("failed to parse any2:", err)
	}
	fmt.Println("player2:", player2)
	fmt.Println("playerHero2:", playerHero2)
}
