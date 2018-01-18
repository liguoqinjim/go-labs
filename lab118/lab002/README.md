### panics

#### 注意点
 - `CapturePanicAndWait`会运行第一个参数传入的函数，要是函数运行有panic，
 `CapturePanicAndWait`会处理错误，然后传错误到sentry，然后defer这个错误。（就是省下了我们自己defer的过程）

#### 参考资料
https://docs.sentry.io/clients/go/