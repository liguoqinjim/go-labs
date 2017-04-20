### 运行示例

#### 运行结果
![Imgur](http://i.imgur.com/tP8ZPou.png)

#### 编写顺序
1. 按照自己的需求，先编写`*.proto`文件，这个文件里面是自己的消息的具体结构
2. 用`protoc.exe`来根据写好的`*.proto`文件来编译出自己要用在哪个语言里面的类。
具体命令:`E:\下载\protoc-3.2.0-win32\bin\protoc.exe -I=E:\Workspace\go-labs\src\lab038\lab001 --go_out=E:\Workspace\go-labs\src\lab
          038\lab001 E:\Workspace\go-labs\src\lab038\lab001\addressbook.proto`

    ![Imgur](http://i.imgur.com/dk19O4B.png)
3. 在自己的程序中导入import生成好的类就可以了

#### 参考资料
https://developers.google.com/protocol-buffers/docs/gotutorial