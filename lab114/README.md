### 得到完整url

#### 注意点
 - golang里面不能直接得到完成的链接，需要自己来拼接出完整的链接
 - 这里是通过request.TLS是否为空来判断是http还是https的，为空的时候是http

#### 运行结果
先访问`http://localhost:9090/index?id=1`

![Imgur](https://i.imgur.com/QRCyk0E.png)

#### 参考资料
http://www.01happy.com/golang-get-full-url/