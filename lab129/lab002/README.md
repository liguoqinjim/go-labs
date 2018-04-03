### sync.Map的range

#### 注意点
 - range不会受store或者delete的影响
 - 一个key最多只会被range到一次
 - 在range的时候，新加入的key是不会被range到的

#### 运行结果
![Imgur](https://i.imgur.com/6Hv4zYE.png)