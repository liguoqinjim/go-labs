### 实验1,`net.InterfaceAddrs`，去除回环地址和ipv6
`net.InterfaceAddrs`会返回本机的所有地址。代码中有去掉回环地址和ipv6

#### 注意点
这个代码虽然去掉了回环地址和ipv6的地址，但是因为我们要是用到了虚拟机和docker，
我们的机器上都会多一个虚拟网卡，这个地址用这个代码是去不掉的。

#### 运行结果
![Imgur](http://i.imgur.com/jsm53Qp.png)

#### 参考资料
http://golangtc.com/t/5111b98b320b5209f4000004
