# os.File

## NOTICE
 - file.Read()方法使用之后，游标会移动，下次Read会从游标当前位置开始读取
 - file.Seek()，这个方法可以移动游标，有两个参数，一个是offset也就是偏移多少，
 还有一个whence表示从什么位置开始偏移，0的时候表示从文件头偏移，具体可以看源代码里的注释

## 运行结果
![Imgur](https://i.imgur.com/BKMfPdf.png)