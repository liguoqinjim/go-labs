# 修改Collector的配置

## NOTICE
 - collector是有默认配置的，同时也可以在创建collector时自定义配置
 - collector的配置不仅可以在创建的时候指定，在任何时候都是可以修改的。在OnRequest方法里面也可以修改(比如可以每次请求前随机一个useragent)。
 - colly用的是golang默认的http client，client也是可以自定义的
 - `AllowURLRevisit`默认为false，为false的时候相同的url第二次就不会访问，还会报`URL already visited`错误

## 运行结果
![Imgur](https://imgur.com/zlvTmb8)

## 参考资料
 - http://go-colly.org/docs/introduction/configuration/