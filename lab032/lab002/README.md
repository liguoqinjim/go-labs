###简单压力测试

####运行结果
`go test -bench=.`不加-bench的时候不会运行BenchmarkXXX的方法。
函数名后面跟的`-8`表示cpu的数量，`200000000`表示循环次数，`1.08 ns/op`表示平均每次执行时间

![Imgur](http://i.imgur.com/lyUEW3b.png)

####参考资料
https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/11.3.md