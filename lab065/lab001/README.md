### 示例程序

#### 运行过程
##### 1.生成`test.manifest`
`rsrc -manifest test.manifest -o rsrc.syso`

##### 2.编译
可以用`go build`编译出来，但是程序运行的时候会有一个cmd窗口在运行。
要是想去掉cmd窗口，带参数使用编译就可以了。
`go build -ldflags="-H windowsgui"`

#### 运行结果
![Imgur](http://i.imgur.com/vQ2iOSM.png)

#### 参考资料
https://github.com/lxn/walk