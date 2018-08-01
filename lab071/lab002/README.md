# must

## 注意点
 - must是在template在parse的时候，验证template是否正确
 - must需要两个参数，一个是template,一个是error。会直接判断是否template是否合法。
 这样会给我们省去一些代码。
 - 但是must会直接引起panic

## 运行结果
![Imgur](https://i.imgur.com/Zr534nz.png)