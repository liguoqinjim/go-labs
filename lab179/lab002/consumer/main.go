package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
)

func main() {
	// make a new reader that consumes from topic-A, partition 0, at offset 42
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"hadoop000:9092"},
		Topic:     "topic-A",
		Partition: 0,
		//MinBytes:  10e3, // 10KB
		//MaxBytes:  10e6, // 10MB
	})
	r.SetOffset(42)

	log.Println("开始接收消息：")
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}

	r.Close()

}
