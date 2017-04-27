### 简单实验4，重连

#### 注意点
客户端在给服务器发了一个数据之后，我们用`clumsy`来模拟丢包延迟的行为。
在客户端的write返回错误之后，客户端会开一条新的连接，发送一个重连请求。
服务器收到新的连接之后，会校对重连数据。要是正确，会验证是否有数据需要重传。要是有重传的数据则会先把这些数据重传。

#### 运行结果
client图片太长就不截图了

#### 参考资料
1. https://github.com/funny/snet
2. http://blog.csdn.net/lc_910927/article/details/37599161?utm_source=tuicool&utm_medium=referral