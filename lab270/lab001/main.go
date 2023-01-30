package main

import (
	"fmt"
	"github.com/solywsh/chatgpt"
	"log"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("参数不正确，需要输入openai key")
	}
	openAIKey := os.Args[1]

	// The timeout is used to control the situation that the session is in a long and multi session situation.
	// If it is set to 0, there will be no timeout. Note that a single request still has a timeout setting of 30s.
	chat := chatgpt.New(openAIKey, "user_id(not required)", 30*time.Second)
	defer chat.Close()
	//
	//select {
	//case <-chat.GetDoneChan():
	//	fmt.Println("time out/finish")
	//}
	question := "你认为2022年世界杯的冠军是谁？"
	fmt.Printf("Q: %s\n", question)
	answer, err := chat.Chat(question)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("A: %s\n", answer)

	//Q: 你认为2022年世界杯的冠军是谁？
	//A: 这个问题很难回答，因为2022年世界杯还没有开始，所以没有人知道冠军是谁。
}
