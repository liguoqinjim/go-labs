# demo

## NOTICE
 - `test.html`里面是没有<html><head>这些标签的，就是两个strong标签。这样的数据在goquery里面读取是没有问题的，
 但是读取出来的结果里面会自动的加上<html><head><body>标签，会把我们的数据嵌套在<body>标签里面
 - `Html()`方法返回的是一个标签里面的数据，如`<div><strong>a</strong></div`，对div这个node调用`Html()`，
 返回的结果会是`<strong>a</strong>`
 - 相对于`Html()`方法，我们要是想得到包含node本身的数据的话，我们调用`goquery.OuterHtml(node)`就可以了

## 运行结果
![Imgur](https://i.imgur.com/bQIkuS7.png)