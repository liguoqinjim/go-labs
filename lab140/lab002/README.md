### 修改Collector的配置

#### 注意点
 - colly的collector是有很多默认配置的，但是我们可以在创建collector直接指定我们的自定义配置
 - collector的配置不仅可以在创建的时候指定，在任何时候都是可以修改的。在OnRequest方法里面也可以修改。
 - colly用的是golang默认的http client，这个client我们也是可以自定义的

#### 运行结果
![Imgur](https://i.imgur.com/TVek2a0.png)

#### 参考资料
 - http://go-colly.org/docs/introduction/configuration/