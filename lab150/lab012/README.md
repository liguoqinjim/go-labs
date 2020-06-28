# middleware的done

## NOTICE
 - 要注意`ctx.Next()`方法，有的时候如果我们返回之后没有调用Next的话，done的middleware是不会被触发的
 - `SetExecutionRules`里面把Done设置为true的话，不管有没有调用Next，最后都会触发Done

## 参考资料
 - https://github.com/kataras/iris/wiki/Routing-middleware