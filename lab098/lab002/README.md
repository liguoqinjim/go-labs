### 解析csv

#### 注意点
 - 两个运行结果有点不一样，goland的那个1前面会多一个空格，但是cmder里面这个空格就没有了。
 - goland里面有空格是因为，这个解析的csv是lab001生成的结果，lab001里面为了要utf8，在开头写了一个字符。这个就是导致空格的问题。不加这个字符生成的csv读取是没问题的。
 - 参考资料2里面，有说不推荐在开头强制写入bom来变成utf8，自己要取舍。

#### 运行结果
![Imgur_goland运行结果](https://i.imgur.com/OQkcJ1I.png)

![Imgur_cmder运行结果](https://i.imgur.com/6cybpw9.png)

#### 参考资料
1. https://www.thepolyglotdeveloper.com/2017/03/parse-csv-data-go-programming-language/
2. https://github.com/golang/go/issues/9588