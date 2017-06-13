### 测试WaitVisible的响应速度
这个实验测试`cdp.WaitVisible`的响应速度。
为什么会有两个测试html文件是因为一开始程序测试的时候一直没有检测到visible，一开始怀疑是不是css属性不一样的问题。
time1里面用的是`visibility`，time2里面用的是`display`。
后面发现并不是这个问题，问题出在`cdp.ByID`，还是要加上这个参数才能很好的运行。

#### 运行结果(只截图time1.html的结果)
![Imgur](http://i.imgur.com/u1wc8Ud.png)

![Imgur](http://i.imgur.com/M8HL5aL.png)
