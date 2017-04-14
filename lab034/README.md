### Golang比较两个slice是否相等

#### 运行结果
`go test -bench=.`

![Imgur](http://i.imgur.com/5AIspfc.png)

#### 注意点
##### slice的len和cap区别
切片中有两个概念：一是len长度，二是cap容量，长度是指已经被赋过值的最大下标+1，可通过内置函数len()获得。容量是指切片目前可容纳的最多元素个数，可通过内置函数cap()获得。切片是引用类型，因此在当传递切片时将引用同一指针，修改值将会影响其他的对象。
```
s0 := make([]int, 5, 10) // len(s0) == 5, cap(s0) == 10
```

##### BCE
有的时候，我们可以故意写一些代码，来让编译器认为是可以BCE的。这样可以提高一点效率，这点可以在我们的最后两个Benchmark的运行结果看出。

#### 参考资料
1. https://kenshinsyrup.github.io/program/2017/04/11/Golang%E6%AF%94%E8%BE%83%E4%B8%A4%E4%B8%AAslice%E6%98%AF%E5%90%A6%E7%9B%B8%E7%AD%89/
2. http://www.tapirgames.com/blog/golang-1.7-bce
3. http://www.cnblogs.com/howDo/archive/2013/04/25/GoLang-Array-Slice.html