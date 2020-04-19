package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/hrpc"
	"log"
	"time"
)

var (
	tableName = []byte("FileTable")
	family01  = "fileInfo"
	family02  = "saveInfo"

	hr = "---------------------------------"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	example()
}

func example() {
	client := gohbase.NewClient("hp001:61002")
	defer client.Close()

	//插入一行数据
	values := map[string]map[string][]byte{family01: {"name": []byte("file3.txt"), "type": []byte("txt")}}
	putReq, err := hrpc.NewPut(context.Background(), tableName, []byte("rowkey3"), values)
	if err != nil {
		log.Fatalf("hrpc.NewPut error:%v", err)
	}
	if result, err := client.Put(putReq); err != nil {
		log.Fatalf("client.Put error:%v", err)
	} else {
		log.Printf("result:%+v", result)
	}

	//插入数据，并设置TTL值
	values = map[string]map[string][]byte{family01: {"name": []byte("file4.txt"), "type": []byte("txt")}}
	putReq, err = hrpc.NewPut(context.Background(), tableName, []byte("rowkey4"), values, hrpc.TTL(time.Second*30))
	if err != nil {
		log.Fatalf("hrpc.NewPut error:%v", err)
	}
	if result, err := client.Put(putReq); err != nil {
		log.Fatalf("client.Put error:%v", err)
	} else {
		log.Printf("result:%+v", result)
	}
}
