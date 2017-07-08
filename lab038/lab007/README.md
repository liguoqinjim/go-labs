### 测试Enum

#### 注意
1. 在protobuf中的enum字段要是没数据的话，golang中会用把这个变量赋值为默认值，也就是enum的第一个值
2. enum可以单独定义，也可以定义在一个message中
3. 用`allow_alias`参数可以设置重复的enum

#### 运行结果
![Imgur](http://i.imgur.com/H6vBqg6.png)

#### 参考资料
https://developers.google.com/protocol-buffers/docs/proto3#enum