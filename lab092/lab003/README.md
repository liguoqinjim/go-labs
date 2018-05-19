### 指定生成文件的路径
绑定static文件夹中的文件，指定到assets文件夹下生成文件。

#### 注意点
 - 两个go-bindata命令，运行的位置是不一样的。运行go-bindata的路径，也就决定了，生成的go文件里面，资源的路径。
 - 读取打包之后的资源，用的路径就是go文件最顶上自动生成的路径。

#### 运行结果
##### go-bindata对static1的命令
![Imgur](https://i.imgur.com/cTVYHay.png)

##### go-bindata对static2的命令
![Imgur](https://i.imgur.com/uhZ0ldy.png)

##### 运行结果
![Imgur](https://i.imgur.com/0Wtk2wp.png)