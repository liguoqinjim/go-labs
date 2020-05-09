/**
* Created by GoLand.
* User: link1st
* Date: 2019-07-31
* Time: 15:17
 */

package task

import (
	"lab225/servers/websocket"
	"log"
	"runtime/debug"
	"time"
)

func Init() {
	Timer(3*time.Second, 30*time.Second, cleanConnection, "", nil, nil)
}

// 清理超时连接
func cleanConnection(param interface{}) (result bool) {
	result = true

	defer func() {
		if r := recover(); r != nil {
			log.Println("ClearTimeoutConnections stop", r, string(debug.Stack()))
		}
	}()

	log.Println("定时任务，清理超时连接", param)
	websocket.ClearTimeoutConnections()

	return
}
