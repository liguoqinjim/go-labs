### glide实验1

#### 使用命令
##### `glide create`
在一个project中初始化vendor文件夹和glide相关文件

##### `glide get xxx`
在glide中安装包

做好上面这两步就可以直接写程序了，之后运行的时候go会自动识别我们装在vendor中的包的

#### 运行结果
![Imgur](http://i.imgur.com/o1bGw69.png)

第一行的输出的`这是vendor文件夹`就是在vendor中一个包中加了一个init函数输出的

#### 参考资料
https://github.com/Masterminds/glide