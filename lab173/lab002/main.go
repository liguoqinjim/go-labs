package main

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"io/ioutil"
	"log"
)

func main() {
	example()
}

func example() {
	conf := readConf()

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": conf.Servers})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	// Produce messages to topic (asynchronously)
	topic := "myTopic"
	for _, word := range []string{"Welcome", "to", "the", "Confluent", "Kafka", "Golang", "client"} {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}, nil)
	}

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)
}

func readConf() *Config {
	data, err := ioutil.ReadFile("../config.json")
	if err != nil {
		log.Fatalf("ioutil.ReadFile error:%v", err)
	}

	var config = &Config{}
	if err := json.Unmarshal(data, config); err != nil {
		log.Fatalf("json.Unmarshal error:%v", err)
	}

	return config
}

type Config struct {
	Servers string `json:"servers"`
}
