### 使用代理

#### 注意点
 - 代理的格式要注意，不能是`http://1.2.3.4:5`这样的格式，而是直接`ip:port`的格式，因为这是设置给浏览器的。
 不是我们直接用代码代理的时候。
 - proxyType要设置为`manual`

#### 参考资料
 - https://github.com/SeleniumHQ/selenium/wiki/DesiredCapabilities#proxy-json-object

#### 运行结果
![Imgur](https://i.imgur.com/2N5MPOI.png)