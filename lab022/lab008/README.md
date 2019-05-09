# closing channels
Closing a channel indicates that no more values will be sent on it. This can be useful to communicate completion to the channel’s receivers.

## NOTICE
 - 两个条件都满足的情况下，1是channel被close了，2是channel里面的数据都读取receive掉了。  
   这时候more就会返回true

## 运行结果
![Imgur](https://i.imgur.com/KEsTAqe.png)

## 资料
https://gobyexample.com/closing-channels