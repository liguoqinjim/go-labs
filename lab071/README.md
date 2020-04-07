# golang/template

|实验|简介|说明|
|---|---|---|
|lab001|text/template例子| |
|lab002|Must| |
|lab003|ParseFiles| |
|lab004|ParseGlob| |
|lab005|遍历slice|{{index . 3}}|
|lab006|遍历map| |

## NOTICE
 - `text/template`和`html/template`差不多，但是在用于网页的时候`html/template`会更安全，可以防止代码注入
 - `html/template`会转义字符，比如+号会改成`&#43:`，要是不想转义就使用`text/template`


## 参考资料
 - https://blog.gopheracademy.com/advent-2017/using-go-templates/
 - github_emoji_generator

## 例子
这个可以遍历一个struct的slice,AttrName就是struct的字段
```
|-|数据|简介|例子|说明|
|:---|---|---|---|---|{{ range $index, $element := .Fields }}
|{{$index}}|{{$element.AttrName}}|{{$element.Desc}}|{{$element.Example}}||{{ end }}
```