### 测试any
在proto中使用any

#### 注意
any的使用，还是要在`.proto`文件里面定义几个message的。比如，我定义了三个message，分别是A、B、C
那么我可以在C里面创建一个any字段，我可以在any里面放A或者B属性，C都可以解析。是这样一个功能。

#### 运行结果
![Imgur](http://i.imgur.com/RII6UFB.png)

#### 参考资料
1. https://medium.com/@pokstad/sending-any-any-thing-in-golang-with-protobuf-3-95f84838028d
