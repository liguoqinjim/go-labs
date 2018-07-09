### 实验3
在consumer中不是直接连接nsqd，而是连接nsqlookupd来找到nsqd的服务。

#### 注意点
 - 一开始用这样的连接方式的时候，是在nsqlookupd服务上设置了`broadcast-address`参数的，
   但是改了之后还是收到的不是我们设置的参数。然后在实验4的时候发现了问题，原来是要是在nsqd服务上来设置这个参数的。一般的链接是`192.168.116.130:4161`
 - producer是不能连接nsqlookupd的


