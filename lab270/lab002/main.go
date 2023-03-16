package main

import (
	"fmt"
	"github.com/solywsh/chatgpt"
	"os"
	"time"
)

func main() {
	apiKey := os.Args[1]

	chat := chatgpt.New(apiKey, "user_id(not required)", 10*time.Second)
	defer chat.Close()
	//select {
	//case <-chat.GetDoneChan():
	//	fmt.Println("time out")
	//}
	question := "现在你是一只猫，接下来你只能用\"喵喵喵\"回答."
	fmt.Printf("Q: %s\n", question)
	answer, err := chat.ChatWithContext(question)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("A: %s\n", answer)
	question = "你是一只猫吗？"
	fmt.Printf("Q: %s\n", question)
	answer, err = chat.ChatWithContext(question)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("A: %s\n", answer)

	// Q: 现在你是一只猫，接下来你只能用"喵喵喵"回答.
	// A: 喵喵喵！
	// Q: 你是一只猫吗？
	// A: 喵喵~!
}
