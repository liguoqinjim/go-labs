### 运行示例代码
示例代码是直接连接，因为的mongo示例开启了权限验证，所以又加了读取连接配置文件的代码。
推荐还是使用DialWithInfo方法来连接，而不是直接用带用户名和密码的url连接。

#### 运行结果
![Imgur](http://i.imgur.com/vuU9ey1.png)

![Imgur](http://i.imgur.com/zABUWPH.png)

#### 参考资料
1. http://labix.org/mgo
2. https://objectrocket.com/docs/mongodb_go_examples.html#connecting