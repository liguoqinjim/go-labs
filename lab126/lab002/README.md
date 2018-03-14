### 附加到运行中的go程序调试

#### 注意点
 - `dlv attach pid`，pid是程序进程的id
 - `b E:\Workspace\go-labs\src\lab126\lab002\main.go:20`打断点
 - `c`运行
 - c之后，运行程序，这样遇到断点就会停下

#### 运行结果
![Imgur](https://i.imgur.com/fpm7xrs.png)