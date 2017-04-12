###test和Benchmark放一起

####运行结果
暂时还不知道为什么，我要是写两个TestXxx在xxx_test.go文件里面，那么`go test -bench=.`的时候也不用运行压力测试
所以函数的命名要注意下。

![Imgur](http://i.imgur.com/p0ZuoiX.png)

####参考资料
http://studygolang.com/articles/5494