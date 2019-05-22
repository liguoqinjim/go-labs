package main

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// make a new reader that consumes from topic-A, partition 0, at offset 42
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"tdh103:9092"},
		Topic:     "HBASE_MAC_HIT",
		GroupID:   "go-labs",
		Partition: 0,
		//MinBytes:  10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
	r.SetOffset(42)

	log.Println("开始接收消息：")

	total := 0
	//go func() {
	//	for {
	//		m, err := r.ReadMessage(context.TODO())
	//		if err != nil {
	//			break
	//		}
	//		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	//
	//		total++
	//	}
	//}()

	//速度快
	go func() {
		for {
			m, err := r.FetchMessage(context.TODO())
			if err != nil {
				break
			}
			_ = m
			total++

			if total%1000 == 0 {
				r.CommitMessages(context.TODO(), m)
				log.Println("total=", total)
			}
		}
	}()

	<-sigs
	r.Close()

	log.Printf("收到消息总数:%d", total)
}
