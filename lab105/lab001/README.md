### gzip

#### 注意点
 - `header fields`是可选的，可填可不填。但是填了的话，会影响压缩结果的长度。
 - 查看压缩后的结果要在write调用了close之后查看

#### 参考资料
https://golang.org/pkg/compress/gzip/

#### 运行结果
![Imgur](https://i.imgur.com/9HIrun2.png)