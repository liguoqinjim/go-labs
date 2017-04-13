### pprof cpu实验
用`go tool pprof`分析一个最简单的程序，来找出`mins`这个单位到底是什么意思。

#### 运行结果
下图是程序运行结果

![Imgur](http://i.imgur.com/8uzBKBi.png)

用`go tool pprof分析`

![Imgur](http://i.imgur.com/m7vnbOr.png)

`web`命令生成的图

![Imgur](http://i.imgur.com/ulDOHUQ.png)

图里面显示的Duration:20.5079308s，这个时间和我们在程序里面计算的总共的运行时间是相同的。
我们可以认为这个Duration就是我们的程序的实际运行时间。
Duration的下面一行380ms of 380ms，现在看来是所有的goroutine的运行时间加起来。这个图片还看不太出来，我们下面贴一张slg游戏服务器的分析图片

![Imgur](http://i.imgur.com/i2368cb.png)

我们在这个图片里面可以到Duration只有25s，但是下面一行确实49.25s of 49.41s。（通过实验，我们一开始的问题也得到了解答，mins到底是什么意思，其实就是分钟的意思）。
因为下面一行是所有的goroutine加起来的，那么时间确实是有可能比Duration大的

`weblist`命令，我们这里也接着使用slg游戏服务器的数据，`weblist gamedata.BastionTeamFight`

![Imgur](http://i.imgur.com/MaR598T.png)

这里面有几个概念，`flat`,`cum`。根据我这段时间看的文章，flat的意思应该是，这个方法或者这行代码运行一次要多少时间。
cum（cumulative）应该是整个程序运行期间，这个代码运行的时间和(可能运行了很多次)。

#### 参考资料
https://groups.google.com/forum/#!topic/golang-nuts/mi5rXPxx6iI