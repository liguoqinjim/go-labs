### 示例4，测试Asset文件时的路径
这个lab里面，static里面还有一个image文件夹，那么我们想asset image路径下的文件的时候使用`image/pic_icon.png`这种方式。
我们只要改下命令里面的路径参数就可以了

#### 注意点
 - `go-bindata.exe -pkg assets -o ..\assets\assets.go ./...` 最后使用`./...`这样的方式可以实现用`image/pic_icon.png`的方式访问资源
 - `go-bindata.exe -pkg assets -o ..\assets\assets.go .` 这样的方式实现不了

