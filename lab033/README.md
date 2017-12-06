### [gops](https://github.com/google/gops)
gops is a command to list and diagnose Go processes currently running on your system.

### gops命令
#### gops
显示所有正在运行的go进程，带星号的表示gops可以跟踪的进程

![Imgur](http://i.imgur.com/c2hkc9u.png)

#### 其他命令
|命令|作用|
|---|---|
|gops stack pid\|addr|堆栈跟踪|
|gops memstats pid\|addr|内存信息|
|gops gc pid\|addr|强制gc|
|gops version pid\|addr|查看进程使用的go版本|
|gops stats pid\|addr|查看goroutines和GOMAXPROCS信息|
|gops pprof -cpu pid\|addr|CPU profile|
|gops pprof -heap pid\|addr|heap profile|
|gops trace pid\|addr|runtime tracer|

#### 参考资料
https://github.com/google/gops


