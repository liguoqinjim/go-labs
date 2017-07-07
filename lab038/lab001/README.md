### 运行demo
运行教程中给出的demo

#### protoc命令
`protoc -I=E:\Workspace\go-labs\src\lab038\lab001\pb --go_out=E:\Workspace\go-labs\src\lab038\lab001\pb E:\Workspace\go-labs\src\lab038\lab001\pb\addressbook.proto`

#### 命令格式
`protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/addressbook.proto`

#### 运行结果
![Imgur](http://i.imgur.com/sRSLmA7.png)

#### 参考资料
https://developers.google.com/protocol-buffers/docs/gotutorial