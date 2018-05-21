### example04

#### 注意点
 - 测试的时候可以使用monkey patch来做mock，有的时候这样可以方便测试
 - 运行main的时候，不会输出，因为`reportExecFailed`方法里面用了`os.Exit()`，这样就比较不好测试了。
 所以我们可以用monkey，在测试里面，把`reportExecFailed`修改，去掉`os.Exit()`

#### 运行结果
![Imgur](https://i.imgur.com/0wSKKGW.png)