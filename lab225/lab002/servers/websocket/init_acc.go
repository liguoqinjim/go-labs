/**
 * Created by GoLand.
 * User: link1st
 * Date: 2019-07-25
 * Time: 16:04
 */

package websocket

import (
	"github.com/gorilla/websocket"
	"github.com/spf13/viper"
	"lab225/helper"
	"lab225/models"
	"log"
	"net/http"
	"time"
)

var (
	clientManager = NewClientManager() // 管理者
	appIds        = []uint32{101, 102} // 全部的平台

	serverIp   string
	serverPort string
)

func GetAppIds() []uint32 {

	return appIds
}

func GetServer() (server *models.Server) {
	server = models.NewServer(serverIp, serverPort)

	return
}

func IsLocal(server *models.Server) (isLocal bool) {
	if server.Ip == serverIp && server.Port == serverPort {
		isLocal = true
	}

	return
}

func InAppIds(appId uint32) (inAppId bool) {

	for _, value := range appIds {
		if value == appId {
			inAppId = true

			return
		}
	}

	return
}

// 启动程序
func StartWebSocket() {
	serverIp = helper.GetServerIp()

	rpcPort := viper.GetString("app.rpcPort")

	serverPort = rpcPort

	http.HandleFunc("/ws", wsPage)

	// 添加处理程序
	go clientManager.start()
	log.Println("WebSocket 启动程序成功", serverIp, serverPort)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("websocket 8080 error:%v", err)
	}
}

func wsPage(w http.ResponseWriter, req *http.Request) {
	log.Println("waPage function start")

	// 升级协议
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		log.Println("升级协议", "ua:", r.Header["User-Agent"], "referer:", r.Header["Referer"])

		return true
	}}).Upgrade(w, req, nil)
	if err != nil {
		http.NotFound(w, req)

		return
	}

	log.Println("webSocket 建立连接:", conn.RemoteAddr().String())

	currentTime := uint64(time.Now().Unix())
	client := NewClient(conn.RemoteAddr().String(), conn, currentTime)

	go client.read()
	go client.write()

	// 用户连接事件
	clientManager.Register <- client

	log.Println("waPage function end")
}
