### 实验7
验证多个consumer连接同一个topic，但是连接不同的channel的时候会发生的情况。

#### 实验结论
要是`Consumer1`和`Consumer2`都连接到同一个的topic的不同的channel的时候，`producer`发送的消息两个人都会收到。
但是要是两个consumer连接到的是同一个channel的时候，nsq会在两个consumer里面随机一个来发送，也就是这个时候，producer发出的消息只会被一个consumer收到。

#### 参考资料
http://nsq.io/overview/design.html#simplifying-configuration-and-administration