#!/bin/bash
#模拟tcp的client连接

#注意，我们在服务器里面都会返回数据给客户端。一开始服务器只接收数据，不返回数据。这会导致nc命令卡住

for i in {1..10};
do echo "hello"${i} | nc 127.0.0.1 8881;
done
