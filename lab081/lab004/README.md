# 给定表示时间的字符串如("2017-02-08")，转换到time
只要指定parse中用的layout参数就可以解析

## NOTICE
 - `ParseInLocation`，相较于`Parse`多了时区的选择。为了保险，一般就用ParseInLocation好了

## 参考资料
https://studygolang.com/articles/3753

## 运行结果
![Imgur](https://i.imgur.com/7cu1VTg.png)