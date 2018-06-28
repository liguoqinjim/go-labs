### 注册服务
这个就和zookeeper里面的临时节点类似。因为etcd使用的是ttl来保持节点健康。

#### 注意点
 - 使用ttl保持节点健康，也就是要隔一段时间就put给服务器。超时没有收到，这个节点就会被删掉。
 - 使用keepalive来保活，不然就会被删掉了
 - keepAlive会返回一个channel，这个channel我们是要处理的，不然keepalive会一直去发送连接给etcd。具体可以看keepAlive源码上面的注释
 - grant方法是去获得一个租约，会有一个lease id，进行put操作的时候要用到这个id，之后也是对这个id进行keep alive

#### 运行结果
![Imgur](https://i.imgur.com/wzibN1s.png)

#### 参考资料
 1. https://www.jianshu.com/p/40a0527b7e9d?hmsr=studygolang.com&utm_medium=studygolang.com&utm_source=studygolang.com
 2. 源码里面的注释