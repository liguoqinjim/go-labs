package main

import (
	ticktick "github.com/ziyixi/go-ticktick"
	"log"
	"os"
)

func main() {
	//读取命令行参数
	args := os.Args
	username := args[1]
	password := args[2]

	client, err := ticktick.NewClient(username, password, "dida365")
	if err != nil {
		log.Fatalf("NewClient error:%v", err)
	}

	err = client.GetToken()
	if err != nil {
		log.Fatalf("GetToken error:%v", err)
	}

	task := &ticktick.TaskItem{
		ProjectId:   "5e16a194a75d5105d64c752f",
		ProjectName: "",
		ParentId:    "",
		Title:       "测试任务",
		IsAllDay:    false,
		Tags:        nil,
		Content:     "",
		Desc:        "",
		AllDay:      false,
		StartDate:   "",
		DueDate:     "",
		TimeZone:    "",
		Reminders:   nil,
		Repeat:      "",
		Priority:    0,
		SortOrder:   0,
		Kind:        "",
		Status:      0,
	}

	t, err := client.CreateTask(task)
	if err != nil {
		log.Fatalf("CreateTask error:%v", err)
	}

	log.Printf("task:%+v", t)
}
