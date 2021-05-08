# +build
做这个实验是因为lab042中用到的库，分别有三个系统自己的文件。做下这个实验确定下是什么来触发这个忽略。

## 实验结果
 1. `hello_windows.go`和`hello_linux.go`这样的文件是可以不加`build tag`，在golang中也会被识别。
但是`hello_bsd.go`没有被识别
 2. 不管是什么文件名，我们可以主动在该文件的第一行加上`build tag`。这样我们也可以指定这个文件可以在哪个系统上被使用，
`// +build linux darwin dragonfly freebsd netbsd openbsd`

## 运行结果
### windows
![Imgur](http://i.imgur.com/FtDssyP.png)

### ubuntu
![Imgur](http://i.imgur.com/zyJE8bR.png)


## 参考资料
 1. https://golang.org/pkg/go/build/#hdr-Build_Constraints
 2. http://studygolang.com/articles/5035