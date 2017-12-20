### 指定生成文件的路径
绑定static文件夹中的文件，指定到assets文件夹下生成文件。

#### 运行命令
在static文件夹路径下执行
` go-bindata.exe -pkg assets -o ..\assets\assets.go .`