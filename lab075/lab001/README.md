### 在一个goroutine中开启另一个gorouine

#### 结论
在一个goroutine中开启另一个goroutine，被开启的goroutine不会因为之前的goroutine退出了而退出。

#### 运行结果
![Imgur](http://i.imgur.com/Abcq6Qy.png)