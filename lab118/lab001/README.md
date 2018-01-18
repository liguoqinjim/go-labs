### 报告错误，error

#### 注意点
 - 设置User信息，`raven.SetUserContext(&raven.User{ID: "123", Email: "136542728@qq.com"})`
 - 设置tag信息，`raven.CaptureErrorAndWait(errors.New("a > 6"), map[string]string{"myErrorCode": "999001"})`

#### 运行结果
##### 设置SetUserContext
![Imgur](https://i.imgur.com/rDryvay.png)

##### 设置tag
![Imgur](https://i.imgur.com/VWycSw7.png)

#### 参考资料
https://blog.sentry.io/2015/07/07/logging-go-errors