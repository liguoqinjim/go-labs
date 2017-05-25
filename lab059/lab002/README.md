### 实验2，`net.Interfaces`，只留下内网地址
`net.Interfaces`，会返回网卡的信息。

#### 注意点
私有地址有3段，ABC类ip各有一段是留给私有的。
```
//A类 10.0.0.0--10.255.255.255
//B类 172.16.0.0--172.31.255.255  B类要特殊处理，不像A类和C类只要判断是否有前缀就可以
//C类 192.168.0.0--192.168.255.255
```

#### 运行结果
![Imgur](http://i.imgur.com/AhGGQ1M.png)

#### 参考资料
https://github.com/toolkits/net/blob/master/ip.go