# [filetype](https://github.com/h2non/filetype)

|实验|简介|说明|
|---|---|---|
|lab001|example| |
|lab002|测试excel文件||

# NOTICE
 - 存在有的xlsx会判断成zip包的情况，参考：`https://github.com/h2non/filetype/issues/38`。
 - https://godoc.org/net/http#DetectContentType，这个函数也可以判断出contentType，而且没有找到的时候回返回`application/octet-stream`。  
 这个特性符合一般对象存储的逻辑。但是如果是1中的xlsx的话，这个函数也是会判断为zip包的