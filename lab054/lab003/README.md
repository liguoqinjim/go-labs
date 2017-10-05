### Path Syntax，用路径的方式得到值

#### 注意
1. gjson用了`#`、`.`、`?`用作特殊符号，但是有的时候我们就是要用来表达原来的意思，这个时候只要在这些符号之前加上。然后在golang中，`\`本身又是要转义的，所以代码里面会是`fav\\.movie`
2. index从1开始

#### 运行结果
![Imgur](https://i.imgur.com/o0SWUJR.png)