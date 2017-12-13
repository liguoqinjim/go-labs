### example02，总结play009(读取)

#### 注意点
 - reflect.Type，和类型相关的都在这个里面找，比如有多少个字段，字段名称，字段类型是什么
 - reflect.Value，和值相关的用这个。比如这个字段的值是多少
 - struct和pointer在使用reflect的时候是不一样的。`FieldByName()`只能在struct使用。pointer的时候要用`Elem()`
 - `Value`和`Type`才有`Kind()`方法，`StructField`是没有的

#### 运行结果
![Imgur](https://i.imgur.com/DH3ESpe.png)

#### 参考资料
1. http://www.jianshu.com/p/0d346577d32f
2. https://github.com/liguoqinjim/go-playground/tree/master/play009