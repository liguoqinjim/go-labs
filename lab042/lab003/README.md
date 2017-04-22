### 实验SO_REUSEPORT是否可以自动分配
在这个实验中，我们一共写了两份服务器代码，server1和server2。两个都是绑定在本机的同一个端口。
这个就是用到了SO_REUSEPORT属性，我们要看下是否系统会自动分配连接分别到两个服务器。

#### 实验结果
我们可以看到下面的截图中，系统是会自动分配理解到两个服务器的。而且非常均匀。

#### 运行结果
![Imgur](http://i.imgur.com/GDklnaE.png)

![Imgur](http://i.imgur.com/eKHFhcy.png)

![Imgur](http://i.imgur.com/J4noXLA.png)

#### 参考资料
1. http://www.blogjava.net/yongboy/archive/2015/02/12/422893.html
2. http://www.blogjava.net/yongboy/archive/2015/02/25/423037.html