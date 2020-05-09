# 测试close chan

## NOTICE
 - `job := <- jobs`这个代码在chan没有close的时候，要是没有数据会阻塞。但是chan被close之后，就不会阻塞了。要注意退出循环。

## 运行结果
![Imgur](https://i.imgur.com/6NLRpqc.png)