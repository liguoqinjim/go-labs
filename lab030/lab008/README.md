### 和logstash的socket连接，并发送数据

#### 注意
要让logstash每次收到消息的时候都单独算，要在发送内容的最后加一个`\n`，不然logstash会算在一起
