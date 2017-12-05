### 生成csv

#### 注意点
 - `f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM`，这个要加在开头，不然在excel里面打开的时候会乱码

#### 参考资料
https://studygolang.com/articles/555