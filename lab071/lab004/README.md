# html/template02

## 注意点
 - 这个里面调用Execute是正确的，是因为我们在生成template的时候没有先New再Parse，而是直接Parse
 - `ParseGlob`可以一次parse多个文件，只要符合pattern就可以了

## 运行结果
![Imgur](https://i.imgur.com/fElP9KI.png)