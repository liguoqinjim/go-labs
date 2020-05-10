# lab002
lab001的基础上精简，删除grpc和分布式

## websocket消息
 - `{"seq":"1588830137763-152805","cmd":"login","data":{"userId":"刘伶","appId":101}}`
 - `{"seq":"1588830167767-643961","cmd":"heartbeat","data":{}}`
 
## NOTICE
 - websocket的连接保存在Client结构体里面，client都存在map里面，通过key来完成相关功能比如广播，私聊