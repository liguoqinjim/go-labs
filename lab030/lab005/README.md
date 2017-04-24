### 运行示例3，模拟心跳连接
感觉这个示例的逻辑不是太好，里面有个`<-time.After()`的函数好像都到达不了

#### 注意点
`SetDeadline`，`SetReadDeadline`，`SetWriteDeadline`用于设置每次socket连接能够维持的最长时间

#### 运行结果
##### client
![Imgur](http://i.imgur.com/YWYPKGp.png)

##### server
![Imgur](http://i.imgur.com/bMmR70U.png)

#### 参考资料
http://blog.csdn.net/ahlxt123/article/details/47726783