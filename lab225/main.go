/**
* Created by GoLand.
* User: link1st
* Date: 2019-07-25
* Time: 09:59
 */

package main

import (
	"github.com/gin-gonic/gin"
	"lab225/lib/redislib"
	"lab225/routers"
	"lab225/servers/grpcserver"
	"lab225/servers/task"
	"lab225/servers/websocket"
	"log"
	"net/http"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	initRedis()

	router := gin.Default()
	// 初始化路由
	routers.Init(router)
	routers.WebsocketInit()

	// 定时任务
	task.Init()

	// 服务注册
	task.ServerInit()

	go websocket.StartWebSocket()
	// grpc
	go grpcserver.Init()

	if err := http.ListenAndServe(":18080", router); err != nil {
		log.Fatalf("listen and serve at 18080 error:%v", err)
	}
}

func initRedis() {
	redislib.ExampleNewClient()
}
