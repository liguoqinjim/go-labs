### 示例1

#### 注意点
 - 运行`go-bindata`的路径，就是在代码里面读取内容的根目录。比如这个实验里面直接写的是test.json，
 因为是直接在data文件夹里面执行的`go-bindata`
 - `go-bindata`生成的go文件的package name默认都是main，要用pkg参数指定，参考lab002

#### 运行结果
##### go-bindata运行截图
![Imgur](https://i.imgur.com/Vxrw2YK.png)

##### 运行结果
![Imgur](https://i.imgur.com/BGgoen7.png)
