### 测试singular
proto3中，默认的Field Rules就是singular，另外的就是repeated。singualar表示可以有0个或1个这个字段，但是也不能超过1个。

#### 注意
实验了一下，要是嵌套的字段。没有赋值的时候就是nil

#### 运行结果
![Imgur](http://i.imgur.com/U8RGy6h.png)

#### 参考资料
https://developers.google.com/protocol-buffers/docs/proto3#specifying-field-rules