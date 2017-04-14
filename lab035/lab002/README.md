### 简单实验2

#### 运行命令
##### 添加本地库
`govendor add github.com/liguoqinjim/pholcus`

这样就可以把本地gopath中的package加入项目中

##### 删除库
`govendor remove github.com/liguoqinjim/pholcus`

##### 有的库加不上的问题
如`github.com/emirpasic/gods`这个库到目前写这个文档为止，用govendor来控制是有点问题的。
初步的猜测装不上的问题是这个库的主目录下是没有文件的。在我自己写了一个一般的main.go文件之后，就可以安装了。但是只安装到了根目录，可能govendor安装的时候是按照import来判断是否要安装其他的库的，而不是只靠判断文件夹层次。

#### 参考资料
https://github.com/kardianos/govendor