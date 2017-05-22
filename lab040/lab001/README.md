### grpc实验1
按照文档运行一次示例代码

#### 安装grpc
 1. `go get google.golang.org/grpc`
 2. https://github.com/google/protobuf/releases下载对应的`protoc-<version>-<platform>.zip`
 3. 下载好的zip解压之后，把bin文件夹的路径加到系统环境变量PATH里面
 4. `go get -u github.com/golang/protobuf/proto`
 5. `go get -u github.com/golang/protobuf/protoc-gen-go`
 6. 需要把GOPATH/bin加入到环境变量PATH里面

#### 运行结果
![Imgur](http://i.imgur.com/bn969af.png)

![Imgur](http://i.imgur.com/aSMnJXm.png)

#### 参考资料
http://www.grpc.io/docs/quickstart/go.html