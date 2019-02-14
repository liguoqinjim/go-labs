package main

import (
	"context"
	"github.com/segmentio/kafka-go"
)

func main() {
	// make a writer that produces to topic-A, using the least-bytes distribution
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"hadoop000:9092"},
		Topic:    "topic-A",
		Balancer: &kafka.LeastBytes{},
	})

	w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("Key-A"),
			Value: []byte("Hello World!"),
		},
		kafka.Message{
			Key:   []byte("Key-B"),
			Value: []byte("One!"),
		},
		kafka.Message{
			Key:   []byte("Key-C"),
			Value: []byte("Two!"),
		},
	)

	w.Close()
}
