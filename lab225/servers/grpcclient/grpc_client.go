/**
* Created by GoLand.
* User: link1st
* Date: 2019-08-03
* Time: 16:43
 */

package grpcclient

import (
	"context"
	"google.golang.org/grpc"
	"lab225/common"
	"lab225/models"
	"lab225/protobuf"
	"log"
	"time"
)

// rpc client
// 给全体用户发送消息
// link::https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_client/main.go
func SendMsgAll(server *models.Server, seq string, appId uint32, userId string, cmd string, message string) (sendMsgId string, err error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(server.String(), grpc.WithInsecure())
	if err != nil {
		log.Println("连接失败", server.String())

		return
	}
	defer conn.Close()

	c := protobuf.NewAccServerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := protobuf.SendMsgAllReq{
		Seq:    seq,
		AppId:  appId,
		UserId: userId,
		Cms:    cmd,
		Msg:    message,
	}
	rsp, err := c.SendMsgAll(ctx, &req)
	if err != nil {
		log.Println("给全体用户发送消息", err)

		return
	}

	if rsp.GetRetCode() != common.OK {
		log.Println("给全体用户发送消息", rsp.String())

		return
	}

	sendMsgId = rsp.GetSendMsgId()
	log.Println("给全体用户发送消息 成功:", sendMsgId)

	return
}

// 获取用户列表
// link::https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_client/main.go
func GetUserList(server *models.Server) (userIds []string, err error) {
	userIds = make([]string, 0)

	conn, err := grpc.Dial(server.String(), grpc.WithInsecure())
	if err != nil {
		log.Println("连接失败", server.String())

		return
	}
	defer conn.Close()

	c := protobuf.NewAccServerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := protobuf.GetUserListReq{}
	rsp, err := c.GetUserList(ctx, &req)
	if err != nil {
		log.Println("获取用户列表 发送请求错误:", err)

		return
	}

	if rsp.GetRetCode() != common.OK {
		log.Println("获取用户列表 返回码错误:", rsp.String())

		return
	}

	userIds = rsp.GetUserId()
	log.Println("获取用户列表 成功:", userIds)

	return
}
