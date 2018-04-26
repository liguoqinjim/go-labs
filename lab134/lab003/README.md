### 连接zookeeper集群

#### 注意点
 - 连接集群的话只要conns写多个就可以了。要是连接的服务器有断开的话，会自动切换到别的服务器上
 - 断开的时候会收到event的

#### 运行结果
![Imgur](https://i.imgur.com/GglLnKn.png)