package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"lab038/lab010/pb"
	"log"
)

func main() {
	itemMap := make(map[int]int)
	for i := 1; i <= 1000; i++ {
		itemMap[i] = i * 100
	}

	mapRes1 := new(pb.MapRes1)
	mapRes1.Id = 1
	itemsString := ""
	for itemId, itemNum := range itemMap {
		itemsString += fmt.Sprintf("%d;%d|", itemId, itemNum)
	}
	itemsString = itemsString[:len(itemsString)-1]
	mapRes1.Items = itemsString

	mapRes2 := new(pb.MapRes2)
	mapRes2.Id = 2
	mapRes2.Items = make(map[int32]int32)
	for k, v := range itemMap {
		mapRes2.Items[int32(k)] = int32(v)
	}

	//解析到pb
	data1, err := proto.Marshal(mapRes1)
	if err != nil {
		log.Println("proto.marshal error1")
	}
	data2, err := proto.Marshal(mapRes2)
	if err != nil {
		log.Println("proto.marshal error2")
	}

	log.Printf("data1.len[%d],data2.len[%d]", len(data1), len(data2))
}
