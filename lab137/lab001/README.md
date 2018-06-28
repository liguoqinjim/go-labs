### 连接etcd

#### 注意点
 - 这个New方法是会发生连接的，会连接到etcd
 - `DialTimeout`参数可以设置连接超时时间，超过这个时间就会报错了

#### 运行结果
##### 连接错误
![Imgur](https://i.imgur.com/7uNwsX5.png)
