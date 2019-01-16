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

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": conf.Servers,
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{"myTopic", "^aRegex.*[Tt]opic"}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	c.Close()
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
