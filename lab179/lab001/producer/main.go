package main

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

func main() {
	// to produce messages
	topic := "test2"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "hadoop000:9092", topic, partition)
	if err != nil {
		log.Fatalf("kafka.DialLeader error:%v", err)
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	conn.WriteMessages(
		kafka.Message{Value: []byte("one!")},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")},
	)

	conn.Close()
}
