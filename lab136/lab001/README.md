###

#### 注意点
 - mockgen的使用：1.先创建mock文件夹，2.`mockgen -source=hellomock.go > mock/mock_Talker.go`
 - `gomock.Eq("王尼玛")`，Eq方法要求测试用例里面的调用参数，一定要等于Eq方法设置的值。这也就是给定参数`张全蛋`的时候会发生错误的原因
 - `gomock.Any()`和`gomock.Eq`相反，也就是可以传入任何值
 - `Return`方法就是给定函数的返回

#### 运行结果
##### `t.Log(company.Meeting("王尼玛"))`
![Imgur](https://i.imgur.com/HD7YcO9.png)

##### `t.Log(company.Meeting("张全蛋"))`
![Imgur](https://i.imgur.com/gT4LGY2.png)