### 实验02，查看两种surf的默认user-agent
`DownloadId`0的时候是默认内核，1的时候是phantomjs的内核

#### 注意
1. 要注意phantomjs可执行文件的路径，在写这个实验的时候，库的代码还不稳定
2. 运行的时候会生成tmp文件夹，然后在tmp文件夹里面生成文件。调用`surfer.DestroyJsFiles()`可以删除生成的文件

#### 运行结果
![Imgur](https://i.imgur.com/mxsNrS1.png)