### 模拟断开连接
在这个实验中，我们模拟连接成功建立之后，我们主动的把一端的conn关掉。
看下在这种情况下，我们接收数据是怎么样的。

#### 注意点
在这个实验中，我们在服务器端故意close掉了连接，但是在我们close掉之后。
客户端还是能发送一次数据(数据肯定没有被服务器接收，但是没有弹出error)，在第二次的时候会返回error了。

#### 运行结果
##### client
![Imgur](http://i.imgur.com/6vEsPoC.png)

##### server
![Imgur](http://i.imgur.com/XzICG9X.png)