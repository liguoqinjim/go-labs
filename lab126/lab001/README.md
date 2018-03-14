### 调试go文件

#### 注意点
 - `dlv debug main.go`进入调试
 - `b main.hi`或者`b E:/Workspace/go-labs/src/lab126/lab001/main.go:19`都可以打断点
 - `n`单步调试
 - `c`运行
 - `p hostName`打印hostName的值
 - `locals`打印所有的本地变量
 - `args`打印出所有的方法参数信息

#### 运行结果
![Imgur](https://i.imgur.com/ba9EChz.png)