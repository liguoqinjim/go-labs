### 调用js文件里面的方法

#### 注意
1. 要调用js文件里面的方法，还是要用`vm.Run()`来调用。run里面写方法名
2. 在compile之后，要把得到script，再run一遍。这样方法才能在vm里面

#### 运行结果
![Imgur](http://i.imgur.com/wJYHuwR.png)