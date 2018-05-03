### example02

#### 注意点
 - 使用mockgen命令：`mockgen.exe -destination spider/mock_spider.go -package spider -source spider\spider.go`，
 -destination指定要生成的文件，-package指定生成文件的包名，-source是指对哪个文件mock
 - `mockSpider.EXPECT().GetBody().Return("go1.8.3")`指定了这个方法的返回

#### 参考资料
 - https://www.jianshu.com/p/598a11bbdafb

#### 运行结果
![Imgur](https://i.imgur.com/inbrXn5.png)