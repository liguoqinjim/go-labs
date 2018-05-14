### 注册服务
这个就和zookeeper里面的临时节点类似。因为etcd使用的是ttl来保持节点健康。

#### 注意点
 - 使用ttl保持节点健康，也就是要隔一段时间就put给服务器。超时没有收到，这个节点就会被删掉。

#### 运行结果
##### 运行结果
![Imgur](https://i.imgur.com/9nhVcVB.png)

##### 查询结果
![Imgur](https://i.imgur.com/fNDjB6U.png)

#### 参考资料
https://www.jianshu.com/p/40a0527b7e9d?hmsr=studygolang.com&utm_medium=studygolang.com&utm_source=studygolang.com