package main

import (
	"context"
	"encoding/json"
	"github.com/sashabaranov/go-openai"
	"log"
	"os"
)

var OpenAIClient *openai.Client

func init() {
	if len(os.Args) != 3 {
		//baseUrl格式:https://api.openai.com/v1
		log.Fatalf("需要输入APIKEY、baseurl")
	}
	apiKey := os.Args[1]
	baseUrl := os.Args[2]

	// 创建client
	clientConfig := openai.DefaultConfig(apiKey)
	clientConfig.BaseURL = baseUrl
	OpenAIClient = openai.NewClientWithConfig(clientConfig)
}

var functionStr = `[
    {
        "name": "open_data_chart",
        "description": "显示数据",
        "parameters": {
            "type": "object",
            "properties": {
                "section": {
                    "type": "string",
                    "description": "要打开的页面，eg. 如果是报表的话，就是report，如果是地图的话，就是map",
                    "enum": [
                        "report",
                        "map"
                    ]
                },
                "location": {
                    "type": "string",
                    "description": "指定的城市，需要在结尾加上市这个字。例子：上海市、北京市、广州市"
                },
                "date": {
                    "type": "string",
                    "description": "当要查看报表的时候，需要指定报表的时间（也就是这个参数是必须的），如果没有指定日期的话就用当前时间的年份。时间格式：2023、2023-05"
                },
                "report_name": {
                    "type": "string",
                    "description": "当要查看报表的时候，这个参数是必须的。",
                    "enum": [
                        "忠诚修理厂的覆盖",
                        "门店覆盖指数",
                        "合格门店的变化分布",
                        "有效修理厂指数",
                        "忠诚修理厂的覆盖",
                        "修理厂扫码指数",
                        "战略产品渗透率"
                    ]
                }
            },
            "required": [
                "section",
                "location"
            ]
        }
    }
]`

func lab001() {
	//解析json
	var functionDefinitions []openai.FunctionDefinition
	if err := json.Unmarshal([]byte(functionStr), &functionDefinitions); err != nil {
		log.Fatalf("解析json失败:%v", err)
	}

	//content := "打开北京的地图"
	content := "打开苏州今年的战略产品渗透率报表"
	resp, err := OpenAIClient.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo0613,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: content,
				},
			},
			Temperature: 0,
			Functions:   functionDefinitions,
		},
	)

	if err != nil {
		log.Printf("ChatCompletion error: %v\n", err)
		return
	}

	log.Println(resp.Choices[0].Message.FunctionCall)
}

func main() {
	lab001()
}
