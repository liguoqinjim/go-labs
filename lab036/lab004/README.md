# 打印出方法名和行号01，demo
logrus本身没有方法可以打印行号，这里要使用hook才行。目前看到的是可以用参考资料1里面的hook

## 注意点
 - 我们可以控制hook的skip参数，也就是hook里面调用runtime.caller的参数
 - hook里面的逻辑是一旦runtime.caller的file不是logrus，就退出逻辑了。如果有需要的话，是要自己改源码的
 - 安装：`go get -u -v github.com/onrik/logrus`

## todo
多个包log的时候，我们用这个hook有的时候无法打印，所以要修改一下代码。现在还是粗改版本

```
func findCaller(skip int) (string, string, int) {
	var (
		pc       uintptr
		file     string
		function string
		line     int
	)
	bn := 0
	for i := 0; i < 10; i++ {
		pc, file, line = getCaller(skip + i)
		fmt.Println(file)
		if !strings.HasPrefix(file, "logrus/") {
			//liguoqinjim
			if bn == 1 {
				break
			}
			bn++
		}
	}
	if pc != 0 {
		frames := runtime.CallersFrames([]uintptr{pc})
		frame, _ := frames.Next()
		function = frame.Function
	}

	return file, function, line
}
```

## 运行结果
![Imgur](https://i.imgur.com/wk1NS7M.png)

## 参考资料
1. https://github.com/onrik/logrus
2. https://github.com/sirupsen/logrus/issues/63