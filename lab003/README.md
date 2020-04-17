# json包中的NewDecoder和NewEncoder
这个方式读取和marshal的读取方式的区别是，这个是读取一个输入流上的内容然后去解析。

## 运行结果
![Imgur](http://i.imgur.com/MiaEaCx.png)

## NOTICE
 - `{"Depth":0,"Ctx":{"page":1,"a":"b"}}`，ctx要是直接unmarshal到`map[string]interface{}`的话，page值会是float64的，因为json包里面的规则就是float64对应json的number，可以看参考资料2里面

## 参考资料
 - http://stackoverflow.com/questions/21197239/decoding-json-in-golang-using-json-unmarshal-vs-json-newdecoder-decode
 - https://golang.org/pkg/encoding/json/#Unmarshal