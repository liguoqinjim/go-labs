### 调用单个js文件
调用一个外部的js文件，但是这个里面也不是直接调用，而是compile这个js文件，然后可以在内部运行这个代码

#### 注意
`vm.Compile("js/test1.js", nil`，只有第二个参数为nil的时候才会去读取指定的这个文件

#### 运行结果
![Imgur](http://i.imgur.com/96tnM9Z.png)
