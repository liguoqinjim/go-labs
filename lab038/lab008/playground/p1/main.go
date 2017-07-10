package main

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/timestamp"
	"log"
)

func main() {
	name := proto.MessageName(&timestamp.Timestamp{})
	log.Printf("Message name of timestamp: %s", name)
}
