package main

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

func main() {
	// to consume messages
	topic := "test2"
	partition := 0

	conn, _ := kafka.DialLeader(context.Background(), "tcp", "hadoop000:9092", topic, partition)

	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	b := make([]byte, 10e3) // 10KB max per message
	for {
		_, err := batch.Read(b)
		if err != nil {
			break
		}
		log.Println(string(b))
	}

	batch.Close()
	conn.Close()
}
