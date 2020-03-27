# [go-zookeeper](https://github.com/samuel/go-zookeeper)
安装：`go get -u -v github.com/samuel/go-zookeeper`

|实验|简介|说明|
|---|---|---|
|lab001|example|关闭zk本身的日志 |
|lab002|创建组，加入组，查看列表，删除组| |
|lab003|连接zookeeper集群| |
|lab004|观察模式| |
|lab005|配置服务| |

## 运行命令
 - `go run main.go -c=127.0.0.1:2181 -c=127.0.0.1:2182`
 
## NOTICE
 - `zk.WithLogInfo(false)`关闭库本身的日志
 - 现在不支持zk的rmr和deleteall命令