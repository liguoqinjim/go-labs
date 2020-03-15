# package内部的测试用例，查看启动顺序

## NOTICE
 - 在运行结果里面我们可以看到，我们是在utils文件夹下，运行`go test`的
 - 可以看到`go test`的输出，是要看utils这个package本身，有没有import别的package。
 所有import的package和utils这个package本身的init方法，都会被调用
 - main是不会调用的

## 运行结果
![Imgur](https://i.imgur.com/vO5WIXc.png)