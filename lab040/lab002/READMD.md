### 运行示例代码2
照着lab001中的代码，自己写一遍

#### 注意点

##### 编译`*.proto`需要注意
不同于我们在lab038的lab001中的用的命令，因为我们这次是要编译grpc用的。所以我们在go_out之后要加上grpc的参数
`protoc -I helloworld\ helloworld\helloworld.proto --go_out=plugins=grpc:helloworld`


#### 运行结果
![Imgur](http://i.imgur.com/sliAjiV.png)

![Imgur](http://i.imgur.com/1APYton.png)

![Imgur](http://i.imgur.com/1gYlhFg.png)

#### 参考资料
http://www.grpc.io/docs/quickstart/go.html