# 实验1：ping set get

## NOTICE
 - set key的时候直接set time.Now，那么redis里面存的是`2020-03-25T14:44:03+08:00`这样的格式。
 get key的时候可以调用Time直接转为time.Time
 

## 参考资料
 - https://github.com/go-redis/redis