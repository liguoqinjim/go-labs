### runtime/pprof

1. 先`go build`，生成可执行文件
2. `lab002.exe -cpuprofile=lab002.prof`，运行并生成prof文件
3. 读取prof文件，`go tool pprof lab002.exe lab002.prof`

#### 参考资料
 - https://github.com/hyper0x/go_command_tutorial/blob/master/0.12.md (还可以记录别的信息)
 - http://www.cnblogs.com/yjf512/archive/2012/12/27/2835331.html