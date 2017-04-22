#!/bin/bash
#模拟tcp的client连接

for i in {1..10};
do echo "hello"${i} | nc 127.0.0.1 8881;
done
