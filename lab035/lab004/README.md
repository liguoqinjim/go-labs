### glide实验2
glide有一个问题是，比如我要装name5566的leaf框架，
leaf框架还import了一个websocket的库。要是直接用`go get`的方式安装，那么websocket也是会自动装好的。
但是`glide get`是不会的

#### 参考资料
https://github.com/Masterminds/glide/issues/468