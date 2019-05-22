# sync.Map使用

## 注意点
 - sync.Map直接可以使用
 - `Load`返回true的时候，表示有这个key。返回false1的饿时候表示没找到这个key
 - `LoadOrStore`返回true的时候，是找到了这个key。false的时候是保存了这个key

## 运行结果
![Imgur](https://i.imgur.com/1Jjq2k8.png)

## 参考资料
1. http://blog.csdn.net/champly/article/details/77622328