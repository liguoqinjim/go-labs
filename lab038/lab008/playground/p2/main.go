package main

//go:generate protoc --go_out=. anything.proto

import (
	"log"
	"reflect"

	"fmt"
	proto "github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	any "github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes/timestamp"
	"lab038/lab008/playground/p2/pb"
)

func main() {
	t1 := &timestamp.Timestamp{
		Seconds: 5, // easy to verify
		Nanos:   6, // easy to verify
	}
	fmt.Println("t1=", t1)

	serialized, err := proto.Marshal(t1)
	if err != nil {
		log.Fatal("could not serialize timestamp")
	}
	fmt.Println("serialized=", serialized)

	// Blue was a great album by 3EB, before Cadgogan got kicked out
	// and Jenkins went full primadonna
	a := anything.AnythingForYou{
		Anything: &any.Any{
			TypeUrl: "example.com/yaddayaddayadda/" + proto.MessageName(t1),
			Value:   serialized,
		},
	}
	fmt.Println("proto.MessageName=", proto.MessageName(t1))

	// marshal to simulate going on the wire:
	serializedA, err := proto.Marshal(&a)
	if err != nil {
		log.Fatal("could not serialize anything")
	}
	fmt.Println("serializedA=", serializedA)

	// unmarshal to simulate coming off the wire
	var a2 anything.AnythingForYou
	if err := proto.Unmarshal(serializedA, &a2); err != nil {
		log.Fatal("could not deserialize anything")
	}

	// unmarshal the timestamp
	var t2 timestamp.Timestamp
	if err := ptypes.UnmarshalAny(a2.Anything, &t2); err != nil {
		log.Fatalf("Could not unmarshal timestamp from anything field: %s", err)
	}
	fmt.Println("t2=", t2)

	// Verify the values are as expected
	if !reflect.DeepEqual(t1, &t2) {
		log.Fatalf("Values don't match up:\n %+v \n %+v", t1, t2)
	}
}
