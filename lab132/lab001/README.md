### ping

#### 注意点
 - 在mac上使用的时候，直接运行会不对。需要使用root，但是`sudo go run main.go`这样的运行的时候会报找不到包。
 最方便的就是go build之后，sudo运行build好的执行文件